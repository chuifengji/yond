package compile

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func Transform() {
	result := api.Transform("let x: number = 1", api.TransformOptions{
		Loader: api.LoaderTS,
	})

	if len(result.Errors) == 0 {
		fmt.Printf("%s", result.Code)
	}
}

func Build() {
	ioutil.WriteFile("in.ts", []byte("let x: number = 1"), 0644)

	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"in.ts"},
		Outfile:     "out.js",
		Write:       true,
	})

	if len(result.Errors) > 0 {
		os.Exit(1)
	}
}

func BuildWithTreeThaking() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"app.js"},
		Bundle:      true,
		TreeShaking: api.TreeShakingIgnoreAnnotations,
	})

	if len(result.Errors) > 0 {
		os.Exit(1)
	}
}
