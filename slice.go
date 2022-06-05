package merge

type Merge[T any] interface {
	Merge(T) T
}

var _ Merge[Slice[string]] = &Slice[string]{}

//Slice usage: merge.Slice[string] same as []string but has method Merge([]string) []string
type Slice[V any] []V

func (left *Slice[V]) Merge(right Slice[V]) Slice[V] {
	if right == nil {
		return *left
	}
	*left = mergeSlice[V, []V](*left, right)
	return *left
}

func mergeSlice[V any, T []V](left, right T) (result T) {
	return append(left, right...)
}
