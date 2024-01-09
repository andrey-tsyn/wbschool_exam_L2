package pattern

import "errors"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Применимость:
 -  Если есть необходимость использовать опциональные параметры для создания класса
 -  Если необходимо создавать сложные объекты с множеством параметров

Плюсы:
 -  Позволяет создавать класс поэтапно
 -  Переиспользование кода для создания различных экземпляров
Минусы:
 -  Усложняет код
 -  Может усложнить внедрение зависимостей
*/

type CarBuilder interface {
	SetEngine(engine *Engine2)
	SetWheelType(wheelType string)
	SetSeatsCount(seatsCount uint8)
	Build() (*Car, error)
}

type ManualCarBuilder struct {
	engine     *Engine2
	wheelType  *string
	seatsCount uint8
}

func NewManualCarBuilder() *ManualCarBuilder {
	return &ManualCarBuilder{}
}

func (c *ManualCarBuilder) SetEngine(engine *Engine2) {
	c.engine = engine
}

func (c *ManualCarBuilder) SetWheelType(wheelType string) {
	c.wheelType = &wheelType
}

func (c *ManualCarBuilder) SetSeatsCount(seatsCount uint8) {
	c.seatsCount = seatsCount
}

func (c *ManualCarBuilder) Build() (*Car, error) {
	if c.engine == nil {
		return nil, errors.New("engine is not set")
	}

	if c.wheelType == nil {
		return nil, errors.New("wheel type is not set")
	}

	if c.seatsCount == 0 {
		return nil, errors.New("seats count is not set")
	}

	return &Car{
		engine:     c.engine,
		wheelType:  *c.wheelType,
		seatsCount: c.seatsCount,
	}, nil
}

type Car struct {
	engine     *Engine2
	wheelType  string
	seatsCount uint8
}

type Engine2 interface {
	GetEngineInfo() string
}

type EngineRealization struct {
	isStarted bool
}

func (e *EngineRealization) GetEngineInfo() string {
	return "I'm basic engine realization."
}
