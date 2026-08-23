package main

import "fmt"

type Calculator struct {
	CurrentInput  string  // Число, которое пользователь ввоодит
	Display       string  // Число, которое горит на экране(ввод/результат)
	PreviousValue float64 // Первое число, которое было введено до начала операций
	Operation     string  // +, -, /, *
	IsNewInput    bool    // Надо ли стереть число(превышает ли его длина 8 цифр)
}

func (v *Calculator) getting() {
	fmt.Scan(&v.CurrentInput)
	v.Display = v.CurrentInput
}
