package quicksort

import net "address/Network"

func QuickSort(arr []net.Network) []net.Network {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less []net.Network
	var greater []net.Network
	for _, element := range arr[1:] {
		if element.GetHosts() >= pivot.GetHosts() {
			less = append(less, element)
		} else {
			greater = append(greater, element)
		}
	}
	less = QuickSort(less)
	greater = QuickSort(greater)
	return append(append(less, pivot), greater...)
}
