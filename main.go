package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	SetDefaultParam()
	IsTargetIpAddr("192.168.0.100")
	IsSameSubnet("192.168.0.100")

	ch := make(chan int, 2)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		ch <- 1
		wg.Add(1)
		go func(index string) {
			stdin := bufio.NewScanner(os.Stdin)
			for stdin.Scan() {
				fmt.Println(index, " : ", stdin.Text())
				break
			}
			<-ch
			wg.Done()
		}(strconv.Itoa(i))
	}

	wg.Wait()
	//EndFlag := 0
	//var DeviceSoc int
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
func EtherRecv() {}

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
