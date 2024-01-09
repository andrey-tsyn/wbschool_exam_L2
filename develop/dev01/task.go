package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	currTime, err := getTime()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("%s\n", err))
		os.Exit(-1)
	}

	h, m, s := currTime.Clock()
	fmt.Printf("Current time: %d:%d:%d", h, m, s)
}

func getTime() (*time.Time, error) {
	currTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		return nil, err
	}

	return &currTime, nil
}
