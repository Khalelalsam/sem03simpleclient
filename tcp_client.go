package main

import (
	"net"
	"log"
	"os"
"github.com/Khalelalsam/is105sem03/mycrypt"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: tcp_client <message>")
	}
kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)
	log.Println("Kryptert melding: ", string(kryptertMelding))
conn, err := net.Dial("tcp", "172.17.0.2:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
_, err = conn.Write([]byte(string(kryptertMelding)))
	if err != nil {
		log.Fatal(err)
	}
buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
response := string(buf[:n])
	log.Printf("Reply from server: %s", response)
}
