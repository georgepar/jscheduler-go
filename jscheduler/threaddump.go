package jscheduler
import (
	"regexp"
	"os/exec"
	"strings"
	"strconv"
)


const THREAD_DESCRIPTOR1 string = `^"(?P<name>[^"]+)".+prio=(?P<prio>[0-9]+)\s+os_prio=(?P<os_prio>[0-9]+)\s+tid=(?P<tid>0x[0-9a-f]+)\s+nid=(?P<nid>0x[0-9a-f]+).+`
const THREAD_DESCRIPTOR2 string = `^"(?P<name>[^"]+)"\s+os_prio=(?P<os_prio>[0-9]+)\s+tid=(?P<tid>0x[0-9a-f]+)\s+nid=(?P<nid>0x[0-9a-f]+).+`

// Use the given regex to decompose a line of the thread dump into the matched fields
func DecomposeTreadDumpLineRe(threadDumpLine string, r *regexp.Regexp) (groups map[string] string, err error) {
	matches := r.FindStringSubmatch(threadDumpLine)
	names := r.SubexpNames()

	groups = make(map[string] string)

	for i,name := range names {
		groups[name] = matches[i]
	}

	return
}

// Match the groups of a thread dump line and get the corresponding fields
func DecomposeTreadDumpLine(threadDumpLine string) (groups map[string] string, err error) {
	//TODO: Optimize/Combine regex. One thread is slipping away when testing on the Intellij process.
	r1 := regexp.MustCompile(THREAD_DESCRIPTOR1)
	r2 := regexp.MustCompile(THREAD_DESCRIPTOR2)


	switch {
	case r1.MatchString(threadDumpLine):
		groups, err = DecomposeTreadDumpLineRe(threadDumpLine, r1)
	case r2.MatchString(threadDumpLine):
		groups, err = DecomposeTreadDumpLineRe(threadDumpLine, r2)
	}

	return
}

// Parse a Java thread dump taken with JStack (or with SIGQUIT)
func ParseThreadDump(threadDump string) (*ThreadList, error) {
	nameToNative := NewThreadList()//NewThreadGroup(THREAD_DESCRIPTOR1)
	lines := strings.Split(threadDump, "\n")

	for _, line := range lines {
		fields, err := DecomposeTreadDumpLine(line)
		if err != nil {
			return &nameToNative, err
		}
		// ParseInt base = 0 -> It is implied to be 16 by the 0x prefix
		val, _ := strconv.ParseInt(fields["nid"], 0, 0)
		if(fields["name"] != "") {
			nameToNative = append(nameToNative, Thread{Name: fields["name"], Tid: int(val)})
		}
	}


	return &nameToNative, nil
}

// Take a thread dump with JStack
//TODO: Can be done natively with syscall.Kill(pid, SIGQUIT) if we find a way to capture the output
func GetThreadDump(pid string) (string, error) {
	out, err := exec.Command("/opt/oracle/jdk1.8.0_51/bin/jstack", "-l", pid).Output()
	return string(out), err
}
