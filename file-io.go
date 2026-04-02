package main

import(
	"os"
	"fmt"
	"log"
)

func main() {
	data := "Welcome to GO Programing Language"
	err := os.WriteFile("./NewFolder/text.txt",[]byte(data),0644)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Done Writing")

	readData, err := os.ReadFile("./NewFolder/text.txt")
	if err!= nil{
		log.Fatal(err)
	} 

	fmt.Printf("%s",readData)

}