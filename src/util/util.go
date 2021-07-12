package util

import "regexp"

var Re_http_import = regexp.MustCompile(`import(\s+\w+)(\sfrom\s+)(['"](http|https)://[\w.]+\/?\S*['"])`)
var Re_between_quotation = regexp.MustCompile(`['"](.*?)['"]`)
var Re_module_name = regexp.MustCompile(`/\w+[/@](\d+(.\d+)*)`)
var Re_Js_Mod = regexp.MustCompile(`require\s+(http|https):\/\/([\w.]+\/?)*`)
var SetModuleUrl = make(map[string]bool)
var RootPath string

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
