package main

import "fmt"

type Item struct {
	Code string
	Sn   string
	Cost float64
}

func main() {
	list1 := []Item{
		{Code: "code1", Sn: "sn1", Cost: 100},
		{Code: "code2", Sn: "sn2", Cost: 200},
	}

	list2 := []Item{
		{Code: "code1", Sn: "sn1", Cost: 50},
		{Code: "code3", Sn: "sn3", Cost: 300},
	}

	itemsMap := make(map[string]Item)

	for _, item := range append(list1, list2...) {
		if existingItem, ok := itemsMap[item.Code]; ok {
			// If the item already exists, simply add the cost
			existingItem.Cost += item.Cost
			itemsMap[item.Code] = existingItem
		} else {
			// If the item doesn't exist yet, add it to the map
			itemsMap[item.Code] = item
		}
	}

	// If you need the deduplicated items in a slice
	var deduplicatedList []Item
	for _, v := range itemsMap {
		deduplicatedList = append(deduplicatedList, v)
	}

	fmt.Println(deduplicatedList)

	fmt.Println(itemsMap)
}
