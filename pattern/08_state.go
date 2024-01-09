package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Применимость:
 - Когда объект может находиться в одном из нескольких состояний и поведение меняется в зависимости от него.
 - Когда код объекта содержит множество условных операторов, проверяющих текущее состояние.
 - Когда нужно скрыть детали реализации алгоритма

Плюсы:
 -  Избавление от больших условных конструкций
 -  Управление состоянием объекта
 -  Простота добавления новых состояний
Минусы:
 -  Может усложнить код, не принося пользы, если состояний немного и они редко меняются/дополняются
*/

type Context struct {
	value int
}

type State interface {
	changeValue(newValue int) error
	doSomething() error
}

type StateRealization1 struct {
	context *Context
}

func (s *StateRealization1) changeValue(newValue int) error {
	s.context.value = newValue
	return nil
}

func (s *StateRealization1) doSomething() error {
	return errors.New("cant do something:'( Invalid state")
}

type StateRealization2 struct {
	context *Context
}

func (s *StateRealization2) changeValue(newValue int) error {
	return errors.New("cant do something:'( Invalid state")
}

func (s *StateRealization2) doSomething() error {
	fmt.Println("Doing something...")
	return nil
}
