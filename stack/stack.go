package stack

import (
	"errors"
	"fmt"
	"reflect"
)

//List 定义栈接口
/*type List interface {
	Push(v interface{}) error
	Pop() (interface{}, error)
	Top() (interface{}, error)
}*/

//Stack 栈结构体
type Stack struct {
	types    string
	elements []interface{}
	top      int
}

//NewStack 初始化函数
func NewStack(t string, v ...interface{}) (*Stack, error) {
	s := &Stack{t, []interface{}{}, -1}
	if len(v) == 0 {
		return s, nil
	}
	for i, j := range v {
		err := s.Push(j)
		if err != nil {
			text := fmt.Sprintf("第%d个元素出现问题: %s", i, err)
			return nil, errors.New(text)
		}
	}
	return s, nil
}

//Push 压栈操作
func (s *Stack) Push(v interface{}) error {
	t := reflect.ValueOf(v)
	if t.Type().String() != s.types {
		return errors.New("元素类型与栈类型不匹配")
	}
	s.top = s.top + 1
	if s.top >= len(s.elements) {
		s.elements = append(s.elements, v)
	} else {
		s.elements[s.top] = v
	}
	return nil
}

//Pop 弹出操作
func (s *Stack) Pop() (interface{}, error) {
	if s.top < 0 {
		err := errors.New("空栈")
		return nil, err
	}
	v := s.elements[s.top]
	s.top = s.top - 1
	return v, nil
}

//Top 返回栈顶
func (s *Stack) Top() (interface{}, error) {
	if s.top < 0 {
		err := errors.New("空栈")
		return nil, err
	}
	v := s.elements[s.top]
	return v, nil
}

//IsEmpty 判断栈空
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}
