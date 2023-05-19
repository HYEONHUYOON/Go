package stack

import (
	"fmt"
)

// 슬라이스를 활용한 스택
type Stack []interface{}

func (stack *Stack) isEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) push(data interface{}) {
	*stack = append(*stack, data)
}

func (stack *Stack) pop() (interface{}, error) {
	if stack.isEmpty() {
		return nil, fmt.Errorf("Empty stack")
	} else {
		top := len(*stack) - 1
		data := (*stack)[top]
		*stack = (*stack)[:top]
		return data, nil
	}
}

func (stack *Stack) peek() (interface{}, error) {
	if stack.isEmpty() {
		return nil, fmt.Errorf("Empty stack")
	} else {
		top := len(*stack) - 1
		data := (*stack)[top]
		return data, nil
	}
}
