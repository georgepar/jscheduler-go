package jscheduler
import (
	"unsafe"
	"syscall"
	"golang.org/x/sys/unix"
	"regexp"
	"net/http/cookiejar"
)

// SetAffinity attend the cpu list to pid,
// note: SetAffinity apply to thread ID only,
// to fully control one process, call SetAffinity for all thread of the process.
// use os.GetThreadIDs() to get all thread of the process
func SetAffinity(pid int, cpus []int) error {
	var mask [1024 / 64]uintptr
	if pid <= 0 {
		pid, _, _ = syscall.RawSyscall(unix.SYS_GETPID, 0, 0, 0)
	}
	for _, cpuIdx := range cpus {
		cpuIndex := uint(cpuIdx)
		mask[cpuIndex / 64] |= 1 << (cpuIndex % 64)
	}
	_, _, err := syscall.RawSyscall(unix.SYS_SCHED_SETAFFINITY, uintptr(pid), uintptr(len(mask) * 8), uintptr(unsafe.Pointer(&mask[0])))
	if err != 0 {
		return err
	}
	return nil
}


func SetAffinityThreadpool(nameTidMap map[string] int, filter *regexp.Regexp, cpus []int) error {
	for name, tid := range nameTidMap {
		if filter.MatchString(name) {
			err := SetAffinity(tid, cpus)
			if err != nil {
				return err
			}
		}
	}
	return nil
}


func SetPriorityThreadpool(nameTidMap map[string] int, filter *regexp.Regexp, prio int) error {
	for name, tid := range nameTidMap {
		if filter.MatchString(name) {
			err := unix.Setpriority(unix.PRIO_PROCESS, tid, prio)
			if err != nil {
				return err
			}
		}
	}
	return nil
}


