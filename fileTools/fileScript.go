package fileTools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	//"regexp"
	cregex "github.com/mingrammer/commonregex"
)

// var fileName = "C:\\Users\\13042\\Desktop\\"
var files = [...]string{
	"C:\\Users\\13042\\Desktop\\1",
	"C:\\Users\\13042\\Desktop\\2",
	"C:\\Users\\13042\\Desktop\\3"}

func RegularMatchFile() {
	ipSet := make(map[string]int)

	for i := 0; i < len(files); i++ {

		// 读取文件内容
		file, err := os.Open(files[i])
		defer file.Close()
		if err != nil {
			fmt.Println("打开文件失败：", err)
			return
		}

		scanner := bufio.NewScanner(file)

		// 正则匹配IP地址
		//re := regexp.MustCompile("((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)")
		re := cregex.IPv4Regex

		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindAllString(line, -1)
			for _, match := range matches {
				ipSet[match] = ipSet[match] + 1
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("读取文件失败：", err)
			return
		}
	}

	// 输出到新文件
	outputFile, err := os.OpenFile("C:\\Users\\13042\\Desktop\\output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return
	}
	defer outputFile.Close()

	_, err01 := outputFile.WriteString("\n")
	for key := range ipSet {
		_, err02 := outputFile.WriteString(key + "\t" + strconv.Itoa(ipSet[key]) + "\n")
		if err01 != nil || err02 != nil {
			fmt.Println("写入文件失败：", err)
			return
		}
	}

	fmt.Println("去重后的IP地址已输出到output.txt文件中")
}
