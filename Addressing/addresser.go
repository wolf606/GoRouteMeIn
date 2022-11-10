package addresser

import (
	bin "address/Binary"
	ip "address/IP"
	net "address/Network"
	quicksort "address/QuickSort"
	"math"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var ip_template map[string]ip.IPv4 = map[string]ip.IPv4{
	"A": ip.IPv4{int64(r.Intn(127 - 0)), 0, 0, 0},
	"B": ip.IPv4{int64(r.Intn(191-128) + 128), int64(r.Intn(255 - 0)), 0, 0},
	"C": ip.IPv4{int64(r.Intn(223-192) + 192), int64(r.Intn(255 - 0)), int64(r.Intn(255 - 0)), 0},
}

var ip_class_hosts map[string]int = map[string]int{
	"A": 16777216,
	"B": 65536,
	"C": 256,
}

func AddressList(arr []net.Network) []net.Network {
	arr = set_first_ip(arr)
	var reserved_bits = calculate_reserved_bits(arr[0].GetHosts())
	arr[0].Mask = ip.CreateIpFromStr(bin.Ones[0:32-reserved_bits] + bin.Zeros[0:reserved_bits])
	for i := 1; i < len(arr); i++ {
		var prev_ip = ip.GetIpBinaryStr(arr[i-1].Ip)
		var new_ip = bin.Sum(prev_ip, "1"+bin.Zeros[0:reserved_bits])

		arr[i].Ip = ip.CreateIpFromStr(new_ip)
		reserved_bits = calculate_reserved_bits(arr[i].GetHosts())
		arr[i].Mask = ip.CreateIpFromStr(bin.Ones[0:32-reserved_bits] + bin.Zeros[0:reserved_bits])
	}
	return arr
}

func set_first_ip(arr []net.Network) []net.Network {
	var networks = quicksort.QuickSort(arr)
	var max_hosts = networks[0].GetHosts()
	var ip_class = select_ip_class(max_hosts)
	networks[0].Ip = ip_template[ip_class]

	return networks
}

func select_ip_class(hosts int) string {
	if hosts > ip_class_hosts["A"]-2 {
		return "A"
	} else if hosts > ip_class_hosts["C"]-2 && hosts < ip_class_hosts["B"]-1 {
		return "B"
	} else if hosts > 0 && hosts < ip_class_hosts["C"]-1 {
		return "C"
	}
	return ""
}

func calculate_reserved_bits(hosts int) int {
	for i := 2; i < 25; i++ {
		if float64(hosts) < (math.Pow(2, float64(i)))-2 {
			return i
		}
	}
	return 0
}
