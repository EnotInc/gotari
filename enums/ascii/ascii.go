package ascii

const (
	ClearAfter   = "\033[0J"
	ClearView    = "\033[2J"
	ClearHistory = "\033[3J"
	MoveToStart  = "\033[0H"

	SaveTerminal  = "\033[?1049h"
	ResetTerminal = "\033[?1049l"

	MoveUp         = "\033[A"
	MoveDown       = "\033[B"
	SaveCursorPos  = "\033[s"
	ResetCursorPos = "\033[u"

	BorderHorisontal = "─"
	BorderVertical   = "│"
	BorderUpperRight = "╮"
	BorderLowerRight = "╯"
	BorderUpperLeft  = "╭"
	BorderLowerLeft  = "╰"
)
