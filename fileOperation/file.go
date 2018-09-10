package fileOperation

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
func copyFileToPath(fromPath string, toPath string) {
	fromP, err := os.Open(fromPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fromP.Close()

	toP, err := os.Create(toPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer toP.Close()

	//按缓冲区大小方式读取
	buf := make([]byte, 1024)
	for {
		n, err := fromP.Read(buf)
		if err == io.EOF {
			break
		}
		toP.Write(buf[:n])
	}

	//buf := bufio.NewReader(fromP)
	//for {
	//按行读取，切片方式追加
	//bufSlice, err := buf.ReadBytes('\n')
	//offset, _ := toP.Seek(0, io.SeekEnd)
	//toP.WriteAt(bufSlice, offset)
	//if err == io.EOF {
	//	fmt.Println(err)
	//	break
	//}

	////按行读取，字符串方式方式添加
	//bufStr, err := buf.ReadString('\n')
	//toP.WriteString(bufStr)
	//if err == io.EOF {
	//	fmt.Println(err)
	//	break
	//}
	//}
}
