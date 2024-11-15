package main

import (
	"fmt"
	"os"
	"time"
)

func group(name string, f func()) {
	fmt.Println(fmt.Sprintf("::group::%s", name))
	f()
	fmt.Println("::endgroup::")
}

func main() {
	group("os.Environ()", func() {
		for _, v := range os.Environ() {
			fmt.Println(v)
		}

	})

	group("Echo the greeting", func() {
		fmt.Println(fmt.Sprintf("Hello %s", os.Getenv("INPUT_WHO-TO-GREET")))
	})

	group("Output", func() {
		fmt.Println("Writing time to output file")
		f, err := os.OpenFile(os.Getenv("GITHUB_OUTPUT"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		fmt.Println("1")
		if err != nil {
			fmt.Println("2")
			fmt.Println(err)
			os.Exit(-1)
		}
		fmt.Println("3")
		defer f.Close()
		fmt.Println("4")
		f.WriteString(fmt.Sprintf("time=%s", time.Now().String()))
		fmt.Println("5")
	})
}
