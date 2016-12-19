package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

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

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Text to send: ")
	// text, _ := reader.ReadString('\n')
	text := "hi mr man"

	pkt := packet.New(packet.MsgData, uint16(1), []byte(text))
	// pkt := packet.New(packet.MsgData, uint16(1), []byte(text))
	log.Println("integrity:", pkt.CheckIntegrity())
	log.Println(pkt.Bytes)
	pc.WriteTo(pkt.Bytes, serverAddr)

	//simple read
	buffer := make([]byte, 1024)
	pc.ReadFrom(buffer)
	recPkt := packet.Decode(buffer)
	fmt.Println(string(recPkt.GetData()))

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
