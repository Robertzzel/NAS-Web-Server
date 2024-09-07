package utils

func First[T any](l []T) Maybe[T] {
	if l == nil || len(l) < 1 {
		return None[T]()
	}
	return Some(l[0])
}

func FirstOr[T any](l []T, defaultValue T) T {
	if l == nil || len(l) < 1 {
		return defaultValue
	}
	return l[0]
}

func LastOr[T any](l []T, defaultValue T) T {
	if l == nil || len(l) < 1 {
		return defaultValue
	}
	return l[len(l)-1]
}

func Last[T any](l []T) Maybe[T] {
	if l == nil || len(l) < 1 {
		return None[T]()
	}
	return Some(l[len(l)-1])
}

func Skip[T any](l []T, number uint) []T {
	if len(l) < int(number) {
		return make([]T, 0)
	}
	result := make([]T, len(l)-int(number))
	copy(result, l[number:])
	return result
}

func Take[T any](l []T, number uint) []T {
	if len(l) < int(number) {
		return make([]T, 0)
	}
	result := make([]T, int(number))
	copy(result, l)
	return result
}

func Map[T any, V any](l []T, f func(T) V) []V {
	result := make([]V, len(l))
	for i, v := range l {
		result[i] = f(v)
	}
	return result
}

func ForEach[T any](l []T, f func(T)) {
	for _, v := range l {
		f(v)
	}
}

func Zip[T any, V any, U any](l []T, l2 []U, f func(T, U) V) []V {
	minSize := len(l)
	if len(l2) < minSize {
		minSize = len(l2)
	}
	result := make([]V, minSize)
	for i := 0; i < minSize; i++ {
		result[i] = f(l[i], l2[i])
	}
	return result
}

func MapWithIndex[T any, V any](l []T, f func(T, int) V) []V {
	result := make([]V, len(l))
	for i, v := range l {
		result[i] = f(v, i)
	}
	return result
}

func Filter[T any](l []T, pred func(T) bool) []T {
	var result = make([]T, 0)
	for _, v := range l {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterSome[T any](l []Maybe[T]) []T {
	var result = make([]T, 0)
	for _, v := range l {
		if v.IsSome() {
			result = append(result, v.Unwrap())
		}
	}
	return result
}

func FilterOk[T any](l []Result[T]) []T {
	var result = make([]T, 0)
	for _, v := range l {
		if v.IsOk() {
			result = append(result, v.Unwrap())
		}
	}
	return result
}

func FilterWithIndex[T any](l []T, pred func(T, int) bool) []T {
	var result = make([]T, 0)
	for index, v := range l {
		if pred(v, index) {
			result = append(result, v)
		}
	}
	return result
}

func ReduceWithIndex[T any](l []T, reducer func(T, T, int) T, initial T) T {
	acc := initial
	for index, v := range l {
		acc = reducer(acc, v, index)
	}
	return acc
}

func Reduce[T any, R any](l []T, reducer func(R, T) R, initial R) R {
	acc := initial
	for _, v := range l {
		acc = reducer(acc, v)
	}
	return acc
}

func Any[T any](l []T, pred func(T) bool) bool {
	for _, v := range l {
		if pred(v) {
			return true
		}
	}
	return false
}

func Empty[T any](l []T) bool {
	return len(l) < 1
}

func All[T any](l []T, pred func(T) bool) bool {
	for _, v := range l {
		if !pred(v) {
			return false
		}
	}
	return true
}

func FlatMap[T any](l []T, f func(T) []T) []T {
	var result []T
	for _, v := range l {
		result = append(result, f(v)...)
	}
	return result
}

func Reverse[T any](l []T) []T {
	listSize := len(l)
	reversed := make([]T, listSize)
	for i, v := range l {
		reversed[listSize-1-i] = v
	}
	return reversed
}

func TakeWhile[T any](lst []T, f func(elem T, index int) bool) []T {
	result := []T{}
	for index, item := range lst {
		if !f(item, index) {
			break
		}
		result = append(result, item)
	}
	return result
}

func SkipUntil[T any](lst []T, f func(elem T, index int) bool) []T {
	var index int
	for i, item := range lst {
		if f(item, i) {
			index = i
		}
	}
	result := make([]T, len(lst[index:]))
	copy(result, lst[index:])
	return result
}

type Maybe[T any] struct {
	value     T
	isPresent bool
}

func Some[T any](value T) Maybe[T] {
	return Maybe[T]{value: value, isPresent: true}
}

func None[T any]() Maybe[T] {
	return Maybe[T]{isPresent: false}
}

func (m Maybe[T]) IsSome() bool {
	return m.isPresent
}

func (m Maybe[T]) IsNone() bool {
	return !m.isPresent
}

func (m Maybe[T]) Unwrap() T {
	if !m.isPresent {
		panic("attempted to unwrap a None value")
	}
	return m.value
}

func (m Maybe[T]) UnwrapOr(defaultValue T) T {
	if !m.isPresent {
		return defaultValue
	}
	return m.value
}

type Result[T any] struct {
	value   T
	error   error
	isError bool
}

func Ok[T any](value T) Result[T] {
	return Result[T]{value: value, isError: false}
}

func Error[T any](err error) Result[T] {
	return Result[T]{error: err, isError: true}
}

func (r Result[T]) IsOk() bool {
	return !r.isError
}

func (r Result[T]) IsError() bool {
	return r.isError
}

func (r Result[T]) Unwrap() T {
	if r.isError {
		panic("attempted to unwrap a Result error value")
	}
	return r.value
}

func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.isError {
		return defaultValue
	}
	return r.value
}

func (r Result[T]) Error() error {
	return r.error
}

func BindResult[T any](value T, err error) Result[T] {
	if err != nil {
		return Error[T](err)
	}
	return Ok(value)
}

type Pair[T any, V any] struct {
	First  T
	Second V
}

func NewPair[T any, V any](first T, second V) Pair[T, V] {
	return Pair[T, V]{first, second}
}
