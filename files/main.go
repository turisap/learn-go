package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// @TODO you will also need to try io and io-utils packages
func main() {
	//basicOperations()
	//cwd()
	//writingToFile()
	//writeBufIO()
	//readFile()
	//readByLine()
	readStdio()
}

func basicOperations() {

	//var newFile *os.File
	//fmt.Printf("%T\n", newFile)

	//nF, err := os.Create("o.txt")
	//
	//if err != nil {
	//	fmt.Println("There was an error", err)
	//	//os.Exit(1)
	//	log.Fatal(err)
	//} else {
	//	fmt.Printf("File created %T %v\n", nF, nF)
	//}
	//
	//nF.Close()

	textFile, err := os.OpenFile("_o.txt", os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("There was an error", err)
		//os.Exit(1)
		log.Fatal(err)
	} else {
		fmt.Printf("File opened  /created %T %v\n", textFile, textFile)
	}

	fmt.Println("_o.txt", textFile)

	//textFile.Close()

	fileInfo, err := os.Stat("_o.txt")

	if err != nil {
		fmt.Println("There was an error", err)
		//os.Exit(1)
		log.Fatal(err)
	} else {
		fmt.Println("File info", fileInfo.Name(), fileInfo.Size(), fileInfo.ModTime(), fileInfo.IsDir())
		fmt.Printf("Permissions: %v\n", fileInfo.Mode())
	}

	f, err := os.OpenFile("_b.txt", os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist", err)
			//log.Fatal("File does not exist", err)
		}
	}

	fmt.Println(f)

	oldPath := "_o.txt"
	newPath := "_aaa.txt"

	err = os.Rename(oldPath, newPath)

	if err != nil {
		if os.IsNotExist(err) {
			//fmt.Println("File does not exist", err)
			//log.Fatal("Could not rename", err)
		}
	}

	//err = os.Remove(newPath)
	//
	//if err != nil {
	//	if os.IsNotExist(err) {
	//		//fmt.Println("File does not exist", err)
	//		log.Fatal("Could not delete", err)
	//	}
	//}
}

func cwd() {
	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatal("Error getting the path", err)
	}

	fmt.Println(currentDir)

}

func writingToFile() {
	path, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting path")
		log.Fatal(err)
	}

	name := path + "/_writable.txt"
	f, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file")
		log.Fatal(err)
	}

	// idiomatic way to close a file
	// defers closing until the
	// surrounding func returns
	defer f.Close()

	// convert a string to bytes
	bytesSlice := []byte("I love golang")

	bytesWritten, err := f.Write(bytesSlice)

	if err != nil {
		fmt.Println("Error writing to file")
		log.Fatal(err)
	}

	fmt.Println("Bytes written", bytesWritten)

	bytesSlice = []byte("Go programming is cool")

	err = os.WriteFile(path+"/_ol.txt", bytesSlice, 0644)

	if err != nil {
		fmt.Println("Error writing the file")
		log.Fatal(err)
	} else {
		fmt.Println("File written")
	}

}

func writeBufIO() {
	path, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting path")
		log.Fatal(err)
	}

	name := path + "/_bufio.txt"
	f, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file")
		log.Fatal(err)
	}

	defer f.Close()

	bufferedWriter := bufio.NewWriter(f)

	bs := []byte{10, 133, 212, 33}

	bytesWritten, err := bufferedWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bytes written", bytesWritten)
	fmt.Println("Bytes available", bufferedWriter.Available())

	bytesWritten, err = bufferedWriter.Write([]byte("\n Just a random stringg"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Buffered bytes %d\n", bufferedWriter.Buffered())

	// @COOL you can write to file after you finished all operations on buffer
	bufferedWriter.Flush()
}

func readFile() {
	currWD, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	path := currWD + "/main.go"

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bs, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("here")
		log.Fatal(err)
	}

	// convert a bite slice into a unicode string
	fmt.Println("Bytes slice", string(bs))
	fmt.Println("Bytes read", len(bs))

	// another way to read a file
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", data)
}

func readByLine() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	path := wd + "/main.go"

	f, err := os.OpenFile(path, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// read line-byline
	scanner := bufio.NewScanner(f)

	ok := scanner.Scan()

	if !ok {
		err := scanner.Err()

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("EOF reached")
		}
	}

	fmt.Println("First line found", scanner.Text())

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readStdio() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		text := scanner.Text()
		bytes := scanner.Bytes()

		fmt.Println("Text:", text)
		fmt.Println("Bytes:", bytes)

		if text == "exit" {
			fmt.Println("Exiting scanning")
			break
		}
	}

	fmt.Println("By...")

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
