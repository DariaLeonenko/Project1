package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Calcul struct {
	CurrentInput  string  // Число, которое пользователь вводит
	PreviousValue float64 // Первое число, которое было введено до начала операций
	Operation     string  // +, -, /, *
	IsNewInput    bool    // Надо ли стереть число(превышает ли его длина 8 цифр)
}

// Логика
func main() {
	ap := app.New()
	window := ap.NewWindow("Calculator")

	c := Calcul{
		IsNewInput: true,
	}

	displayLabel := widget.NewLabel("0") // Дисплей
	// Тап по цифре
	pressDigit := func(digit string) {

		if c.IsNewInput {
			c.CurrentInput = ""
			c.IsNewInput = false
		}

		if len(c.CurrentInput) < 9 {
			c.CurrentInput += digit
			displayLabel.SetText(c.CurrentInput)
		}
	}
	// Тап по операции
	pressOper := func(op string) {
		val, err := strconv.ParseFloat(c.CurrentInput, 64)
		if err != nil {
			return
		}
		c.PreviousValue = val
		c.Operation = op
		c.IsNewInput = true
		displayLabel.SetText(op)
	}
	// Тап по равно (=)
	equalBtn := widget.NewButton("=", func() {
		// Проверки на некорректность
		if c.Operation == "" {
			return
		}
		currentVal, err := strconv.ParseFloat(c.CurrentInput, 64)
		if err != nil {
			return
		}

		// Выполнение операции
		var result float64
		switch c.Operation {
		case "+":
			result = c.PreviousValue + currentVal
		case "-":
			result = c.PreviousValue - currentVal
		case "*":
			result = c.PreviousValue * currentVal
		case ":":
			if currentVal == 0 {
				displayLabel.SetText("Error")
				c.IsNewInput = true
				return
			}
			result = c.PreviousValue / currentVal
		}
		// Числовой результат обратно в строку и затем в дисплей
		resultEnd := strconv.FormatFloat(result, 'f', -1, 64)
		displayLabel.SetText(resultEnd)
		c.CurrentInput = resultEnd
		c.Operation = ""
		c.IsNewInput = true
	})
	// Стирание одного значения
	pressBack := func() {
		if c.IsNewInput {
			return
		}
		length := len(c.CurrentInput)
		if length > 1 {
			c.CurrentInput = c.CurrentInput[:length-1]
			displayLabel.SetText(c.CurrentInput)
		} else {
			c.CurrentInput = ""
			displayLabel.SetText("0")
		}
	}
	// Кнопки
	grid := container.NewGridWithColumns(4,
		widget.NewButton("1", func() { pressDigit("1") }),
		widget.NewButton("2", func() { pressDigit("2") }),
		widget.NewButton("3", func() { pressDigit("3") }),
		widget.NewButton("+", func() { pressOper("+") }),

		widget.NewButton("4", func() { pressDigit("4") }),
		widget.NewButton("5", func() { pressDigit("5") }),
		widget.NewButton("6", func() { pressDigit("6") }),
		widget.NewButton("-", func() { pressOper("-") }),

		widget.NewButton("7", func() { pressDigit("7") }),
		widget.NewButton("8", func() { pressDigit("8") }),
		widget.NewButton("9", func() { pressDigit("9") }),
		widget.NewButton("*", func() { pressOper("*") }),

		widget.NewButton("Back", func() { pressBack() }),
		equalBtn,
		widget.NewButton("0", func() { pressDigit("0") }),
		widget.NewButton(":", func() { pressOper(":") }),
	)
	// Склейка Дисплея и Кнопок
	mainLayout := container.NewVBox(
		displayLabel,
		grid,
	)
	window.SetContent(mainLayout)
	window.ShowAndRun()
}
