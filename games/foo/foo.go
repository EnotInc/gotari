package foo

import (
	"fmt"
)

type FooGame struct {
	list   []string
	cursor int
}

func Init() *FooGame {
	return &FooGame{
		list: []string{"foo", "bar", "baz"},
	}
}

func (fg *FooGame) Render() []*string {
	var str []*string

	for i, l := range fg.list {
		c := "  "
		if fg.cursor == i {
			c = "> "
		}
		line := fmt.Sprintf("%s%s", c, l)
		str = append(str, &line)
	}

	return str
}

func (fg *FooGame) Handle(key rune) {
	switch key {
	case 'j':
		if fg.cursor < len(fg.list)-1 {
			fg.cursor += 1
		}
	case 'k':
		if fg.cursor > 0 {
			fg.cursor -= 1
		}
	}
}
