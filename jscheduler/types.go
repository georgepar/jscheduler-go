package jscheduler

import (
	"regexp"
)

type CpuPool []int

func NewCpuPool() CpuPool {
	return make([]int, 0)
}

type ThreadSpecification struct {
	Filter *regexp.Regexp
	Prio   int
	Cpus   CpuPool
}

type Thread struct {
	Name    string
	Tid     int
	Prio    int
	Cpus    CpuPool
	HasSpec bool
}

func NewThread(name string, tid int) Thread {
	return Thread{
		Name:    name,
		Tid:     tid,
		Prio:    0,
		Cpus:    NewCpuPool(),
		HasSpec: false,
	}
}

func (t *Thread) SetSpec(spec *ThreadSpecification) {
	if spec.Filter.MatchString(t.Name) {
		t.Prio = spec.Prio
		t.Cpus = spec.Cpus
		t.HasSpec = true
	}
}

type ThreadList []Thread

func NewThreadList() ThreadList {
	return make([]Thread, 0)
}
