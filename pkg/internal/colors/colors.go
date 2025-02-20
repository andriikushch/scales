package colors

const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

const (
	BGBlack   = "\033[40m"
	BGGrey    = "\033[100m"
	BGRed     = "\033[41m"
	BGGreen   = "\033[42m"
	BGYellow  = "\033[43m"
	BGBlue    = "\033[44m"
	BGMagenta = "\033[45m"
	BGCyan    = "\033[46m"
	BGWhite   = "\033[47m"
	End       = "\033[0m"
)

var bgColors = []string{
	"\033[1;41m",  // Bright Red
	"\033[1;42m",  // Bright Green
	"\033[1;43m",  // Bright Yellow
	"\033[1;44m",  // Bright Blue
	"\033[1;45m",  // Bright Magenta
	"\033[1;46m",  // Bright Cyan
	"\033[1;47m",  // Bright White
	"\033[1;100m", // Bright Grey
	"\033[1;101m", // Bright Light Red
	"\033[1;102m", // Bright Light Green
	"\033[1;103m", // Bright Light Yellow
	"\033[1;104m", // Bright Light Blue
}

func GetColor(i int) string {
	if i < 0 {
		return BGGrey
	}
	i = i % 12

	return bgColors[i]
}
