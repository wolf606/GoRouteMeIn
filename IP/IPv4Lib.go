package ipv4lib

import (
	"fmt"
	"strconv"
)

type IPv4 struct {
	Oct1 int64
	Oct2 int64
	Oct3 int64
	Oct4 int64
}

func NewIp(Oct1 int64, Oct2 int64, Oct3 int64, Oct4 int64) *IPv4 {
	octets := [4]int64{Oct1, Oct2, Oct3, Oct4}
	for index, element := range octets {
		if !check_octet(element) {
			fmt.Println("Invalid octet: ", index)
			return nil
		}
	}
	return &IPv4{Oct1, Oct2, Oct3, Oct4}
}

func check_octet(octet int64) bool {
	if octet >= 0 && octet <= 255 {
		return true
	}
	return false
}

func (ip *IPv4) StrIp() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.Oct1, ip.Oct2, ip.Oct3, ip.Oct4)
}

func CreateIpFromStr(str string) IPv4 {
	//fmt.Println("STR LENGTH: ", len(str))
	var oct1, err1 = strconv.ParseInt(str[0:8], 2, 64)
	var oct2, err2 = strconv.ParseInt(str[8:16], 2, 64)
	var oct3, err3 = strconv.ParseInt(str[16:24], 2, 64)
	var oct4, err4 = strconv.ParseInt(str[24:32], 2, 64)
	if !check_ip_binary_str_length(str) {
		fmt.Println(err1, err2, err3, err4)
	}
	return IPv4{oct1, oct2, oct3, oct4}
}

func check_ip_binary_str_length(str string) bool {
	if len(str) == 32 {
		return true
	}
	return false
}

func GetIpBinaryStr(ip IPv4) string {
	return fmt.Sprintf("%08b%08b%08b%08b", ip.Oct1, ip.Oct2, ip.Oct3, ip.Oct4)
}
