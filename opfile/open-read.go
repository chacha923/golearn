package opfile

import (
	"os"
	"github.com/golang/glog"
	"bufio"
	"fmt"
)

func OpenFileRead(){
	file, err := os.OpenFile("file.txt", os.O_RDONLY,0)
	if err != nil {
		glog.Fatal(err.Error())
	}
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err.Error() != "EOF"{
			glog.Fatal(err.Error())
		}

		if err != nil && err.Error() == "EOF" {
			glog.Info("end of file")
			return
		}
		fmt.Println(string(line))

	}


}