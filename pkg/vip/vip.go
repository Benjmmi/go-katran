package vip

import (
	"github.com/go-katran/pkg/bootstrap"
	"github.com/go-katran/pkg/ch_helpers"
)

type ModifyAction int

const (
	ADD ModifyAction = 0
	DEL ModifyAction = 1
)

type RealPos struct {
	Real int `json:"real"`
	Pos  int `json:"pos"`
}

type UpdateReal struct {
	action  ModifyAction        `json:"action"`
	Enpoint ch_helpers.Endpoint `json:"enpoint"`
}

/**
 * VipRealMeta struct which is used by Vip class to store real's related metadata
 * such as real's weight and hash
 */
type VipRealMeta struct {
	weight int `json:"weight"`
	hash   int `json:"hash"`
}

type Vip struct {
	VipNum     int                    `json:"vipNum"`     // number which uniquely identifies this vip (also used as an index inside forwarding table)
	VipFlags   int                    `json:"vipFlags"`   // vip related flags (such as "dont use src port for hashing" etc)
	ChRingSize int                    `json:"chRingSize"` // size of ch ring
	Reals      map[int]VipRealMeta    `json:"reals"`      // map of reals (theirs opaque id). the value is a real's related metadata (weight and per real hash value).
	ChRing     []int                  `json:"chRing"`     //ch ring which is used for this vip. we are going to use it for delta computation (between old and new ch rings)
	Chash      bootstrap.HashFunction `json:"chash"`      // hash function to generate hash ring
}

func NewVip(vipNum int, vipFlags int, ringSize int, hashFunc bootstrap.HashFunction) *Vip {
	return &Vip{
		VipNum:     vipNum,
		VipFlags:   vipFlags,
		ChRingSize: ringSize,
		Chash:      hashFunc,
	}
}

func (v Vip) ClearVipFlags() {
	v.VipFlags = 0
}
