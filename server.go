package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func checkError(err error){
	if err != nil{
		log.Fatal(err)
	}
}

var (
	allClients = make(map[net.Conn]bool)
	newClient = make(chan net.Conn)
	exitClient = make(chan net.Conn)
)


func main(){
	fmt.Println("Chat Server has started...")

	// create chat server
	listener, err := net.Listen("tcp", "localhost:8080")
	checkError(err)

	go func(){
		for{
			conn, err := listener.Accept()
			checkError(err)

			allClients[conn] = true
			newClient <- conn
		}
	}()

	for {
		select {
			case conn := <-newClient:
				go broadcaster(conn)

			case conn := <- exitClient:
				for item := range newClient{
					if item == conn{
						break
					}
				}

				delete(allClients,conn)
			}
		}


}


func broadcaster(conn net.Conn){
	for{
		reader := bufio.NewReader(conn)
		message, err:= reader.ReadString('\n')
		if err != nil{
			break
		}
		for item := range allClients{
			if item != conn{
				item.Write([]byte(message))
			}
		}
	}

	exitClient <- conn
}


