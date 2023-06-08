package design

type Cloneable interface {
	Clone() *SunWuKong
}

type SunWuKong struct {
	name  string
	real  bool
	skill []string
}

func (wukong SunWuKong) Clone() *SunWuKong {
	return &SunWuKong{
		name:  "keith",
		real:  true,
		skill: []string{"swimming", "fly"},
	}
}
