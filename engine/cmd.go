package engine

type CMD interface {
	String() string
}

type cmdstring struct {
	s string
}

func (c *cmdstring) String() string {
	return c.s
}

func newCMD(text string) CMD {
	return &cmdstring{s: text}
}

var SelectGame = newCMD("SelectGame")
var CloseGame = newCMD("CloseGame")
var QuitGotari = newCMD("QuitGotari")
