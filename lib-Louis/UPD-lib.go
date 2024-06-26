package liblouis

type UDPPackage struct {
	sourcePort uint16

	destinationPort uint16

	//length(bytes) of UDP header + UDP data, minimum 8 bytes(UDP header)
	length uint16

	checksum uint16

	data []uint8
}

func UnpackUPD(rawUDP []uint8) UDPPackage {

	//rawUDP
	//[16bits - source port	][16bits - destination port	]
	//[16bits - length		][16bits - data				]
	//[			length bytes - data						]

	var header UDPPackage

	header.sourcePort = uint16(rawUDP[0] + rawUDP[1])

	header.destinationPort = uint16(rawUDP[2] + rawUDP[3])

	header.length = uint16(rawUDP[4] + rawUDP[5])

	header.checksum = uint16(rawUDP[6] + rawUDP[7])

	header.data = make([]uint8, header.length)

	//fill data of size header.length with chuncks of 1 byte
	for i := 0; i < len(header.data); i++ {
		header.data[i] = uint8(rawUDP[8+i])
	}

	return header
}

func PackUDP()
