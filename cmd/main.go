package main

import (
	"fmt"
	"time"

	"github.com/ecVictor7/go-htmx-websockets-example/internal/hardware"
)

func main() {
	fmt.Println("Starting my System monitor...")
	go func() {
		for {
			systemSection, err := hardware.GetSystemSection()
			if err != nil {
				fmt.Println(err)
			}

			diskSection, err := hardware.GetDiskSection()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(systemSection)
			fmt.Println(diskSection)
			fmt.Println(cpuSection)

			time.Sleep(3 * time.Second)
		}
	}()
	//sleep the main threead
	time.Sleep(5 * time.Minute)
}
