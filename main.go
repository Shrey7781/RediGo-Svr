package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main(){
	ln,err := net.Listen("tcp", "localhost:6379")
	if err!=nil{
		log.Fatal("cannot listen on port 6379")
	}
	defer ln.Close()

	 conn,err := ln.Accept()
	 if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	 }
	 defer conn.Close()

	 buf:= make([]byte, 1024)
	 conn.Read(buf)

	fmt.Println(string(buf))

	conn.Write([]byte("+OK\r\n"))

	



}