package packet

import "encoding/binary"

func decodeUInt16(inputByte []byte) uint16 {
	return binary.LittleEndian.Uint16(inputByte)
}

func encodeUInt16(inputUInt16 uint16) []byte {
	resultByte := make([]byte, 2)
	binary.LittleEndian.PutUint16(resultByte, inputUInt16)
	return resultByte
}

func decodeUInt32(inputByte []byte) uint32 {
	return binary.LittleEndian.Uint32(inputByte)
}

func encodeUInt32(inputUInt32 uint32) []byte {
	resultByte := make([]byte, 4)
	binary.LittleEndian.PutUint32(resultByte, inputUInt32)
	return resultByte
}
