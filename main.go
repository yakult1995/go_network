package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	IPTTL   int    `toml:"IP-TTL"`
	MTU     int    `toml:"MTU"`
	Device  string `toml:"device"`
	Vmac    string `toml:"vmac"`
	Vip     string `toml:"vip"`
	Gateway string `toml:"gateway"`
}

func main() {
	// デフォルトパラメータ受け取り
	var config Config
	_, err := toml.DecodeFile("eth_conf.toml", &config)
	if err != nil {
		panic(err)
	}
	fmt.Println("パラメータ : ", config)

	EndFlag := 0
	var DeviceSoc int
}

// デフォルトパラメータのセット
func SetDefaultParam() {}

// パラメータの読み込み
func ReadParam() {}

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
