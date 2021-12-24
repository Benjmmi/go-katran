package bootstrap

import "github.com/go-katran/pkg/ch_helpers"

type PcapStorageFormat int

const (
	FILE  PcapStorageFormat = 1
	IOBUF PcapStorageFormat = 2
	PIPE  PcapStorageFormat = 3
)

type KatranFeatureEnum int

const (
	SrcRouting                KatranFeatureEnum = 1 << 0
	InlineDecap               KatranFeatureEnum = 1 << 1
	Introspection             KatranFeatureEnum = 1 << 2
	GueEncap                  KatranFeatureEnum = 1 << 3
	DirectHealthchecking      KatranFeatureEnum = 1 << 4
	LocalDeliveryOptimization KatranFeatureEnum = 1 << 5
	FlowDebug                 KatranFeatureEnum = 1 << 6
)

// KatranMonitorConfig katran monitoring config. being used if katran's bpf code was
// build w/ introspection enabled (-DKATRAN_INTROSPECTION)
type KatranMonitorConfig struct {
	NCpus      int                    `json:"nCpus"`     //NCpus number of cpus
	Pages      int                    `json:"pages"`     //Pages number of pages for even pipe shared memory
	MapFd      int                    `json:"mapFd"`     //MapFd descriptor of event pipe map
	QueueSize  int                    `json:"queueSize"` //QueueSize size of mpmc queue between readers and pcap writer
	PcktLimit  int                    `json:"pcktLimit"` //PcktLimit to capture, 0 - no limit
	SnapLen    int                    `json:"snapLen"`   //SnapLen maximum number of bytes from packet to write.
	Events     map[string]interface{} `json:"events"`    //Events maximum supported events/pcap writers
	Path       string                 `json:"path"`      //Oath where pcap outputs are going to be stored
	Storage    PcapStorageFormat      `json:"storage"`
	BufferSize int                    `json:"bufferSize"`
}

// 	KatranConfig struct which contains all configurations for KatranLB
//	note about rootMapPath and rootMapPos:
//	katran has two modes of operation.
//	the first one is "standalone":
//	when it register itself as one and only xdp prog; this is
//	default. for this mode to work rootMapPath must be equal to "".
//	and we dont evaluate rootMapPos (so it could be any value).
//
//	the second mode of operation - "shared" -
//	is when we have root xdp prog: which is
//	just doing bpf_tail_call for other xdp's progs, which must registers
//	(put their fd's into predifiened position inside rootMap).
//	in this case rootMapPath must be path to "pinned" map, which has been
//	used by root xdp prog, and rootMapPos is a position (index) of
//	katran's fd inside this map.
//
//	by default, if hcInterface is not specified we are going to attach
//	healthchecking bpf program to the mainInterfaces
type KatranConfig struct {
	MainInterface          string                  `json:"mainInterface"`          // MainInterface  name where to attach bpf prog (e.g eth0)
	V4TunInterface         string                  `json:"v4TunInterface"`         // V4TunInterface  name for ipip encap (for healtchecks)
	V6TunInterface         string                  `json:"v6TunInterface"`         // V6TunInterface  name for ip(6)ip6 encap (for healthchecks)
	BalancerProgPath       string                  `json:"balancerProgPath"`       // BalancerProgPath  path to bpf prog for balancer
	HealthcheckingProgPath string                  `json:"healthcheckingProgPath"` // HealthcheckingProgPath path to bpf prog for healthchecking
	DefaultMac             []int                   `json:"defaultMac"`             // DefaultMac defaultMac mac address of default router
	Priority               int                     `json:"priority"`               //Priority tc priority of healtchecking task
	RootMapPath            string                  `json:"rootMapPath"`            // RootMapPath rootMapPath path to pinned map from root xdp prog
	RootMapPos             int                     `json:"rootMapPos"`             // RootMapPos position inside rootMap
	EnableHc               bool                    `json:"enableHc"`               // EnableHc flag, if set - we will load healthchecking bpf prog
	TunnelBasedHCEncap     bool                    `json:"tunnelBasedHCEncap"`     // TunnelBasedHCEncap flag, if set - katran will redirect packets to v4TunInterface and v6TunInterface to encap v4 and v6 packets respectively using the bpf prog to healthcheck backend reals.
	DisableForwarding      bool                    `json:"disableForwarding"`      // DisableForwarding  flag - if set, we don't load the forwarding (xdp) bpf program
	MaxVips                int                     `json:"maxVips"`                // MaxVips  maximum allowed vips to configure
	MaxReals               int                     `json:"maxReals"`               // MaxReals maximum allowed reals to configure
	ChRingSize             int                     `json:"chRingSize"`             // ChRingSize size of ch ring for each real
	Testing                bool                    `json:"testing"`                // Testing flag, if true - don't program forwarding
	LruSize                int64                   `json:"lruSize"`                // LruSize size of connection table
	ForwardingCores        []int                   `json:"forwardingCores"`        // ForwardingCores  responsible for forwarding
	NumaNodes              []int                   `json:"numaNodes"`              // NumaNodes  mapping of cores to NUMA nodes
	MaxLpmSrcSize          int                     `json:"maxLpmSrcSize"`          // MaxLpmSrcSize maximum size of map for src based routing
	MaxDecapDst            int                     `json:"maxDecapDst"`            // MaxDecapDst maximum number of destinations for inline decap
	HcInterface            string                  `json:"hcInterface"`            // HcInterface interface where we want to attach hc bpf prog
	XdpAttachFlags         int                     `json:"xdpAttachFlags"`         // XdpAttachFlags
	MonitorConfig          int                     `json:"monitorConfig"`          // MonitorConfig for katran introspection
	MemlockUnlimited       bool                    `json:"memlockUnlimited"`       // MemlockUnlimited  should katran set memlock to unlimited by default
	KatranSrcV4            string                  `json:"katranSrcV4"`            // KatranSrcV4  string ipv4 source address for GUE packets
	KatranSrcV6            string                  `json:"katranSrcV6"`            // KatranSrcV6 string ipv6 source address for GUE packets
	LocalMac               []int                   `json:"localMac"`               // LocalMac localMac mac address of local server
	FlowDebug              bool                    `json:"flowDebug"`              // FlowDebug hashFunction to create hash ring
	HashFunction           ch_helpers.HashFunction `json:"hashFunction"`           // HashFunction if set, creates and populates extra debugging maps
}

