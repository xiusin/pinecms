package annotations

import (
	"fmt"
	"testing"

	"github.com/MarcGrol/golangAnnotations/generator"
	"github.com/MarcGrol/golangAnnotations/generator/rest"
	"github.com/MarcGrol/golangAnnotations/parser"
)

var excludeMatchPattern = "^" + generator.GenfilePrefix + ".*.go$"
var inputDir = "/Users/xiusin/projects/src/github.com/xiusin/pinecms/annotations"




func Test_runAll(t *testing.T) {

	parsedSources, err := parser.New().ParseSourceDir(inputDir, "^.*.go$", excludeMatchPattern)

	fmt.Println(parsedSources)

	//registry := annotation.NewRegistry(restAnnotation.Get())

	//a, ok := registry.ResolveAnnotation(`// @RestOperation( Method = "GET", path = "/person/:uid" )`)
	//
	//fmt.Println(a, ok)

	if err != nil {
		panic(err)
	}
	if err := rest.NewGenerator().Generate(inputDir, parsedSources); err != nil {
			panic(err)
	}

}
