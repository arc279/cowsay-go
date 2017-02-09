package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	cowsay "github.com/arc279/cowsay-go/lib"
)

type CliOption struct {
	cow     cowsay.CowOption
	face    cowsay.FaceOption
	cowfile string
	list    bool
	random  bool
}

func cmd(opt CliOption) {
	flag.StringVar(&opt.cow.Eyes, "e", cowsay.DefaultCowOption.Eyes, "eyes")
	flag.StringVar(&opt.cowfile, "f", "", "cowfile")
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

	cowlists := func(pattern string) []string {
		files, err := filepath.Glob(pattern)
		if err != nil {
			log.Fatal(err)
		}
		return files
	}(path.Join(cowsay.COWS_DIR, "*.cow"))

	if opt.list {
		for _, name := range cowlists {
			fmt.Println(path.Base(name))
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
			if len(cowlists) == 0 {
				return strings.NewReader(cowsay.DEFAULT_COW)
			}
			rand.Seed(time.Now().UnixNano())
			i := rand.Intn(len(cowlists))
			opt.cowfile = path.Base(cowlists[i])
		}

		if opt.cowfile == "" {
			return strings.NewReader(cowsay.DEFAULT_COW)
		}

		fp, err := os.Open(path.Join(cowsay.COWS_DIR, opt.cowfile))
		if err != nil {
			log.Fatal(err)
		}
		return fp
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
