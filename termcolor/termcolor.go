package termcolor

import "fmt"

// Read https://en.wikipedia.org/wiki/ANSI_escape_code to know how these crazy looking character combinations work

func BgColor(str string, col *[3]int) string {
	result := fmt.Sprintf("\033[48;2;%d;%d;%dm%s", col[0], col[1], col[2], str)
	return result
}

func FgColor(str string, col *[3]int) string {
	result := fmt.Sprintf("\033[38;2;%d;%d;%dm%s", col[0], col[1], col[2], str)
	return result
}

func FullColor(str string, fg *[3]int, bg *[3]int) string {
	result := fmt.Sprintf(
		"\033[38;2;%d;%d;%dm\033[48;2;%d;%d;%dm%s", fg[0], fg[1], fg[2], bg[0], bg[1], bg[2], str,
	)
	return result
}
