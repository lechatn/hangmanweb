package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr := "127.0.0.1:8080"
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Erreur lors de la connexion au serveur:", err)
		return
	}
	defer conn.Close()

	serverIP, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	fmt.Println("Adresse IP du serveur:", serverIP)
	fmt.Println("Port du serveur:", serverAddr)
}
