package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Применимость:
 -  Если нужна возможность отменить действие
 -  Если нужно откладывать выполнение действия

Плюсы:
 -  Описаны в применимости
 -  Вызов одной и той же функции разными способами(кнопка, шорткат...)
 -  Отделение отправителя от получателя
Минусы:
 -  Усложняет код из-за увеличения кол-ва структур
*/

type Button struct {
	command Command
}

func NewButton(command Command) *Button {
	return &Button{command: command}
}

func (b *Button) Press() {
	b.command.Exec()
}

type Command interface {
	Exec()
}

type ExitCommand struct{}

func (c *ExitCommand) Exec() {
	fmt.Println("Exiting...")
}

type DoSomethingElseCommand struct{}

func (c *DoSomethingElseCommand) Exec() {
	fmt.Println("Do something else...")
}
