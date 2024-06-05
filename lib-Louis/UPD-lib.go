package liblouis

type UDPHeader struct {
	sourcePort      [2]byte
	destinationPort [2]byte
	length          [2]byte
	checksum        [2]byte
	data            []byte
}

func UnpackUPD(bytes []byte) UDPHeader {

	var header UDPHeader

	header.sourcePort[0] = bytes[0]
	header.sourcePort[1] = bytes[1]

	header.destinationPort[0] = bytes[2]
	header.destinationPort[1] = bytes[3]

	header.length[0] = bytes[4]
	header.length[1] = bytes[5]

	header.checksum[0] = bytes[6]
	header.checksum[1] = bytes[7]

	var dataLength int = int(header.length[0]) + int(header.length[1])
	header.data = make([]byte, dataLength)

	return header
}
