package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
	"unsafe"
)

func main() {
	conn, err := net.Dial("tcp4", ":7000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	start := time.Now()
	for i := 0; i < 150000; i++ {
		data := createmessage(MessageTypeText, "hello from client")
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("write error:", err)
		}
	}
	end := time.Since(start)
	fmt.Printf("duration time: %s", end)

	for {
		select {}
	}

}

const (
	MessageTypeJSON = 1
	MessageTypeText = 2
	MessageTypeXML  = 3
)

/*
0 1 2 3 | 4 5 6 7 | 8 N+
uint32  | uint32  | string
type    | length  | data
*/
func createmessage(mtype int, data string) []byte {
	buf := make([]byte, 512)
	binary.LittleEndian.PutUint32(buf[0:], uint32(mtype))
	binary.LittleEndian.PutUint32(buf[4:], uint32(len(data)))
	copy(buf[8:], []byte(data))
	return buf
}

func readmessage(data []byte) (mtype, mlen uint32, msg string) {
	mtype = binary.LittleEndian.Uint32(data[0:])
	mlen = binary.LittleEndian.Uint32(data[4:])
	//msg = string(data[8:])
	msgptr := data[8:]
	msg = *(*string)(unsafe.Pointer(&msgptr))
	return mtype, mlen, msg
}
