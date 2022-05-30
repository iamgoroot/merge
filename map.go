package merge

//Map usage: merge.Map[string, struct{}] same as map[string]struct{} but has method Merge(merge.Map[string, struct{}]) merge.Map[string, struct{}]
type Map[K comparable, V any] map[K]V

func (left Map[K, V]) Merge(right Map[K, V]) Map[K, V] {
	return mergeMap[K, V, map[K]V](left, right)
}

func mergeMap[K comparable, V any, T map[K]V](left, right T) T {
	if left == nil {
		left = map[K]V{}
	}
	for k, v := range right {
		left[k] = v
	}
	return left
}
