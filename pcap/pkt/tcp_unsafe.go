// Copyright 2013 The golibpcap Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !safe,!appengine

package pkt

/*
#include <netinet/in.h>
#include <netinet/tcp.h>
#include "wrappers.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// The TcpHdr struct is a wrapper for the tcphdr struct in <netinet/tcp.h>.
type TcpHdr struct {
	cptr    *C.struct_tcphdr // see <net/tcp.h> struct tcphdr
	Source  uint16           // source port
	Dest    uint16           // destination port
	Seq     uint32           // sequence number
	AckSeq  uint32           // acknowledgement number
	Doff    uint8            // The length of the TCP header (data offset) in 32 bit words.
	Flags   uint16           // TCP flags per RFC 793, September, 1981
	Window  uint16           // window advertisement
	Check   uint16           // checksum
	UrgPtr  uint16           // urgent pointer
	payload unsafe.Pointer
}

// With an unsafe.Pointer to the block of C memory NewTcpHdr returns a filled in TcpHdr struct.
func NewTcpHdr(p unsafe.Pointer) (*TcpHdr, unsafe.Pointer) {
	tcpHead := &TcpHdr{
		cptr: (*C.struct_tcphdr)(p),
		// Since cgo does not provide access to bit fields in a struct
		// we index 12 octets in and then shift out the unneeded bits.
		Doff: *(*byte)(unsafe.Pointer(uintptr(p) + uintptr(12))) >> 4,
	}
	tcpHead.Source = uint16(C.tcphdr_source_ntohs(tcpHead.cptr))
	tcpHead.Dest = uint16(C.tcphdr_dest_ntohs(tcpHead.cptr))
	tcpHead.Seq = uint32(C.tcphdr_seq_ntohl(tcpHead.cptr))
	tcpHead.AckSeq = uint32(C.tcphdr_ack_seq_ntohl(tcpHead.cptr))
	// A this time (and there are no plans to support it) cgo does not
	// provide access to bit fields in a struct so this is what we are stuck
	// with.  We index 12 octets in and then use a bit mask.
	tcpHead.Flags = uint16(C.ntohs(C.uint16_t(
		*(*uint16)(unsafe.Pointer(uintptr(p) + uintptr(12)))))) & uint16(0x01FF)
	tcpHead.Window = uint16(C.tcphdr_window_ntohs(tcpHead.cptr))
	tcpHead.Check = uint16(C.tcphdr_check_ntohs(tcpHead.cptr))
	tcpHead.UrgPtr = uint16(C.tcphdr_urg_ptr_ntohs(tcpHead.cptr))
	tcpHead.payload = unsafe.Pointer(uintptr(p) + uintptr(tcpHead.Doff*4))
	return tcpHead, tcpHead.payload
}

// GetPayloadBytes returns the bytes from the packet's payload.  This is a Go
// slice backed by the C bytes.  The result is that the Go slice uses very
// little extra memory.
func (h *TcpHdr) GetPayloadBytes(pl uint16) []byte {
	l := int(h.PayloadLen(pl))
	if l <= 0 {
		return []byte{}
	}
	var b []byte
	sh := (*reflect.SliceHeader)((unsafe.Pointer(&b)))
	sh.Cap = l
	sh.Len = l
	sh.Data = uintptr(h.payload)
	return b
}
