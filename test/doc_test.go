package test

import (
	"fmt"
	"github.com/cilium/ebpf/link"
	"testing"
)

func Test_func(t *testing.T) {
	fmt.Println("Helllllll")
	link.AttachRawLink()
}
