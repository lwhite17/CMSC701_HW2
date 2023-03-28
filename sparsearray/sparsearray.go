package sparsearray

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/lwhite17/CMSC701_HW2/bitvector"
	"github.com/lwhite17/CMSC701_HW2/bvrank"
	"github.com/lwhite17/CMSC701_HW2/bvselect"
	"github.com/lwhite17/CMSC701_HW2/utils"
	"github.com/tmthrgd/go-popcount"
)

type SparseArray struct {
	bitv    *bitvector.BitVector
	rank    *bvrank.Rank
	sel     *bvselect.Select
	strings []string
}

func Create(size uint64) *SparseArray {
	sa := SparseArray{}

	// Make bitvector of given size of all 0s
	bvbytes := make([]byte, size)
	for i := 0; uint64(i) < size; i++ {
		bvbytes[i] = 0
	}

	// Save bitvector to SparseArray
	sa.bitv = bitvector.NewBitVector(bvbytes, size)
	// Initialize SparseArray's array of strings
	sa.strings = make([]string, 0, size/2)

	return &sa

}

func (sa *SparseArray) Append(elem string, pos uint64) *SparseArray {
	// If input position too large, do nothing / print error
	if pos >= sa.bitv.Length() {
		fmt.Println("Error: Append failed. Entered position is larger than given size. ")
		return sa
	}

	// Get current rank up to input position
	stringind := popcount.CountBytes(sa.bitv.Bytes()[:pos])
	// Change bitvector to 1 at input position
	sa.bitv.Set(pos, 1)

	// Update list of strings
	if stringind >= uint64(len(sa.strings)) {
		// Add to end of array
		sa.strings = append(sa.strings, elem)
	} else if stringind == 0 {
		// Add to beginning of array
		sa.strings = append([]string{elem}, sa.strings...)
	} else {
		// Add to middle of array
		sa.strings = append(sa.strings[:stringind+1], sa.strings[stringind:]...)
		sa.strings[stringind] = elem
	}

	return sa
}

func (sa *SparseArray) Finalize() *SparseArray {
	// Create Rank struct and save to SparseArray
	sa.rank = bvrank.NewRankStruct(sa.bitv)
	// sa.rank = sa.rank.PopulateRanks()

	// Create Select struct and save to SparseArray
	sa.sel = bvselect.NewSelectStruct(sa.bitv, sa.rank)

	return sa
}

func (sa *SparseArray) GetAtRank(r uint64) (*string, bool) {
	// If input rank larger than number of strings in array,
	// return nil and false
	if r >= uint64(len(sa.strings)) {
		return nil, false
	}

	// Otherwise, return pointer to string at input index in string array
	return &sa.strings[r], true
}

func (sa *SparseArray) GetAtIndex(r uint64) (*string, bool) {
	// Get bitvector value at index r
	bval := sa.bitv.Bytes()[r]

	// If bitvector value at index r is 1,
	// return pointer to string at rank1(r)
	if bval == 1 {
		rrank := sa.rank.Rank1(r)
		return &sa.strings[rrank], true
	}

	// Otherwise, return nil and false
	return nil, false
}

func (sa *SparseArray) GetIndexOf(r uint64) uint64 {
	// If input r is larger than the length of strings,
	// return max uint64 value
	if r > uint64(len(sa.strings)) {
		return math.MaxUint64
	}

	// Otherwise, return Select1(r-1), which gives the
	// index at the previous rank-increase (accounts for
	// change of index from rth value to indexing starting from 0)
	ind := sa.sel.Select1(r - 1)
	return ind
}

func (sa *SparseArray) NumElemAt(r uint64) uint64 {
	// Returns the rank of the input index + 1
	return sa.rank.Rank1(r + 1)
}

func (sa *SparseArray) Size() uint64 {
	// Returns the size of the sparse array
	return sa.bitv.Length()
}

func (sa *SparseArray) NumElem() uint64 {
	// Returns the number of elements in the sparse array
	return uint64(len(sa.strings)) + 1
}

func (sa *SparseArray) Save(filename string) {
	/*

		Saves sparse array to file

		Specifically, saves underlying structures
			Bitvector
			Rank
			Select
		and array of strings

		All to separate files

	*/

	fmt.Println(filename)

	// Save bitvector and rank struct
	sa.rank.Save(filename)

	// Save strings
	fname := strings.Split(filename, ".")[0] + "_sparsearray.txt"

	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	utils.CheckError(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	for i := 0; i < len(sa.strings); i++ {
		w.WriteString(sa.strings[i] + "\n")
	}

	w.Flush()

}

func Load(filename string) *SparseArray {
	/*
		Assumes that the filename points to a file saved using the
		SparseArray.Save() method.

		Loads SparseArray from saved files

		Loads underlying structures
			Bitvector
			Rank
			Select
		and loads string array

	*/

	sa := &SparseArray{}

	r := bvrank.Load(filename)
	sa.rank = r
	sa.bitv = r.Bitvector
	sa.sel = bvselect.NewSelectStruct(sa.bitv, sa.rank)

	// Load strings
	fname := strings.Split(filename, ".")[0] + "_sparsearray.txt"

	file, err := os.Open(fname)
	utils.CheckError(err)
	defer file.Close()

	strings := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	sa.strings = strings

	return sa
}

// Delete later
func (sa *SparseArray) GetBytes() []byte {
	return sa.bitv.Bytes()
}

// Delete later
func (sa *SparseArray) GetStrings() []string {
	return sa.strings
}
