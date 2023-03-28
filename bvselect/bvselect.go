package bvselect

import (
	"math"

	"github.com/lwhite17/CMSC701_HW2/bitvector"
	"github.com/lwhite17/CMSC701_HW2/bvrank"
)

type Select struct {
	b *bitvector.BitVector
	r *bvrank.Rank
}

func NewSelectStruct(b *bitvector.BitVector, r *bvrank.Rank) *Select {
	s := Select{}
	s.b = b
	s.r = r
	return &s
}

func (s *Select) Select1(i uint64) uint64 {

	// Get length of bitvector
	n := s.b.Length()

	// Initialize left, right, and center indices.
	l := uint64(0)
	r := n - 1
	ind := uint64((l + r) / 2)

	// Initialize max number of iterations and max rank
	maxiter := uint64(math.Ceil(math.Log2(float64(n))))
	maxrank := s.r.Rank1(s.b.Length())
	if i >= maxrank-1 {
		return s.b.Length() - 1
	}

	var j uint64

	// Binary search
	for j = 0; j < maxiter; j++ {
		indrank := s.r.Rank1(ind)
		if indrank > i { // Search left
			if r == uint64(ind)+1 && l == uint64(ind-1) {
				return uint64(ind) - 1
			}
			r = uint64(ind)

		} else { // Search right
			if r == uint64(ind)+1 && l == uint64(ind-1) {
				if i == indrank {
					return uint64(ind)
				} else {
					return uint64(ind) + 1
				}
			}

			l = uint64(ind)

		}

		// Update center index
		ind = uint64((l + r) / 2)

	}

	return uint64(ind)
}

func (s *Select) Overhead() uint64 {
	// Since Select struct just points to Rank struct,
	// return size of Rank struct
	return s.r.Overhead()
}

func Load(filename string) *Select {
	// Load underlying rank and bitvector structs
	s := &Select{}
	s.r = bvrank.Load(filename)
	s.b = s.r.Bitvector
	return s
}

func (s *Select) Save(filename string) {
	// Save underlying structures
	s.r.Save(filename)
}
