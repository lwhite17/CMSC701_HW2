package run

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lwhite17/CMSC701_HW2/bitvector"
	"github.com/lwhite17/CMSC701_HW2/bvrank"
	"github.com/lwhite17/CMSC701_HW2/bvselect"
	"github.com/lwhite17/CMSC701_HW2/sparsearray"
)

func RunTask1() {
	/* Runs task 1

	For varying lengths of bitvector,
		Generate bitvector
		Generate rank structure
		Time multiple calls of Rank1
	Print results

	*/

	fmt.Println("running task 1....")

	lengths := []int{500, 1000, 2500, 5000, 7500, 10000, 12500, 15000, 17500, 25000, 50000, 75000, 100000, 250000, 500000, 750000, 1000000}
	var timer float64

	times := make([]float64, len(lengths))
	avgtimes := make([]float64, len(lengths))
	overheads := make([]uint64, len(lengths))

	// bv := bitvector.NewBitVector(make([]byte, 0), 0)
	var r *bvrank.Rank

	for lengthind := 0; lengthind < len(lengths); lengthind++ {
		// fmt.Println(lengths[lengthind])
		i := lengths[lengthind]

		// Make bit vector of length lengths[k] of zeros
		by := make([]byte, i)
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

		// Create and populate rank struct
		r = bvrank.NewRankStruct(bv)
		r = r.PopulateRanks()

		// Save rank struct
		// savefname := fmt.Sprintf(basefname, lengths[lengthind])
		// r.Save(savefname)

		// initialize timer
		timer = float64(0)

		// set number of rank calls to perform
		numberofcalls := 2000

		// Do rank1 operations
		for k := 0; k < numberofcalls; k++ {
			ind := uint64(rand.Intn(int(n)))
			start := time.Now()
			r.Rank1(ind)
			end := time.Now()
			timer += end.Sub(start).Seconds()
		}
		// Track results
		times[lengthind] = timer
		avgtimes[lengthind] = timer / float64(numberofcalls)
		overheads[lengthind] = r.Overhead()

	}
	// Print results to command line
	fmt.Println("\n    Total times per N:\n", times)
	fmt.Println("\n    Average times per query N:\n", avgtimes)
	fmt.Println("\n    Overhead per N:\n", overheads)
}

func RunTask2() {
	/* Runs task 1

	For varying lengths of bitvector,
		Generate bitvector
		Generate rank structure
		Generate select structure
		Time multiple calls of Select1
	Print results

	*/

	fmt.Println("running task 2....")

	lengths := []int{500, 1000, 2500, 5000, 7500, 10000, 12500, 15000, 17500, 25000, 50000, 75000, 100000, 250000, 500000, 750000, 1000000}
	var timer float64

	times := make([]float64, len(lengths))
	avgtimes := make([]float64, len(lengths))
	overheads := make([]uint64, len(lengths))

	for lengthind := 0; lengthind < len(lengths); lengthind++ {
		// fmt.Println(lengths[lengthind])
		i := lengths[lengthind]

		// Make bit vector of length lengths[k] of zeros
		by := make([]byte, i)
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

		// Create and populate rank struct
		r := bvrank.NewRankStruct(bv)
		r = r.PopulateRanks()

		// Create select struct
		s := bvselect.NewSelectStruct(r.Bitvector, r)

		// Initialize timer
		timer = float64(0)

		// Initialize number of calls
		numberofcalls := 2000

		// Find max possible rank to bound random inputs to select
		maxrank := int(r.Rank1(r.Bitvector.Length()))

		// fmt.Println("Call selects.....")

		// Do rank1 operations
		for k := 0; k < numberofcalls; k++ {
			ind := uint64(rand.Intn(maxrank))
			start := time.Now()
			s.Select1(ind)
			end := time.Now()
			timer += end.Sub(start).Seconds()
		}

		// Print results
		fmt.Println("Average time (seconds) per call of Rank1:", timer/float64(numberofcalls))
		fmt.Println("\n ")
		times[lengthind] = timer
		avgtimes[lengthind] = timer / float64(numberofcalls)
		overheads[lengthind] = r.Overhead()

	}
	// Print results
	fmt.Println("\n    Total times per N:\n", times)
	fmt.Println("\n    Average times per query N:\n", avgtimes)
	fmt.Println("\n    Overhead per N:\n", overheads)
}

