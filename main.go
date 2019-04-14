package main

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"os"
	"strconv"
	"sync"
	"syscall"
)

var EndFlag = 0
var DeviceSoc int32
var Timeout = 500

type IPv4Flag uint8
type IPv4 struct {
	Version    uint8
	IHL        uint8
	TOS        uint8
	Length     uint16
	Id         uint16
	Flags      IPv4Flag
	FragOffset uint16
	TTL        uint8
	Protocol   uint8
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
			fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
			if err != nil {
				panic(err)
			}
			defer syscall.Close(fd)

			file := os.NewFile(uintptr(fd), "")

			//var Nready int
			// https://forum.golangbridge.org/t/unix-poll-help/6834

			for EndFlag == 0 {
				buffer := make([]byte, 1024)
				_, _ = file.Read(buffer)
				IpHeaderDecode(buffer)

				//EtherRecv(buffer[:num])
			}

			// 処理終了のお知らせ
			<- ch
			wg.Done()
		}(strconv.Itoa(i))
	}

	wg.Wait()
}

// IPヘッダー解析
func IpHeaderDecode(IpBuff []byte) {
	var Ip IPv4
	Ip.Version = IpBuff[0] >> 4
	Ip.IHL = IpBuff[0] & 0x0F
	Ip.TOS = IpBuff[1]
	Ip.Length = binary.BigEndian.Uint16(IpBuff[2:4])
	Ip.Id = binary.BigEndian.Uint16(IpBuff[4:6])
	Ip.Flags = IPv4Flag(binary.BigEndian.Uint16(IpBuff[6:8]) >> 13)
	Ip.FragOffset = binary.BigEndian.Uint16(IpBuff[6:8]) & 0x1FFF
	Ip.TTL = IpBuff[8]
	Ip.Protocol = IpBuff[9]
	Ip.Checksum = binary.BigEndian.Uint16(IpBuff[10:12])
	Ip.SrcIP = IpBuff[12:16]
	Ip.DstIP = IpBuff[16:20]
	Ip.Options = Ip.Options[:0]
	Ip.Padding = nil
	fmt.Println(Ip)
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
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		panic(err)
	}
	//var ifReq int

	//fmt.Println(unix.IoctlSetInt(fd, unix.SIOCGIFFLAGS, ifReq))
	fmt.Println(fd, unix.SIOCGIFFLAGS)
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
