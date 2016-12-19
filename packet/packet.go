package packet

import "hash/crc32"

// Packet type contains a MsgType, SequenceNumber and the Data. You can get each
// element and do a crc32 integrity check on the packet.
type Packet struct {
	bytes []byte
}

// GetSequenceNo returns the packet sequence number
func (pk *Packet) GetSequenceNo() uint16 {
	return decodeUInt16(pk.bytes[1:3])
}

// GetMsgType returns the packet MsgType
func (pk *Packet) GetMsgType() byte {
	return pk.bytes[0]
}

// GetData return the packet data
func (pk *Packet) GetData() []byte {
	return pk.bytes[3 : len(pk.bytes)-4]
}

// Get returns the packet MsgType, sequence number and data
func (pk *Packet) Get() (msgType byte, seqNo uint16, data []byte) {
	return pk.bytes[0], decodeUInt16(pk.bytes[1:3]), pk.bytes[3 : len(pk.bytes)-4]
}

// CheckIntegrity returns true if the sent hash and newly calculated hash match correctly
func (pk *Packet) CheckIntegrity() bool {
	hash := crc32.ChecksumIEEE(pk.bytes[:len(pk.bytes)-4])
	return hash == decodeUInt32(pk.bytes[len(pk.bytes)-4:])
}

// New creates a newly encoded packet when passed a MsgType, sequence number and data
func New(msgType byte, seqNo uint16, data []byte) Packet {
	bt := []byte{msgType}
	bt = append(bt, encodeUInt16(seqNo)...)
	bt = append(bt, data...)

	hash := crc32.ChecksumIEEE(bt)
	bt = append(bt, encodeUInt32(hash)...)

	return Packet{bt}
}
