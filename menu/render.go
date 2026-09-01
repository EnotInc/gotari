package menu

import "strings"

func (c *card) render(selected bool) []*string {
	var render []*string
	amount := 20
	color := gray
	if selected {
		amount = 25
		color = blue
	}

	border := strings.Repeat(horisontal, amount)

	var upperBorder strings.Builder
	var lowerBorder strings.Builder
	var middle strings.Builder

	upperBorder.WriteString(color)
	upperBorder.WriteString(border)
	upperBorder.WriteString(uppreCorner)
	upperBorder.WriteString(reset)

	lowerBorder.WriteString(color)
	lowerBorder.WriteString(border)
	lowerBorder.WriteString(lowerCorner)
	lowerBorder.WriteString(reset)

	gameName := c.name
	nameAmount := amount - len(gameName)
	if nameAmount <= 0 {
		nameAmount = 0
		gameName = c.name[:amount]
	}
	nameSpace := strings.Repeat(" ", nameAmount)
	middle.WriteString(nameSpace)
	middle.WriteString(gameName)
	middle.WriteString(color)
	middle.WriteString(vertical)
	middle.WriteString(reset)

	u := upperBorder.String()
	m := middle.String()
	l := lowerBorder.String()

	render = append(render, &u)
	render = append(render, &m)
	render = append(render, &l)

	return render
}
