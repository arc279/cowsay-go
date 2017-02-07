package cowsay

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"text/template"
)

const DEFAULT_COW = `        {{.Thoughts}}   ^__^
         {{.Thoughts}}  ({{.Eyes}})\_______
            (__)\       )\/\
             {{.Tongue}} ||----w |
                ||     ||
`

type CowOption struct {
	Thoughts string
	Eyes     string
	Tongue   string
	Columns  uint
	Border   BallonBorder
}

var DefaultCowOption = CowOption{
	Thoughts: `\`,
	Eyes:     `oo`,
	Tongue:   `  `,
	Columns:  DEFAULT_BALLOON_WIDTH,
	Border:   CowSayBallonBorder,
}

func makeCow(cow string, opt CowOption) (string, error) {
	tmpl, err := template.New("cow").Parse(cow)
	if err != nil {
		return "", err
	}

	var doc bytes.Buffer
	err = tmpl.Execute(&doc, opt)
	if err != nil {
		return "", err
	}

	return doc.String(), nil
}

func CowSay(cowreader io.Reader, message string, opt CowOption, w io.Writer) error {
	pairs, maxW := WrapLines(strings.NewReader(message), opt.Columns)
	msg := strings.Join(makeBalloon(pairs, opt.Border, maxW), "")

	contents, err := ioutil.ReadAll(cowreader)
	if err != nil {
		return err
	}
	cow, err := makeCow(string(contents), opt)
	if err != nil {
		return err
	}

	bw := bufio.NewWriter(w)
	bw.WriteString(msg)
	bw.WriteString(cow)
	return bw.Flush()
}
