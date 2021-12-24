package maglev

import (
	"github.com/go-katran/pkg/ch_helpers"
	"github.com/go-katran/pkg/maglev"
	"sort"
	"testing"
)

// 检查是否与 C++ 实现的结果一致
func Test(t *testing.T) {
	endpoints := []maglev.Endpoint{}
	freq := make([]uint32, 400)
	var endpoint maglev.Endpoint
	var n1 float64 = 0
	var n2 float64 = 0

	for i := 0; i < 400; i++ {
		endpoint.Num = uint32(i)
		endpoint.Hash = uint64(10 * i)
		if i%1 == 0 {
			endpoint.Weight = 100
		} else {
			endpoint.Weight = 1
		}
		endpoints = append(endpoints, endpoint)
	}
	hash_func := ch_helpers.Maglev

	maglev_hashing := ch_helpers.CHFactoryMake(hash_func)
	ch1 := maglev_hashing.GenerateHashRing(endpoints, 65537)

	deleted_real_num := 400 - 1
	endpoints = endpoints[:1]
	ch2 := maglev_hashing.GenerateHashRing(endpoints, 65537)
	for i := 0; i < len(ch1); i++ {
		freq[ch1[i]]++
	}
	sorted_freq := []uint32{}
	sorted_freq = append(sorted_freq, freq...)
	sort.Slice(sorted_freq, func(i, j int) bool { return sorted_freq[i] < sorted_freq[j] })
	t.Log("min freq is ", sorted_freq[0], " max freq is ", sorted_freq[len(sorted_freq)-1])
	t.Log("p95 w: ", sorted_freq[(len(sorted_freq)/20)*19], "\np75 w: ", sorted_freq[(len(sorted_freq)/20)*15], "\np50 w: ", sorted_freq[len(sorted_freq)/2], "\np25 w: ", sorted_freq[len(sorted_freq)/4], "\np5 w: ", sorted_freq[len(sorted_freq)/20])
	for i := 0; i < len(ch1); i++ {
		if ch1[i] != ch2[i] {
			if ch1[i] == deleted_real_num {
				n1++
				continue
			}
			n2++
		}
	}
	t.Log("changes for affected real: ", n1, "; and for not affected ", n2)
	t.Log(" this is: ", (int(n2) / len(ch1) * 100), "%")
}
