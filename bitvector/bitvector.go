package bitvector

import (
	"fmt"
	"os"

	"github.com/lwhite17/CMSC701_HW2/utils"
)

type BitVector struct {
	data   []byte
	length uint64
}

func NewBitVector(data []byte, length uint64) *BitVector {
	// Return a pointer to the created bitvector
	return &BitVector{
		data:   data,
		length: length}
}

func (b *BitVector) Bytes() []byte {
	// Return the bitvector's data
	return b.data
}

func (b *BitVector) Length() uint64 {
	// Return the length of the bitvector

	return b.length
}

func (b *BitVector) Get(i int) byte {
	// Get byte at index from vector

	return b.data[i]
}

func (b *BitVector) GetSlice(i int, j int) []byte {
	// Get bytes [i,j) from bitvector

	return b.data[i:j]
}

func (b BitVector) Save(filename string) {
	// Save bitvector to file

	err := os.WriteFile(filename, b.Bytes(), 0777)
	utils.CheckError(err)

}

func Load(filename string) *BitVector {
	// Load bitvector from file

	var bb []byte
	bb, err := os.ReadFile(filename)
	utils.CheckError(err)

	b := NewBitVector(bb, uint64(len(bb)))

	return b
}

func (b *BitVector) Set(ind uint64, val byte) *BitVector {
	// Sets the bitvector at index ind to value val
	// if val is 0 or 1

	if val != 0 && val != 1 {
		fmt.Println("Error: please input 0 or 1 as value")
		return b
	}
	b.data[ind] = val
	return b
}
