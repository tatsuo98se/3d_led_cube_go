package ledlib

import (
	"encoding/json"
	"errors"
	"fmt"
)

type LedBlockRenderer interface {
	Abort()
	Start()
	Terminate()
	Show(blocks map[string]interface{})
}

func NewLedBlockRenderer() LedBlockRenderer {
	return &ledBockRendererImpl{
		make(chan map[string]interface{}),
		make(chan struct{})}
}

func getParam2(m map[string]interface{}, key string, d interface{}) interface{} {
	if val, ok := m[key]; ok {
		return val
	}
	return d
}
func getParam(m interface{}, key string, defaults interface{}) (interface{}, error) {
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
		if val, err := getParam(orders[i], "id", nil); err == nil {
			if val.(string) == "ctrl-loop" {
				ordersInLoop, err := getOrdersInLoop(orders, i+1)
				if err != nil {
					return nil, errors.New("invalid order format")
				}
				count, _ := getParam(orders[i], "count", 3)
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
	orderCh chan map[string]interface{}
	doneCh  chan struct{}
}

func (l *ledBockRendererImpl) Terminate() {
	close(l.orderCh)
	<-l.doneCh
}

func (l *ledBockRendererImpl) Abort() {
	l.orderCh <- map[string]interface{}{
		"id": "ctrl-filter-clear",
	}
}

func (l *ledBockRendererImpl) Show(blocks map[string]interface{}) {
	l.orderCh <- blocks
}

func (l *ledBockRendererImpl) Start() {

	go func() {
		defer func() { close(l.doneCh) }()
		for {
			select {
			case t := <-l.orderCh:
				if t == nil {
					fmt.Println("terminated")
					return
				} else {
					//new order arrival
					// clean and setup up drawing frames next
				}
			default:
				fmt.Println("...")
				// drawframe
			}
		}
	}()
}
