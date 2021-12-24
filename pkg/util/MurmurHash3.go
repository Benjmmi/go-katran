package util

func rotl64(x uint64, r int8) uint64 {
	return (x << r) | (x >> (64 - r))
}

func MurmurHash3_x64_64(A, B uint64, seed uint32) uint64 {
	var h1 = uint64(seed)
	var h2 = uint64(seed)

	var c1 uint64 = 0x87c37b91114253d5

	var c2 uint64 = 0x4cf5ad432745937f

	//----------
	// body

	k1 := A
	k2 := B

	k1 *= c1
	k1 = rotl64(k1, 31)
	k1 *= c2
	h1 ^= k1

	h1 = rotl64(h1, 27)
	h1 += h2
	h1 = h1*5 + 0x52dce729

	k2 *= c2
	k2 = rotl64(k2, 33)
	k2 *= c1
	h2 ^= k2

	h2 = rotl64(h2, 31)
	h2 += h1
	h2 = h2*5 + 0x38495ab5

	//----------
	// finalization

	h1 ^= 16
	h2 ^= 16

	h1 += h2
	h2 += h1

	h1 ^= h1 >> 33
	h1 *= 0xff51afd7ed558ccd

	h1 ^= h1 >> 33
	h1 *= 0xc4ceb9fe1a85ec53

	h1 ^= h1 >> 33

	h2 ^= h2 >> 33
	h2 *= 0xff51afd7ed558ccd

	h2 ^= h2 >> 33
	h2 *= 0xc4ceb9fe1a85ec53

	h2 ^= h2 >> 33

	h1 += h2

	return h1
}
