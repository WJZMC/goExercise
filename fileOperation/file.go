package fileOperation

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func FileExercesie() {

	//writeRepaceFile("./bin/text.text")
	//appendFileContent("./bin/text.text")
	//appendFileContent("./bin/text1.text")

	//read := bufio.NewReader(os.Stdin) //从键盘接收输入
	//str, err := read.ReadString('\n') //接收字符串类型的文本以'\n'结束
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//writeToFile("./bin/text.text", str)

	//copyFileToPath("./bin/text.text", "./bin/copy.text")

	//f , err := os.Create("./file.txt")
	//if err != nil{
	//	fmt.Println("create Err",err)
	//	return
	//}
	//defer  f.Close()
	//
	//f.WriteString("I am jack ")
	//
	//fread , err := os.Open("/Users/jackwei/Desktop/未命名.mov")
	//if err != nil{
	//	fmt.Println("open Err",err)
	//	return
	//}
	//defer fread.Close()
	//
	//f, err := os.Create("/Users/jackwei/Desktop/test.mov")
	////f, err := os.OpenFile("/Users/jackwei/Desktop/",os.O_RDWR,0666)
	//if err != nil{
	//	fmt.Println("open Err",err)
	//	return
	//}
	//defer f.Close()
	//
	////offset,_ := f.Seek(0,io.SeekEnd)
	////f.WriteAt([]byte("\nhello word"),offset)
	//
	//var buf []byte = make([]byte,1024)
	//reader := bufio.NewReader(fread)
	//for{
	//	buf,err = reader.ReadBytes('\n')//如果没有'\n'将会出现漏掉信息
	//	f.Write(buf)
	//	if err ==io.EOF{
	//		fmt.Println(err)
	//		break
	//	}
	//}
	////for{
	////	n,err := fread.Read(buf)
	////	if err==io.EOF{
	////		break
	////	}
	////	f.Write(buf[:n])
	////}

	//findPath := "/Users/jackwei/Desktop/code/goCode/src/studyProject"
	//suffix := ".go"
	//指定目录检索特定文件：
	//从用户给出的目录中，找出所有的 .jpg 文件。
	//result := findFileWithPathAndSuffix(findPath,suffix)
	//for _,v := range result{
	//	fmt.Println(v)
	//}

	//指定目录拷贝特定文件：
	//从用户给出的目录中，拷贝 .mp3文件到指定目录中。
	//aimPath := "/Users/jackwei/Desktop"
	//result := findFileWithPathAndSuffix(findPath, suffix)
	//for _, v := range result {
	//	copyFileToPath(findPath+"/"+v, aimPath+"/"+v)
	//}

	//统计指定目录内单词出现次数：
	//统计指定目录下，所有.txt文件中，“Love”这个单词 出现的次数。
	//sum := 0
	//result := findFileWithPathAndSuffix(findPath, suffix)
	//for _, v := range result {
	//	count := sumFileWordCount(findPath+"/"+v, "main")
	//	fmt.Println(v, ":", count)
	//	sum += count
	//}
	//fmt.Println(sum)

}

//根据路径查找路径中的指定后缀的文件，并返回对应的文件名字
func findFileWithPathAndSuffix(filePath, suffix string) []string {
	result := make([]string, 0)

	path, err := os.OpenFile(filePath, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("open :", err)
		return result
	}
	defer path.Close()
	fileSlice, err := path.Readdir(-1)
	if err != nil {
		fmt.Println("read :", err)
		return result
	}
	for _, file := range fileSlice {
		if strings.HasSuffix(file.Name(), suffix) {
			result = append(result, file.Name())
		}
	}
	return result

}
func copyFileToPath(fromPath, toPath string) {
	fP, err := os.Open(fromPath)
	if err != nil {
		fmt.Println("open\t", err)
		return
	}
	defer fP.Close()
	toP, err := os.Create(toPath)
	if err != nil {
		fmt.Println("create\t", err)
		return
	}
	defer toP.Close()

	buf := make([]byte, 4096)
	for {
		n, err := fP.Read(buf)
		if err == io.EOF {
			break
		}
		toP.Write(buf[:n])
	}

}
func sumFileWordCount(filePath, word string) int {
	result := 0

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open", err)
		return result
	}
	defer f.Close()

	buf := make([]byte, 4096)
	reader := bufio.NewReader(f)
	for {
		buf, err = reader.ReadBytes('\n')
		str := string(buf)
		result += strings.Count(str, word)
		if err != nil && err == io.EOF {
			break
		}

	}
	return result

}

func writeRepaceFile(filePath string) {
	dFile, err := os.Create(filePath) //以覆盖的方式创建文件
	if err != nil {
		fmt.Println(err)
		return
	}
	read := bufio.NewReader(os.Stdin) //从键盘接收输入
	str, err := read.ReadString('\n') //接收字符串类型的文本以'\n'结束
	if err != nil {
		fmt.Println(err)
		return
	}
	//_, err = dFile.WriteString(str)//写入字符串
	_, err = dFile.Write([]byte(str)) //写入字符切片
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dFile.Close()
}

func appendFileContent(filePath string) {
	read := bufio.NewReader(os.Stdin)
	slice, err := read.ReadBytes('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	//已只写的方式打开文件
	dfile, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		//判断错误类型是否是文件存在的错误
		if os.IsExist(err) {
			fmt.Println(err)
			return
		} else {
			//文件不存在，创建文件
			dfile, err = os.Create(filePath)
			if os.IsExist(err) {
				fmt.Println(err)
				return
			} else {
				dfile.Write(slice)
			}
		}

	} else {
		//获取偏移量
		offset, err := dfile.Seek(0, io.SeekEnd)
		if err != nil {
			fmt.Println(err)
			return
		}
		//向指定偏移量的位置追加字符切片
		_, err = dfile.WriteAt(slice, offset)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	defer dfile.Close()
}

func writeToFile(filePath string, content string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 6)

	if err != nil {
		if !os.IsExist(err) {
			file, err = os.Create(filePath)
			_, err = file.WriteString(content)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println(err)
			return
		}
	}
	//_, err = file.WriteString(content)

	count, _ := file.Seek(0, io.SeekEnd)
	_, err = file.WriteAt([]byte(content), count)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

}
