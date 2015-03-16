package main

import (
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"strconv"
)

func main() {
	crc32Hash := crc32.NewIEEE()
	crc32Hash.Write([]byte("test"))
	crc32v := crc32Hash.Sum32()

	adlerHash := adler32.New()
	adlerHash.Write([]byte("test"))
	adler32v := adlerHash.Sum32()

	crc64Table := crc64.MakeTable(crc64.ECMA)

	checksum64 := crc64.Checksum([]byte("test"), crc64Table)

	fmt.Printf("crc32 : %x\n", crc32v)
	fmt.Printf("alder32 : %x\n", adler32v)
	fmt.Printf("crc64 checksum : %x\n", checksum64)
	fmt.Printf("crc64 checksum : %v\n", string(checksum64))
	fmt.Printf("%s\n", strconv.FormatUint(checksum64, 10))
}
