package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"syscall"
)

var EndFlag = 0
var DeviceSoc int32
var Timeout = 500

func main() {
	SetDefaultParam()
	IsTargetIpAddr("192.168.0.100")
	IsSameSubnet("192.168.0.100")

	// 同時実行スレッド数
	ch := make(chan int, 2)
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
				num, _ := file.Read(buffer)

				//fmt.Printf("% X\n", buffer[:num])
				EtherRecv(buffer[:num])
			}

			// 処理終了のお知らせ
			<- ch
			wg.Done()
		}(strconv.Itoa(i))
	}

	wg.Wait()
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
