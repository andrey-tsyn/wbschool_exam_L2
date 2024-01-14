package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"task3/sorter"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки DONE
-n — сортировать по числовому значению DONE
-r — сортировать в обратном порядке DONE
-u — не выводить повторяющиеся строки DONE

Дополнительное

Поддержать ключи

TODO: Сделать, если сделаю остальные и будет время
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	needToReverse := flag.Bool("r", false, "Reverses the result of comparisons.")
	compareColumn := flag.Int("k", 0, "Start a key at POS1 (origin 1), end it at POS2 (default end of line)")
	uniqueOutput := flag.Bool("u", false, "Unique processing to suppress all but one in each set of lines having equal keys.")
	numericSort := flag.Bool("n", false, "Compares according to string numerical value.")

	flag.Parse()

	args := flag.Args()

	sorter, err := sorting.NewStringSorter(
		sorting.WithReverse(*needToReverse),
		sorting.WithUniqueOutput(*uniqueOutput),
		sorting.WithComparingColumn(*compareColumn),
		sorting.WithNumericSort(*numericSort),
	)

	if err != nil {
		fmt.Printf("Error occuried: %s", err.Error())
		return
	}

	arr, err := getTextFileLines(args[0])
	if err != nil {
		fmt.Printf("Error occuried: %s", err.Error())
		return
	}

	arr = sorter.Sort(arr)
	fmt.Printf("%s\n", strings.Join(arr, "\n"))
}

func getTextFileLines(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), "\n"), nil
}
