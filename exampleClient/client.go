package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/charlesworth/udpConnection/packet"
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

	// simple write
	text := getStdIn()
	sendPkt := packet.New(packet.MsgData, uint16(1), []byte(text))
	pc.WriteTo(sendPkt.Bytes, serverAddr)

	//simple read
	buffer := make([]byte, 1024)
	pc.ReadFrom(buffer)
	recPkt := packet.Decode(buffer)
	fmt.Println(string(recPkt.GetData()))

}

func getStdIn() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')
	return text
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
