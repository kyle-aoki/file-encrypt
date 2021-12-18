package operate

import (
	"file-encrypt/args"
	"file-encrypt/ex"
	"file-encrypt/mathe"
	"github.com/gtank/cryptopasta"
	"os"
)

type OperationType int

const (
	ENCRYPT = iota
	DECRYPT
)

const (
	ChunkSize = 2500
)

func Operate(ot OperationType) {
	fname := args.Poll()
	key := args.PollBytes()
	keyArray32 := byteArray32(key)

	file, err := os.Open(fname)
	ex.Check(err)
	defer file.Close()

	stat, err := file.Stat()
	ex.Check(err)

	var newFileName string
	encryptedPrefix := "encrypted_"

	if ot == ENCRYPT {
		newFileName = encryptedPrefix + fname
	} else {
		if len(fname) > len(encryptedPrefix) && encryptedPrefix == fname[:len(encryptedPrefix)] {
			newFileName = "de" + fname[2:]
		} else {
			newFileName = "decrypted_" + fname
		}
	}

	newFile, err := os.Create(newFileName)
	ex.Check(err)
	defer newFile.Close()

	fileBytes := stat.Size()
	var index int64 = 0

	for index < fileBytes {
		bytesToRead := mathe.Min(ChunkSize, fileBytes-index)

		b := make([]byte, bytesToRead)
		_, err := file.Read(b)
		ex.Check(err)

		var newBytes []byte

		if ot == ENCRYPT {
			newBytes, err = cryptopasta.Encrypt(b, &keyArray32)
			ex.Check(err)
		} else if ot == DECRYPT {
			newBytes, err = cryptopasta.Decrypt(b, &keyArray32)
			ex.Check(err)
		}

		_, err = newFile.Write(newBytes)
		ex.Check(err)

		index += bytesToRead
	}
}
