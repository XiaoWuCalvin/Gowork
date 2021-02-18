package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
			断点续传：
			  文件传递：本质上也就是文件复制

		      D:/Go语言学习/go相关/files/guliang.jpg

			  复制到当前工程下
			思路：边复制，边记录复制的总量

	*/
	srcFile := "D:/Go语言学习/go相关/files/guliang.jpg"
	destFile := srcFile[strings.LastIndex(srcFile, "/")+1:] //获取文件名称  guliang.jpg
	fmt.Println(destFile)
	tempFile := destFile + "temp.txt"
	fmt.Println(tempFile) //guliang.jpgtemp.txt

	file1, err := os.Open(srcFile)
	HandErr(err)
	file2, err := os.OpenFile(destFile,os.O_CREATE| os.O_RDWR, os.ModePerm)
	HandErr(err)
	file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	HandErr(err)

	defer file1.Close()
	defer file2.Close()

	//step1:先读取临时文件中的数据，再seek
	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n1, err := file3.Read(bs)
	//HandErr(err) //第一次在读取的时候没有任何文件数据，不需要处理

	countStr := string(bs[:n1])
	count, err := strconv.ParseInt(countStr, 10, 64)
	//HandErr(err) //第一次在读取的时候无任何数据，所有count为空字符串
	fmt.Println(count)

	//step2:设置读写的位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	data := make([]byte, 100, 100)
	n2 := -1 //读取的数据量
	n3 := -1 //写出的数据量
	total := int(count)

	//step3:复制文件
	for {
		n2, err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件复制完毕")
			file3.Close()
			os.Remove(tempFile)
			break
		}
		n3, err = file2.Write(data[:n2])
		total += n3
		fmt.Println(total)
		//将复制的总量，存储到临时文件中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

		fmt.Printf("total:%d\n",total)

		//假装程序断电
		//if total>1194600{
		//	panic("程序断点遇到突发异常。。。")
		//}

	}
}

func HandErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
