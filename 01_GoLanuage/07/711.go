package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	copylen, err := copyFile("dst.txt", "711.go")
	if err != nil {
		return
	} else {
		fmt.Println(copylen)
	}

}

func copyFile(dstName, srcName string) (copylen int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close() //当return 时就会调用src.Close() 把源文件关闭

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
