package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf(t.Format("2006年01月02日"))
	switch t.Weekday().String() {
	case "Sunday":
		fmt.Printf("星期天\n")
	case "Saturday":
		fmt.Printf("星期六\n")
	}
}
