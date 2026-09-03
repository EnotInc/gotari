package cmd

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

// -==[ Game Commands ]==-
var SelectGame = newCMD("SelectGame")
var CloseGame = newCMD("CloseGame")

// -==[ Menu Commands ]==-
var QuitGotari = newCMD("QuitGotari")
