package ledlib

import (
	"errors"
	"fmt"
	"ledlib/util"
	"reflect"
)

func CreateObject(order map[string]interface{}, ledCanvas LedCanvas) (interface{}, float64, error) {
	value, err := GetJSONValue(order, "id")
	ilifetime, _ := GetJSONValueOrDefault(order, "lifetime", float64(10.0))
	if err != nil {
		return nil, 0, err
	}
	if id, ok := value.(string); ok {

		if lifetime, ok := ilifetime.(float64); ok {
			switch id {
			/*
				Objects
			*/
			case "object-blank":
				return NewObjectFill(util.NewFromRGB(0, 0, 0)), lifetime, nil
			case "object-rocket":
				return NewObjectRocket(), lifetime, nil
				/*
					Filters
				*/
			case "filter-rolling":
				return NewFilterRolling(ledCanvas), 0, nil
			case "filter-skewed":
				return NewFilterSkewed(ledCanvas), 0, nil
			case "filter-snows":
				return NewFilterSnows(ledCanvas), 0, nil
			case "filter-mountain":
				return NewFilterMountain(ledCanvas), 0, nil
			case "filter-grass":
				return NewFilterGrass(ledCanvas), 0, nil

			default:
				return nil, 0, errors.New("Unnown Object Id")
			}
		} else {
			return nil, 0, fmt.Errorf("Unknown Error: %s", reflect.TypeOf(ilifetime))
		}
	} else {
		return nil, 0, fmt.Errorf("Unknown Error: %s", reflect.TypeOf(value))
	}

	return nil, 0, errors.New("Unnown Error")
}
