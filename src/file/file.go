package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"xx/util"

	"github.com/otiai10/copy"
)

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func ReadFile(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}

func CreateDir(path string) {
	os.Mkdir(path, os.ModePerm)
}

func WriteFile(path string, content string) {
	data := []byte(content)
	if ioutil.WriteFile(path, data, 0666) == nil {
		fmt.Println("write successfully:", path)
	}
}

func CloneDir(path string) {
	fmt.Println(path)
	if Exists(path + "\\dev") {
		return
	}
	err := copy.Copy(path+"\\src", path+"\\dev")
	if err != nil {
		log.Println(err)
	}
}

func GetFiles(rootPath string) []string {
	fmt.Println(rootPath)
	var result []string
	usefulSuffix := []string{".js", ".ts"}
	err := filepath.Walk(rootPath, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		var filesuffix = path.Ext(filePath)
		if util.Find(usefulSuffix, filesuffix) {
			result = append(result, filePath)
		}
		if err != nil {
			log.Println(err)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	return result
}

func AppendToJsMod(url string) {
	filePath := util.RootPath + "/js.mod"
	if !Exists(filePath) {
		_, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString("require " + `"` + url + `"` + "\n")
	//Flush将缓存的文件真正写入到文件中
	write.Flush()

}
