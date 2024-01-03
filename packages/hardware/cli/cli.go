package cli

import (
	"bufio"
	"fmt"
	"hardware/device"
	"os"
	"strconv"
	"sync"
)

func ReadInCommand(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func PrintHelp() {
	fmt.Println("help Info")
	fmt.Println("new monitor(nm)")
	fmt.Println("new bell(nb)")
	fmt.Println("list devices(l)")
	fmt.Println("run monitor(rm)")
	fmt.Println("run bell(rb)")
	fmt.Println("exit(q)")
	fmt.Println("help(?)")
}

var Monitors sync.Map
var Bells sync.Map

func Run() {
	for {
		command := ReadInCommand(">>")
		if command == "q" {
			break
		}
		if command == "?" {
			PrintHelp()
		}
		if command == "nm" {
			bed_id := ReadInCommand("bed id: ")
			m := &device.Monitor{
				BedID: func() (i uint64) {
					i, _ = strconv.ParseUint(bed_id, 10, 64)
					return
				}(),
			}
			Monitors.Store(bed_id, m)
		}
		if command == "l" {
			Monitors.Range(func(key, value interface{}) bool {
				fmt.Println(key, value)
				return true
			})
			Bells.Range(func(key, value interface{}) bool {
				fmt.Println(key, value)
				return true
			})
		}
		if command == "rm" {
			bed_id := ReadInCommand("bed id: ")
			m, ok := Monitors.Load(bed_id)
			if ok {
				go m.(*device.Monitor).Run()
			} else {
				fmt.Println("device not found")
			}
		}
		if command == "nb" {
			department_id := ReadInCommand("department id: ")
			b := &device.Bell{
				DepartmentID: func() (i uint64) {
					i, _ = strconv.ParseUint(department_id, 10, 64)
					return
				}(),
			}
			Bells.Store(department_id, b)
		}
		if command == "rb" {
			department_id := ReadInCommand("department id: ")
			b, ok := Bells.Load(department_id)
			if ok {
				go b.(*device.Bell).Run()
			} else {
				fmt.Println("device not found")
			}
		}
		if command == "sm" {
			bed_id := ReadInCommand("bed id: ")
			m, ok := Monitors.Load(bed_id)
			if ok {
				m.(*device.Monitor).Stop()
			} else {
				fmt.Println("device not found")
			}
		}
		if command == "sb" {
			department_id := ReadInCommand("department id: ")
			b, ok := Bells.Load(department_id)
			if ok {
				b.(*device.Bell).Stop()
			} else {
				fmt.Println("device not found")
			}
		}
		if command == "c" {
			bed_id := ReadInCommand("bed id: ")
			m, ok := Monitors.Load(bed_id)
			if ok {
				m.(*device.Monitor).Call()
			} else {
				fmt.Println("device not found")
			}
		}
	}
}
