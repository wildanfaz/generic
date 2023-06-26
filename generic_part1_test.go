package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGeneric(t *testing.T) {
	v := NonGeneric("foo")

	v2 := GenericTypeParameter("bar")

	number := GenericTypeParameter[int](123)

	// non generic
	fmt.Println(Hello(v.(string)))

	// generic
	fmt.Println("\nType Parameter")
	fmt.Println(Hello(v2))

	fmt.Println(number)

	fmt.Println("\nMultiple Parameter")
	GenericMultipleParameter[string, int]("ID :", 123)

	GenericMultipleParameter[int, string](200, "OK")

	fmt.Println("\nMultiple Type Parameter")
	GenericMultipleTypeParameter[string]("Hello")

	GenericMultipleTypeParameter[int](1)

	fmt.Println("\nType Sets")
	GenericTypeSets(123.321)

	fmt.Println("\nType Approximation")
	GenericTypeApproximation[Id](Id(123))

	fmt.Println("\nComparable")
	// var test interface{} = 123
	GenericComparable("Hello", "World")

	fmt.Println("\nType Inheritance")
	impl := &ImplSayHello{"Foo"}

	GenericTypeInheritance(impl)
}

// Non Generic
func NonGeneric(input any) any {
	return input
}

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

// Generic
// Type Parameter
func GenericTypeParameter[T any](input T) T {
	return input
}

// Multiple Parameter
func GenericMultipleParameter[T1, T2 any](input T1, input2 T2) {
	fmt.Println(input, input2)
}

// Multiple Type Parameter
func GenericMultipleTypeParameter[T string | int](input T) {
	fmt.Println(input)
}

type IntOrFloat interface {
	int | float64
}

// Type Sets
func GenericTypeSets[T IntOrFloat](input T) {
	fmt.Println(reflect.TypeOf(input), ":", input)
}

type Number interface {
	~int | int8 | float64
}

type Id int

// Type Approximation
func GenericTypeApproximation[T Number](input T) {
	fmt.Println(input)
}

// Comparable
func GenericComparable[T comparable](input T, input2 T) {
	if input == input2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

type SayHello interface {
	HelloName() string
}

type ImplSayHello struct {
	Name string
}

func (v *ImplSayHello) HelloName() string {
	v.Name = "Hello " + v.Name

	return v.Name
}

// Type Inheritance
func GenericTypeInheritance[T SayHello](input T) {
	// hello := T.HelloName(input)

	hello := input.HelloName()

	fmt.Println(hello)
}
