package main

import (
	"fmt"

	"./task"
)

func main() {
	defer func() {
		fmt.Println("test done")
	}()

	// fmt.Println("get command is", snmpGetPrefix+rmPostFix(systemContact))
	// _, b4 := utils.ShellExec(snmpGetPrefix + rmPostFix(systemContact))
	// fmt.Println("set command is", snmpSetPrefix+rmPostFix(systemContact))
	// cmd := "snmpwalk -v 3 -u walter -l authPriv -a MD5 -A 123456789 -x DES -X 123456789 192.168.15.10 1.3.6.1.4.1.37072.302.2.3.2.1.1.1.1"
	// _, setResult := utils.ShellExec(cmd)
	// fmt.Println("result is " + setResult)
	taskEntry := task.GetTaskEntry()
	for _, val := range taskEntry {
		val.Exec()
	}
	task.PrintStats()
}
