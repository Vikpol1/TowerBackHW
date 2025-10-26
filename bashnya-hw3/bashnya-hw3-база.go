package main

import "fmt"

type Stack struct {
	arr []int
}

func (s *Stack) Push(elem int) {
	s.arr = append(s.arr, elem)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	index := len(s.arr) - 1
	elem := s.arr[index]
	s.arr = s.arr[:index]
	return elem
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return 0
	}
	return s.arr[len(s.arr)-1]
}

func (s *Stack) IsEmpty() bool {
	if len(s.arr) == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	length := len(s.arr)
	return length
}

func (s *Stack) Clear() {
	s.arr = nil
}

func main() {
	st := Stack{}
	fmt.Println("Начальное состояние стека")
	fmt.Println("Пустой ли стек?", st.IsEmpty())
	fmt.Println("Размер стека:", st.Size())
	fmt.Println("Содержимое стека:", st.arr)
	fmt.Println("Верхний элмент стека (без удаления):", st.Peek())
	fmt.Println("Добавляем элемент 10")
	st.Push(10)
	fmt.Println("Верхний элемент стека (без удаления):", st.Peek())
	fmt.Println("Добавляем элементы 11, 12, 13")
	st.Push(11)
	st.Push(12)
	st.Push(13)
	fmt.Println("Верхний элемент стека (без удаления):", st.Peek())
	fmt.Println("Содержимое стека:", st.arr)
	fmt.Println("Пустой ли стек?", st.IsEmpty())
	fmt.Println("Извлекаем элемент из стека")
	fmt.Println("Результат Pop():", st.Pop())
	fmt.Println("Размер стека после Pop():", st.Size())
	fmt.Println("Содержимое стека после Pop():", st.arr)
	fmt.Println("Добавляем элемент 14")
	st.Push(14)
	fmt.Println("Верхний элемент стека (без удаления):", st.Peek())
	fmt.Println("Размер стека после Push(14):", st.Size())
	fmt.Println("Еще раз извлекаем элемент")
	fmt.Println("Результат Pop():", st.Pop())
	fmt.Println("Содержимое стека:", st.arr)
	st.Clear()
	fmt.Println("Размер стека после очистки:", st.Size())
	fmt.Println("Содержимое стека после очискти:", st.arr)
}
