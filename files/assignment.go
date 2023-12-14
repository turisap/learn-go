package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("_info.txt")

	if err != nil {
		log.Fatal(err)
	} else {

		bw, err := f.Write([]byte("The Go gopher is an iconic mascot!"))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Bytes written", bw)
	}

	defer f.Close()
}
