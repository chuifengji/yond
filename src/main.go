//目前是同步链路版
package main

import (
	"xx/file"
	"xx/precompile"
	"xx/util"
)

func init() {
	content := file.ReadFile("../test/temp/js.mod")
	matches := util.Re_Js_Mod.FindAllString(content, -1)
	for _, value := range matches {
		util.SetModuleUrl[util.Re_between_quotation.FindString(value)] = true
	}
}

func main() {
	file.CloneDir(`../test`)
	matchPaths := file.GetFiles(`../test/temp`)
	// fmt.Println(matchPaths)
	targetFiles := precompile.GetSources(matchPaths)
	precompile.TransformHttpImport(targetFiles)
}
