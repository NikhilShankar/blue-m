package main 

import(
	"crypto/sha1"
	"hash"
	"fmt"
	"encoding/binary"
)

type Bluem struct {
	blueBitArray  []bool
	blueBitArraySize  uint 
	blueHashSize  uint
	blueHash  hash.Hash
}

type bluem interface {
	Add(b bluem_input)
	Check(b bluem_input) bool
}

type bluem_input []byte

func New(m uint, k uint) *Bluem { 

	return &Bluem {
		blueBitArray : make([]bool,m),
		blueBitArraySize : m,
		blueHashSize : k,	
		blueHash : sha1.New(),
	}

}

func (b *Bluem) Add(in bluem_input) {
	fmt.Println("Add ", in)
	for i:= uint(0); i<b.blueHashSize; i++ {
		setBit(b.blueBitArray, getHash(b.blueHash, in, []byte(string(i)), b.blueBitArraySize))
	}
}

func (b *Bluem) Check(in bluem_input) bool {
	fmt.Println("Check ", in)
	for i:= uint(0); i<b.blueHashSize; i++ {
		pos := getHash(b.blueHash, in, []byte(string(i)), b.blueBitArraySize)
		if !b.blueBitArray[pos] {
			return false
		} 
	}
	return true
}

func getHash(blueHash hash.Hash, data []byte, addendum []byte, m uint) uint {
	blueHash = sha1.New()
	blueHash.Write(data)
	resultInt := binary.BigEndian.Uint64(blueHash.Sum(addendum))
	fmt.Println("HASH ", uint(resultInt) % m)
	return uint(resultInt) % m
}

func setBit(bitArray []bool, pos uint) {
	bitArray[pos] = true
}
