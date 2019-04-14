package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"net"
)

type Config struct {
	MyMac   int    `toml:"MyMac"`
	MyIp    string `toml:"MyIP"`
	IpTTL   int    `toml:"IP-TTL"`
	MTU     int    `toml:"MTU"`
	Device  string `toml:"device"`
	Vmac    string `toml:"vmac"`
	Vip     string `toml:"vip"`
	Vmask   string `toml:"vmask"`
	Gateway string `toml:"gateway"`
}

var DefaultPingSixe = 64
var DummyWaitMs = 100
var RetryCount = 3
var Param Config

// デフォルトパラメータのセット
func SetDefaultParam() int {
	// デフォルトパラメータ受け取り
	_, err := toml.DecodeFile("eth_conf.toml", &Param)
	if err != nil {
		panic(err)
	}
	fmt.Println("パラメータ : ", Param)
	return 0
}

// パラメータの読み込み
func ReadParam() int {
	return 0
}

// 仮想IPとの一致判定
func IsTargetIpAddr(Addr string) int {
	fmt.Println("IsTargetIpAddr() : ", Param.Vip == Addr)
	if net.ParseIP(Param.Vip).String() == Addr {
		return 1
	}
	return 0
}

// 同一サブネットかの判定
func IsSameSubnet(Addr string) int {
	// maskの作成
	Mask := net.IPMask(net.ParseIP(Param.Vmask).To4())
	Network := net.ParseIP(Param.Vip).Mask(Mask).String()
	CheckIpNetwork := net.ParseIP(Addr).Mask(Mask).String()
	fmt.Println("IsSameSubnet() : ", Network == CheckIpNetwork)
	if Network == CheckIpNetwork {
		return 1
	}
	return 0
}
