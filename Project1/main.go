package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Calcul struct {
	CurrentInput  string  // Число, которое пользователь ввоодит
	Display       string  // Число, которое горит на экране(ввод/результат)
	PreviousValue float64 // Первое число, которое было введено до начала операций
	Operation     string  // +, -, /, *
	IsNewInput    bool    // Надо ли стереть число(превышает ли его длина 8 цифр)
}

func main() {
	ap := app.New()
	window := ap.NewWindow("Calculator")

	label := widget.NewLabel("0")

	btn1 := widget.NewButton("1", func() {
		label.SetText("1")
	})

	btn1, err := strconv.Atoi("1")
	if err != nil {

	}

	window.SetContent(widget.NewVBox(label, btn1))

	window.Show()
	ap.Run()
}
