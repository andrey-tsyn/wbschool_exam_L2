package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Применимость:
 - Когда имеется более одного объекта, который может выполнить запрос, но неизвестно, какой объект конкретно это будет.
 - Когда набор объектов для обработки запроса может изменяться динамически.
 - Когда важна последовательность обработки

Плюсы:
 -  Уменьшается зависимость клиента от обработчиков
 -  Простота добавления новых обработчиков
Минусы:
 -  Усложняет код из-за увеличения кол-ва структур
*/

type Payload struct{}

type Handler interface {
	setNext(*Handler)
	exec(*Payload)
}

type Handler1 struct {
	next *Handler
}

func (h Handler1) setNext(handler *Handler) {
	h.next = handler
}

func (h Handler1) exec(payload Payload) {
	// Do some important things
	if h.next != nil {
		h.exec(payload)
	}
}

type Handler2 struct {
	next *Handler
}

func (h Handler2) setNext(handler *Handler) {
	h.next = handler
}

func (h Handler2) exec(payload Payload) {
	// Do some important things
	if h.next != nil {
		h.exec(payload)
	}
}
