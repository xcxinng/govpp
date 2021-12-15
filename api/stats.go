//  Copyright (c) 2019 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package api

// StatsProvider provides methods for retrieving statistics.
type StatsProvider interface {
	GetSystemStats(*SystemStats) error
	GetNodeStats(*NodeStats) error
	GetInterfaceStats(*InterfaceStats) error
	GetErrorStats(*ErrorStats) error
	GetBufferStats(*BufferStats) error
}

// SystemStats represents global system statistics.
type SystemStats struct {
	VectorRate          uint64
	NumWorkerThreads    uint64
	VectorRatePerWorker []uint64
	InputRate           uint64
	LastUpdate          uint64
	LastStatsClear      uint64
	Heartbeat           uint64
}

// NodeStats represents per node statistics.
type NodeStats struct {
	Nodes []NodeCounters
}

// NodeCounters represents node counters.
type NodeCounters struct {
	NodeIndex uint32
	NodeName  string // requires VPP 19.04+

	Clocks   uint64
	Vectors  uint64
	Calls    uint64
	Suspends uint64
}

// InterfaceStats represents per interface statistics.
type InterfaceStats struct {
	Interfaces []InterfaceCounters
}

// InterfaceCounters represents interface counters.
type InterfaceCounters struct {
	InterfaceIndex uint32 `json:"interface_index"`
	InterfaceName  string `json:"interface_name"` // requires VPP 19.04+

	Rx InterfaceCounterCombined `json:"rx"`
	Tx InterfaceCounterCombined `json:"tx"`

	RxErrors uint64 `json:"rx_errors"`
	TxErrors uint64 `json:"tx_errors"`

	RxUnicast   InterfaceCounterCombined `json:"rx_unicast"`
	RxMulticast InterfaceCounterCombined `json:"rx_multicast"`
	RxBroadcast InterfaceCounterCombined `json:"rx_broadcast"`
	TxUnicast   InterfaceCounterCombined `json:"tx_unicast"`
	TxMulticast InterfaceCounterCombined `json:"tx_multicast"`
	TxBroadcast InterfaceCounterCombined `json:"tx_broadcast"`

	Drops   uint64 `json:"drops"`
	Punts   uint64 `json:"punts"`
	IP4     uint64 `json:"ip4"`
	IP6     uint64 `json:"ip6"`
	RxNoBuf uint64 `json:"rx_no_buf"`
	RxMiss  uint64 `json:"rx_miss"`
	Mpls    uint64 `json:"mpls"`
}

// InterfaceCounterCombined defines combined counters for interfaces.
type InterfaceCounterCombined struct {
	Packets uint64 `json:"packets"`
	Bytes   uint64 `json:"bytes"`
}

// ErrorStats represents statistics per error counter.
type ErrorStats struct {
	Errors []ErrorCounter
}

// ErrorCounter represents error counter.
type ErrorCounter struct {
	CounterName string

	Value uint64
}

// BufferStats represents statistics per buffer pool.
type BufferStats struct {
	Buffer map[string]BufferPool
}

// BufferPool represents buffer pool.
type BufferPool struct {
	PoolName string

	Cached    float64
	Used      float64
	Available float64
}
