package cowsay

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

var expect = `        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`

func TestCowString(t *testing.T) {
	var s = `        {{.Thoughts}}   ^__^
         {{.Thoughts}}  ({{.Eyes}})\_______
            (__)\       )\/\
            {{.Tongue}}  ||----w |
                ||     ||
`

	opt := CowOption{Thoughts: `\`, Eyes: `oo`, Tongue: `  `}

	result, err := makeCow(s, opt)
	if err != nil {
		t.Error(err)
	}
	if result != expect {
		t.Error("invalid cow")
	}
}

func TestCowFile(t *testing.T) {
	var cowname = "default.cow"
	fp, err := os.Open(path.Join(COWS_DIR, cowname))
	if err != nil {
		t.Error(err)
	}
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		t.Error(err)
	}

	opt := CowOption{Thoughts: `\`, Eyes: `oo`, Tongue: `  `}

	s := string(data)
	result, err := makeCow(s, opt)
	if err != nil {
		t.Error(err)
	}
	if result != expect {
		t.Error("invalid cow")
	}
}
