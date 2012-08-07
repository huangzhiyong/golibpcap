// Copyright 2012 The golibpcap Authors. All rights reserved.                      
// Use of this source code is governed by a BSD-style                              
// license that can be found in the LICENSE file.

package pkt

import (
	"fmt"
	"strconv"
	"strings"
)

// The point of a HttpHdr is to ease the mapping of application level logs to
// pcap traces.  Capturing the actual HTTP headers has not been implemented to
// keep this fast.
//
// Since HTTP requests and responses can span multiple packets this is not
// perfect and as such your millage may vary.
type HttpHdr struct {
	Proto      string // e.g. "HTTP/1.0"
	Method     string // GET, POST, PUT, etc.
	RequestURI string // The unmodified Request-URI
	StatusCode int64  // e.g. 200
	Status     string // e.g. "OK"
}

// Given the payload of a transport layer packet NewHttpHdr will return a
// *HttpHdr struct is the bytes contain a valid HTTP header, or nil otherwise.
func NewHttpHdr(b []byte) *HttpHdr {
	h := &HttpHdr{}
	s := string(b)
	ln := strings.SplitN(s, "\r\n", 2)
	if len(ln) < 2 {
		return nil
	}
	lnp := strings.SplitN(ln[0], " ", 3)
	if len(lnp) < 3 {
		return nil
	}

	m := strings.ToUpper(lnp[0])

	switch m {
	case "GET", "POST", "HEAD", "PUT", "DELETE", "TRACE", "OPTIONS", "CONNECT", "PATCH":
		h.Method = m
		h.RequestURI = lnp[1]
		p := strings.ToUpper(lnp[2])
		if strings.HasPrefix(p, "HTTP/") {
			h.Proto = p
		}
	default:
		if strings.HasPrefix(m, "HTTP/") {
			h.Proto = m
		} else {
			return nil
		}
		var err error
		h.StatusCode, err = strconv.ParseInt(lnp[1], 10, 0)
		if err != nil {
			return nil
		}
		h.Status = lnp[2]
	}
	//NOTE: If we want to implement parsing the HTTP headers and the
	// payload then we will need to make sure that this method fails fast so
	// that we don't spend too much time trying to reconstruct a incomplete
	// HTTP request.
	return h
}

// JsonElement returns a JSON encoding of the HttpHdr struct.
func (h *HttpHdr) JsonElement() string {
	if h.StatusCode == 0 {
		return fmt.Sprintf("\"httphdr\":{\"proto\":\"%s\",\"method\":\"%s\",\"request\":\"%s\"}",
			h.Proto,
			h.Method,
			h.RequestURI)
	}
	return fmt.Sprintf("\"httphdr\":{\"proto\":\"%s\",\"statusCode\":%d}",
		h.Proto,
		h.StatusCode)
}

// String returns a minimal encoding of the HttpHdr struct.
func (h *HttpHdr) String() string {
	if h.StatusCode == 0 {
		return fmt.Sprintf("%s %s %s",
			h.Proto,
			h.Method,
			h.RequestURI)
	}
	return fmt.Sprintf("%s %d",
		h.Proto,
		h.StatusCode)
}