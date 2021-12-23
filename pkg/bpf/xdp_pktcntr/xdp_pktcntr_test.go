package xdp_pktcntr

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"testing"
)

func Test_Netlink_XDP(t *testing.T) {

	fmt.Println("Fa")

	objs := bpfObjects{}
	if err := loadBpfObjects(&objs, nil); err != nil {
		t.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	link, err := netlink.LinkByName("eth0")
	if err != nil {
		t.Fatal(err)
	}

	if err := netlink.LinkSetXdpFd(link, objs.bpfPrograms.Pktcntr.FD()); err != nil {
		t.Fatal(err)
	}
}
