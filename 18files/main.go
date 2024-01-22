package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Welcome to files")
	content := "This needs to go in a file"

	file,err := os.Create("myFile.txt")

	if err != nil {
		fmt.Println("Error creating file")
		return
	}

	length,err := io.WriteString(file,content)

	if err != nil {
		fmt.Println("Error writing to file")
		return
	}
	fmt.Println("Length is ",length)
	readFiles("myFile.txt")
	defer file.Close()
}


func readFiles(filname string){
	databyte,err := ioutil.ReadFile(filname)

	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside file is\n",string(databyte))
}