package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
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
		fmt.Println("Random must Integer Error:%s", err)
		return
	}

	if max := len(random); max > 4 {
		fmt.Println("Random lenght must less than 4")
		return
	}

	random = formatNumber(random)

	hexString := blockHash + random
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Errorf("Hex String To Bytes Error: %s", err)
	}

	shaBytes := blockGuessSha256(hexBytes)
	xorbytes := hfBytes(shaBytes)
	xorbytes = hfBytes(xorbytes)
	data := binary.BigEndian.Uint64(xorbytes)

	resultNumber := formatNumber(fmt.Sprintf("%d", data%1000))
	fmt.Println(resultNumber)
}

func blockGuessSha256(plaintext []byte) []byte {
	hash := sha256.New()
	hash.Write(plaintext)
	md := hash.Sum(nil)
	hash2 := sha256.New()
	hash2.Write(md)
	res := hash2.Sum(nil)
	return res
}

func hfBytes(data []byte) []byte {
	lenght := len(data)
	xorbytes := make([]byte, lenght/2)

	xORBytes(xorbytes, data[0:lenght/2], data[lenght/2:lenght])
	return xorbytes
}

func xORBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return n
}

func formatNumber(number string) string {
	maxN := len(number)
	for i := 0; i <= 3-maxN; i++ {
		number = "0" + number
	}
	return number
}
