package precompile

import (
	"math"
	"strings"
	"yond/file"
	"yond/load"
	"yond/util"
)

// var result []string

type TargetFile struct {
	Path   string
	Source string
}

func GetModuleName(path string) string {
	temp := strings.Replace(util.Re_module_name.FindString(path), ".", "_", -1)
	moduleName := strings.Replace(temp, "/", "_", -1)
	return moduleName
}

func GetSources(paths []string) []TargetFile {
	var targetFiles []TargetFile
	for _, path := range paths {
		var targetFile TargetFile
		targetFile.Path = path
		targetFile.Source = file.ReadFile(path)
		targetFiles = append(targetFiles, targetFile)
	}
	return targetFiles
}

func getAbsPath(rootPath string, filePath string, moduleName string) string {
	prefix := ``
	temp := filePath[len(rootPath):]
	count := int(math.Floor(float64(strings.Count(temp, "\\")) / 2))
	for i := 0; i < count; i++ {
		prefix += `../`
	}
	return prefix + "modules/" + moduleName + ".js"
}

func TransformHttpImport(sources []TargetFile) {

	for i := 0; i < len(sources); i++ {
		var content = sources[i].Source
		var filePath = sources[i].Path
		import_matches := util.Re_http_import.FindAllString(content, -1)
		//var http_matches []string
		for _, value := range import_matches {
			//exp. 'https://cdnjs.cloudflare.com/ajax/libs/axios/0.21.1/axios.min.js'
			var httpUrl = util.Re_between_quotation.FindString(value)
			//regexp 不支持断言和回溯。。。。
			httpUrl = strings.Replace(strings.Replace(httpUrl, `'`, "", -1), `"`, "", -1)
			var moduleName = GetModuleName(httpUrl)
			//replace url to moduleName
			content = strings.Replace(content, httpUrl, getAbsPath(util.RootPath, filePath, moduleName), -1)
			//distribute a http task
			load.DistributeHttpTask(moduleName, httpUrl)
			//http_matches = append(http_matches, httpUrl)
		}
		//writeFile
		file.WriteFile(sources[i].Path, content)
		// result = append(result, http_matches...)
	}

}
