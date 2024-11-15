package main

import (
	"fmt"
	"log"
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
		f, err := os.OpenFile(os.Getenv("GITHUB_OUTPUT"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		f.WriteString(fmt.Sprintf("time=%s", time.Now().String()))
	})
}
