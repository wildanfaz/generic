package main

import (
	"fmt"
	"testing"
)

func TestGeneric2(t *testing.T) {
	// generic
	fmt.Println("Type Inference")
	GenericTypeInference("Hello")

	fmt.Println("\nGeneric Type")
	GenericType(Names[string]{"Foo", "Bar"})

	fmt.Println("\nGeneric Struct")
	obj := Object[int, string]{Id: 123, Name: "Book"}

	GenericStruct(obj)

	fmt.Println("\nGeneric Method")
	obj.GetId()

	obj.ChangeName("Game")

	fmt.Println("\nGeneric Interface")
	impl := &ImplGet[string]{"Foo"}

	GenericInterface[string](impl)

	fmt.Println("\nGeneric In Line Type Constraint")
	GenericInLineTypeConstraint("Foo")

	fmt.Println("\nGeneric In Type Parameter")
	GenericInTypeParameter([]string{"Foo", "Bar"})
}

// Type Inference
func GenericTypeInference[T any](input T) {
	fmt.Println(input)
}

type Names[T any] []T

// Generic Type
func GenericType[T any](input Names[T]) {
	for _, v := range input {
		fmt.Println("Hello", v)
	}
}

type Object[T1, T2 any] struct {
	Id   T1
	Name T2
}

// Generic Struct
func GenericStruct[T1, T2 any](input Object[T1, T2]) {
	fmt.Println(input)
}

// Generic Method
func (o *Object[_, _]) GetId() {
	fmt.Println("Id :", o.Id)
}

// Generic Method
func (o *Object[_, T2]) ChangeName(input T2) {
	o.Name = input

	fmt.Println("Name :", o.Name)
}

type Get[T any] interface {
	Get() T
}

type ImplGet[T any] struct {
	Name T
}

func (g *ImplGet[T]) Get() T {
	return g.Name
}

// Generic Interface
func GenericInterface[T any](input Get[T]) {
	result := input.Get()

	fmt.Println(result)
}

// Generic In Line Type Constraint
func GenericInLineTypeConstraint[T interface{ int | string }](input T) {
	fmt.Println(input)
}

// Generic In Type Parameter
func GenericInTypeParameter[T []E, E any](input T) {
	fmt.Println(input)
}

// Golang Experimental
// https://pkg.go.dev/golang.org/x/exp
