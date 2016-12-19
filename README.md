Virtual UDP connection, for real time data transfer.

Aimed features:
  - crc32 hash on every datagram to check that the packet hasn't been corrupted on transmittal
  - full duplex connection initialization and destruction
  - variable keep alive messages per second limit , connections will close if messages sent per second drops below this number, with both sides able to send empty keep alive when there's no data available
  - message sequence counter to measure RTT and congestion
