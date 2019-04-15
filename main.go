package main

import (
	"encoding/binary"
	"fmt"
	_ "golang.org/x/sys/unix"
	"os"
	"strconv"
	"sync"
	"syscall"
)

var EndFlag = 0
var DeviceSoc int32
var Timeout = 500

// https://github.com/fridolin-koch/grnvs/blob/master/ndp/ndisc.go
func Htons(n uint16) uint16 {
	var (
		high = n >> 8
		ret  = n<<8 + high
	)
	return ret
}

func main() {
	SetDefaultParam()
	IsTargetIpAddr("192.168.0.100")
	IsSameSubnet("192.168.0.100")

	// 同時実行スレッド数
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}

	// 特に指定がなければ無限ループ
	for i := 0; i < 3; i++ {
		ch <- 1
		wg.Add(1)
		go func(index string) {
			// これでPreAmbleは入ってこない。ちなみにmacOSでは`AF_PACKET`, `ETH_P_IP`が認識されないので動きませぬ。
			fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(Htons(syscall.ETH_P_ALL)))
			if err != nil {
				fmt.Println("Socket")
				panic(err)
			}
			defer syscall.Close(fd)

			file := os.NewFile(uintptr(fd), "")

			for EndFlag == 0 {
				fmt.Println()
				buffer := make([]byte, 1024)
				num, err := file.Read(buffer)
				if err != nil {
					fmt.Println("Buffer")
					panic(err)
				}

				// Ethernetを全て監視しているのでとりあえずEthernetFrame解析
				EthernetFrameDecode(buffer, num)
			}

			// 処理終了のお知らせ
			<-ch
			wg.Done()
		}(strconv.Itoa(i))
	}

	wg.Wait()
}

// イーサネットフレーム解析
// LAN内通信だとイーサネットフレーム付かないのか！
func EthernetFrameDecode(EthernetFrameBuff []byte, num int) {
	fmt.Println("EthernetFrameDecode()")
	var EthernetFrame Ethernet
	EthernetFrame.DstMac = EthernetFrameBuff[0:6]
	EthernetFrame.SrcMac = EthernetFrameBuff[6:12]
	EthernetFrame.Type = EthernetProtocol[binary.BigEndian.Uint16(EthernetFrameBuff[12:14])]
	fmt.Println("EthernetFrame : ", EthernetFrame)

	// プロトコル別に場合分け
	if EthernetFrame.Type == "IPv4" {
		IpHeaderDecode(EthernetFrameBuff[14:], num-14)
	} else if EthernetFrame.Type == "ARP" {
		ArpHeaderDecode(EthernetFrameBuff[14:], num-14)
	}
}

// IPヘッダー解析
func IpHeaderDecode(IpBuff []byte, num int) {
	fmt.Println("IpHeaderDecode()")
	var Ip IPv4
	Ip.Version = IpBuff[0] >> 4
	Ip.IHL = IpBuff[0] & 0x0F
	Ip.TOS = IpBuff[1]
	Ip.Length = binary.BigEndian.Uint16(IpBuff[2:4])
	Ip.Id = binary.BigEndian.Uint16(IpBuff[4:6])
	Ip.Flags = IPv4Flag(binary.BigEndian.Uint16(IpBuff[6:8]) >> 13)
	Ip.FlagOffset = binary.BigEndian.Uint16(IpBuff[6:8]) & 0x1FFF
	Ip.TTL = IpBuff[8]
	Ip.Protocol = IPProtocol[int(IpBuff[9])]
	Ip.Checksum = binary.BigEndian.Uint16(IpBuff[10:12])
	Ip.SrcIP = IpBuff[12:16]
	Ip.DstIP = IpBuff[16:20]
	Ip.Options = Ip.Options[:0]
	Ip.Padding = nil
	fmt.Println("IP : ", Ip)

	// プロトコル別に場合分け
	if Ip.Protocol == "ICMP" {
		IcmpDecode(IpBuff[20:], num-20)
	} else if Ip.Protocol == "TCP" {
		TcpDecode(IpBuff[20:], num-20)
	}
}

