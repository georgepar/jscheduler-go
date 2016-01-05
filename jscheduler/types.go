package jscheduler

import (
	"regexp"
	"runtime"
	"fmt"
	"strings"
	"strconv"
)

type CpuPool []int

func NewEmptyCpuPool() CpuPool {
	return make([]int, 0)
}

func NewCpuPool(numCpus int) CpuPool {
	pool := make([]int, numCpus)

	for i := 0; i < numCpus; i++ {
		pool[i] = i
	}

	return pool
}

func ParseCpuPool(pool string) CpuPool{
	elements := strings.Split(pool, ",")
	cpus := NewEmptyCpuPool()
	cpuSet := make(map[int]struct{})
	for _, el := range elements {
		step := 1
		cpuRange := strings.Split(el, "-")
		if len(cpuRange) > 1 {
			rangeSplit := strings.Split(cpuRange[1], ":")
			if len(rangeSplit) > 1 {
				step, _ = strconv.Atoi(rangeSplit[1])
			}
			c1, _ := strconv.Atoi(cpuRange[0])
			c2, _ := strconv.Atoi(rangeSplit[0])
			for c:=c1; c<=c2; c+=step {
				if _, cpuExists := cpuSet[c]; !cpuExists {
					cpus = append(cpus, c)
					cpuSet[c] = struct{}{}
				}
			}
		} else {
			c, _ := strconv.Atoi(cpuRange[0])
			if _, cpuExists := cpuSet[c]; !cpuExists {
				cpus = append(cpus, c)
				cpuSet[c] = struct{}{}
			}
		}
	}
	return cpus
}

type ThreadSpecification struct {
	Filter *regexp.Regexp
	Prio   int
	Cpus   CpuPool
}

func NewThreadSpecification() ThreadSpecification {
	return ThreadSpecification{
		Filter: &regexp.Compile(""),
		Prio: 0,
		Cpus: NewEmptyCpuPool(),
	}
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
		Cpus:    NewCpuPool(runtime.NumCPU()),
		HasSpec: false,
	}
}

func (t *Thread) FilterAndSetSpec(spec *ThreadSpecification) {
	if spec.Filter.MatchString(t.Name) {
		t.SetSpec(spec)
	}
}

func (t *Thread) SetSpec(spec *ThreadSpecification) {
	t.Prio = spec.Prio
	t.Cpus = spec.Cpus
	t.HasSpec = true
}

type ThreadList []Thread

func NewThreadList() ThreadList {
	return make([]Thread, 0)
}


type ThreadSpecArgList struct {
	FlagValue []ThreadSpecification
}

func NewThreadSpecArgList() ThreadSpecArgList {
	return ThreadSpecArgList{
		FlagValue: make([]ThreadSpecification, 0),
	}
}

func (lst *ThreadSpecArgList) String() string {
	strLst := make([]string, 0)

	for _, el := range lst.FlagValue {
		filter := el.Filter.String()
		prio := el.Prio
		cpus := strings.Join(el.Cpus[:], ",")
		strLst = append(strLst, fmt.Sprintf("\"%s\";%d;%s", filter, prio, cpus))
	}

	return strings.Join(strLst[:], "::")
}


func (lst *ThreadSpecArgList) Set(s string) error {
	strLst := strings.Split(s, "::")
	lst.FlagValue = make([]ThreadSpecification, 0)
	for _, el := range strLst {
		ts := NewThreadSpecification()
		tsEl := strings.Split(el, ";")
		if tsEl[0] != "" {
			ts.Filter = &regexp.Compile(tsEl[0])
		}
		if tsEl[1] != "" {
			ts.Prio = int(tsEl[1])
		}
		if tsEl[2] != "" {
			ts.Cpus = ParseCpuPool(tsEl[2])
		}
		lst.FlagValue = append(lst.FlagValue, ts)
	}
	return nil
}

func (lst *ThreadSpecArgList) Get() []ThreadSpecification {
	return lst.FlagValue
}

func (lst *ThreadSpecArgList) IsSet() bool {
	return len(lst.FlagValue) > 0
}