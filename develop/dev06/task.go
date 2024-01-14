package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	delimiter := flag.String("d", "\t", "delimiter(default=TAB")
	separatedOnly := flag.Bool("s", false, "только строки с разделителем")
	colNum := flag.Int("f", 0, "выбрать поля (колонки)")

	flag.Parse()

	if *colNum == 0 {
		println("you must specify field")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			println(err.Error())
			return
		}

		split := strings.Split(str, *delimiter)

		if *separatedOnly && len(split) == 1 {
			continue
		}

		if len(split) < *colNum {
			println()
		} else {
			println(split[*colNum-1])
		}
	}
}
