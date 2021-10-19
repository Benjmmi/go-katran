package xdp_root

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -cflags "-fdebug-prefix-map=/ebpf=." bpf ../../lib/bpf/xdp_root.c -- -I../../lib/linux_includes

