package healthchecking_kern

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -cflags "-fdebug-prefix-map=/ebpf=." bpf ../../lib/bpf/healthchecking_kern.c -- -I../../lib/linux_includes




