package maglev

import (
	"github.com/go-katran/pkg/ch_helpers"
	"github.com/go-katran/pkg/util"
)

const (
	kHashSeed0 = 0
	kHashSeed1 = 2307
	kHashSeed2 = 42
	kHashSeed3 = 2718281828
)

type MaglevBase interface {
	ch_helpers.ConsistentHash
	GenMaglevPermutation(permutation []uint32, endpoint ch_helpers.Endpoint, pos, ring_size uint32)
}

type MaglevBaseImpl struct {
	MaglevBase
}

func (m MaglevBaseImpl) GenMaglevPermutation(permutation []uint32,
	endpoint ch_helpers.Endpoint, pos, ring_size uint32) {
	offset_hash := util.MurmurHash3_x64_64(endpoint.Hash, kHashSeed2, kHashSeed0)
	offset := offset_hash % uint64(ring_size)
	skip_hash := util.MurmurHash3_x64_64(endpoint.Hash, kHashSeed3, kHashSeed1)
	skip := (skip_hash % uint64((ring_size - 1))) + 1
	permutation[2*pos] = uint32(offset)
	permutation[2*pos+1] = uint32(skip)
}
