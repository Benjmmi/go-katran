package healthchecking_ipip

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -cflags "-fdebug-prefix-map=/ebpf=." bpf ../../lib/bpf/healthchecking_ipip.c -- -I../../lib/linux_includes
