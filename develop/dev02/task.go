package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	IncorrectString = errors.New("provided string is incorrect")
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackString(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	runes := []rune(str)

	if unicode.IsDigit(runes[0]) {
		return "", IncorrectString
	}

	builder := strings.Builder{}
	runesLen := len(runes)
	backSlashRepeated := false

	for i := 0; i < runesLen; i++ {
		symbol := runes[i]

		// Проверка на обратный слеш, если он есть, то переключаемся на следующий символ
		if symbol == '\\' && !backSlashRepeated {
			i++
			if i == runesLen {
				return "", IncorrectString
			}
			symbol = runes[i]
			// Если следующий символ также обратный слеш, то мы ставим флаг, чтобы при следующей операции работать с
			// ним, как с обычным
			if symbol == '\\' {
				backSlashRepeated = true
			}
		} else {
			backSlashRepeated = false
		}

		// Получаем длину последовательности цифр(если есть)
		startDigitsIndex, endDigitsIndex := i+1, 0
		for i+1 != runesLen && unicode.IsDigit(runes[i+1]) {
			i++
			endDigitsIndex = i
		}

		// Если цифры были найдены, то парсим их в int и добавляем строку с повторяющимся символом заданное
		// кол-во раз
		if endDigitsIndex != 0 {
			count, err := strconv.Atoi(string(runes[startDigitsIndex : endDigitsIndex+1]))
			if err != nil {
				// Паника, так как этого не должно произойти в программе, это не ошибка пользователя
				panic("Function must provide correct rune array with only digits!")
			}
			builder.WriteString(strings.Repeat(string(symbol), count))
			continue
		}

		builder.WriteRune(symbol)
	}

	return builder.String(), nil
}
