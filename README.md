Virtual UDP connection, for real time data transfer.

# Target Interface:
I want to have interface parity with the net package's implementation of a tcp connection, supplying:
  - Listener Type: used by the server to receive connection requests
  - Connection Type: used by the server and client to send datagrams over a accepted connection
  - Dial: to connect to a server and supply a connection
  - Listen: to create a server, and await incoming connection requests
  - Accept: used to accept (or not) an incomming connection request from a client and return a connection

# Aimed features:
  - crc32 hash on every datagram to check that the packet hasn't been corrupted on transmittal
  - full duplex connection initialization and destruction
  - variable "keep alive messages per second" limit , connections will close if messages sent per second drops below this number, with both sides able to send empty keep alive messages
  - message sequence counter to measure RTT and congestion
  - simple network congestion controls
