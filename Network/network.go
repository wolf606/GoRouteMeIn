package network

import (
	ip "address/IP"
)

type Network struct {
	Ip    ip.IPv4
	Mask  ip.IPv4
	hosts int
}

func NewNetwork(hosts int) *Network {
	return &Network{hosts: hosts}
}

func (net *Network) GetHosts() int {
	return net.hosts
}
