package colors

const Black = "\033[30m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const White = "\033[37m"

const BGBlack = "\033[40m"
const BGRed = "\033[41m"
const BGGreen = "\033[42m"
const BGYellow = "\033[43m"
const BGBlue = "\033[44m"
const BGMagenta = "\033[45m"
const BGCyan = "\033[46m"
const BGWhite = "\033[47m"
const End = "\033[0m"

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
		return BGBlack
	}
	i = i % 12

	return bgColors[i]
}
