package algo

// GroupedItem represents a group of items sharing the same key.
type GroupedItem[K comparable, T any] struct {
	Key   K
	Items []T
}

// GroupBy organizes items into groups based on a key function.
// Items with the same key are collected into the same group while maintaining their order.
//
// Example:
//
//	orders := []Order{
//	    {ID: 1, UserID: 100, Item: "Book"},
//	    {ID: 2, UserID: 101, Item: "Pen"},
//	    {ID: 3, UserID: 100, Item: "Paper"},
//	}
//
//	groups := GroupBy(orders, func(o Order) int { return o.UserID })
//	// Results in groups by UserID: [100: [Order{1}, Order{3}], 101: [Order{2}]]
func GroupBy[T any, K comparable](data []T, keyFunc func(T) K) []GroupedItem[K, T] {
	groupMap := make(map[K][]T)
	for _, item := range data {
		key := keyFunc(item)
		groupMap[key] = append(groupMap[key], item)
	}

	groupedItems := make([]GroupedItem[K, T], 0, len(groupMap))
	for key, items := range groupMap {
		groupedItems = append(groupedItems, GroupedItem[K, T]{Key: key, Items: items})
	}

	return groupedItems
}
