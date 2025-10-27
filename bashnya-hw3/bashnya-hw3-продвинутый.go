package main

import "fmt"

type Deque struct {
	arr []int
}

func (d *Deque) IsEmpty() bool {
	if len(d.arr) == 0 {
		return true
	}
	return false
}

func (d *Deque) PushFront(elem int) {
	d.arr = append([]int{elem}, d.arr...)
}

func (d *Deque) PushBack(elem int) {
	d.arr = append(d.arr, elem)
}

func (d *Deque) PopFront() *int {
	if d.IsEmpty() {
		return nil
	}
	elem := d.arr[0]
	d.arr = d.arr[1:]
	return &elem
}

func (d *Deque) PopBack() *int {
	if d.IsEmpty() {
		return nil
	}
	index := len(d.arr) - 1
	elem := d.arr[index]
	d.arr = d.arr[:index]
	return &elem
}

func (d *Deque) Back() *int {
	if d.IsEmpty() {
		return nil
	}
	return &d.arr[len(d.arr)-1]
}

func (d *Deque) Front() *int {
	if d.IsEmpty() {
		return nil
	}
	return &d.arr[0]
}
func (d *Deque) Size() int {
	length := len(d.arr)
	return length
}

func (d *Deque) Clear() {
	d.arr = nil
}

func main() {
	dq := Deque{}
	fmt.Println("Начальное состояние дека")
	fmt.Println("Пустой ли дек?", dq.IsEmpty())
	fmt.Println("Размер дека:", dq.Size())
	fmt.Println("Содержимое дека:", dq.arr)
	fmt.Println("Проверка Front и Back на пустом деке")
	if front := dq.Front(); front != nil {
		fmt.Println("Первый элемент:", *front)
	} else {
		fmt.Println("Первый элемент: <nil>")
	}
	if back := dq.Back(); back != nil {
		fmt.Println("Последний элемент:", *back)
	} else {
		fmt.Println("Последний элемент: <nil>")
	}
	fmt.Println("Добавляем 10 в конец")
	dq.PushBack(10)
	if front := dq.Front(); front != nil {
		fmt.Println("Первый элемент:", *front)
	} else {
		fmt.Println("Первый элемент: <nil>")
	}
	if back := dq.Back(); back != nil {
		fmt.Println("Последний элемент:", *back)
	} else {
		fmt.Println("Последний элемент: <nil>")
	}
	fmt.Println("Добавляем элементы в дек")
	fmt.Println("Добавляем 11 в конец")
	dq.PushBack(11)
	fmt.Println("Добавляем 12 в начало")
	dq.PushFront(12)
	fmt.Println("Добавляем 13 в конец")
	dq.PushBack(13)
	fmt.Println("Добавляем 14 в начало")
	dq.PushFront(14)
	fmt.Println("Добавляем 15 в конец")
	dq.PushBack(15)
	fmt.Println("Состояние дека после добавления элементов")
	if front := dq.Front(); front != nil {
		fmt.Println("Первый элемент:", *front)
	} else {
		fmt.Println("Первый элемент: <nil>")
	}
	if back := dq.Back(); back != nil {
		fmt.Println("Последний элемент:", *back)
	} else {
		fmt.Println("Последний элемент: <nil>")
	}
	fmt.Println("Содержимое дека:", dq.arr)
	fmt.Println("Пустой ли дек?", dq.IsEmpty())
	fmt.Println("Извлекаем элемент из начала")
	if popFront := dq.PopFront(); popFront != nil {
		fmt.Println("PopFront (извлеченный элемент):", *popFront)
	} else {
		fmt.Println("PopFront (извлеченный элемент): <nil>")
	}
	fmt.Println("Размер дека:", dq.Size())
	fmt.Println("Содержимое дека :", dq.arr)
	fmt.Println("Извлекаем элемент из конца")
	if popBack := dq.PopBack(); popBack != nil {
		fmt.Println("PopBack (извлеченный элемент):", *popBack)
	} else {
		fmt.Println("PopBack (извлеченный элемент): <nil>")
	}
	fmt.Println("Размер дека:", dq.Size())
	fmt.Println("Содержимое дека:", dq.arr)
	fmt.Println("Добавляем 16 в начало")
	dq.PushFront(16)
	if front := dq.Front(); front != nil {
		fmt.Println("Первый элемент:", *front)
	} else {
		fmt.Println("Первый элемент: <nil>")
	}
	if back := dq.Back(); back != nil {
		fmt.Println("Последний элемент:", *back)
	} else {
		fmt.Println("Последний элемент: <nil>")
	}
	fmt.Println("Размер дека:", dq.Size())
	fmt.Println("Еще раз извлекаем из начала")
	if popFront := dq.PopFront(); popFront != nil {
		fmt.Println("PopFront (извлеченный элемент):", *popFront)
	} else {
		fmt.Println("PopFront (извлеченный элемент): <nil>")
	}
	fmt.Println("Содержимое дека:", dq.arr)
	fmt.Println("Еще раз извлекаем из конца")
	if popBack := dq.PopBack(); popBack != nil {
		fmt.Println("PopBack (извлеченный элемент):", *popBack)
	} else {
		fmt.Println("PopBack (извлеченный элемент): <nil>")
	}
	fmt.Println("Содержимое дека:", dq.arr)
	dq.Clear()
	fmt.Println("Размер дека после очистки:", dq.Size())
	fmt.Println("Содержимое дека после очискти:", dq.arr)
}
