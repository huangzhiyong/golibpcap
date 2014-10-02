// Copyright 2014 The golibpcap Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <arpa/inet.h>
#include <netinet/tcp.h>
#include <netinet/udp.h>
#include "wrappers.h"

uint16_t tcphdr_source(struct tcphdr* h) {
  return h->source;
}

uint16_t tcphdr_dest(struct tcphdr* h) {
  return h->dest;
}

uint32_t tcphdr_seq(struct tcphdr* h) {
  return h->seq;
}

uint32_t tcphdr_ack_seq(struct tcphdr* h) {
  return h->ack_seq;
}

uint16_t tcphdr_source_ntohs(struct tcphdr* h) {
  return ntohs(h->source);
}

uint16_t tcphdr_dest_ntohs(struct tcphdr* h) {
  return ntohs(h->dest);
}

uint32_t tcphdr_seq_ntohl(struct tcphdr* h) {
  return ntohl(h->seq);
}

uint32_t tcphdr_ack_seq_ntohl(struct tcphdr* h) {
  return ntohl(h->ack_seq);
}

uint16_t tcphdr_window_ntohs(struct tcphdr* h) {
  return ntohs(h->window);
}

uint16_t tcphdr_check_ntohs(struct tcphdr* h) {
  return ntohs(h->check);
}

uint16_t tcphdr_urg_ptr_ntohs(struct tcphdr* h) {
  return ntohs(h->urg_ptr);
}

uint16_t udphdr_source_ntohs(struct udphdr* h) {
  return ntohs(h->source);
}

uint16_t udphdr_dest_ntohs(struct udphdr* h) {
  return ntohs(h->dest);
}

uint16_t udphdr_check_ntohs(struct udphdr* h) {
  return ntohs(h->check);
}

uint16_t udphdr_len_ntohs(struct udphdr* h) {
  return ntohs(h->len);
}
