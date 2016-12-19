package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"udpConnection/packet"
)

func main() {
	udpOps(":8081")
}

func udpOps(port string) {
	serverAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatalln(err)
	}

	pc, err := net.ListenPacket("udp", ":8091")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')

	packet := packet.New(packet.MsgData, uint16(1), []byte(text))
	pc.WriteTo(packet.Bytes, serverAddr)

	//simple read
	buffer := make([]byte, 1024)
	pc.ReadFrom(buffer)
	fmt.Println(string(buffer))

}

func getUDPPort() string {
	resp, err := http.Get("http://localhost:8080/udp")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
