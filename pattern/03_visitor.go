package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Применимость:
 -  Если необходимо добавить новую функциональность для группы объектов, не изменяя их классы.
 -  Если операции, выполняемые над объектами, могут быть вынесены, чтобы избежать загрязнения основных классов.
 -  Если структура объектов сложна и возможности для изменения ограничены

Плюсы:
 -  Легкость добавления новых операций
 -  Отделение новой функциональности от основной структуры
Минусы:
 -  Введение класса-посетителя и его методов может усложнить структуру программы.
*/

type BuildingMaterial uint8

const (
	Stone = iota
	Wood
)

// Building interface

type Building interface {
	getBuildingType() BuildingMaterial
	accept(Visitor)
}

// Building implementations

type WoodBuilding struct{}

func (b *WoodBuilding) getBuildingType() BuildingMaterial {
	return Wood
}

func (b *WoodBuilding) accept(visitor Visitor) {
	visitor.visitForWoodBuilding(b)
}

type StoneBuilding struct{}

func (b *StoneBuilding) getBuildingType() BuildingMaterial {
	return Stone
}

func (b *StoneBuilding) accept(visitor Visitor) {
	visitor.visitForStoneBuilding(b)
}

type Visitor interface {
	visitForWoodBuilding(building *WoodBuilding)
	visitForStoneBuilding(building *StoneBuilding)
}

type XMLExporterVisitor struct {
}

func (X XMLExporterVisitor) visitForWoodBuilding(building *WoodBuilding) {
	fmt.Println("Exporting wood building...")
}

func (X XMLExporterVisitor) visitForStoneBuilding(building *StoneBuilding) {
	fmt.Println("Exporting stone building...")
}
