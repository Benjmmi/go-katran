package netlink

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"testing"
)

func Test_LoadBPFProg(t *testing.T) {
	fmt.Println("fa")

	f := netlink.BpfFilter{}
	netlink.FilterAdd(&f)
}
