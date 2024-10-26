package main

import (
	"fmt"
	"strconv"
)

// DataProcessor struct
type DataProcessor struct{}

// Process method processes different types of data
func (dp *DataProcessor) Process(data interface{}) string {
	switch v := data.(type) {
	case string:
		return "Processed string: " + v
	case int:
		return "Processed integer: " + strconv.Itoa(v)
	case float64:
		return fmt.Sprintf("Processed float: %.2f", v)
	default:
		return "Unknown type"
	}
}

func main() {
	dp := &DataProcessor{}

	// Different types of data
	dataItems := []interface{}{
		"Hello, Go!",
		42,
		3.14,
		true, // An unsupported type for this processor
	}

	// Process each data item
	for _, item := range dataItems {
		result := dp.Process(item)
		fmt.Println(result)
	}
}
