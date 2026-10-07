package scheme

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func QuickSort(arr []float64, start int, end int) {
	if start >= end {
		return
	}
	pivot := arr[end]
	i := start - 1
	j := start
	for j < end {
		if arr[j] <= pivot {
			i++
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
		j++
	}
	arr[end] = arr[i+1]
	arr[i+1] = pivot
	QuickSort(arr, start, i)
	QuickSort(arr, i+2, end)
}

func Hash(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func Ping(address string) (string, error) {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return "", errors.New(err.Error())
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	message := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   1234,
			Seq:  1,
			Data: []byte("hello"),
		},
	}
	bytes, err := message.Marshal(nil)
	if err != nil {
		return "", err
	}
	dst := &net.IPAddr{
		IP: net.ParseIP(address),
	}
	//measure time before sending
	start := time.Now()

	_, err = conn.WriteTo(bytes, dst)
	if err != nil {
		return "", err
	}
	buffer := make([]byte, 1500)
	n, _, err := conn.ReadFrom(buffer)
	if err != nil {
		return "", err
	}
	rtt := time.Since(start)

	reply, err := icmp.ParseMessage(1, buffer[:n])
	if err != nil {
		return "", err
	}
	if reply.Type != ipv4.ICMPTypeEchoReply {
		return "", errors.New("recieved non-echo-reply ICMP packet")
	}
	body, ok := reply.Body.(*icmp.Echo)
	if !ok {
		return "", errors.New("Error reading body of response")
	}
	if body.ID != 1234 || body.Seq != 1 {
		return "", errors.New("recieved unrelated ICMP reply")
	}
	return fmt.Sprintf("Reply from %s: time=%v", address, rtt), nil

}
