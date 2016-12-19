package main

import (
	"log"
	"net"

	"github.com/charlesworth/udpConnection/packet"
)

var connectionMgr = newConnectionManager()
var discoveryPort = ":8080"
var udpPort = ":8081"
var MTU = 1000

func main() {
	socket := initSocket(udpPort)
	defer socket.Close()
	log.Println("UDP Server started on port ", udpPort, ", listening for connections")
	listenForUDP(socket)
	//startPortDiscoveryServer()
}

func initSocket(port string) *net.UDPConn {
	localAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal("error resolving local IP port ", udpPort, ":", err)
	}

	socket, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Fatal("Error starting listenUDP server:", err)
	}

	return socket
}

func listenForUDP(socket *net.UDPConn) {
	for {
		buffer := make([]byte, MTU)
		bytes, ip, err := socket.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal("unable to read UDP packet payload due to error:", err)
		}
		if bytes == 0 {
			log.Fatal("unable to read UDP packet payload due to error:", err)
		}

		log.Println("Packet from IP[", ip.String(), "], ", bytes, " recieved")
		recPkt := packet.Decode(buffer)
		if !recPkt.CheckIntegrity() {
			log.Println("Corrupt Packet, discarding")
		}

		data := recPkt.GetData()
		log.Println(string(data))
		log.Println(data)

		sendPtk := packet.New(packet.MsgData, uint16(2), []byte("hi"))

		// send a plain UDP message back
		remoteAddr, err := net.ResolveUDPAddr("udp", ip.String())
		socket.WriteToUDP(sendPtk.Bytes, remoteAddr)
	}
}
