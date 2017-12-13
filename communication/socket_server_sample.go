package communication

import (
	"net"
	"log"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {

	}
	defer l.Close()
	log.Println("listen ok")

	var i int
	for {
		time.Sleep(time.Second * 10)
		conn, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			return
		}
		i++
		log.Printf("%d: accept a new connection\n", i)
		bytes := make([]byte, 0)
		conn.Read(bytes)
	}
}

