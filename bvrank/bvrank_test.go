package bvrank_test

import (
	"fmt"
	"testing"

	"github.com/lwhite17/CMSC701_HW2/bitvector"
	"github.com/lwhite17/CMSC701_HW2/bvrank"
	"github.com/tmthrgd/go-popcount"
)

func TestRank(t *testing.T) {
	// bs := []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	// bs := []byte{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	// bv := bitvector.NewBitVector(bs, uint64(len(bs)))

	bv := bitvector.NewBitVector([]byte{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0}, uint64(32))

	r := bvrank.NewRankStruct(bv)

	fmt.Println("    Length of bytes: ", bv.Length())
	fmt.Println("    Number of chunks: ", r.NumberOfChunks())
	fmt.Println("        Length of chunk: ", r.ChunkLength())
	fmt.Println("    Number of sub chunks: ", r.NumberOfSubChunks())
	fmt.Println("        Lengh of sub chunks: ", r.SubChunkLength())

	r.PopulateRanks()

	// fmt.Println(bv.Bytes())
	// fmt.Println(r.GetRanks())

	testints := []uint64{0, 1, 10, 11, 22, 33, 44, 55, 66, 77, 88, 99, 150, 194, 195, 196, 197, 198, 199, 200, 201}
	for i := 0; i < len(testints); i++ {
		_ = r.Rank1(testints[i])
		// ranki := r.Rank1(testints[i])
		// fmt.Println("rank at ind ", testints[i], " is ", ranki)
	}
}

func TestDevRank(t *testing.T) {
	// bv := bitvector.NewBitVector([]byte{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0}, uint64(16))
	bv := bitvector.NewBitVector([]byte{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0}, uint64(32))
	// bv := bitvector.NewBitVector([]byte{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0}, uint64(64))
	// bv := bitvector.NewBitVector([]byte{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0}, uint64(64))
	r := bvrank.NewRankStruct(bv)
	// r = r.PopulateRanks()

	// fmt.Println(r.Ranks())

	nchunks := r.NumberOfChunks()
	chlength := r.ChunkLength()
	nschunks := r.NumberOfSubChunks()
	schlength := r.SubChunkLength()

	fmt.Println("chunk length: ", chlength)
	fmt.Println("number of chunks: ", nchunks)
	fmt.Println("sub chunk length: ", schlength)
	fmt.Println("number of sub chunks: ", nschunks)

	chranks := make([]uint64, nchunks)
	schranks := make([][]uint64, nchunks)

	var chunk []byte

	// remainder := bv.Length() - (nchunks-1)*chlength
	// fmt.Println("remainder: ", remainder)

	// fmt.Println("sub chunk length: ", schlength)
	// remaindersch := uint64(math.Ceil(float64(remainder) / float64(schlength)))
	// fmt.Println("sub chunks in remainder chunk: ", remaindersch)

	vector := bv.Bytes()

	var singlesubchranks []uint64

	var i uint64
	var j uint64
	for i = 0; i < nchunks; i++ {
		if i == uint64(0) {
			chranks[i] = 0
		} else {
			lastsubchunk := chunk[j*schlength:]
			chranks[i] = chranks[i-1] + schranks[i-1][j-1] + popcount.CountBytes(lastsubchunk)
		}

		if i == nchunks-1 {
			chunk = vector[i*chlength:]
			singlesubchranks = make([]uint64, r.RemainderNumberOfSubChunks())
			fmt.Println("Empty chunk subranks: ", singlesubchranks)
		} else {
			chunk = vector[i*chlength : (i+1)*chlength]
			singlesubchranks = make([]uint64, nschunks)
			fmt.Println("Empty chunk subranks: ", singlesubchranks)
		}

		fmt.Println("chunk: ", chunk)

		for j = 0; j < uint64(len(singlesubchranks)); j++ {
			fmt.Println("j: ", j)
			if j == uint64(0) {
				singlesubchranks[j] = uint64(0)
			} else {

				prevsubchunk := chunk[(j-1)*schlength : (j)*schlength]
				fmt.Println("    schind: ", j, "prev sub chunk: ", prevsubchunk)

				singlesubchranks[j] = singlesubchranks[j-1] + popcount.CountBytes(prevsubchunk)
				fmt.Println("  update: ", singlesubchranks)
			}
		}

		fmt.Println("Done j loop")

		schranks[i] = singlesubchranks

	}

	fmt.Println("chunk ranks: ", chranks)
	fmt.Println("sub chunk ranks: ", schranks)

}

func TestRank400(t *testing.T) {
	bv := bitvector.NewBitVector(make([]byte, 0), 0)

	fmt.Println("bitvector length: ", bv.Length())

	r := bvrank.NewRankStruct(bv)
	r = r.PopulateRanks()
	fmt.Println(r.GetRanks16())
	fmt.Print(r.GetRanks8())

	// fmt.Println("\n", r.Overhead())

	// for i := 0; i < 401; i += 10 {
	// 	fmt.Println("i, rank: ", i, r.Rank1(i))
	// }

}

func TestRank20000(t *testing.T) {
	// Make bit vector of length lengths[k] of zeros
	by := make([]byte, 400)
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
	// r = r.PopulateRanks()
	fmt.Println(r.GetRanks16())

	chr, _ := r.GetRanks16()
	fmt.Println(len(chr), chr)

	fmt.Println(r.ChunkLength() * uint64(len(chr)))

	_, schr := r.GetRanks8()
	fmt.Println(schr)

	fmt.Println(r.Overhead())

	var i uint64
	for i = 19000; i < 20000; i++ {
		fmt.Println("\n\ni, rank: ", i, r.Rank1(i))
	}
}

func TestRank750000(t *testing.T) {
	// Make bit vector of length lengths[k] of zeros
	by := make([]byte, 750000)
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
	// fmt.Println(r.GetRanks16())
	var i uint64
	for i = 748979; i < 750000; i++ {
		fmt.Println(i, r.Rank1(i))
	}
}

func TestRank100000(t *testing.T) {
	// Make bit vector of length lengths[k] of zeros
	by := make([]byte, 100000)
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
	// fmt.Println(r.GetRanks16())
	// for i := 50000; i < 100000; i++ {
	// 	fmt.Println(i, r.Rank1(i))
	// }

	fmt.Println(r.GetRanks16())

}

func TestConsecutive(t *testing.T) {
	for k := 10; k < 100; k++ {
		fmt.Println("\n", k)
		//
		// Make bit vector of length lengths[k]

		by := make([]byte, k)
		n := uint64(len(by))

		for j := 0; j < k; j++ {
			if j%23 == 10 {
				by[j] = 1
			} else {
				by[j] = 0
			}
		}

		// fmt.Println(by, n)

		b := bitvector.NewBitVector(by, n)

		r := bvrank.NewRankStruct(b)
		r.PopulateRanks()
	}
}

func TestAccuracy(t *testing.T) {
	b := []byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1}
	bv := bitvector.NewBitVector(b, uint64(len(b)))

	fmt.Println(bv.Length())

	r := bvrank.NewRankStruct(bv)
	r = r.PopulateRanks()

	var i uint64
	for i = 0; i < bv.Length(); i++ {
		fmt.Println("i, rank:", i, r.Rank1(i))
	}
	// i := 379
	// fmt.Println("i, rank:", i, r.Rank1(i))

}
