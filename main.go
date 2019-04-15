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

			//var Nready int
			// https://forum.golangbridge.org/t/unix-poll-help/6834

			for EndFlag == 0 {
				fmt.Println()
				buffer := make([]byte, 1024)
				num, err := file.Read(buffer)
				if err != nil {
					fmt.Println("Buffer")
					panic(err)
				}
				//IpHeaderDecode(buffer, num)
				//EtherRecv(buffer[:num])
				EthernetFrameDecode(buffer, num)
			}

			// 処理終了のお知らせ
			<- ch
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
	EthernetFrame.Type = EtherProtocol[binary.BigEndian.Uint16(EthernetFrameBuff[12:14])]
	fmt.Println("EthernetFrame : ", EthernetFrame)
	IpHeaderDecode(EthernetFrameBuff[14:], num - 14)
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
	Ip.Protocol = Protcol[int(IpBuff[9])]
	Ip.Checksum = binary.BigEndian.Uint16(IpBuff[10:12])
	Ip.SrcIP = IpBuff[12:16]
	Ip.DstIP = IpBuff[16:20]
	Ip.Options = Ip.Options[:0]
	Ip.Padding = nil
	fmt.Println("IP : ", Ip)

	IcmpBuff := IpBuff[20:]
	var Icmp ICMP
	Icmp.Type = IcmpBuff[0]
	Icmp.Code = IcmpBuff[1]
	Icmp.Checksum = binary.BigEndian.Uint16(IcmpBuff[2:4])
	Icmp.Length = IcmpBuff[5]
	Icmp.Data = IcmpBuff[8:num]
	fmt.Println("ICMP : ", Icmp)
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
func MyEthThread() {
	//var nready int
}

// イーサネットフレーム受信処理
func EtherRecv(buff []byte) {
	fmt.Println(buff)
}

// ARPパケット受信処理
func ArpRecv() {}

// ターゲットIPアドレスの判定
func IsTragetAddr() {}

// ARPテーブルへの追加
func ArpAddTable() {}

// ARPパケットの送信
func ArpSend() {}

// イーサネットフレーム送信
func EtherSend() {}

// IPパケット受信処理
func IpRecv() {}

// IP受信バッファへの追加
func IpRecvBufAdd() {}

// ICMPパケット受信処理
func IcmpRecv() {}

// ICMPエコーリプライパケットの送信
func IcmpSendEchoReply() {}

//
