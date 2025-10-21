package main

import "fmt"

func main() {
	var n int
	mark := true
	fmt.Println("Введите цклое число меньшее 12307:")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Нужно ввести целое число!")
		return
	}
	if n >= 12307 {
		fmt.Println("Попросили же ввести меньше 12307(")
	} else {
		for n < 12307 {
			switch {
			case n < 0:
				n *= -1
			case n%7 == 0:
				n *= 39
			case n%9 == 0:
				n = n*13 + 1
				continue
			default:
				n = (n + 2) * 3
			}
			if n%9 == 0 && n%13 == 0 {
				fmt.Println("service error")
				mark = false
				break
			} else {
				n += 1
			}

		}
		if mark == true {
			fmt.Printf("Финальное значение: %d\n", n)
			fmt.Printf("Шестнадцатеричное представление: 0x%X\n", n)
			fmt.Printf("Двоичное представление: 0b%b\n", n)
		}
	}
}