// TCP解析
func TcpDecode(TcpBuff []byte, num int) {
	fmt.Println("TcpDecode()")
	var TcpHeader TCP
	TcpHeader.SrcPort = binary.BigEndian.Uint16(TcpBuff[0:2])
	TcpHeader.DstPort = binary.BigEndian.Uint16(TcpBuff[2:4])
	TcpHeader.SequenceNum = binary.BigEndian.Uint32(TcpBuff[4:8])
	TcpHeader.AcknowledgementNumber = binary.BigEndian.Uint32(TcpBuff[8:12])
	TcpHeader.HeaderLength = TcpBuff[12] >> 4
	TcpHeader.Reserved = 0
	TcpHeader.URG = int(binary.BigEndian.Uint16(TcpBuff[12:14]) >> 5 & 0x01)
	TcpHeader.ACK = int(binary.BigEndian.Uint16(TcpBuff[12:14]) >> 4 & 0x01)
	TcpHeader.PSH = int(TcpBuff[14] >> 3 & 0x01)
	TcpHeader.RST = int(TcpBuff[14] >> 2 & 0x01)
	TcpHeader.SYN = int(TcpBuff[14] >> 1 & 0x01)
	TcpHeader.FIN = int(TcpBuff[14] & 0x01)
	TcpHeader.WindowSize = binary.BigEndian.Uint16(TcpBuff[14:16])
	TcpHeader.CheckSum = binary.BigEndian.Uint16(TcpBuff[16:18])
	TcpHeader.UrgentPointer = binary.BigEndian.Uint16(TcpBuff[18:20])
	fmt.Println("TCP Header : ", TcpHeader)
}

// ICMP解析
func IcmpDecode(IcmpBuff []byte, num int) {
	fmt.Println("IcmpDecode()")
	var Icmp ICMP
	Icmp.Type = IcmpBuff[0]
	Icmp.Code = IcmpBuff[1]
	Icmp.Checksum = binary.BigEndian.Uint16(IcmpBuff[2:4])
	Icmp.Length = IcmpBuff[5]
	//Icmp.Data = IcmpBuff[8:num]
	fmt.Println("ICMP : ", Icmp)
}

// ARPヘッダー解析
func ArpHeaderDecode(ArpBuff []byte, num int) {
	fmt.Println("ArpHeaderDecode()")
	var Arp ARP
	Arp.HardwareType = binary.BigEndian.Uint16(ArpBuff[0:2])
	Arp.ProtocolType = binary.BigEndian.Uint16(ArpBuff[2:4])
	Arp.HardwareLength = ArpBuff[4]
	Arp.ProtocolLength = ArpBuff[5]
	Arp.Operation = binary.BigEndian.Uint16(ArpBuff[6:8])
	Arp.SrcMacAddress = ArpBuff[8:14]
	Arp.SrcIP = ArpBuff[14:18]
	Arp.DstMacAddress = ArpBuff[18:24]
	Arp.DstIP = ArpBuff[24:28]
	fmt.Println("ARP : ", Arp)
}

// IP受信バッファの初期化
func IpRecvBufInit() {}

// ソケット初期化
func InitSocket() {}

// インターフェース初期化
func ShowIfreq() {}

// MACアドレス調査
func GetMacAddress() {}

// 終了シグナルハンドラ
func SigTerm() {}

// 送受信スレッド
func MyEthThread() {}

// ターゲットIPアドレスの判定
func IsTragetAddr() {}

// ARPテーブルへの追加
func ArpAddTable() {}

// ARPパケットの送信
func ArpSend() {}

// イーサネットフレーム送信
func EtherSend() {}

// IP受信バッファへの追加
func IpRecvBufAdd() {}

// ICMPエコーリプライパケットの送信
func IcmpSendEchoReply() {}
