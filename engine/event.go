package engine

type Event interface {
	isEvent()
}

type KeyEvent struct {
	Key rune
}

type MouseEvent struct {
	X int
	Y int
}

func (KeyEvent) isEvent()   {}
func (MouseEvent) isEvent() {}
