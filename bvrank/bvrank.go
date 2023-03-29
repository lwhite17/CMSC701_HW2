package bvrank

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/lwhite17/CMSC701_HW2/bitvector"
	"github.com/lwhite17/CMSC701_HW2/utils"
	"github.com/tmthrgd/go-popcount"
)

type Rank struct {
	Bitvector  *bitvector.BitVector
	chcode     uint8
	schcode    uint8
	chranks8   []uint8
	schranks8  [][]uint8
	chranks16  []uint16
	schranks16 [][]uint16
	chranks32  []uint32
	schranks32 [][]uint32
	chranks64  []uint64
	schranks64 [][]uint64
}

func NewRankStruct(b *bitvector.BitVector) *Rank {
	// Initialize rank structure with bitvector
	r := Rank{}
	r.Bitvector = b

	// Compute cumulative chunk ranks and
	// relative sub chunk ranks
	rp := r.PopulateRanks()
	return rp
}

func (r *Rank) PopulateRanks() *Rank {
	// Fills out the rank structure's cumulative chunk ranks
	// and relative subchunk ranks

	// Get number of chunks and subchunks
	nchunks := r.NumberOfChunks()
	chlength := r.ChunkLength()

	// Get length of chunk and subchunk
	nschunks := r.NumberOfSubChunks()
	schlength := r.SubChunkLength()

	// Initialize empty chunk ranks and subchunk ranks
	chranks64 := make([]uint64, nchunks)
	schranks64 := make([][]uint64, nchunks)

	// Get bitvector bytes
	vector := r.Bitvector.Bytes()

	var chunk []byte
	var singlesubchranks []uint64

	var i uint64
	var j uint64
	for i = 0; i < nchunks; i++ {
		if i == uint64(0) {
			// If first chunk, rank = 0
			chranks64[i] = 0

		} else {
			// Otherwise, rank = rank of previous chunk + popCount(last subchunk)
			lastsubchunk := chunk[(j-1)*schlength:]
			chranks64[i] = chranks64[i-1] + schranks64[i-1][j-1] + popcount.CountBytes(lastsubchunk)

		}

		if i == nchunks-1 {
			// If last chunk,
			// 		Chunk bitvector is full bitvector from index to end
			// 		Initialize subchunks array to have length of remaining subchunks
			chunk = vector[i*chlength:]
			singlesubchranks = make([]uint64, r.RemainderNumberOfSubChunks())

		} else {
			// If not last chunk,
			// 		Chunk bitvector is full bitvector from i*chlength to (i+1)*chlength
			// 		Initialize subchunks array to have length regular number of subchunks
			chunk = vector[i*chlength : (i+1)*chlength]
			singlesubchranks = make([]uint64, nschunks)

		}

		// Populate subchunk ranks
		for j = 0; j < uint64(len(singlesubchranks)); j++ {
			if j == uint64(0) {
				// If first subchunk, rank is 0
				singlesubchranks[j] = uint64(0)
			} else {
				// Otherwise, add popcount of subchunk's bitvector to previous rank
				prevsubchunk := chunk[(j-1)*schlength : (j)*schlength]
				singlesubchranks[j] = singlesubchranks[j-1] + popcount.CountBytes(prevsubchunk)

			}
		}

		schranks64[i] = singlesubchranks

	}

	r = r.ConvertToSmallest(chranks64, schranks64)

	return r
}

func (r *Rank) Rank1(i uint64) uint64 {
	// Return the rank at the input index

	// Get lengths of chunks and subchunks
	chunklength := r.ChunkLength()
	subchunklength := r.SubChunkLength()
	// Get number of chunks and subchunks
	nchunks := r.NumberOfChunks()
	nschunks := r.NumberOfSubChunks()

	// Get length of bitvector
	n := r.Bitvector.Length()

	if i < 1 {
		// If i is 0, return 0
		return uint64(0)
	} else if uint64(i) > n {
		// If i is greater than the length of the bitvector
		// set i to the length of the bitvector
		i = uint64(n)
	}

	if uint64(i) == n {
		// If i is the length of the bitvector,

		// Use the last chunk
		chind := r.NumberOfChunks() - 1

		// Use the last sub chunk
		schind := r.RemainderNumberOfSubChunks() - 1

		// Get rank of last chunk
		chrank := r.accesschunkrank(chind)

		// Get rank of last subchunk
		schrank := r.accesssubchunkrank(chind, schind)

		// Get remaining sub-bitvector
		check := r.Bitvector.Bytes()[chind*chunklength+(schind)*subchunklength:]

		// Return sum of ranks and popcount of sub-bitvector
		return uint64(chrank + schrank + popcount.CountBytes(check))
	}

	// Get chunk and subchunk indices
	chunkind := uint64(i) / chunklength
	subchunkind := (uint64(i) - chunkind*chunklength) / subchunklength

	// Correct edge cases
	if chunkind >= nchunks {
		chunkind = nchunks - 1
		if subchunkind >= r.RemainderNumberOfSubChunks() {
			subchunkind = r.RemainderChunkLength() - 1
		}
	} else if subchunkind >= nschunks {
		subchunkind = subchunklength - 1
	}

	// Get chunk and subchunk ranks
	var chunkrank uint64 = r.accesschunkrank(chunkind)
	var subchunkrank uint64 = r.accesssubchunkrank(chunkind, subchunkind)

	// Get sub-bitvector
	check := r.Bitvector.Bytes()[chunkind*chunklength+(subchunkind)*subchunklength : i]

	// Return sum of ranks and popcount of sub-bitvector
	return uint64(chunkrank + subchunkrank + popcount.CountBytes(check))
}

