package cowsay

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/donatj/wordwrap"
	"github.com/mattn/go-runewidth"
)

const DEFAULT_BALLOON_WIDTH uint = 40

type BallonBorder string

func (self BallonBorder) Bounds(category int) (uint8, uint8) {
	return self[category], self[category+1]
}

const CowSayBallonBorder = `/\||\/<>`
const CowThinkBallonBorder = `()()()()`

const (
	BORDER_TOP    = 0
	BORDER_MIDDLE = 2
	BORDER_BOTTOM = 4
	BORDER_SINGLE = 6
)

type StringWithTerminalWidth struct {
	msg  string
	msgW int
}

func WrapLines(reader io.Reader, wrapW uint) ([]StringWithTerminalWidth, int) {
	scanner := bufio.NewScanner(reader)

	pairs := []StringWithTerminalWidth{}
	maxW := -1

	for scanner.Scan() {
		s := strings.Replace(scanner.Text(), "\t", " ", -1)
		s2 := wordwrap.WrapString(s, wrapW)

		for _, l := range strings.Split(s2, "\n") {
			w := runewidth.StringWidth(l)
			pairs = append(pairs, StringWithTerminalWidth{msg: l, msgW: w})

			if w > maxW {
				maxW = w
			}
		}
	}

	return pairs, maxW
}

func makeBalloon(pairs []StringWithTerminalWidth, border BallonBorder, maxW int) []string {
	ret := []string{}

	ret = append(ret, fmt.Sprintf(" %s \n", strings.Repeat("_", maxW+2)))
	for i, p := range pairs {
		open, close := border.Bounds(BORDER_MIDDLE)
		switch i {
		case 0:
			if len(pairs) == 1 {
				open, close = border.Bounds(BORDER_SINGLE)
			} else {
				open, close = border.Bounds(BORDER_TOP)
			}
		case len(pairs) - 1:
			open, close = border.Bounds(BORDER_BOTTOM)
		}
		ret = append(ret, fmt.Sprintf("%c %s%s %c\n", open, p.msg, strings.Repeat(" ", maxW-p.msgW), close))
	}
	ret = append(ret, fmt.Sprintf(" %s \n", strings.Repeat("-", maxW+2)))

	return ret
}
