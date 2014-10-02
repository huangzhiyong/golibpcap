// Copyright 2014 The golibpcap Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <netinet/tcp.h>
#include <netinet/udp.h>

uint16_t tcphdr_source(struct tcphdr*);
uint16_t tcphdr_dest(struct tcphdr*);
uint32_t tcphdr_seq(struct tcphdr*);
uint32_t tcphdr_ack_seq(struct tcphdr*);

uint16_t tcphdr_source_ntohs(struct tcphdr*);
uint16_t tcphdr_dest_ntohs(struct tcphdr*);
uint32_t tcphdr_seq_ntohl(struct tcphdr*);
uint32_t tcphdr_ack_seq_ntohl(struct tcphdr*);
uint16_t tcphdr_window_ntohs(struct tcphdr*);
uint16_t tcphdr_check_ntohs(struct tcphdr*);
uint16_t tcphdr_urg_ptr_ntohs(struct tcphdr*);

uint16_t udphdr_source_ntohs(struct udphdr*);
uint16_t udphdr_dest_ntohs(struct udphdr*);
uint16_t udphdr_check_ntohs(struct udphdr*);
uint16_t udphdr_len_ntohs(struct udphdr*);
