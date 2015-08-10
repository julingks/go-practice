package main

import (
	"fmt"
	"reflect"
	"sort"
)

func Resemble(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
OutLoop:
	for _, v1 := range a {
		for i, v2 := range b {
			if v1 == v2 {
				b = append(b[:i], b[i+1:]...)
				continue OutLoop
			}
		}
		return false
	}
	return true
}

func Any(a []int, b ...int) bool {
	for _, v1 := range a {
		for _, v2 := range b {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func All(a []int, b int) bool {
	for _, v := range a {
		if v != b {
			return false
		}
	}
	return true
}

func Frequency(a []int, b int) int {
	c := 0
	for _, v := range a {
		if v == b {
			c++
		}
	}
	return c
}

type IntType []int

func (p IntType) Len() int           { return len(p) }
func (p IntType) Less(i, j int) bool { return p[i] < p[j] }
func (p IntType) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Sort(a []int) {
	sort.Sort(IntType(a))
}

func SortReverse(a []int) {
	sort.Sort(sort.Reverse(IntType(a)))
}

func Concat(a ...[]int) []int {
	b := make([]int, 0)
	for _, v := range a {
		b = append(b, v...)
	}
	return b
}

func IndexOf(a []int, x int) int {
	for i, v := range a {
		if v == x {
			return i
		}
	}
	return -1
}

func Partition(a []int, f func(idx int, value int) bool) ([]int, []int) {
	b := make([]int, 0)
	c := make([]int, 0)
	for i, v := range a {
		if f(i, v) {
			b = append(b, v)
		} else {
			c = append(c, v)
		}
	}
	return b, c
}

type IntS []int

func (a IntS) Filter(filter func(idx int, value int) bool) IntS {
	b := make(IntS, 0)
	for i, v := range a {
		if filter(i, v) {
			b = append(b, v)
		}
	}
	return b
}

func (a IntS) Val() []int {
	return []int(a)
}

func (a IntS) Contains(b ...int) bool {
	if len(a) < len(b) {
		return false
	}
OutLoop:
	for _, v1 := range a {
		for i, v2 := range b {
			if v1 == v2 {
				b = append(b[:i], b[i+1:]...)
				if len(b) == 0 {
					return true
				}
				continue OutLoop
			}
		}
	}
	return false
}

type IntValueInterface interface {
	Value() int
}

func Extract(arr interface{}, p interface{}) []int {

	a := make([]int, 0)
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return a
	}
	if v.IsNil() {
		return a
	}
	for i := 0; i < v.Len(); i++ {
		//e := v.Index(i)
	}
	return a
}

func ToIntSlice(a []int32) []int {
	b := make([]int, 0)
	for _, v := range a {
		b = append(b, int(v))
	}
	return b
}

type Foo struct {
	Name  string
	Value int
}

type intFilter struct{}

var IntFilter = intFilter{}

func (_ intFilter) GreaterThan(f int) func(idx int, value int) bool {
	return func(idx int, value int) bool {
		if value >= f {
			return true
		}
		return false
	}
}

func main() {
	//foos := []Foo{Foo{"a", 1}, Foo{"b", 2}, Foo{"c", 3}}

	/*
		seg := capnp.NewBuffer(nil)
		list := gamedata.NewCookies(seg, 1)
		cookie := gamedata.NewCookie(seg)
		list.Set(0, cookie)
		cookie.Name()

		l := list.Filter(func(e gamedata.Cookie) bool {
			return true
		}).ToStringSlice(gamedata.Cookie.Name)
	*/

	a := []int{7, 1, 2, 3, 5, 5}
	b := []int{5, 5, 2}
	c := []int{5, 5, 2}
	fmt.Println(IntS(a).Contains(b...))
	fmt.Println(Frequency(a, 5))
	Sort(a)
	fmt.Println("a", a)
	SortReverse(a)
	fmt.Println("a", a)
	fmt.Println(Concat(a, b, c))
	a1, a2 := Partition(a, func(i int, v int) bool {
		if v <= 3 {
			return true
		}
		return false
	})
	fmt.Println(a1)
	fmt.Println(a2)

	i := IntS(a).Filter(IntFilter.GreaterThan(1)).Filter(IntFilter.GreaterThan(3))
	fmt.Println(i)
}
