package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	log "github.com/sirupsen/logrus"
	pb "github.com/zbaylab/openbytes/api/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type CapturesImpl struct {
	pb.UnimplementedCapturesServer
}

func (s *CapturesImpl) Device(context.Context, *emptypb.Empty) (*structpb.ListValue, error) {
	lv := &structpb.ListValue{}
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	for _, device := range devices {
		if len(device.Addresses) > 0 && !strings.HasPrefix(device.Name, "lo") {
			lv.Values = append(lv.Values, structpb.NewStringValue(device.Name))
		}
	}
	return lv, nil
}

func (s *CapturesImpl) List(in *pb.Capture, stream pb.Captures_ListServer) error {
	// Open iface
	handle, err := pcap.OpenLive(in.Iface, 65535, true, pcap.BlockForever)
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(in.Filter); err != nil {
		log.Errorln(err)
		return err
	}
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

var pointChan = make(chan pb.Point, 100)

func init() {
	go fetch()
}

func fetch() {
	handle, err := pcap.OpenLive("en0", 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	source := gopacket.NewPacketSource(handle, handle.LinkType())
	ticker := time.NewTicker(time.Second * 1)
	tcpCount, udpCount, otherCount := 0, 0, 0
	comeIn := false
	for packet := range source.Packets() {
		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			tcpCount++
		} else if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
			udpCount++
		} else {
			otherCount++
		}
		go func() {
			if !comeIn {
				for _ = range ticker.C {
					//println("test")
					pointChan <- pb.Point{
						Label:      time.Now().Format(time.TimeOnly),
						TcpCount:   int32(tcpCount),
						UdpCount:   int32(udpCount),
						OtherCount: int32(otherCount),
					}
					tcpCount = 0
					udpCount = 0
					otherCount = 0
					comeIn = false
				}
			}
		}()
	}
}

func (s *CapturesImpl) Traffic(in *pb.Capture, stream pb.Captures_TrafficServer) error {
	for v := range pointChan {
		if err := stream.Send(&v); err != nil {
			log.Errorln(err)
			return err
		}
	}

	return nil
}

func (s *CapturesImpl) Copy(ctx context.Context, request *pb.CopyRequest) (*wrapperspb.StringValue, error) {
	fmt.Println(request.String())
	return nil, status.Errorf(codes.Unimplemented, "method Copy not implemented")
}
