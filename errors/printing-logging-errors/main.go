package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt") // for f there will be no error
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("error occured", err) // this will write it to a file
		log.Fatalf(err.Error())           // this needs to be err.Error() as it expects a string
		log.Panicf(err.Error())           // it will exit and will never get to panic
		panic(err)                        // again it will not get there as it has already exited
	}
	defer f2.Close()

	fmt.Println("check the file log.txt in the directory")
}
