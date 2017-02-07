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

type BallonBorder struct {
	TopL, TopR       string
	MiddleL, MiddleR string
	BottomL, BottomR string

	SingleL, SingleR string
}

var CowSayBallonBorder = BallonBorder{
	TopL: `/`, TopR: `\`,
	MiddleL: `|`, MiddleR: `|`,
	BottomL: `\`, BottomR: `/`,
	SingleL: `<`, SingleR: `>`,
}

var CowThinkBallonBorder = BallonBorder{
	TopL: `(`, TopR: `)`,
	MiddleL: `(`, MiddleR: `)`,
	BottomL: `(`, BottomR: `)`,
	SingleL: `(`, SingleR: `)`,
}

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
		open, close := border.MiddleL, border.MiddleR
		switch i {
		case 0:
			if len(pairs) == 1 {
				open, close = border.SingleL, border.SingleR
			} else {
				open, close = border.TopL, border.TopR
			}
		case len(pairs) - 1:
			open, close = border.BottomL, border.BottomR
		}
		ret = append(ret, fmt.Sprintf("%s %s%s %s\n", open, p.msg, strings.Repeat(" ", maxW-p.msgW), close))
	}
	ret = append(ret, fmt.Sprintf(" %s \n", strings.Repeat("-", maxW+2)))

	return ret
}
