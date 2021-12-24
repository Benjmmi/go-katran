package ch_helpers

import "github.com/go-katran/pkg/maglev"

type hashFunction int

const (
	Maglev   hashFunction = 1
	MaglevV2 hashFunction = 2
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

func CHFactoryMake(functoin hashFunction) ConsistentHash {
	switch functoin {
	case Maglev:
		return &maglev.MaglevHash{}
	case MaglevV2:
		return &maglev.MaglevHashV2{}
	}
	return nil
}
