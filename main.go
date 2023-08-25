package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack[T any] struct {
	array []T
	len   int
}

func (s *Stack[T]) Len() int {
	return s.len
}

func NewStack[T any](size int) Stack[T] {
	return Stack[T]{make([]T, size), 0}
}

func (s *Stack[T]) Push(val T) {
	if s.len > cap(s.array) {
		return
	}
	s.array[s.len] = val
	s.len++
}

func (s *Stack[T]) Pop() (ret T) {
	ret = s.array[0]
	if s.len < 1 {
		return
	}
	ret = s.array[s.len-1]
	s.len--
	return
}

func Opposite(r rune) (ret rune) {
	switch r {
	case '{':
		ret = '}'
	case '}':
		ret = '{'
	case '(':
		ret = ')'
	case ')':
		ret = '('
	case '[':
		ret = ']'
	case ']':
		ret = '['
	default:
		ret = '-'
	}
	return
}

func CheckString(str string) bool {
	tmp := NewStack[rune](len(str))
	for _, val := range str {
		switch val {
		case '{':
			fallthrough
		case '(':
			fallthrough
		case '[':
			tmp.Push(val)
		case '}':
			fallthrough
		case ')':
			fallthrough
		case ']':
			if Opposite(tmp.Pop()) != val {
				return false
			}
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: string not readed")
		}
		if CheckString(strings.Trim(str, "\r\n")) {
			fmt.Println("Скобки сбалансированы")
		} else {
			fmt.Println("Скобки не сбалансированы")
		}
	}
}
