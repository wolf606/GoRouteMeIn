package main

import (
	addresser "address/Addressing"
	net "address/Network"
	"fmt"
	"strconv"
)

func main() {
	var net1 = net.NewNetwork(224)
	var net2 = net.NewNetwork(13)
	var net3 = net.NewNetwork(1424)

	var networks = []net.Network{*net1, *net2, *net3}
	networks = addresser.AddressList(networks)

	fmt.Println("Addressing")

	for index, element := range networks {
		fmt.Println("Network "+strconv.Itoa(index+1)+": \n",
			"	Hosts: "+strconv.FormatInt(int64(element.GetHosts()), 10),
			" \n	Network IP: "+element.Ip.StrIp(),
			" \n	Mask: "+element.Mask.StrIp())
	}
}
