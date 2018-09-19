package ledlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type LedBlockRenderer interface {
	Abort()
	Start()
	Terminate()
	Show(blocks string)
}

func NewLedBlockRenderer() LedBlockRenderer {
	return &ledBockRendererImpl{
		make(chan string),
		make(chan struct{})}
}

func GetJSONValue(m interface{}, key string) (interface{}, error) {
	value, _ := GetJSONValueOrDefault(m, key, nil)
	if value == nil {
		return nil, errors.New("invalid json format")
	} else {
		return value, nil
	}
}

func GetJSONValueOrDefault(m interface{}, key string, defaults interface{}) (interface{}, error) {
	if mm, ok := m.(map[string]interface{}); ok {
		if val, ok := mm[key]; ok {
			return val, nil
		}
	}
	return defaults, errors.New("invalid json format")
}

func getOrdersFromJson(rawJson string) ([]interface{}, error) {
	var ordersMap interface{}
	err := json.Unmarshal([]byte(rawJson), &ordersMap)
	if err != nil {
		return nil, err
	}

	if val, ok := ordersMap.(map[string]interface{}); ok {
		if val, ok := val["orders"]; ok {
			if val, ok := val.([]interface{}); ok {
				return val, nil
			}
		}
	}
	return nil, errors.New("invalid json format")
}

func getOrdersInLoop(orders []interface{}, start int) ([]interface{}, error) {
	ordersInLoop := make([]interface{}, 0)
	for i := start; i < len(orders); i++ {
		order := orders[i]
		mapOrder := order.(map[string]interface{})
		if val, ok := mapOrder["id"]; ok {
			if val.(string) == "ctrl-loop" {
				return ordersInLoop, nil
			}
			ordersInLoop = append(ordersInLoop, orders[i])
		} else {
			return nil, errors.New("invalid json format")
		}

	}
	return ordersInLoop, nil
}

func expands(orders []interface{}, count int) []interface{} {
	newOrders := make([]interface{}, 0)
	for i := 0; i < count; i++ {
		newOrders = append(newOrders, orders...)
	}
	return newOrders
}

func flattenOrders(orders []interface{}) ([]interface{}, error) {
	flatten := make([]interface{}, 0)

	for i := 0; i < len(orders); i++ {
		if val, err := GetJSONValueOrDefault(orders[i], "id", nil); err == nil {
			if val.(string) == "ctrl-loop" {
				ordersInLoop, err := getOrdersInLoop(orders, i+1)
				if err != nil {
					return nil, errors.New("invalid order format")
				}
				count, _ := GetJSONValueOrDefault(orders[i], "count", 3)
				flatten = append(flatten, expands(ordersInLoop, count.(int))...)
				i += len(ordersInLoop) + 1
			} else {
				flatten = append(flatten, orders[i])
			}
		} else {
			return nil, errors.New("invalid order array. key: id not found")
		}
	}
	return flatten, nil
}

type ledBockRendererImpl struct {
	orderCh chan string
	doneCh  chan struct{}
}

func (l *ledBockRendererImpl) Terminate() {
	close(l.orderCh)
	<-l.doneCh
}

func (l *ledBockRendererImpl) Abort() {
	l.orderCh <- `{"orders":[{"id":"object-blank", "lifetime":1}]}`
}

func (l *ledBockRendererImpl) Show(blocks string) {
	l.orderCh <- blocks
}

func (l *ledBockRendererImpl) Start() {

	go func() {
		defer func() { close(l.doneCh) }()

		baseCanvas := NewLedCanvas()
		var object LedObject
		var filters LedCanvas
		var orders []interface{}
		var lifetime float64
		var startTime int64
		isExpired := true

		for {
			select {
			case t := <-l.orderCh:
				switch t {
				case "":
					fmt.Println("terminated")
					return
				default:
					fmt.Println(t)
					if arrayOrders, err := getOrdersFromJson(t); err != nil {
						//error
					} else if flattenOrders, err := flattenOrders(arrayOrders); err != nil {
						//error
					} else {
						object, filters, lifetime, orders, err = GetFilterAndObject(flattenOrders, baseCanvas)
						if err != nil {
							isExpired = true
						} else {
							isExpired = false
							startTime = time.Now().Unix()
						}
					}
				}
			default:
				if lifetime != 0 &&
					(time.Now().Unix()-startTime) > int64(lifetime) {
					var err error
					object, filters, lifetime, orders, err = GetFilterAndObject(orders, filters)
					if err != nil {
						isExpired = true
					}
				}

				if !isExpired {
					filters.PreShow()
					object.Draw(filters)
				} else {
					// idle
				}
			}
		}
	}()
}

func GetFilterAndObject(iOrders []interface{}, canvas LedCanvas) (LedObject, LedCanvas, float64, []interface{}, error) {

	filter := canvas
	var object LedObject
	jsonOrders := iOrders
	for {
		if len(jsonOrders) == 0 {
			// invalid order
			return nil, nil, 0, iOrders, errors.New("invalid order format")
		}
		rawOrder := jsonOrders[0]
		jsonOrders = jsonOrders[1:]

		if jsonOrder, ok := rawOrder.(map[string]interface{}); ok {
			order, lifetime, err := CreateObject(jsonOrder, filter)
			if err != nil {
				return nil, nil, 0, iOrders, err
			}

			switch v := order.(type) {
			case LedObject:
				object = v
				return object, filter, lifetime, jsonOrders, nil
			case LedCanvas:
				filter = v
			default:
				return nil, nil, 0, iOrders, errors.New("invalid order format")
			}
		}
	}
}
