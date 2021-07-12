package funcs

import (
	"yond/file"
	"yond/precompile"
)

func Init() {

}

func BuildDev(rootPath string) {
	file.CloneDir(rootPath)
	matchPaths := file.GetFiles(rootPath + "\\dev")
	targetFiles := precompile.GetSources(matchPaths)
	precompile.TransformHttpImport(targetFiles)
}
