package cowsay

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

func (self FaceOption) MakeUp(opt *CowOption) {
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
