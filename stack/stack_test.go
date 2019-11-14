package stack

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/shigzz/gokit/tree"
)

func TestStack(t *testing.T) {
	st, err := NewStack("int", 1, 2)
	if err != nil {
		fmt.Println(err)
	}
	err = st.Push(499)
	if err != nil {
		fmt.Println(err)
	}
	v, err := st.Pop()
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(st)
}

func TestStackWithTree(t *testing.T) {
	head := tree.CreateTree([]int{1, 2, 3, 4, 5})
	st, err := NewStack(reflect.TypeOf(head).String(), head, 1)
	if err != nil {
		t.Log(err)
		return
	}
	err = st.Push(head)
	if err != nil {
		t.Log(err)
	}
	t.Log(st)
}
