package main

import (
	"log"
	"net"
)

var connectionMgr = newConnectionManager()
var discoveryPort = ":8080"
var udpPort = ":8081"
var MTU = 1500

func main() {
	socket := initSocket(udpPort)
	go listenForUDP(socket)
	log.Println("UDP Server started, listening for connections")
	startPortDiscoveryServer()
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
	defer socket.Close()

	//for each message
	//check crc32, if corrupt, throw away
	//decode
	//if init MsgType->pass
	//if not known IP, throw away
	//switch on msgType
	for {
		buffer := make([]byte, MTU)
		bytes, ip, err := socket.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal("unable to read UDP packet payload")
		}
		//
		// log.Println("UDP packet recieved [IP:", ip, "] [bytes:", bytes, "]")
		// uncorrupt, packet := crc32CheckTruncatePacket(buffer)
		// if !uncorrupt {
		// 	continue
		// }
		//
		// //packetSeqNo, msgType, data := decodePacket(packet)
		// _, msgType, data := decodePacket(packet)
		// if !connectionMgr.containsIP(ip.String()) {
		// 	if msgType == msgInit {
		// 		//pass to init here
		// 	}
		// 	continue
		// }
		//
		// //switch on msgType
		// switch msgType {
		// case msgData:
		// 	log.Println(data)
		// 	//do something
		// default:
		// 	//error
		// }

		// send a message back
		remoteAddr, err := net.ResolveUDPAddr("udp", ip.String())
		socket.WriteToUDP([]byte("hello from reciever"), remoteAddr)
	}
}
