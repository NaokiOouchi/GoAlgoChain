package algo

type GroupedItem[K comparable, T any] struct {
	Key   K
	Items []T
}

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
