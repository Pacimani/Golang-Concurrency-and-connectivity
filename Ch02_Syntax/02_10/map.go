package main

import "fmt"

func main() {

	stocks := map[string]float64{
		"MSFT": 100.00,
		"GOOG": 200.00,
		"FB":   300.00,
	}

	fmt.Println(stocks)
	fmt.Println(len(stocks))

	// GET value from the map
	fmt.Println(stocks["MSFT"])

	// get zero value if not found
	fmt.Println(stocks["TSLA"])

	// use two value to see if value is present in the map

	value, ok := stocks["TSLA"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("TSLA not found")
	}

	value, found := stocks["MSFT"]
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("MSFT not found")
	}

	// set or add new key and value
	stocks["TSLA"] = 400.00
	fmt.Println(stocks)

	// delete a stock
	delete(stocks, "TSLA")
	fmt.Println(stocks)

	// loop through the map
	for key, value := range stocks {
		fmt.Println(key, value)
	}

	// convert map to slice
	keys := make([]string, 0, len(stocks))
	for k := range stocks {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	// this is to convert slice to map
	stocks2 := make(map[string]float64)
	for i := 0; i < len(keys); i++ {
		stocks2[keys[i]] = stocks[keys[i]]
	}
	fmt.Println(stocks2)
}
