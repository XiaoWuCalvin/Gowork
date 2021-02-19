package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	/*
		1.递归实现遍历所有文件夹下子文件
	*/

	dirname := "D:/Go语言学习"
	listFiles(dirname,0)
}

//加层级
func listFiles(dirname string,level int) {
	//level用来记录当前递归的层次，生成带有层次感的空格
	s:="|--"
	for i:=0;i<level ;i++  {
		s="|	"+s
	}
	fileInfo, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	//遍历文件
	for _, fi := range fileInfo {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n",s, filename)
		if fi.IsDir() {
			//递归调用自己本身
			listFiles(filename,level+1)
		}
	}
}
