package main

import (
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"

	"golang.org/x/crypto/pbkdf2"
)

var (
	blockHash string
	random    string
	help      bool
)

func init() {
	flag.StringVar(&blockHash, "blockhash", "000000000000000001ed86134bcee0ad3f879f88e4cc3b27138d5c738de04fa9", "BCH block hash")
	flag.StringVar(&random, "random", "0123", "random.org random number")
	flag.BoolVar(&help, "h", false, "Print Help Info")
}

func main() {
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	if _, err := strconv.Atoi(random); err != nil {
		fmt.Printf("Random must Integer Error:%s", err)
		return
	}

	if max := len(random); max > 4 {
		fmt.Println("Random lenght must less than 4")
		return
	}

	random = formatNumber(random)

	blockhashBytes, err := hex.DecodeString(blockHash)
	if err != nil {
		fmt.Printf("Block Hash String To Bytes Error: %s", err)
	}

	randomBytes, err := hex.DecodeString(random)
	if err != nil {
		fmt.Printf("Random String To Bytes Error: %s", err)
	}

	xorbytes := kdfBytes(removeZeroByte(blockhashBytes), randomBytes, 2)
	data := binary.BigEndian.Uint16(xorbytes)

	resultNumber := formatNumber(fmt.Sprintf("%d", data%1000))
	fmt.Println(resultNumber)
}

func kdfBytes(blockhash, random []byte, outlen int) []byte {
	return pbkdf2.Key(blockhash, random, 1<<14, outlen, sha512.New)
}

func removeZeroByte(blockhash []byte) []byte {
	index := 0
	for i, b := range blockhash {
		if b != 0x00 {
			index = i
			break
		}
	}
	return blockhash[index:len(blockhash)]
}

func formatNumber(number string) string {
	maxN := len(number)
	for i := 0; i <= 3-maxN; i++ {
		number = "0" + number
	}
	return number
}