//KatranMonitorStats struct which contains stats from katran monitor
type KatranMonitorStats struct {
	Limit      int `json:"limit"`  //Limit of packet writer. how many packets we would write before we stop
	Amount     int `json:"amount"` //Amount of packets which has been written so far
	BufferFull int `json:"bufferFull"`
}

// KatranBpfMapStats generic bpf map stats
type KatranBpfMapStats struct {
	MaxEntries     int `json:"maxEntries"`     //MaxEntries size of the bpf map in the kernel
	CurrentEntries int `json:"currentEntries"` //CurrentEntries number of entries we are managing
}

// KatranLbStats generic userspace related stats to track internals of katran library
// such as number of failed bpf syscalls (could happens if we are trying to add
// to many vips etc)
type KatranLbStats struct {
	BpfFailedCalls       int `json:"bpfFailedCalls"`       // BpfFailedCalls number of failed syscalls
	AddrValidationFailed int `json:"addrValidationFailed"` // AddrValidationFailed times provided ipaddress was invalid
}

// HealthCheckProgStats  struct to record packet level counters for events in health-check program
// NOTE: this must be kept in sync with 'hc_stats' in healthchecking_ipip.c
type HealthCheckProgStats struct {
	PacketsProcessed int `json:"packetsProcessed"` // PacketsProcessed  number of packets processed for the healthcheck prog
	PacketsDropped   int `json:"packetsDropped"`   // PacketsDropped total number of packets dropped
	PacketsSkipped   int `json:"packetsSkipped"`   // PacketsSkipped total number of packets without action taken
	PacketsTooBig    int `json:"packetsTooBig"`    // PacketsTooBig total number of packets larger than prespecified max size for a packet
}

//
type KatranFeatures struct {
	SrcRouting                bool `json:"srcRouting"`                //SrcRouting  flag which indicates that source based routing feature has been enabled/compiled in bpf forwarding plane
	InlineDecap               bool `json:"inlineDecap"`               //InlineDecap flag which indicates that inline decapsulation feature has been enabled/compiled in bpf forwarding plane
	Introspection             bool `json:"introspection"`             //Introspection flag which indicates that katran introspection is enabled
	GueEncap                  bool `json:"gueEncap"`                  //GueEncap flag which indicates that GUE instead of IPIP should be used
	DirectHealthchecking      bool `json:"directHealthchecking"`      //DirectHealthchecking flag which indicates that hc encapsulation would be directly created instead of using tunnel interfaces
	LocalDeliveryOptimization bool `json:"localDeliveryOptimization"` //LocalDeliveryOptimization flag which indicates that local delivery would be optimized by passing (xdp_pass) local traffic
	FlowDebug                 bool `json:"flowDebug"`
}

type VipKey struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Proto   int    `json:"proto"`
}
