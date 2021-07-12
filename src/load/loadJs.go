package load

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"xx/file"
	"xx/util"
)

//实际上仅仅在第一次初始化时用，后面判断应该从 js.mod 文件中读取
var setTemp = make(map[string]bool)

func DistributeHttpTask(moduleName string, url string) {
	exists := setTemp[moduleName] || util.SetModuleUrl[url]
	if !exists {
		createHttpTask(moduleName, url)
		setTemp[moduleName] = true
		//main.SetModuleUrl[url] = true
	}
}

func createHttpTask(moduleName string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	addToModules(moduleName, string(body), url)
}

func addToModules(moduleName string, content string, url string) {
	if !file.Exists("../test/temp/modules") {
		file.CreateDir("../test/temp/modules")
	}
	if !file.Exists("../test/temp/modules/" + moduleName + ".js") {
		file.WriteFile("../test/temp/modules/"+moduleName+".js", content)
		file.AppendToJsMod(url)
	}

}
