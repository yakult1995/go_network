package main

func main() {
	SetDefaultParam()
	IsSameSubnet("192.168.0.100")

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
