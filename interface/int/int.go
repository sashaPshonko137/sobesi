package main

import (
	"fmt"
)

func printNumber(ptrToNumber interface{}) {
	// Проверяем, что переданное значение является указателем на int
	if ptrToNumber != nil {
		if numPtr, ok := ptrToNumber.(*int); ok {
			if numPtr != nil {
				fmt.Println(*numPtr)
			} else {
				fmt.Println("nil")
			}
		} else {
			fmt.Println("Not a pointer to int")
		}
	} else {
		fmt.Println("nil")
	}
}