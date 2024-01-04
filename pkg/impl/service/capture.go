package service

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	log "github.com/sirupsen/logrus"
	pb "github.com/zbaylab/openbytes/api/go"
)

type CapturesImpl struct {
	pb.UnimplementedCapturesServer
}

func (s *CapturesImpl) List(in *pb.Capture, stream pb.Captures_ListServer) error {
	// Open iface
	handle, err := pcap.OpenLive("en0", 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer handle.Close()
	source := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range source.Packets() {
		vPacket := &pb.Packet{Src: &pb.Address{}, Dst: &pb.Address{}, Data: packet.String()}

		for _, v := range packet.Layers() {
			vPacket.Layers = vPacket.Layers + v.LayerType().String() + ","
		}
		if ethernetLayer := packet.Layer(layers.LayerTypeEthernet); ethernetLayer != nil {
			ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
			vPacket.Src.Mac = ethernetPacket.SrcMAC.String()
			vPacket.Dst.Mac = ethernetPacket.DstMAC.String()
		}
		if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			vPacket.Src.Ip = ip.SrcIP.String()
			vPacket.Dst.Ip = ip.DstIP.String()
			vPacket.Protocol = ip.Protocol.String()
		}
		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			vPacket.Src.Port = int32(tcp.SrcPort)
			vPacket.Dst.Port = int32(tcp.DstPort)
		}
		if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
			udp, _ := udpLayer.(*layers.UDP)
			vPacket.Src.Port = int32(udp.SrcPort)
			vPacket.Dst.Port = int32(udp.DstPort)
		}
		if appLayer := packet.ApplicationLayer(); appLayer != nil {
			vPacket.Payload = packet.ApplicationLayer().Payload()
		}

		if err := stream.Send(vPacket); err != nil {
			log.Errorln(err)
			return err
		}
	}

	return nil
}
