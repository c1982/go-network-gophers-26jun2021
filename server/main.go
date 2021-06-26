package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/signal"
	"unsafe"
)

func main() {
	ls, err := net.Listen("tcp4", ":7000")
	if err != nil {
		panic(err)
	}

	fmt.Println("connection ready!")

	go func() {
		for {
			conn, err := ls.Accept()
			if err != nil {
				fmt.Println("connection error:", err)
				continue
			}
			go handler(conn)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func handler(conn net.Conn) {
	fmt.Println("connection accepted:", conn.RemoteAddr().String())

	for {
		header := make([]byte, 512)
		_, err := conn.Read(header[:])
		if err != nil {
			fmt.Println("read error:", err)
			conn.Close()
			break
		}

		_, _, _ = readmessage(header)
		/*
			mlen := binary.LittleEndian.Uint32(header[4:])
			databuf := make([]byte, mlen)
			_, err = conn.Read(databuf[:])
			if err != nil {
				fmt.Println("read error:", err)
				conn.Close()
				break
			}

			var messagebuf []byte
			messagebuf = append(messagebuf, header...)
			messagebuf = append(messagebuf, databuf...)
			_, _, _ = readmessage(messagebuf)

			//fmt.Printf("type: %d, len: %d, msg: %s\n", mtype, mlen, msg)
		*/
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
	buf := make([]byte, 4+4+len(data))
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
