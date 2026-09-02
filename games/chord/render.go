package chord

import (
	"fmt"
	"math"
	"strings"

	"github.con/enotinc/gotari/enums/ascii"
	color "github.con/enotinc/gotari/enums/color"
	"github.con/enotinc/gotari/enums/keys"
)

func (c *Chord) render() []*string {
	if c.ended {
		return c.renderStats()
	} else {
		return c.renderInputBox()
	}
}

func (c *Chord) renderStats() []*string {
	var render []*string
	t := c.end.Sub(c.begin).Seconds()
	round := math.Round(t*100) / 100

	result := fmt.Sprintf(" ended in: %s%.2f%s", color.LightBlue, round, color.Reset)
	typed := fmt.Sprintf(" typed %s%d%s symbols with %s%d%s error(s)", color.DarkBlue, c.typed, color.Reset, color.Red, c.errors, color.Reset)
	speed := fmt.Sprintf(" average speed: %s%.2f%s", color.Purple, float64(c.typed)/round*60, color.Reset)
	restart := " press 'r' restart"

	render = append(render, &result)
	render = append(render, &typed)
	render = append(render, &speed)
	render = append(render, &restart)

	return render
}

func (c *Chord) renderInputBox() []*string {
	var render []*string

	var upper strings.Builder
	var middle strings.Builder
	var lower strings.Builder

	border := strings.Repeat(ascii.BorderHorisontal, len(c.text))

	upper.WriteString(ascii.BorderUpperLeft)
	upper.WriteString(border)
	upper.WriteString(ascii.BorderUpperRight)
	upper.WriteString(color.Reset)

	lower.WriteString(ascii.BorderLowerLeft)
	lower.WriteString(border)
	lower.WriteString(ascii.BorderLowerRight)
	lower.WriteString(color.Reset)

	var line strings.Builder
	for i, ch := range c.input {
		if byte(ch) != c.text[i] {
			line.WriteString(color.Red)
			if ch == keys.Space {
				ch = '\u00b7'
			}
		} else {
			line.WriteString(color.DarkBlue)
		}
		line.WriteRune(ch)
	}
	line.WriteString(color.Gray)
	line.WriteString(c.text[len(c.input):])

	middle.WriteString(ascii.BorderVertical)
	middle.WriteString(line.String())
	middle.WriteString(color.Reset)
	middle.WriteString(ascii.BorderVertical)

	u := upper.String()
	m := middle.String()
	l := lower.String()

	render = append(render, &u)
	render = append(render, &m)
	render = append(render, &l)

	return render
}
