package ledlib

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJSONValue1(t *testing.T) {
	rawOrder := `{"id":"test"}`
	var order interface{}
	json.Unmarshal([]byte(rawOrder), &order)

	actual, err := GetJSONValueOrDefault(order, "id", nil)
	assert.Nil(t, err)
	assert.Equal(t, "test", actual.(string))

}

func TestGetJSONValueErrorCase1(t *testing.T) {
	rawOrder := `{"id":"test"}`
	var order interface{}
	json.Unmarshal([]byte(rawOrder), &order)

	actual, err := GetJSONValueOrDefault(order, "key", "default")
	assert.NotNil(t, err)
	assert.Equal(t, "default", actual.(string))
}

func TestGetJSONValueErrorCase2(t *testing.T) {
	rawOrder := `aaaaa`
	var order interface{}
	json.Unmarshal([]byte(rawOrder), &order)

	actual, err := GetJSONValueOrDefault(order, "key", "default")
	assert.NotNil(t, err)
	assert.Equal(t, "default", actual.(string))
}

func TestOrderInLoop1(t *testing.T) {

	rawOrders := `[{"id":"test"},{"id":"ctrl-loop"},{"id":"test2"},{"id":"ctrl-loop"},{"id":"test"}]`
	var orders interface{}
	json.Unmarshal([]byte(rawOrders), &orders)

	actual, err := getOrdersInLoop(orders.([]interface{}), 0)

	assert.Nil(t, err)
	assert.Len(t, actual, 1)
	order := actual[0].(map[string]interface{})
	assert.Equal(t, "test", order["id"].(string))
}

func TestOrderInLoop2(t *testing.T) {

	rawOrders := `[{"id":"test"},{"id":"ctrl-loop"},{"id":"test2"},{"id":"ctrl-loop"},{"id":"test"}]`
	var orders interface{}
	json.Unmarshal([]byte(rawOrders), &orders)

	actual, err := getOrdersInLoop(orders.([]interface{}), 1)

	assert.Nil(t, err)
	assert.Len(t, actual, 0)
}

func TestOrderInLoop3(t *testing.T) {

	rawOrders := `[{"id":"test"},{"id":"ctrl-loop"},{"id":"test2"},{"id":"ctrl-loop"},{"id":"test"}]`
	var orders interface{}
	json.Unmarshal([]byte(rawOrders), &orders)

	actual, err := getOrdersInLoop(orders.([]interface{}), 2)

	assert.Nil(t, err)
	assert.Len(t, actual, 1)
	order := actual[0].(map[string]interface{})
	assert.Equal(t, "test2", order["id"].(string))
}

func TestGetOrdersFromJson(t *testing.T) {
	orders := `{"orders":[{"id":"test"},{"id":"ctrl-loop"},{"id":"test2"},{"id":"ctrl-loop"},{"id":"test"}]}`
	if target, err := getOrdersFromJson(orders); err == nil {
		assert.NotEqual(t, 0, len(target))
	} else {
		t.Fail()
	}
}
func TestGetOrdersFromJson_ErrorCase1(t *testing.T) {
	orders := `{"xxxxx":[{"id":"test"},{"id":"ctrl-loop"},{"id":"test2"},{"id":"ctrl-loop"},{"id":"test"}]}`
	if _, err := getOrdersFromJson(orders); err == nil {
		t.Fail()
	}
}

func TestGetOrdersFromJson_ErrorCase2(t *testing.T) {
	orders := `{"orders":"test"}`
	if _, err := getOrdersFromJson(orders); err == nil {
		t.Fail()
	}
}

func TestExpands(t *testing.T) {
	rawOrders := `[{"id":"test"},{"id":"test2"}]`
	var orders interface{}
	json.Unmarshal([]byte(rawOrders), &orders)

	target := expands(orders.([]interface{}), 3)

	assert.Len(t, target, 3*2)

	var order map[string]interface{}

	for i := 0; i < 3; i += 2 {
		order = target[i].(map[string]interface{})
		assert.Equal(t, "test", order["id"].(string))
		order = target[i+1].(map[string]interface{})
		assert.Equal(t, "test2", order["id"].(string))

	}
}

func testFlatternOrder(t *testing.T, orders string, expectIDs []string) {
	arrayOrders, _ := getOrdersFromJson(orders)
	flattenOrders, err := flattenOrders(arrayOrders)
	assert.Nil(t, err)

	assert.Len(t, flattenOrders, len(expectIDs))

	for i, expectID := range expectIDs {
		actual, _ := GetJSONValueOrDefault(flattenOrders[i], "id", nil)
		assert.Equal(t, actual.(string), expectID)
	}
}

