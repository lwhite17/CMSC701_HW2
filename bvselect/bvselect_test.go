package bvselect_test

import (
	"fmt"
	"testing"

	"github.com/lwhite17/CMSC701_HW2/bitvector"
	"github.com/lwhite17/CMSC701_HW2/bvrank"
	"github.com/lwhite17/CMSC701_HW2/bvselect"
)

func TestNewSelect(t *testing.T) {

	b := []byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1}
	bv := bitvector.NewBitVector(b, uint64(len(b)))

	r := bvrank.NewRankStruct(bv)
	r = r.PopulateRanks()

	s := bvselect.NewSelectStruct(bv, r)
	// i := uint64(10)

	maxrank := r.Rank1(bv.Length())

	var i uint64
	for i = 0; i < maxrank+1; i++ {
		fmt.Println("i, ind:", i, s.Select1(i))
	}

	// fmt.Println("Last index with rank", i, "is", s.Select1(i))
	// fmt.Println("\n\nfinal answer:", s.Select1(0))
	// fmt.Println("\n\nfinal answer:", s.Select1(1))
	// fmt.Println("\n\nfinal answer:", s.Select1(2))
	// fmt.Println("\n\nfinal answer:", s.Select1(3))
	// fmt.Println("\n\nfinal answer:", s.Select1(4))
	// fmt.Println("\n\nfinal answer:", s.Select1(5))
	// fmt.Println("\n\nfinal answer:", s.Select1(6))
	// fmt.Println("\n\nfinal answer:", s.Select1(7))

}

func TestSelect20000(t *testing.T) {
	by := make([]byte, 20000)
	n := uint64(len(by))

	// Populate bitvector with some 1s
	for j := 0; uint64(j) < n; j++ {
		if j%23 == 10 {
			by[j] = 1
		} else {
			by[j] = 0
		}
	}

	bv := bitvector.NewBitVector(by, n)
	fmt.Println("bitvector length: ", bv.Length())

	r := bvrank.NewRankStruct(bv)
	r = r.PopulateRanks()

	s := bvselect.NewSelectStruct(bv, r)

	maxrank := r.Rank1(bv.Length())

	var i uint64
	for i = 0; i < maxrank+1; i++ {
		fmt.Println("\nhighest ind with rank i: ", i, s.Select1(i))
	}
}
