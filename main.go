package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("LNick Generator")
	w.Resize(fyne.NewSize(500, 300))

	generateNickLabel := widget.NewLabel("Сгенерируй свой уникальный никнейм!")

	nick1Field := widget.NewEntry()
	nick1Field.PlaceHolder = "Слово 1"

	nick2Field := widget.NewEntry()
	nick2Field.PlaceHolder = "Слово 2"

	nickLenField := widget.NewEntry()
	nickLenField.PlaceHolder = "Длина никнейма"

	resultLabel := canvas.NewText("", color.White)
	resultLabel.TextStyle = fyne.TextStyle{Bold: true}
	resultLabel.TextSize = 20
	btnGen := widget.NewButton("Сгенерировать", func() {
		lengthNick, _ := strconv.Atoi(nickLenField.Text)
		resultLabel.Text = cutString(generateNickname(nick1Field.Text, nick2Field.Text), lengthNick)
	})
	w.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			container.New(layout.NewCenterLayout(), generateNickLabel),
			nick1Field,
			nick2Field,
			nickLenField,
			btnGen,
			container.New(layout.NewCenterLayout(), resultLabel),
		),
	)
	w.ShowAndRun()
}

func cutString(str string, length int) string {
	if length == 0 {
		return ""
	} else if length <= 0 {
		return ""
	} else if length > len(str) {
		return str
	}

	res := ""

	count := 0
	for _, char := range str {
		res += string(char)

		count++
		if count >= length {
			break
		}
	}

	return res
}

func generateNickname(nick1 string, nick2 string) string {
	allVowelsInNick := make([]string, 0)
	allConsonantsInNick := make([]string, 0)

	if len(nick1) < len(nick2) {
		nick1, nick2 = nick2, nick1
	}

	nick1_rune := []rune(nick1)
	nick2_rune := []rune(nick2)

	for letter_i := range nick1_rune {
		if isVowel(nick1_rune[letter_i]) {
			allVowelsInNick = append(allVowelsInNick, string(nick1_rune[letter_i]))
		} else {
			allConsonantsInNick = append(allConsonantsInNick, string(nick1_rune[letter_i]))
		}

		if letter_i < len(nick2_rune) {
			if isVowel(nick2_rune[letter_i]) {
				allVowelsInNick = append(allVowelsInNick, string(nick2_rune[letter_i]))
			} else {
				allConsonantsInNick = append(allConsonantsInNick, string(nick2_rune[letter_i]))
			}
		}
	}

	max_i := max(len(allVowelsInNick), len(allConsonantsInNick))

	fullNick := ""

	for i := 0; i < max_i; i++ {
		if i < len(allConsonantsInNick) {
			fullNick += allConsonantsInNick[i]
		}
		if i < len(allVowelsInNick) {
			fullNick += allVowelsInNick[i]
		}
	}

	return fullNick
}

func isVowel(letter rune) bool {
	return letter == 'а' || letter == 'у' || letter == 'о' || letter == 'ы' || letter == 'и' || letter == 'э' || letter == 'я' || letter == 'ю' || letter == 'ё' || letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' || letter == 'y'
}
