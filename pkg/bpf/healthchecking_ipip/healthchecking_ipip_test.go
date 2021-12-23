package healthchecking_ipip

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	"testing"
)

func Test_Netlink_Filter(t *testing.T) {
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

	bpfFilter := netlink.BpfFilter{
		FilterAttrs: netlink.FilterAttrs{
			LinkIndex: link.Attrs().Index,
			Parent:    netlink.HANDLE_MIN_EGRESS,
			Handle:    netlink.MakeHandle(0, 1),
			Protocol:  unix.ETH_P_ALL,
			Priority:  1,
		},
		ClassId:      netlink.HANDLE_ROOT,
		Fd:           objs.bpfPrograms.Healthchecker.FD(),
		Name:         "",
		DirectAction: true,
		Id:           1,
		Tag:          "",
	}

	if err := netlink.FilterAdd(&bpfFilter); err != nil {
		t.Fatal(err)
	}
}
