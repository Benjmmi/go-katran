package maglev

import "github.com/go-katran/pkg/ch_helpers"

type MaglevHashV2 struct {
	ch_helpers.ConsistentHash
}

func (m MaglevHash) GenerateHashRing(endpoints []ch_helpers.Endpoint, ring_size int) []int {
	return nil
}
