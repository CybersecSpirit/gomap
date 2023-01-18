package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func checkTCP(ip string, port string) {
	// Si le port est égal à "*", on scanne tous les ports TCP
	if port == "*" {
		startPort := 1
		endPort := 65535
		for i := startPort; i <= endPort; i++ {
			address := fmt.Sprintf("%s:%d", ip, i)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				continue
			} else {
				fmt.Printf("Port %d: open\n", i)
				conn.Close()
			}
		}
		// Si le port contient un "-", on considère qu'il s'agit d'une plage de ports
	} else if strings.Contains(port, "-") {
		portRange := strings.Split(port, "-")
		startPort, _ := strconv.Atoi(portRange[0])
		endPort, _ := strconv.Atoi(portRange[1])
		if startPort < 1 || startPort > 65535 || endPort < 1 || endPort > 65535 {
			fmt.Println("Invalid port range")
			return
		}
		for i := startPort; i <= endPort; i++ {
			address := fmt.Sprintf("%s:%d", ip, i)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				continue
			} else {
				fmt.Printf("Port %d: open\n", i)
				conn.Close()
			}
		}
		// Si le port contient une ",", on considère qu'il s'agit d'une sélection multiple de ports
	} else if strings.Contains(port, ",") {
		individualPorts := strings.Split(port, ",")
		for _, p := range individualPorts {
			portInt, _ := strconv.Atoi(p)
			if portInt < 1 || portInt > 65535 {
				fmt.Printf("Port %d: invalid\n", portInt)
				continue
			}
			address := fmt.Sprintf("%s:%d", ip, portInt)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("Port %d: closed\n", portInt)
			} else {
				fmt.Printf("Port %d: open\n", portInt)
				conn.Close()
			}
		}
		// Sinon on considère qu'il s'agit d'un seul port
	} else {
		portInt, _ := strconv.Atoi(port)
		if portInt < 1 || portInt > 65535 {
			fmt.Printf("Port %d: invalid\n", portInt)
			return
		}
		address := fmt.Sprintf("%s:%d", ip, portInt)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("Port %d: closed\n", portInt)
		} else {
			fmt.Printf("Port %d: open\n", portInt)
			conn.Close()
		}
	}
	// vérification de la validité de l'adresse IP
	if net.ParseIP(ip) == nil {
		fmt.Println("Invalid IP address")
		return
	}
}

func main() {
	// utilisation des arguments en ligne de commande pour récupérer l'adresse IP et le port
	ip := os.Args[1]
	port := os.Args[2]
	checkTCP(ip, port)
}
