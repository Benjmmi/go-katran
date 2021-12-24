package ch_helpers

import "github.com/go-katran/pkg/maglev"

type hashFunction int

const (
	Maglev   hashFunction = 1
	MaglevV2 hashFunction = 2
)

func CHFactoryMake(functoin hashFunction) maglev.ConsistentHash {
	switch functoin {
	case Maglev:
		return &maglev.MaglevHash{}
	case MaglevV2:
		return &maglev.MaglevHashV2{}
	}
	return nil
}
