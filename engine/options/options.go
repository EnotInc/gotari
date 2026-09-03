package options

type Options struct {
	Fullscreen  bool
	EnableMouse bool
}

func GetOptions(fullscreen bool, enableMouse bool) *Options {
	return &Options{
		Fullscreen:  fullscreen,
		EnableMouse: enableMouse,
	}
}
