package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Применимость:
 - Когда есть несколько вариантов алгоритма
 - Когда необходимо менять алгоритмы в рантайме
 - Когда нужно скрыть детали реализации алгоритма

Плюсы:
 -  Динамическая замена алгоритмов
 -  Сокрытие реализации от других классов/структур
Минусы:
 -  Усложнение кода из-за дополнительных классов
 -  Клиенту необходимо знать о разнице между стратегиями для выбора подходящей
*/

type Product struct {
	Name string
}

// Strategy interface

type DeliveryStrategy interface {
	deliver(Product)
}

// Strategy implementations

type CarDeliveryStrategy struct{}

func (c *CarDeliveryStrategy) deliver(product Product) {
	fmt.Printf("Deliver product with name '%s' by car.\n", product.Name)
}

type BicycleDeliveryStrategy struct{}

func (b *BicycleDeliveryStrategy) deliver(product Product) {
	fmt.Printf("Deliver product with name '%s' by bicycle.\n", product.Name)
}

// Using strategies

type DeliveryService struct {
	strategy DeliveryStrategy
}

func (ds *DeliveryService) DeliverProduct(product Product) {
	ds.strategy.deliver(product)
}

func (ds *DeliveryService) SetStrategy(strategy DeliveryStrategy) {
	ds.strategy = strategy
}
