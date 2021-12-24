package maglev

import "github.com/go-katran/pkg/ch_helpers"

type MaglevHashV2 struct {
	MaglevBase
}

func (m MaglevHashV2) GenerateHashRing(endpoints []ch_helpers.Endpoint, ring_size int) []int {
	result := []int{ring_size, -1}

	if len(endpoints) == 0 {
		return result
	} else if len(endpoints) == 1 {
		for i := range result {
			result[i] = int(endpoints[0].Num)
		}
		return result
	}

	var max_weight uint32 = 0
	for _, endpoint := range endpoints {
		if endpoint.Weight > max_weight {
			max_weight = endpoint.Weight
		}
	}
	runs := uint32(0)
	permutation := []uint32{uint32(len(endpoints) * 2), 0}
	next := []uint32{uint32(len(endpoints)), 0}
	cum_weight := []uint32{uint32(len(endpoints)), 0}

	for i := 0; i < len(endpoints); i++ {
		m.GenMaglevPermutation(permutation, endpoints[i], uint32(i), uint32(ring_size))
	}
	for {
		for i := 0; i < len(endpoints); i++ {
			cum_weight[i] += endpoints[i].Weight
			if cum_weight[i] >= max_weight {
				cum_weight[i] -= max_weight
				offset := permutation[2*i]
				skip := permutation[2*i+1]
				cur := (offset + next[i]*skip) % uint32(ring_size)
				for result[cur] >= 0 {
					next[i] += 1
					cur = (offset + next[i]*skip) % uint32(ring_size)
				}
				result[cur] = int(endpoints[i].Num)
				next[i] += 1
				runs++
				if runs == uint32(ring_size) {
					return result
				}
			}
		}
	}
}
