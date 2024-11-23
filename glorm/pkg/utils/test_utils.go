package utils

import (
	"fmt"
	"reflect"
)

func CompareInterfaceSlices(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if b[i] == nil {
			return false // Key does not exist in the second map
		}

		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func CompareMaps(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for key, aValue := range a {
		bValue, exists := b[key]
		if !exists {
			return false // Key does not exist in the second map
		}

		// Use reflect.DeepEqual to compare values of different types
		if !reflect.DeepEqual(aValue, bValue) {
			return false // Values are not equal
		}
	}

	// All keys and values are equal
	return true
}

func ConvertMap(rv reflect.Value) (map[string]interface{}, error) {
	if rv.Kind() != reflect.Map {
		return nil, fmt.Errorf("expected a map, got %s", rv.Kind())
	}

	// Create a new map[string]interface{}
	newMap := make(map[string]interface{})

	// Iterate over the map's keys and values
	for _, key := range rv.MapKeys() {
		// Assert that the key is of type string
		if key.Kind() != reflect.String {
			return nil, fmt.Errorf("map keys must be of type string, got %s", key.Kind())
		}
		// Get the value corresponding to the key
		value := rv.MapIndex(key)

		// Store in the new map
		newMap[key.String()] = value.Interface()
	}

	return newMap, nil
}
