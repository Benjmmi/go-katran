package maglev

import "github.com/go-katran/pkg/ch_helpers"

const (
	kHashSeed0 = 0
	kHashSeed1 = 2307
	kHashSeed2 = 42
	kHashSeed3 = 2718281828
)

type MaglevBase interface {
	ch_helpers.ConsistentHash
	GenMaglevPermutation(permutation []int, endpoint ch_helpers.Endpoint, pos int, ring_size int) []int
}
