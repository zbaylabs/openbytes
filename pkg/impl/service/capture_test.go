package service

import (
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/google/gopacket/pcap"
	log "github.com/sirupsen/logrus"
)

func TestDevice(t *testing.T) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	// 打印设备信息
	fmt.Println("Devices found:")
	for _, device := range devices {
		if len(device.Addresses) != 0 {
			fmt.Println("\nName: ", device.Name)
			fmt.Println("Description: ", device.Description)
			fmt.Println("Devices flag: ", device.Flags)
			// if device.Flags &net.FlagLoopback != 0 && device.Flags&net.FlagUp != 0 {
			// 	//return &ifi
			// 	fmt.Println("loopback", device.Name, device.Flags)
			// }
			for _, address := range device.Addresses {
				fmt.Println("- IP address: ", address.IP)
				fmt.Println("- Subnet mask: ", address.Netmask)
				fmt.Println("- Broadaddr:  ", address.Broadaddr)
				fmt.Println("- P2P:  ", address.P2P)
			}
		}
		//fmt.Println("\nName: ", device.Name)
		//fmt.Println("Description: ", device.Description)
		//fmt.Println("Devices flag: ", device.Flags)
		// if device.Flags &net.FlagLoopback != 0 && device.Flags&net.FlagUp != 0 {
		// 	//return &ifi
		// 	fmt.Println("loopback", device.Name, device.Flags)
		// }
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
			fmt.Println("- Broadaddr:  ", address.Broadaddr)
			fmt.Println("- P2P:  ", address.P2P)
		}
	}

	t.Error("done")
}

func TestDevice2(t *testing.T) {
	// devices, err := pcap.FindAllDevs()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // 打印设备信息
	// fmt.Println("Devices found:")
	// for _, device := range devices {

	// 	fmt.Println("\nName: ", device.Name)
	// 	fmt.Println("Description: ", device.Description)
	// 	fmt.Println("Devices flag: ", device.Flags)
	// 	for _, address := range device.Addresses {
	// 		fmt.Println("- IP address: ", address.IP)
	// 		fmt.Println("- Subnet mask: ", address.Netmask)
	// 		fmt.Println("- Broadaddr:  ", address.Broadaddr)
	// 		fmt.Println("- P2P:  ", address.P2P)
	// 	}
	// }
	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	fmt.Println("Available network interfaces on this machine : ")
	for _, ifi := range interfaces {
		//fmt.Printf("Name : %v \n", i.Name)
		//fmt.Printf("Full object : %+v \n", i)
		//if net.FlagLoopback
		//if add,_:=i.Addrs();add!=nil && add[0].
		if ifi.Flags&net.FlagLoopback != 0 && ifi.Flags&net.FlagUp != 0 {
			//return &ifi
			fmt.Println("loopback", ifi.Name, ifi.Flags)
		}
		fmt.Println(ifi.Name, ifi.Flags)
	}
	t.Error("done")
}
