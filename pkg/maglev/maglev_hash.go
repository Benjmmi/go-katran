package maglev

type MaglevHash struct {
	MaglevBaseImpl
}

func (m MaglevHash) GenerateHashRing(endpoints []Endpoint, ring_size int) []int {
	result := make([]int, ring_size)
	for i := range result {
		result[i] = -1
	}

	if len(endpoints) == 0 {
		return result
	} else if len(endpoints) == 1 {
		for i := range result {
			result[i] = int(endpoints[0].Num)
		}
		return result
	}
	runs := uint32(0)
	permutation := make([]uint32, len(endpoints)*2)
	next := make([]uint32, len(endpoints))

	for i := 0; i < len(endpoints); i++ {
		m.GenMaglevPermutation(permutation, endpoints[i], uint32(i), uint32(ring_size))
	}
	for {
		for i := 0; i < len(endpoints); i++ {
			offset := permutation[2*i]
			skip := permutation[2*i+1]
			// our realization of "weights" for maglev's hash.
			for j := uint32(0); j < endpoints[i].Weight; j++ {
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
			endpoints[i].Weight = 1
		}
	}
}
