package main

import (
    "fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
    handle, err := pcap.OpenOffline(os.Args[1])
    if err != nil {
        panic(err)
    }
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

    res := map[int]int{}

	// TODO: This could be parallelized, but it's fast enough already
    for packet := range packetSource.Packets() {
		content := packet.TransportLayer().LayerPayload()
		for _, b := range content {
			res[int(b)] += 1
		}
    }

	fp, err := os.Create("/tmp/plot.data")
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	for i := 0; i < 256; i++ {
		//TODO: This should be a buffer, but it's fast enough already
		_, err := fp.WriteString(fmt.Sprintf("%d\t%d\n", i, res[i]))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("set title \"Packet Payload Histogram for " + os.Args[1] + "\"")
	fmt.Println("set xlabel \"Byte Values\"")
	fmt.Println("set ylabel \"Frequency\"")
	fmt.Println("set autoscale")
	fmt.Println("set terminal png")
	fmt.Println("set output \"output.png\"")
	fmt.Println("set yrange [0:*]")
	fmt.Println("set xrange [0:255]")
	fmt.Println("set format x \"%02x\"")
	fmt.Println("set nokey")
	fmt.Println("plot \"/tmp/plot.data\" lt 1")
	fmt.Println("quit")
}
