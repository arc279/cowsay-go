package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	cowsay "github.com/arc279/cowsay-go/lib"
)

type CliOption struct {
	cow     cowsay.CowOption
	face    cowsay.FaceOption
	cowname string
	list    bool
	random  bool
}

func cmd(opt CliOption) {
	flag.StringVar(&opt.cow.Eyes, "e", cowsay.DefaultCowOption.Eyes, "eyes")
	flag.StringVar(&opt.cowname, "f", "", "cowname")
	flag.StringVar(&opt.cow.Tongue, "T", cowsay.DefaultCowOption.Tongue, "tongue")
	flag.UintVar(&opt.cow.Columns, "W", cowsay.DEFAULT_BALLOON_WIDTH, "columns")
	flag.BoolVar(&opt.list, "l", false, "list cows")
	flag.BoolVar(&opt.random, "random", false, "random select")

	flag.BoolVar(&opt.face.Borg, "b", false, "face borg")
	flag.BoolVar(&opt.face.Dead, "d", false, "face dead")
	flag.BoolVar(&opt.face.Greedy, "g", false, "face greedy")
	flag.BoolVar(&opt.face.Paranoid, "p", false, "face paranoid")
	flag.BoolVar(&opt.face.Stoned, "s", false, "face stoned")
	flag.BoolVar(&opt.face.Tired, "t", false, "face tired")
	flag.BoolVar(&opt.face.Wired, "w", false, "face wired")
	flag.BoolVar(&opt.face.Young, "y", false, "face young")
	flag.Parse()
	//fmt.Println(opt)
	opt.face.MakeUp(&opt.cow)

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
		if opt.random {
			rand.Seed(time.Now().UnixNano())
			cows := cowsay.AssetNames()
			i := rand.Intn(len(cows))
			opt.cowname = cows[i]
		}

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
