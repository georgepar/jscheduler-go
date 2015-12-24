package jscheduler

type Thread struct {
	Name string
	Tid int
}
type ThreadList []Thread

type ThreadGroup struct {
	Descriptor string
	Threads ThreadList
}

func NewThreadGroup(f string) *ThreadGroup {
	return &ThreadGroup{Descriptor:f, Threads:make([]Thread, 0)}
}


func NewThreadList() ThreadList {
	return make([]Thread, 0)
}