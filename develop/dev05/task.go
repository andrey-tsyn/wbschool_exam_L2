package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "around" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Params struct {
	after        int
	before       int
	around       int
	printCount   bool
	ignoreCase   bool
	invert       bool
	fixed        bool
	printLineNum bool
}

func main() {
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	printCount := flag.Bool("c", false, "печатать только кол-во совпадений")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "инвертировать вывод")
	fixed := flag.Bool("F", false, "точное совпадение со строкой")
	printLineNum := flag.Bool("n", false, "печатать номер строки")

	flag.Parse()

	params := Params{
		after:        *after,
		before:       *before,
		around:       *context,
		printCount:   *printCount,
		ignoreCase:   *ignoreCase,
		invert:       *invert,
		fixed:        *fixed,
		printLineNum: *printLineNum,
	}

	args := flag.Args()

	if len(args) < 2 {
		fmt.Printf("путь к файлу должен быть обязательно!\n")
		return
	}

	pattern := args[0]
	lines, err := getTextFileLines(args[1])
	if err != nil {
		fmt.Printf("Error occuried: %s", err.Error())
		return
	}

	result := grep(lines, pattern, params)

	if params.printCount {
		fmt.Printf("Matched lines cound: %d\n", len(result))
		return
	}

	if params.invert {
		slices.Reverse(result)
	}

	println(strings.Join(result, "\n"))
}

// TODO Rename
func grep(lines []string, pattern string, params Params) []string {
	var result []string

	lastEndAddedIndex := 0

	if params.around > 0 {
		params.before = params.around
		params.after = params.around
	} else {
		if params.before < 0 {
			params.before = 0
		}
		if params.after < 0 {
			params.after = 0
		}
	}

	if params.ignoreCase {
		pattern = strings.ToLower(pattern)
	}

	for i, line := range lines {
		if params.ignoreCase {
			line = strings.ToLower(line)
		}

		if params.fixed && strings.Contains(line, pattern) {
			linesToAdd := getFormatedStrings(i, &lastEndAddedIndex, lines, params)
			fmt.Printf("%+v\n", linesToAdd)
			result = append(result, linesToAdd...)
		} else if !params.fixed {
			match, _ := regexp.Match(pattern, []byte(line))
			if match {
				linesToAdd := getFormatedStrings(i, &lastEndAddedIndex, lines, params)
				fmt.Printf("%+v\n", linesToAdd)
				result = append(result, linesToAdd...)
			}
		}
	}

	return result
}

func getFormatedStrings(currIndex int, lastEndAddedIndex *int, lines []string, params Params) []string {
	firstLineIndex, lastLineIndex := currIndex, currIndex

	firstLineIndex -= params.before
	if firstLineIndex < 0 {
		firstLineIndex = 0
	}
	if firstLineIndex < *lastEndAddedIndex {
		firstLineIndex = *lastEndAddedIndex
	}
	if firstLineIndex >= len(lines) {
		return nil
	}

	lastLineIndex += params.after + 1
	if lastLineIndex > len(lines) {
		lastLineIndex = len(lines)
	}

	*lastEndAddedIndex = lastLineIndex

	fmt.Printf("s: %d, e: %d", firstLineIndex, lastLineIndex)

	linesToAdd := lines[firstLineIndex:lastLineIndex]

	if params.printLineNum {
		linesToAdd = addNumPrefixToStringSlice(firstLineIndex, linesToAdd)
	}

	return linesToAdd
}

func addNumPrefixToStringSlice(firstLineNum int, lines []string) []string {
	for i, lineNum := 0, firstLineNum+1; i < len(lines); i, lineNum = i+1, lineNum+1 {
		lines[i] = fmt.Sprintf("#%d %s", lineNum, lines[i])
	}

	return lines
}

func getTextFileLines(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), "\n"), nil
}
