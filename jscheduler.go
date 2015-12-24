package main
import (
	"./jscheduler"
	"fmt"
)



func main() {
	/*
		example := `
		"Reference Handler" #2 daemon prio=10 os_prio=0 tid=0x00007f1ff8409000 nid=0xb14 in Object.wait()
		basghahav
		blahhhh
		basghahipj
		asghash
		"Gang worker#18 (Parallel GC Threads)" os_prio=0 tid=0x00007f1ff803a000 nid=0xb07 runnable
		325235235
		235
		51235
		135
		gwsdgwsedg
		"Concurrent Mark-Sweep GC Thread" os_prio=0 tid=0x00007f1ff8159800 nid=0xb12 runnable
		sdhah
		asedh
		"Concurrent Mark-Sweep GC Thread" os_prio=0 tid=0x00007f1ff8159800 nid=0xb12 runnable
		`
	*/
	//g1, _ := IsThreadDescription("\"Reference Handler\" #2 daemon prio=10 os_prio=0 tid=0x00007f1ff8409000 nid=0xb14 in Object.wait()")
	//g2, _ := IsThreadDescription("\"Gang worker#18 (Parallel GC Threads)\" os_prio=0 tid=0x00007f1ff803a000 nid=0xb07 runnable")
	//g3, _ := IsThreadDescription("\"Concurrent Mark-Sweep GC Thread\" os_prio=0 tid=0x00007f1ff8159800 nid=0xb12 runnable")

	//fmt.Println(g1["nid"])
	//fmt.Println(g2["nid"])
	//fmt.Println(g3["nid"])

	//out, _ := exec.Command("/opt/oracle/jdk1.8.0_51/bin/jstack", "-l", "10188").Output()



	out, _ := jscheduler.GetThreadDump("34768")

	fmt.Println(out)

	parsed, _ := jscheduler.ParseThreadDump(out)

	fmt.Println(parsed)

	fmt.Println(len(parsed))

}
