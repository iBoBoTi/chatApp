package main

import (
	"bufio"
	"fmt"
	"log"

	//"log"
	"net"
	"os"
	"strings"
)


func checkErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", ":8080")
	checkErr(err)

	fmt.Println("Please, enter your name:")
	inputReader := bufio.NewReader(os.Stdin)
	clientName, _ := inputReader.ReadString('\n')
	clientName = strings.Trim(clientName, "\r\n")

	fmt.Printf("Hi %s\n ", clientName)
	fmt.Println("----------------------------------------------------")
	fmt.Println("To exit chat, type \"exit\" ")
	fmt.Println("----------------------------------------------------")

	go read(conn)
	write(conn, clientName)


}

func read(conn net.Conn){
	for{
		reader := bufio.NewReader(conn)
		msg,_ := reader.ReadString('\n')

		fmt.Println(msg)
		fmt.Println("----------------------------------------------------")
	}
}


func write(conn net.Conn, clientName string){
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")

		if trimmedInput == "exit" {
			return
		}
		conn.Write([]byte(fmt.Sprintf("%s : %s\n", clientName, trimmedInput)))

	}
}