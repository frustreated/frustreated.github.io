// main.go
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	htmlfiles := list.New()
	files, _ := ioutil.ReadDir("./")
	// url := list.New()
	for _, f := range files {
		if path.Ext(f.Name()) == ".html" {
			htmlfiles.PushBack(f.Name())
		}
	}

	filePath := "./README.md"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()
	//写入内容
	writer := bufio.NewWriter(file)
	writer.WriteString("双击bat-del.bat可删除更新历史")
	writer.WriteString("\n")
	writer.WriteString("双击bat.bat不删除更新历史")
	writer.WriteString("\n")
	for htmlf := htmlfiles.Front(); htmlf != nil; htmlf = htmlf.Next() {
		url := fmt.Sprintf("[%s](http://frustreated.github.io/%s)", htmlf.Value, htmlf.Value)
		writer.WriteString(url)
		writer.WriteString("\n")
		writer.Flush()
	}

}
