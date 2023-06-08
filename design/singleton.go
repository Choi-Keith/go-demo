package design

type Singleton struct {
	name string
}

var instance *Singleton

func (a *Singleton) getInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{name: "keith"}
	}
	return instance
}
