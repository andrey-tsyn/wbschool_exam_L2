package main

import (
	"flag"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	var timeout time.Duration

	flag.DurationVar(&timeout, "timeout", 10*time.Second, "timeout")

	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		println("host and port must be provided")
		return
	}

	url := args[0]
	port := args[1]

	// Сигнал для graceful shutdown
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, os.Interrupt)

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(url, port), timeout)
	if err != nil {
		println(err.Error())
		return
	}
	defer conn.Close()

	go func() {
		_, err := io.Copy(os.Stdout, conn)
		// Ctrl+D в shell отправляет EOF, io.Copy возвращает ошибку io.EOF, если он обнаружен
		if err != nil {
			interrupt <- os.Interrupt
		}
	}()
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			interrupt <- os.Interrupt
		}
	}()

	<-interrupt
}
