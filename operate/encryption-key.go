package operate

func byteArray32(key []byte) [32]byte {
	if len(key) > 32 {
		key = key[:32]
	}

	var bytes [32]byte

	for i := 0; i < len(key); i++ {
		bytes[i] = key[i]
	}

	return bytes
}
