// Copyright 2012 The golibpcap Authors. All rights reserved.                      
// Use of this source code is governed by a BSD-style                              
// license that can be found in the LICENSE file.

package pkt

/*
#include <net/ethernet.h>
#include <netinet/ether.h>
#include <netinet/in.h>
*/
import "C"
import (
	"fmt"
	"net"
	"unsafe"
)

// The EthHdr struct is a wrapper for the ether_header struct in <net/ethernet.h>.
type EthHdr struct {
	cptr      *C.struct_ether_header // C pointer to ether_header
	SrcAddr   net.HardwareAddr       // the sender's MAC address
	DstAddr   net.HardwareAddr       // the receiver's MAC address
	EtherType uint16                 // packet type ID field
	payload   unsafe.Pointer
}

// With an unsafe.Pointer to the block of C memory NewEthHdr returns a filled in EthHdr struct.
func NewEthHdr(p unsafe.Pointer) (*EthHdr, unsafe.Pointer) {
	ethHdr := &EthHdr{
		cptr:    (*C.struct_ether_header)(p),
		payload: unsafe.Pointer(uintptr(p) + uintptr(C.ETHER_HDR_LEN)),
	}
	ethHdr.SrcAddr = net.HardwareAddr(C.GoBytes(unsafe.Pointer(&ethHdr.cptr.ether_shost), C.ETH_ALEN))
	ethHdr.DstAddr = net.HardwareAddr(C.GoBytes(unsafe.Pointer(&ethHdr.cptr.ether_dhost), C.ETH_ALEN))
	ethHdr.EtherType = uint16(C.ntohs(C.uint16_t(ethHdr.cptr.ether_type)))
	return ethHdr, ethHdr.payload
}

// JsonElement returns a JSON encoding of the EthHdr struct.
func (h *EthHdr) JsonElement() string {
	return fmt.Sprintf("\"ether_header\":{\"ether_shost\":\"%s\",\"ether_dhost\":\"%s\",\"ether_type\":0x%x}",
		h.SrcAddr.String(),
		h.DstAddr.String(),
		h.EtherType)
}

// String returns a minimal encoding of the EthHdr struct.
func (h *EthHdr) String() string {
	return fmt.Sprintf("%s->%s 0x%x",
		h.SrcAddr.String(),
		h.DstAddr.String(),
		h.EtherType)
}