package bitvector_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lwhite17/CMSC701_HW2/bitvector"
)

func TestSave(t *testing.T) {
	testb := []byte{0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}
	testbl := uint64(len(testb))

	b := bitvector.NewBitVector(testb, testbl)
	b.Save("testsave.data")

}

func TestRead(t *testing.T) {
	// bb := bitvector.NewBitVector([]byte{}, 0)
	// bb = bb.Load("testread.data")
	// fmt.Println(bb.Bytes(), bb.Length())
	// fmt.Printf("bb: %s\n", reflect.TypeOf(bb))

	b := bitvector.Load("/Users/leahwhite/Documents/UMD/SP23/CMSC701/HWs/HW2/Testing/Test400.data")
	fmt.Println("\n", b.Bytes(), b.Length())
	fmt.Printf("b: %s\n", reflect.TypeOf(b))
}

func TestSaveMany(t *testing.T) {
	testb := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	testbl := uint64(len(testb))

	b := bitvector.NewBitVector(testb, testbl)
	b.Save("Test400.data")
}

func TestGenerate1(t *testing.T) {

	// lengths := []int{25, 50, 100, 150, 250, 500, 1000, 1500, 2000, 2500, 5000, 10000, 20000, 25000, 50000, 100000, 500000, 1000000}
	lengths := []int{100, 250, 500, 750, 1000, 1500, 2000, 2500, 5000, 7500, 10000, 12500, 15000, 17500, 25000, 50000, 75000, 100000, 250000, 500000, 750000, 1000000}

	for k := 1; k < len(lengths); k++ {
		i := lengths[k]
		//
		// Make bit vector of length lengths[k]

		by := make([]byte, i)
		n := uint64(len(by))

		for j := 0; uint64(j) < n; j++ {
			if j%23 == 10 {
				by[j] = 1
			} else {
				by[j] = 0
			}
		}

		// fmt.Println(by, n)

		b := bitvector.NewBitVector(by, n)

		filename := fmt.Sprintf("/Users/leahwhite/Documents/UMD/SP23/CMSC701/HWs/HW2/Analysis/Length%d.data", i)
		// fmt.Println(filename)

		b.Save(filename)
		// fmt.Printf(fmt.Sprintf("/Users/leahwhite/Documents/UMD/SP23/CMSC701/HWs/HW2/Analysis/Length%d.data", i))

	}

}
