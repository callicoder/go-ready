package snowflake

import (
	"encoding/binary"
	"errors"
	"net"
	"sync"
	"time"
)

const (
	// Custom Epoch (Thursday, November 8, 2018 7:17:50 AM, UTC)
	epoch        int64 = 1541661470000
	totalBits          = 64
	epochBits          = 42
	nodeIdBits         = 10
	sequenceBits       = 12

	maxNodeId   = -1 ^ (-1 << nodeIdBits)
	maxSequence = -1 ^ (-1 << sequenceBits)
)

// SnowFlake is a structure which holds snowflake-specific data.
type SnowFlake struct {
	lastTimestamp uint64
	sequence      uint32
	nodeId        uint32
	lock          sync.Mutex
}

// NewSnowFlake initializes the generator.
func NewSnowFlake() (*SnowFlake, error) {
	nodeId, err := createNodeId()

	if err != nil {
		return nil, err
	}

	return &SnowFlake{nodeId: nodeId}, nil
}

// Create Unique node id
func createNodeId() (uint32, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}
	return (binary.BigEndian.Uint32(ip) & nodeIdBits), nil
}

func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("No private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

// Get current timestamp in milliseconds, adjust for the custom epoch.
func timestamp() uint64 {
	return uint64(time.Now().UnixNano()/int64(1000000) - epoch)
}

// Next generates the next unique ID.
func (sf *SnowFlake) Next() uint64 {
	sf.lock.Lock()
	defer sf.lock.Unlock()

	ts := timestamp()
	if ts < sf.lastTimestamp {
		panic("invalid system clock")
	}

	if ts == sf.lastTimestamp {
		sf.sequence = (sf.sequence + 1) & maxSequence
		if sf.sequence == 0 {
			// Sequence Exhausted, wait till next millisecond.
			ts = sf.waitNextMilli(ts)
		}
	} else {
		// reset sequence to start with zero for the next millisecond
		sf.sequence = 0
	}

	sf.lastTimestamp = ts
	return sf.pack()
}

// Pack bits into a snowflake value.
func (sf *SnowFlake) pack() uint64 {
	return (sf.lastTimestamp << (totalBits - epochBits)) |
		(uint64(sf.nodeId) << (totalBits - epochBits - nodeIdBits)) |
		(uint64(sf.sequence))
}

// Sequence exhausted. Wait till the next millisecond.
func (sf *SnowFlake) waitNextMilli(ts uint64) uint64 {
	for ts == sf.lastTimestamp {
		time.Sleep(100 * time.Microsecond)
		ts = timestamp()
	}
	return ts
}
