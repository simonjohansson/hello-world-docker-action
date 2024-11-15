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
		f, err := os.Open(os.Getenv("GITHUB_OUTPUT"))
		if err != nil {
			panic(err)
		}

		f.WriteString(fmt.Sprintf("time=%s", time.Now().String()))
	})
}
