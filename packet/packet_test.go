package packet

import (
	"bytes"
	"hash/crc32"
	"testing"
)

func TestPacket_New(t *testing.T) {
	testData := []byte("testData")
	testPacket := New(MsgData, uint16(1), testData)

	if testPacket.Bytes[0] != MsgData {
		t.Error("New packet encoding error, encoded the incorrect msgType")
	}

	if !(decodeUInt16(testPacket.Bytes[1:3]) == uint16(1)) {
		t.Error("New packet encoding error, encoded the incorrect sequence number")
	}

	if !bytes.Equal(testPacket.Bytes[3:len(testPacket.Bytes)-4], testData) {
		t.Log("should be:", testData)
		t.Log("is:", testPacket.Bytes[3:len(testPacket.Bytes)-4])
		t.Error("New packet encoding error, encoded the incorrect data")
	}

	if !(crc32.ChecksumIEEE(testPacket.Bytes[:len(testPacket.Bytes)-4]) == decodeUInt32(testPacket.Bytes[len(testPacket.Bytes)-4:])) {
		t.Error("New packet encoding error, encoded crc32 checksum does not match recalculated checksum")
	}
}

func TestPacket_GetSequence(t *testing.T) {
	testSequence := uint16(1234)
	testPacket := New(MsgData, testSequence, []byte("testData"))
	if testPacket.GetSequenceNo() != testSequence {
		t.Error("GetSequence is not passing back the correct sequence number")
	}
}

func TestPacket_GetMsgType(t *testing.T) {
	testMsg := MsgFin
	testPacket := New(testMsg, uint16(1324), []byte("testData"))
	if testPacket.GetMsgType() != testMsg {
		t.Error("GetMsgType is not passing back the correct MsgType")
	}
}

func TestPacket_GetData(t *testing.T) {
	testData := []byte("testData")
	testPacket := New(MsgData, uint16(1324), []byte("testData"))
	if !bytes.Equal(testPacket.GetData(), testData) {
		t.Error("GetData is not passing back the correct Data")
	}
}

func TestPacket_Get(t *testing.T) {
	testMsg := MsgInit
	testSequence := uint16(1234)
	testData := []byte("testData")
	testPacket := New(testMsg, testSequence, testData)

	resultMsg, resultSequence, resultData := testPacket.Get()

	if resultSequence != testSequence {
		t.Error("Get is not passing back the correct sequence number")
	}

	if resultMsg != testMsg {
		t.Error("Get is not passing back the correct MsgType")
	}

	if !bytes.Equal(resultData, testData) {
		t.Error("Get is not passing back the correct Data")
	}
}

func TestPacket_CheckIntegrity(t *testing.T) {
	testPacket := New(MsgData, uint16(1324), []byte("testData"))
	if !testPacket.CheckIntegrity() {
		t.Error("CheckIntegrity failed on a non-corrupt packet")
	}

	if testPacket.Bytes[0] != 0xF {
		testPacket.Bytes[0] = 0xF
	} else {
		testPacket.Bytes[0] = 0x1
	}

	if testPacket.CheckIntegrity() {
		t.Error("CheckIntegrity passed on a corrupt packet")
	}
}
