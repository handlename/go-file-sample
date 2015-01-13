package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/k0kubun/pp"
)

var errorOccured = false

func main() {
	if !doMain() {
		os.Exit(1)
	}
}

func doMain() bool {
	// temporaryなディレクトリを作る
	tempDirPath, err := ioutil.TempDir("", "")
	if err != nil {
		fmt.Println("error on create temp dir:", err)
		return false
	}

	// 消すのはioutil.TempDirを呼び出した側の責任
	// 今回はファイルに書かれた内容を確認したいので消さない
	// defer func() {
	// 	err := os.Remove(tempDirPath)
	// 	if err != nil {
	// 		fmt.Println("error on remove temp dir:", err)
	// 	}
	// }()

	filePath := tempDirPath + "/create.txt"

	// ファイルを作る
	file, err := os.Create(filePath)
	// ファイルを開く
	// file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error on Open:", err)
		return false
	}
	defer file.Close()

	pp.Println(file)

	// ファイルに書き込む
	_, err = file.WriteString("hello\n")
	if err != nil {
		fmt.Println("error on write to file:", err)
		return false
	}

	// // NewFile
	// file2 := os.NewFile(file.Fd(), tempDirPath+"/new_file.txt")
	// defer file2.Close()
	// file2.Sync()

	return true
}
