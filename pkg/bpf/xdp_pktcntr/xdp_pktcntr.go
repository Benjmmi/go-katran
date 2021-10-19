package xdp_pktcntr

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -cflags "-fdebug-prefix-map=/ebpf=." bpf ../../lib/bpf/xdp_pktcntr.c -- -I../../lib/linux_includes





