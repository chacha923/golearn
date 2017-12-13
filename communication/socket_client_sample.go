package communication

import "net"

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		//handle error
	}
	// read or write on conn


}
