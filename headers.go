package main

import "net"

type Ethernet struct {
	DstMac []byte
	SrcMac []byte
	Type uint16
	Data []byte
}

var Protcol = map[int] string {
	1 : "ICMP",
	6 : "TCP",
	17 : "UDP",
}

type IPv4Flag uint8
type IPv4 struct {
	Version    uint8
	IHL        uint8
	TOS        uint8
	Length     uint16
	Id         uint16
	Flags      IPv4Flag
	FlagOffset uint16
	TTL        uint8
	Protocol   string
	Checksum   uint16
	SrcIP      net.IP
	DstIP      net.IP
	Options    []IPv4Option
	Padding    []byte
}
type IPv4Option struct {
	OptionType   uint8
	OptionLength uint8
	OptionData   []byte
}

type ICMP struct {
	Type uint8
	Code uint8
	Checksum uint16
	Length uint8
	Data []byte
}
