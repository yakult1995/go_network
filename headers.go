package main

import "net"

var EthernetProtocol = map[uint16]string{
	2048:  "IPv4",
	34525: "IPv6",
	2054:  "ARP",
}

type Ethernet struct {
	DstMac []byte
	SrcMac []byte
	Type   string
	Data   []byte
}

var IPProtocol = map[int]string{
	1:  "ICMP",
	6:  "TCP",
	17: "UDP",
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

type ARP struct {
	HardwareType   uint16
	ProtocolType   uint16
	HardwareLength uint8
	ProtocolLength uint8
	Operation      uint16
	SrcMacAddress  []byte
	SrcIP          net.IP
	DstMacAddress  []byte
	DstIP          net.IP
}

type ICMP struct {
	Type     uint8
	Code     uint8
	Checksum uint16
	Length   uint8
	Data     []byte
}

type TCP struct {
	SrcPort               uint16
	DstPort               uint16
	SequenceNum           uint32
	AcknowledgementNumber uint32
	HeaderLength          uint8
	Reserved              uint8
	URG                   int
	ACK                   int
	PSH                   int
	RST                   int
	SYN                   int
	FIN                   int
	WindowSize            uint16
	CheckSum              uint16
	UrgentPointer         uint16
	Options               []byte
	Padding               []byte
}
