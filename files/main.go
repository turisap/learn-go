package main

import (
	"fmt"
	"log"
	"os"
)

// @TODO you will also need to try io and io-utils packages
func main() {
	//basicOperations()
	//cwd()
	//writingToFile()
	writeBufIO()
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

}