func (r *Rank) Overhead() uint64 {
	// Calculate the overhead of this structure

	nchunks := r.NumberOfChunks()
	nsubchunks := r.NumberOfChunks()
	nremainder := r.RemainderNumberOfSubChunks()

	var chm uint64
	var schm uint64
	if r.chcode == 0 {
		chm = 1
	} else if r.chcode == 1 {
		chm = 2
	} else if r.chcode == 2 {
		chm = 4
	} else {
		chm = 8
	}

	if r.schcode == 0 {
		schm = 1
	} else if r.chcode == 1 {
		schm = 2
	} else if r.chcode == 2 {
		schm = 4
	} else {
		schm = 8
	}

	// Calculate number of chunks * size of integer used to store chunk rank +
	// 		number of subchunks * size of integer used to store subchunk rank

	nchranks := nchunks * chm
	nschranks := ((nsubchunks * nsubchunks) + nremainder) * schm

	// Add 2 for the codes that store the integer type
	return nchranks + nschranks + 2

	// // Naive calc
	// n := r.bitvector.Length()
	// cumrank := math.Log2(float64(n)) * float64(n) / math.Pow(math.Log2(float64(n)), 2)
	// relcumrank := float64(n) * math.Log2(math.Pow(math.Log2(float64(n)), 2)) / math.Log2(float64(n))
	// return uint64(cumrank + relcumrank)
}

