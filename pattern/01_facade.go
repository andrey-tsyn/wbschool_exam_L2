package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Применимость: следует использовать, если нужно предоставить простой интерфейс для выполнения сложных действий с
использованием нескольких классов.

Плюсы:
 -  Даёт возможность пользоваться сложной системе в упрощённом виде. Таким образом фасад и его подсистемы может разрабатывать
	одна команда, а другие использовать в большинстве случаев фасад, избавляя от необходимости вникать в детали реализации
	и позволяя сконцентрироваться на основной задаче
Минусы:
 -  Фасад может стать god классом, начать отвечать за слишком многое

Примеры использования:
 -  Конвертер видео. Вместо использования нескольких классов можно использовать фасад с
	методом convert(file []byte, resultFormat string)
*/

// ===================================================
type CarFacade struct {
	engine *Engine
	doors  []Door
}

func (c *CarFacade) PrepareCarForDriving() {
	c.engine.StartEngine()

	for _, door := range c.doors {
		door.UnlockDoor()
	}
}

func (c *CarFacade) LeaveCar() {
	c.engine.StopEngine()

	for _, door := range c.doors {
		door.LockDoor()
	}
}

// ===================================================
type Engine struct {
	isStarted bool
}

func (e *Engine) StartEngine() {
	e.isStarted = true
}

func (e *Engine) StopEngine() {
	e.isStarted = false
}

// ===================================================
type Door struct {
	isLocked bool
}

func (d *Door) LockDoor() {
	d.isLocked = true
}

func (d *Door) UnlockDoor() {
	d.isLocked = false
}
