package cowsay

import (
	"strings"
	"testing"
)

func TestBalloon1(t *testing.T) {
	var s = `僕は死を軽んずることを大したことだとは思わない。
その死がもし、自ら引き受けた責任の観念に深く根ざしていないかぎり、
それは単に貧弱さの表れ、若気のいたりにしか過ぎない。
		--サン・テグジュペリ`
	var expect = ` ____________________________ 
/ 僕は死を軽んずることを大し \
| たことだとは思わない。     |
| その死がもし、自ら引き受け |
| た責任の観念に深く根ざして |
| いないかぎり、             |
| それは単に貧弱さの表れ、若 |
| 気のいたりにしか過ぎない。 |
\   --サン・テグジュペリ     /
 ---------------------------- 
`

	pairs, maxW := WrapLines(strings.NewReader(s), DEFAULT_BALLOON_WIDTH)
	if maxW != 26 {
		t.Error("Max terminal width should be 26")
	}

	msg := strings.Join(makeBalloon(pairs, CowSayBallonBorder, maxW), "")
	if msg != expect {
		t.Error("Invalid balloon message.")
	}
}

func TestBalloon2(t *testing.T) {
	var s = `		--サン・テグジュペリ`
	var expect = ` ________________________ 
<   --サン・テグジュペリ >
 ------------------------ 
`

	pairs, maxW := WrapLines(strings.NewReader(s), DEFAULT_BALLOON_WIDTH)
	if maxW != 22 {
		t.Error("Max terminal width should be 22")
	}

	msg := strings.Join(makeBalloon(pairs, CowSayBallonBorder, maxW), "")
	if msg != expect {
		t.Error("Invalid balloon message.")
	}
}