func RunTask3() {
	/* Runs task 1

	For varying lengths of sparse arrays,
		Generate sparse array
		Time multiple calls of GetAtRank
		Time multiple calls of GetAtIndex
		Time multiple calls of GetIndexOf
	Print results

	*/
	fmt.Println("running task 3....")

	sparsities := []int{1, 5, 10}
	lengths := []int{1000, 10000, 100000, 1000000}

	fmt.Println("lengths:", lengths)

	nfcalls := 2000
	fmt.Println("number of function calls:", nfcalls)

	var skipfact float64
	var timer float64

	for s := 0; s < len(sparsities); s++ {
		sparsity := sparsities[s]
		fmt.Println("\n\nSPARSITY:", sparsity, "%")

		timesGAR := make([]float64, len(lengths))
		avgtimesGAR := make([]float64, len(lengths))
		timesGAI := make([]float64, len(lengths))
		avgtimesGAI := make([]float64, len(lengths))
		timesGIO := make([]float64, len(lengths))
		avgtimesGIO := make([]float64, len(lengths))

		numelems := make([]uint64, len(lengths))

		if s == 0 {
			skipfact = 0.01
		} else if s == 1 {
			skipfact = 0.05
		} else {
			skipfact = .10
		}

		for l := 0; l < len(lengths); l++ {
			length := lengths[l]

			// Generate sparse array
			sa := sparsearray.Create(uint64(length))
			nels := int(float64(length) * skipfact)
			numelems[l] = uint64(nels)

			for j := 0; j < nels; j++ {
				val := fmt.Sprintf("elementnumber%d", j)
				pos := uint64(float64(j*length/nels) + 5)
				sa = sa.Append(val, pos)
			}

			sa.Finalize()

			// Call many GetAtRank
			timer = float64(0)
			for k := 0; k < nfcalls; k++ {
				maxval := nels
				ind := uint64(rand.Intn(maxval))
				start := time.Now()
				sa.GetAtRank(ind)
				end := time.Now()
				timer += end.Sub(start).Seconds()
			}
			// Track times
			timesGAR[l] = timer
			avgtimesGAR[l] = timer / float64(nfcalls)

			// Call many GetAtIndex
			timer = float64(0)
			for k := 0; k < nfcalls; k++ {
				maxval := sa.Size()
				ind := uint64(rand.Intn(int(maxval)))
				start := time.Now()
				sa.GetAtIndex(ind)
				end := time.Now()
				timer += end.Sub(start).Seconds()
			}
			// Track times
			timesGAI[l] = timer
			avgtimesGAI[l] = timer / float64(nfcalls)

			// Call many GetAtIndex
			timer = float64(0)
			for k := 0; k < nfcalls; k++ {
				maxval := nels
				ind := uint64(rand.Intn(maxval))
				if ind == 0 {
					ind += 1
				}
				start := time.Now()
				sa.GetIndexOf(ind)
				end := time.Now()
				timer += end.Sub(start).Seconds()
			}
			// Track times
			timesGIO[l] = timer
			avgtimesGIO[l] = timer / float64(nfcalls)

		}

		// Print results to command line
		fmt.Println("timesGAR:", timesGAR)
		fmt.Println("avgtimesGAR:", avgtimesGAR)
		fmt.Println("timesGAI:", timesGAI)
		fmt.Println("avgtimesGAI:", avgtimesGAI)
		fmt.Println("timesGIO:", timesGIO)
		fmt.Println("avgtimesGIO:", avgtimesGIO)
		fmt.Println("num elements:", numelems)

	}

	fmt.Println("\n\n ")

	// // // USED BENCHMARKING:
	// // How much memory are we saving?

	// // Length: 1,000
	// var l uint64 = 1000
	// s := 0.01
	// nel := uint64(float64(l) * s)
	// emptystrings := make([]string, l-nel)
	// fmt.Println(l - nel)
	// for j := 0; j < len(emptystrings); j++ {
	// 	emptystrings[j] = ""
	// }
	// fmt.Println(unsafe.Sizeof(emptystrings) + unsafe.Sizeof(emptystrings))

	// l = 10000
	// nel = uint64(float64(l) * s)
	// emptystrings = make([]string, l-nel)
	// fmt.Println(l - nel)
	// for j := 0; j < len(emptystrings); j++ {
	// 	emptystrings[j] = ""
	// }
	// fmt.Println(unsafe.Sizeof(emptystrings) + unsafe.Sizeof(emptystrings))

	// l = 100000
	// s = 0.01
	// nel = uint64(float64(l) * s)
	// fmt.Println(l - nel)
	// emptystrings = make([]string, l-nel)
	// for j := 0; j < len(emptystrings); j++ {
	// 	emptystrings[j] = ""
	// }

}
