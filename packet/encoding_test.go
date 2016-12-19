package packet

import "testing"

func TestEncodingUInt16(t *testing.T) {
	var maxUInt16 uint16 = 65535
	resultByte := encodeUInt16(maxUInt16)
	if (resultByte[0] != 0xFF) || (resultByte[1] != 0xFF) {
		t.Error("encoding 65535 returned: ", resultByte[0], resultByte[1])
	}

	resultUInt16 := decodeUInt16(resultByte)
	if resultUInt16 != 65535 {
		t.Error("decoding 0xFFFF returned: ", resultUInt16)
	}

	var minUInt16 uint16 = 0
	resultByte = encodeUInt16(minUInt16)
	if (resultByte[0] != 0x00) || (resultByte[1] != 0x00) {
		t.Error("encoding 0 returned: ", resultByte[0], resultByte[1])
	}

	resultUInt16 = decodeUInt16(resultByte)
	if resultUInt16 != 0 {
		t.Error("decoding 0x0000 returned: ", resultUInt16)
	}
}
