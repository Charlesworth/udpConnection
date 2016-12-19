package packet

import "hash/crc32"

// Packet type contains a MsgType, SequenceNumber and the Data. You can get each
// element and do a crc32 integrity check on the packet.
type Packet struct {
	Bytes []byte
}

// GetSequenceNo returns the packet sequence number
func (pk *Packet) GetSequenceNo() uint16 {
	return decodeUInt16(pk.Bytes[1:3])
}

// GetMsgType returns the packet MsgType
func (pk *Packet) GetMsgType() byte {
	return pk.Bytes[0]
}

// GetData return the packet data
func (pk *Packet) GetData() []byte {
	return pk.Bytes[3 : len(pk.Bytes)-4]
}

// Get returns the packet MsgType, sequence number and data
func (pk *Packet) Get() (msgType byte, seqNo uint16, data []byte) {
	return pk.Bytes[0], decodeUInt16(pk.Bytes[1:3]), pk.Bytes[3 : len(pk.Bytes)-4]
}

// CheckIntegrity returns true if the sent hash and newly calculated hash match correctly
func (pk *Packet) CheckIntegrity() bool {
	hash := crc32.ChecksumIEEE(pk.Bytes[:len(pk.Bytes)-4])
	return hash == decodeUInt32(pk.Bytes[len(pk.Bytes)-4:])
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

func Decode(bytes []byte) *Packet {
	return &Packet{bytes}
}