func (r *Rank) Save(filename string) {
	// Save rank structure and underlying bitvector to files

	basefname := strings.Split(filename, ".")[0]
	rankfname := basefname + "_rank.txt"
	bitvectorfname := basefname + "_bitvector.data"

	r.Bitvector.Save(bitvectorfname)

	f, err := os.OpenFile(rankfname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	utils.CheckError(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	w.WriteString(strconv.FormatUint(r.Bitvector.Length(), 10) + "\n")

	chcode := r.chcode
	schcode := r.schcode

	nchunks := r.NumberOfChunks()

	w.WriteString(strconv.Itoa(int(chcode)) + "\n")
	w.WriteString(strconv.Itoa(int(schcode)) + "\n")
	w.WriteString("\n")

	chranks64, schranks64 := r.ConvertToUint64()

	var j uint64
	for j = 0; j < nchunks; j++ {
		w.WriteString(strconv.FormatUint(chranks64[j], 10) + "\n")
		for i := 0; i < len(schranks64[j]); i++ {
			w.WriteString(strconv.FormatUint(schranks64[j][i], 10) + "\n")

		}
		w.WriteString("\n")
	}
	w.Flush()
}

func Load(filename string) *Rank {
	// Load bitvector and rank structure
	// Assumes bvrank.Save() was used to save out files

	basefname := strings.Split(filename, ".")[0]

	// Load bitvector
	bitvectorfname := basefname + "_bitvector.data"

	bv := bitvector.Load(bitvectorfname)

	r := &Rank{}
	r.Bitvector = bv

	// Load rank
	rankfname := basefname + "_rank.txt"

	r.chcode = 0
	r.schcode = 0
	r.chranks8 = make([]uint8, 0)
	r.schranks8 = make([][]uint8, 0)
	r.chranks16 = make([]uint16, 0)
	r.schranks16 = make([][]uint16, 0)
	r.chranks32 = make([]uint32, 0)
	r.schranks32 = make([][]uint32, 0)
	r.chranks64 = make([]uint64, 0)
	r.schranks64 = make([][]uint64, 0)

	file, err := os.Open(rankfname)
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	n, err := strconv.ParseUint(scanner.Text(), 10, 64)
	utils.CheckError(err)
	fmt.Println(n)

	scanner.Scan()
	chcode, err := strconv.Atoi(scanner.Text())
	utils.CheckError(err)

	scanner.Scan()
	schcode, err := strconv.Atoi(scanner.Text())
	utils.CheckError(err)

	r.chcode = uint8(chcode)
	r.schcode = uint8(schcode)

	nchunks := uint64(math.Ceil(float64(n) / (math.Pow(math.Log2(float64(n)), 2))))
	nschunks := 2 * uint64(math.Log2(float64(n)))

	chranks64 := make([]uint64, nchunks)
	schranks64 := make([][]uint64, nchunks)

	var i uint64
	var j uint64
	maxi := nschunks
	for j = 0; j < nchunks; j++ {
		scanner.Scan()
		scanner.Scan()
		chranks64[j], err = strconv.ParseUint(string(scanner.Text()), 10, 64)
		utils.CheckError(err)

		if j == nchunks-1 {
			maxi = uint64(math.Ceil(float64(r.RemainderChunkLength()) / float64(r.SubChunkLength())))
		}

		schranks64[j] = make([]uint64, maxi)

		for i = 0; i < maxi; i++ {
			scanner.Scan()
			schranks64[j][i], err = strconv.ParseUint(string(scanner.Text()), 10, 64)
			utils.CheckError(err)
		}
	}

	r = r.ConvertToSmallest(chranks64, schranks64)

	return r

}

func (r *Rank) ChunkLength() uint64 {
	// Returns the length of a chunk's sub-bitvector
	n := uint64(r.Bitvector.Length())
	return uint64(math.Pow(math.Log2(float64(n)), 2))
}

func (r *Rank) NumberOfChunks() uint64 {
	// Returns the number of chunks
	n := uint64(r.Bitvector.Length())
	return uint64(math.Ceil(float64(n) / (math.Pow(math.Log2(float64(n)), 2))))
}

func (r *Rank) SubChunkLength() uint64 {
	// Returns the length of a subchunk's sub-bitvector
	n := uint64(r.Bitvector.Length())
	return uint64(.5 * math.Log2(float64(n)))
}

func (r *Rank) NumberOfSubChunks() uint64 {
	// Returns the number of subchunks
	n := uint64(r.Bitvector.Length())
	return 2 * uint64(math.Log2(float64(n)))
}

func (r *Rank) RemainderChunkLength() uint64 {
	// Returns the length of the last (/remainder) chunk
	return r.Bitvector.Length() - (r.NumberOfChunks()-1)*r.ChunkLength()
}

func (r *Rank) RemainderNumberOfSubChunks() uint64 {
	// Returns the number of subchunks in the last (/remainder) chunk
	return uint64(math.Ceil(float64(r.RemainderChunkLength()) / float64(r.SubChunkLength())))
}

func (r *Rank) accesschunkrank(i uint64) uint64 {
	// Returns the uint64 chunk rank at the given index
	if r.chcode == 0 {
		return uint64(r.chranks8[i])
	} else if r.chcode == 1 {
		return uint64(r.chranks16[i])
	} else if r.chcode == 2 {
		return uint64(r.chranks32[i])
	} else {
		return r.chranks64[i]
	}
}

func (r *Rank) accesssubchunkrank(i uint64, j uint64) uint64 {
	// Returns the uint64 subchunk rank at the given indices
	if r.schcode == 0 {
		return uint64(r.schranks8[i][j])
	} else if r.schcode == 1 {
		return uint64(r.schranks16[i][j])
	} else if r.schcode == 2 {
		return uint64(r.schranks32[i][j])
	} else {
		return r.schranks64[i][j]
	}
}

func (r *Rank) ConvertToSmallest(chranks64 []uint64, schranks64 [][]uint64) *Rank {
	// Converts the chunk ranks and subchunk ranks into arrays of the
	// smallest possible integers and updates the respective structure field

	nchunks := r.NumberOfChunks()
	nschunks := r.NumberOfSubChunks()

	maxchrank := chranks64[len(chranks64)-1]

	var maxschrank uint64 = uint64(math.Max(float64(nschunks), float64(r.RemainderNumberOfSubChunks())))

	var i uint64

	if maxchrank < uint64(math.Pow(2, 8)) {
		chranks8 := make([]uint8, nchunks)
		for i = 0; i < nchunks; i++ {
			chranks8[i] = uint8(chranks64[i])
		}
		r.chcode = byte(0)
		r.chranks8 = chranks8

	} else if maxchrank < uint64(math.Pow(2, 16)) {
		chranks16 := make([]uint16, nchunks)
		for i = 0; i < nchunks; i++ {
			chranks16[i] = uint16(chranks64[i])
		}
		r.chcode = byte(1)
		r.chranks16 = chranks16

	} else if maxchrank < uint64(math.Pow(2, 32)) {
		chranks32 := make([]uint32, nchunks)
		for i = 0; i < nchunks; i++ {
			chranks32[i] = uint32(chranks64[i])
		}
		r.chcode = byte(2)
		r.chranks32 = chranks32

	} else {
		r.chcode = byte(3)
		r.chranks64 = chranks64
	}

	var j uint64

	if maxschrank < uint64(math.Pow(2, 8)) {
		schranks8 := make([][]uint8, nchunks)
		for i = 0; i < nchunks; i++ {
			lsc := uint64(len(schranks64[i]))
			schranks8[i] = make([]uint8, lsc)
			for j = 0; j < lsc; j++ {
				schranks8[i][j] = uint8(schranks64[i][j])
			}
		}
		r.schcode = byte(0)
		r.schranks8 = schranks8

	} else if maxschrank < uint64(math.Pow(2, 16)) {
		schranks16 := make([][]uint16, nchunks)
		for i = 0; i < nchunks; i++ {
			lsc := uint64(len(schranks64[i]))
			schranks16[i] = make([]uint16, lsc)
			for j = 0; j < lsc; j++ {
				schranks16[i][j] = uint16(schranks64[i][j])
			}
		}
		r.schcode = byte(1)
		r.schranks16 = schranks16

	} else if maxschrank < uint64(math.Pow(2, 32)) {
		schranks32 := make([][]uint32, nchunks)
		for i = 0; i < nchunks; i++ {
			lsc := uint64(len(schranks64[i]))
			schranks32[i] = make([]uint32, lsc)
			for j = 0; j < lsc; j++ {
				schranks32[i][j] = uint32(schranks64[i][j])
			}
		}
		r.schcode = byte(2)
		r.schranks32 = schranks32

	} else {
		r.schcode = byte(3)
		r.chranks64 = chranks64
	}

	return r
}

func (r *Rank) ConvertToUint64() ([]uint64, [][]uint64) {
	// Converts the current compactly stored chunk ranks and subchunk ranks
	// to uint64 arrays / slices

	nchunks := r.NumberOfChunks()

	var chranks64 []uint64 = make([]uint64, nchunks)
	var schranks64 [][]uint64 = make([][]uint64, nchunks)

	if r.chcode == 0 {
		chranks8 := r.chranks8
		for i := 0; uint64(i) < nchunks; i++ {
			chranks64[i] = uint64(chranks8[i])
		}

	} else if r.chcode == 1 {
		chranks16 := r.chranks16
		for i := 0; uint64(i) < nchunks; i++ {
			chranks64[i] = uint64(chranks16[i])
		}

	} else if r.chcode == 2 {
		chranks32 := r.chranks32
		for i := 0; uint64(i) < nchunks; i++ {
			chranks64[i] = uint64(chranks32[i])
		}

	} else {
		chranks64 = r.chranks64

	}

	if r.schcode == 0 {
		schranks8 := r.schranks8
		for i := 0; uint64(i) < nchunks; i++ {
			if uint64(i) == nchunks-1 {
				schranks64[i] = make([]uint64, r.RemainderNumberOfSubChunks())
			} else {
				schranks64[i] = make([]uint64, len(schranks8[i]))
			}
			for j := 0; j < len(schranks8[i]); j++ {
				schranks64[i][j] = uint64(schranks8[i][j])
			}
		}

	} else if r.schcode == 1 {
		schranks16 := r.schranks16
		for i := 0; uint64(i) < nchunks; i++ {
			if uint64(i) == nchunks-1 {
				schranks64[i] = make([]uint64, r.RemainderNumberOfSubChunks())
			} else {
				schranks64[i] = make([]uint64, len(schranks16[i]))
			}
			for j := 0; j < len(schranks16[i]); j++ {
				// fmt.Println("j", j)
				schranks64[i][j] = uint64(schranks16[i][j])
			}
		}

	} else if r.schcode == 2 {
		schranks32 := r.schranks32
		for i := 0; uint64(i) < nchunks; i++ {
			if uint64(i) == nchunks-1 {
				schranks64[i] = make([]uint64, r.RemainderNumberOfSubChunks())
			} else {
				schranks64[i] = make([]uint64, len(schranks32[i]))
			}
			for j := 0; j < len(schranks32[i]); j++ {
				// fmt.Println("j", j)
				schranks64[i][j] = uint64(schranks32[i][j])
			}
		}

	} else {
		schranks64 = r.schranks64

	}
	return chranks64, schranks64
}

func (r *Rank) GetRanks8() ([]uint8, [][]uint8) {
	return r.chranks8, r.schranks8
}

func (r *Rank) GetRanks16() ([]uint16, [][]uint16) {
	return r.chranks16, r.schranks16
}

func (r *Rank) GetRanks32() ([]uint32, [][]uint32) {
	return r.chranks32, r.schranks32
}

func (r *Rank) GetRanks64() ([]uint64, [][]uint64) {
	return r.chranks64, r.schranks64
}
