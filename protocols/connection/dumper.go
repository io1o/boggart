package connection

import (
	"encoding/hex"
	"log"
)

type dumper struct {
	Conn
}

func NewDumper(conn Conn) Conn {
	return &dumper{
		Conn: conn,
	}
}

func (d *dumper) Read(p []byte) (n int, err error) {
	n, err = d.Conn.Read(p)
	log.Printf("Read <<< n: %d err: %v\n", n, err)

	if err == nil && n > 0 {
		log.Printf("\n%s\n", hex.Dump(p[:n]))
	}

	return n, err
}

func (d *dumper) Write(p []byte) (n int, err error) {
	log.Printf("Write >>> err: %v\n", err)
	log.Printf("\n%s\n", hex.Dump(p))

	return d.Conn.Write(p)
}

func (d *dumper) Invoke(request []byte) (response []byte, err error) {
	log.Println("Write >>>")
	log.Printf("\n%s\n", hex.Dump(request))

	if i, ok := d.Conn.(Invoker); ok {
		response, err = i.Invoke(request)
	}

	log.Printf("Read <<< err: %v\n", err)
	log.Printf("\n%s\n", hex.Dump(response))

	return response, err
}
