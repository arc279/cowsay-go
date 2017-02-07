package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	cowsay "github.com/arc279/cowsay-go/lib"
)

type CliOption struct {
	cow     cowsay.CowOption
	cowname string
	list    bool
}

type FaceOption struct {
	Borg     bool
	Dead     bool
	Greedy   bool
	Paranoid bool
	Stoned   bool
	Tired    bool
	Wired    bool
	Young    bool
}

func (self FaceOption) MakeUp(opt *cowsay.CowOption) {
	if self.Borg {
		opt.Eyes = "=="
	}
	if self.Dead {
		opt.Eyes = "xx"
		opt.Tongue = "U "
	}
	if self.Greedy {
		opt.Eyes = "$$"
	}
	if self.Paranoid {
		opt.Eyes = "@@"
	}
	if self.Stoned {
		opt.Eyes = "**"
		opt.Tongue = "U "
	}
	if self.Tired {
		opt.Eyes = "--"
	}
	if self.Wired {
		opt.Eyes = "OO"
	}
	if self.Young {
		opt.Eyes = ".."
	}
}

func cmd(opt CliOption) {
	face := FaceOption{}

	flag.StringVar(&opt.cow.Eyes, "e", cowsay.DefaultCowOption.Eyes, "eyes")
	flag.StringVar(&opt.cowname, "f", "", "cowname")
	flag.StringVar(&opt.cow.Tongue, "T", cowsay.DefaultCowOption.Tongue, "tongue")
	flag.UintVar(&opt.cow.Columns, "W", cowsay.DEFAULT_BALLOON_WIDTH, "columns")
	flag.BoolVar(&opt.list, "l", false, "list cows")

	flag.BoolVar(&face.Borg, "b", false, "face borg")
	flag.BoolVar(&face.Dead, "d", false, "face dead")
	flag.BoolVar(&face.Greedy, "g", false, "face greedy")
	flag.BoolVar(&face.Paranoid, "p", false, "face paranoid")
	flag.BoolVar(&face.Stoned, "s", false, "face stoned")
	flag.BoolVar(&face.Tired, "t", false, "face tired")
	flag.BoolVar(&face.Wired, "w", false, "face wired")
	flag.BoolVar(&face.Young, "y", false, "face young")
	flag.Parse()
	//fmt.Println(opt)
	face.MakeUp(&opt.cow)

	if opt.list {
		for _, name := range cowsay.AssetNames() {
			fmt.Println(name)
		}
		os.Exit(2)
	}

	msg := func() string {
		if flag.NArg() == 0 {
			reader := bufio.NewReader(os.Stdin)
			msg, err := ioutil.ReadAll(reader)
			if err != nil {
				log.Fatal(err)
			}
			return string(msg)
		} else {
			return strings.Join(flag.Args(), " ")
		}
	}()

	cow := func() io.Reader {
		if len(opt.cowname) == 0 {
			return strings.NewReader(cowsay.DEFAULT_COW)
		} else {
			data, err := cowsay.Asset(opt.cowname)
			if err != nil {
				log.Fatal(err)
			}
			return bytes.NewReader(data)
		}
	}()

	err := cowsay.CowSay(cow, msg, opt.cow, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	selfname := path.Base(os.Args[0])

	opt := CliOption{}
	if selfname == "cowthink" {
		opt.cow.Thoughts = "o"
		opt.cow.Border = cowsay.CowThinkBallonBorder
	} else {
		opt.cow.Thoughts = cowsay.DefaultCowOption.Thoughts
		opt.cow.Border = cowsay.CowSayBallonBorder
	}

	cmd(opt)
}
