package menu

import (
	"strings"

	color "github.con/enotinc/gotari/enums/color"
)

func (c *card) render(selected bool) []*string {
	var render []*string
	amount := 20
	clr := color.Gray
	if selected {
		amount = 25
		clr = color.LightBlue
	}

	border := strings.Repeat(horisontal, amount)

	var upperBorder strings.Builder
	var lowerBorder strings.Builder
	var middle strings.Builder

	upperBorder.WriteString(clr)
	upperBorder.WriteString(border)
	upperBorder.WriteString(uppreCorner)
	upperBorder.WriteString(color.Reset)

	lowerBorder.WriteString(clr)
	lowerBorder.WriteString(border)
	lowerBorder.WriteString(lowerCorner)
	lowerBorder.WriteString(color.Reset)

	gameName := c.name
	nameAmount := amount - len(gameName)
	if nameAmount <= 0 {
		nameAmount = 0
		gameName = c.name[:amount]
	}
	nameSpace := strings.Repeat(" ", nameAmount)
	middle.WriteString(nameSpace)
	middle.WriteString(gameName)
	middle.WriteString(clr)
	middle.WriteString(vertical)
	middle.WriteString(color.Reset)

	u := upperBorder.String()
	m := middle.String()
	l := lowerBorder.String()

	render = append(render, &u)
	render = append(render, &m)
	render = append(render, &l)

	return render
}
