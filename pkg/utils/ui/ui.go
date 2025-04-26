package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func StartUI() {
	// Cria a aplicação
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculadora de Soma")
	myWindow.Resize(fyne.NewSize(400, 300))

	// Elementos da interface
	label1 := widget.NewLabel("Digite o primeiro valor:")
	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("Número 1")

	label2 := widget.NewLabel("Digite o segundo valor:")
	entry2 := widget.NewEntry()
	entry2.SetPlaceHolder("Número 2")

	btnCalcular := widget.NewButton("Calcular Soma", func() {
		// Converter valores para float
		val1, err1 := strconv.ParseFloat(entry1.Text, 64)
		val2, err2 := strconv.ParseFloat(entry2.Text, 64)

		if err1 != nil || err2 != nil {
			dialog.ShowError(fmt.Errorf("Valores inválidos!\nDigite números válidos."), myWindow)
			return
		}

		// Calcular e mostrar resultado
		resultado := val1 + val2
		dialog.ShowInformation("Resultado",
			fmt.Sprintf("A soma de %.2f e %.2f é: %.2f", val1, val2, resultado),
			myWindow)
	})

	// Layout da janela
	content := container.NewVBox(
		label1,
		entry1,
		label2,
		entry2,
		btnCalcular,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
