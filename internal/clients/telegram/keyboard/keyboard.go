package keyboaed

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CreateInlineKeyboardButtonList(values []string) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	for _, value := range values {
		button := tgbotapi.NewInlineKeyboardButtonData(value, value)
		row := []tgbotapi.InlineKeyboardButton{button}
		rows = append(rows, row)
	}

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func CreateNumericKeyboard(numbers []string) tgbotapi.ReplyKeyboardMarkup {
	var rows [][]tgbotapi.KeyboardButton

	// Splitting numbers into groups of 3 digits.
	for i := 0; i < len(numbers); i += 3 {
		end := i + 3
		if end > len(numbers) {
			end = len(numbers)
		}
		row := tgbotapi.NewKeyboardButtonRow()
		for _, number := range numbers[i:end] {
			button := tgbotapi.NewKeyboardButton(number)
			row = append(row, button)
		}
		rows = append(rows, row)
	}

	return tgbotapi.NewReplyKeyboard(rows...)
}

func CreateMenuKeyboard(buttons []string) tgbotapi.ReplyKeyboardMarkup {
	var keyboardButtons []tgbotapi.KeyboardButton
	for _, button := range buttons {
		keyboardButtons = append(keyboardButtons, tgbotapi.NewKeyboardButton(button))
	}

	// Creating rows of buttons.
	var keyboardRows [][]tgbotapi.KeyboardButton
	for i := 0; i < len(keyboardButtons); i += 3 {
		end := i + 3
		if end > len(keyboardButtons) {
			end = len(keyboardButtons)
		}
		keyboardRows = append(keyboardRows, tgbotapi.NewKeyboardButtonRow(keyboardButtons[i:end]...))
	}

	return tgbotapi.NewReplyKeyboard(keyboardRows...)
}
