package opfile

import (
	"os"
	"github.com/golang/glog"
	"fmt"
)

func init() {

}
func OpenFileWrite(){
	file, err := os.OpenFile("file.txt",os.O_APPEND | os.O_RDWR,0)
	if err != nil {
		glog.Fatal(err.Error())
		return
	}
	n, err := file.WriteString("hello")
	if err != nil {
		glog.Fatal(err.Error())
		return
	}
	fmt.Println(n)
}