func TestFlattenOrders1(t *testing.T) {

	orders := `{"orders":[{"id":"test"},{"id":"ctrl-loop"},{"id":"test2"},{"id":"ctrl-loop"},{"id":"test"}]}`
	expectIDs := []string{"test", "test2", "test2", "test2", "test"}
	testFlatternOrder(t, orders, expectIDs)
}

func TestFlattenOrders2(t *testing.T) {

	orders := `{"orders":[{"id":"test"},{"id":"ctrl-loop"},{"id":"test2"},{"id":"test2"},{"id":"test"}]}`
	expectIDs := []string{"test", "test2", "test2", "test", "test2", "test2", "test", "test2", "test2", "test"}
	testFlatternOrder(t, orders, expectIDs)
}

func TestFlattenOrders3(t *testing.T) {

	orders := `{"orders":[{"id":"test"},{"id":"ctrl-loop"},{"id":"ctrl-loop"},{"id":"test2"},{"id":"test"}]}`
	expectIDs := []string{"test", "test2", "test"}
	testFlatternOrder(t, orders, expectIDs)
}

func TestFlattenOrders4(t *testing.T) {

	orders := `{"orders":[{"id":"test"},{"id":"test2"},{"id":"test"},{"id":"test"},{"id":"ctrl-loop"}]}`
	expectIDs := []string{"test", "test2", "test", "test"}
	testFlatternOrder(t, orders, expectIDs)
}

func TestFlattenOrders5(t *testing.T) {

	orders := `{"orders":[{"id":"ctrl-loop"}]}`
	expectIDs := []string{}
	testFlatternOrder(t, orders, expectIDs)
}

func TestFlattenOrders6(t *testing.T) {

	orders := `{"orders":[{"id":"ctrl-loop"},{"id":"ctrl-loop"},{"id":"ctrl-loop"},{"id":"ctrl-loop"},{"id":"ctrl-loop"},{"id":"ctrl-loop"}]}`
	expectIDs := []string{}
	testFlatternOrder(t, orders, expectIDs)
}

func TestFlattenOrders7(t *testing.T) {

	orders := `{"orders":[{"id":"ctrl-loop"},{"id":"test"}]}`
	expectIDs := []string{"test", "test", "test"}
	testFlatternOrder(t, orders, expectIDs)
}

func TestGetFilterAndObject1(t *testing.T) {
	orders := `{"orders":[{"id":"filter-skewed"},{"id":"filter-rolling"},{"id":"object-rocket", "lifetime":11}]}`
	arrayOrders, _ := getOrdersFromJson(orders)
	flattenOrders, _ := flattenOrders(arrayOrders)

	object, filters, lifetime, newOrders, _, err := GetFilterAndObject(flattenOrders, nil, NewLedCanvasParam())
	assert.Nil(t, err)
	assert.NotNil(t, object)
	assert.NotNil(t, filters)
	assert.Equal(t, 0, len(newOrders))
	assert.Equal(t, 11.0, lifetime)

}

func TestGetFilterAndObject2(t *testing.T) {
	orders := `{"orders":[{"id":"filter-skewed"},{"id":"filter-rolling"},{"id":"object-rocket", "lifetime":11},{"id":"filter-skewed"},{"id":"filter-rolling"},{"id":"object-rocket", "lifetime":11}]}`
	arrayOrders, _ := getOrdersFromJson(orders)
	flattenOrders, _ := flattenOrders(arrayOrders)

	object, filters, lifetime, newOrders, _, err := GetFilterAndObject(flattenOrders, nil, NewLedCanvasParam())
	assert.Nil(t, err)
	assert.NotNil(t, object)
	assert.NotNil(t, filters)
	assert.Equal(t, 3, len(newOrders))
	assert.Equal(t, 11.0, lifetime)

	object, filters, lifetime, newOrders, _, err = GetFilterAndObject(flattenOrders, nil, NewLedCanvasParam())
	assert.Nil(t, err)
	assert.NotNil(t, object)
	assert.NotNil(t, filters)
	assert.Equal(t, 3, len(newOrders))
	assert.Equal(t, 11.0, lifetime)
}

func TestGetFilterAndObjectErrorCase1(t *testing.T) {
	orders := `{"orders":[{"id":"filter-skewed"},{"id":"filter-rolling"}]}`
	arrayOrders, _ := getOrdersFromJson(orders)
	flattenOrders, _ := flattenOrders(arrayOrders)

	object, filters, lifetime, newOrders, _, err := GetFilterAndObject(flattenOrders, nil, NewLedCanvasParam())
	assert.NotNil(t, err)
	assert.Nil(t, object)
	assert.Nil(t, filters)
	assert.Equal(t, 2, len(newOrders))
	assert.Equal(t, 0.0, lifetime)

}
