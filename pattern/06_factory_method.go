package pattern

import "errors"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Применимость:
 - Когда необходимо избежать привязки клиента к конкретным классам, которые он создает.

Плюсы:
 -  Простота добавления новых типов не изменяя код, а только дополняя
 -  Клиент работает с абстрактными классами, что упрощает его расширение и поддержку.
Минусы:
 -  Трудности в поддержке сложных иерархий классов
*/

type Transport interface {
	DoSomeTransportingThings()
}

type Airplane struct{}

func NewAirplane() *Airplane {
	return &Airplane{}
}

func (a Airplane) DoSomeTransportingThings() {
	// Logic...
}

type Ship struct{}

func NewShip() *Ship {
	return &Ship{}
}

func (s Ship) DoSomeTransportingThings() {
	// Logic...
}

// Factory
func getTransport(transportType string) (Transport, error) {
	if transportType == "ship" {
		return NewShip(), nil
	}
	if transportType == "airplane" {
		return NewAirplane(), nil
	}

	return nil, errors.New("wrong transport type provided")
}
