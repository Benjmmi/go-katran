package maglev

import (
	"github.com/go-katran/pkg/util"
)

const (
	kHashSeed0 = 0
	kHashSeed1 = 2307
	kHashSeed2 = 42
	kHashSeed3 = 2718281828
)

/**
 * struct which describes backend, each backend would have unique number,
 * weight (the measurment of how often we would see this endpoint
 * on CH ring) and hash value, which will be used as a seed value
 * (it should be unique value per endpoint for CH to work as expected)
 */
type Endpoint struct {
	Num    uint32 `json:"num"`
	Weight uint32 `json:"weight"`
	Hash   uint64 `json:"hash"`
}

/**
 * ConsistentHash implements interface, which is used by CHFactory class to
 * generate hash ring
 */
type ConsistentHash interface {
	// GenerateHashRing endpoints, ring_size default 65537
	GenerateHashRing(endpoints []Endpoint, ring_size int) []int
}

type MaglevBase interface {
	ConsistentHash
	GenMaglevPermutation(permutation []uint32, endpoint Endpoint, pos, ring_size uint32)
}

type MaglevBaseImpl struct {
	MaglevBase
}

func (m MaglevBaseImpl) GenMaglevPermutation(permutation []uint32,
	endpoint Endpoint, pos, ring_size uint32) {
	offset_hash := util.MurmurHash3_x64_64(endpoint.Hash, kHashSeed2, kHashSeed0)
	offset := offset_hash % uint64(ring_size)
	skip_hash := util.MurmurHash3_x64_64(endpoint.Hash, kHashSeed3, kHashSeed1)
	skip := (skip_hash % uint64((ring_size - 1))) + 1
	permutation[2*pos] = uint32(offset)
	permutation[2*pos+1] = uint32(skip)
}
