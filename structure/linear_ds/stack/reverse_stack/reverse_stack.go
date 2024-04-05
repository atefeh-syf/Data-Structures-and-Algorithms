package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	elementValue int
}

type Stack struct {
	elements     []*Element
	elementCount int
}

type ListNode []int

func (element *Element) String() string {
	return strconv.Itoa(element.elementValue)
}

func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount = stack.elementCount + 1
}

func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}

	var length int = len(stack.elements)
	var element *Element = stack.elements[length-1]

	if length > 1 {
		stack.elements = stack.elements[:length-1]
	} else {
		stack.elements = stack.elements[0:]
	}
	stack.elementCount = len(stack.elements)
	return element
}

func reverseList(head *ListNode) *ListNode {
	var stack *Stack = &Stack{}
	stack.New()
	for node := range *head {
		stack.Push( &Element{node})
	}

	var resultHead = ListNode{}
	for i := 0; i < len(stack.elements); i++ {
		element := stack.Pop()
		resultHead = append(resultHead, element.elementValue)
	}
	return &resultHead
}

func main() {
	var head = &ListNode{1,2,3,4,5}
	var resultHead = reverseList(head)
	fmt.Println(resultHead)

}
