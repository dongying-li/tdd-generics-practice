package generics

import (
    "reflect"
    "testing"
)

func TestStack(t *testing.T) {
    t.Run("Push and GetElements with int", func(t *testing.T) {
        intStack := Stack[int]{}
        intStack.Push(1)
        intStack.Push(2)

        expect := []int{1, 2}
        actual := intStack.GetElements()

        if !reflect.DeepEqual(expect, actual) {
            t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
        }
    })

    t.Run("Push and GetElements with string", func(t *testing.T) {
        stringStack := Stack[string]{}
        stringStack.Push("hello")
        stringStack.Push("world")

        expect := []string{"hello", "world"}
        actual := stringStack.GetElements()

        if !reflect.DeepEqual(expect, actual) {
            t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
        }
    })

    t.Run("Push and GetElements with struct", func(t *testing.T) {
        type Person struct {
            Name string
            Age  int
        }

        personStack := Stack[Person]{}
        personStack.Push(Person{Name: "Alice", Age: 30})
        personStack.Push(Person{Name: "Bob", Age: 25})

        expect := []Person{
            {Name: "Alice", Age: 30},
            {Name: "Bob", Age: 25},
        }
        actual := personStack.GetElements()

        if !reflect.DeepEqual(expect, actual) {
            t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
        }
    })

    t.Run("Pop with string", func(t *testing.T) {
        stringStack := Stack[string]{}
        stringStack.Push("hello")
        stringStack.Push("world")

        element, err := stringStack.Pop()
        if err != nil {
            t.Errorf("Unexpected error: %s", err)
        }
        if element != "world" {
            t.Errorf("\nexpect: world\nactual: %s\n", element)
        }

        // Check remaining elements
        expect := []string{"hello"}
        actual := stringStack.GetElements()
        if !reflect.DeepEqual(expect, actual) {
            t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
        }
    })

    t.Run("IsEmpty with struct", func(t *testing.T) {
        type Person struct {
            Name string
            Age  int
        }

        personStack := Stack[Person]{}

        if !personStack.IsEmpty() {
            t.Errorf("expect: true\nactual: false")
        }

        personStack.Push(Person{Name: "Alice", Age: 30})
        if personStack.IsEmpty() {
            t.Errorf("expect: false\nactual: true")
        }
    })
}
