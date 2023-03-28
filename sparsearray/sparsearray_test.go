package sparsearray_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/lwhite17/CMSC701_HW2/sparsearray"
)

func TestSparseArray(t *testing.T) {
	fmt.Println("testing")

	sa := sparsearray.Create(uint64(10))
	fmt.Println(len(sa.GetBytes()))

	sa = sa.Append("hi", 2)
	sa = sa.Append("hey", 9)
	sa = sa.Append("bonjour", 4)
	sa = sa.Append("hola", 0)
	fmt.Println(sa.GetBytes())
	fmt.Println(sa.GetStrings())

	sa = sa.Finalize()

	// To get string back from pointer in golang, use *elem

	// // Test GetAtRank
	// elem, b := sa.GetAtRank(0)
	// if b {
	// 	fmt.Printf("\n%v %p \n", *elem, elem)
	// }
	// fmt.Println(b)

	// elem, b = sa.GetAtRank(3)
	// if b {
	// 	fmt.Printf("\n%v %p \n", *elem, elem)
	// }
	// fmt.Println(b)

	// elem, b = sa.GetAtRank(10)
	// if b {
	// 	fmt.Printf("\n%v %p \n", *elem, elem)
	// }
	// fmt.Println(b)

	// // Test GetAtIndex
	// elem, b := sa.GetAtIndex(0)
	// if b {
	// 	fmt.Printf("\n0, %v %p \n", *elem, elem)
	// }
	// fmt.Println(b)

	// elem, b = sa.GetAtIndex(9)
	// if b {
	// 	fmt.Printf("\n9, %v %p \n", *elem, elem)
	// }
	// fmt.Println(b)

	elem, b := sa.GetAtIndex(4)
	if b {
		fmt.Printf("\n4, %v %p \n", *elem, elem)
	}
	fmt.Println(b)

	elem, b = sa.GetAtIndex(5)
	if b {
		fmt.Printf("\n5, %v %p \n", *elem, elem)
	}
	fmt.Println(b)

	// Test GetIndexof

	ind := sa.GetIndexOf(1)
	fmt.Println("1", ind)

	ind = sa.GetIndexOf(2)
	fmt.Println("2", ind)

	ind = sa.GetIndexOf(3)
	fmt.Println("3", ind)

	ind = sa.GetIndexOf(4)
	fmt.Println("4", ind)

	ind = sa.GetIndexOf(5)
	fmt.Println("5", ind)

	// // Test NumElemAt
	// ind = sa.NumElemAt(0)
	// fmt.Println("0", ind)

	// ind = sa.NumElemAt(9)
	// fmt.Println("9", ind)

	// ind = sa.NumElemAt(10)
	// fmt.Println("10", ind)

	// ind = sa.NumElemAt(2)
	// fmt.Println("2", ind)

	// ind = sa.NumElemAt(3)
	// fmt.Println("3", ind)

	// ind = sa.NumElemAt(5)
	// fmt.Println("5", ind)

	// ind = sa.NumElemAt(7)
	// fmt.Println("7", ind)

}

func TestAccuracy(t *testing.T) {

	sa := sparsearray.Create(10)
	sa = sa.Append("foo", 1)
	sa = sa.Append("bar", 5)
	sa = sa.Append("baz", 9)
	fmt.Println(sa.GetBytes())
	fmt.Println(sa.GetStrings())

	sa.Finalize()

	val, b := sa.GetAtRank(1)
	fmt.Println(b)
	if b {
		fmt.Println(*val)
	}

	val, b = sa.GetAtIndex(3)
	fmt.Println(b)
	if b {
		fmt.Println(*val)
	}

	val, b = sa.GetAtIndex(5)
	fmt.Println(b)
	if b {
		fmt.Println(*val)
	}

	fmt.Println("GIO 2:", sa.GetIndexOf(2))

	fmt.Println("size:", sa.Size())
	fmt.Println(sa.NumElem())

	fmt.Println(strconv.Itoa(int(0)))
	fmt.Println(string(byte(1)))

	fname := "/Users/leahwhite/Documents/UMD/SP23/CMSC701/HWs/HW2/Testing/SparseArrayTest.txt"
	sa.Save(fname)

	sa1 := sparsearray.Load(fname)

	val, b = sa1.GetAtRank(1)
	fmt.Println(b)
	if b {
		fmt.Println(*val)
	}

	val, b = sa1.GetAtIndex(3)
	fmt.Println(b)
	if b {
		fmt.Println(*val)
	}

	val, b = sa1.GetAtIndex(5)
	fmt.Println(b)
	if b {
		fmt.Println(*val)
	}
	// fmt.Println(sa.GetIndexOf(2))

	fmt.Println("index of 3rd el:", sa.GetIndexOf(3))

	fmt.Println("size:", sa.Size())
	fmt.Println(sa.NumElem())

	fmt.Println(strconv.Itoa(int(0)))
	fmt.Println(string(byte(1)))
}

func TestLarge(t *testing.T) {
	skipfact := 0.01

	length := 1000
	fmt.Println("	length:", length)

	// Generate sparse array
	sa := sparsearray.Create(uint64(length))
	nels := int(float64(length) * skipfact)

	fmt.Println("		number of elements", nels)
	for j := 0; j < nels; j++ {
		val := fmt.Sprintf("elementnumber%d", j)
		pos := uint64(float64(j*length/nels) + 5)
		// fmt.Println(pos)
		sa = sa.Append(val, pos)
	}

	sa = sa.Finalize()

	fmt.Println(sa.GetBytes()[900:910])

	var k uint64
	for k = 1; k < uint64(nels+1); k++ {
		val := sa.GetIndexOf(k)
		fmt.Println(k, val)
	}
}
