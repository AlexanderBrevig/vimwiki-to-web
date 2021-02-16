package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	vwtw "github.com/AlexanderBrevig/vimwiki-to-web/pkg"
	"github.com/AlexanderBrevig/vimwiki-to-web/pkg/gemini"
)

var vimwikiPath string
var filePath string
var geminiOutPath string

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	wikihome := filepath.Join(usr.HomeDir, "vimwiki")
	defaultOut := filepath.Join(usr.HomeDir, "vimwiki-gemini")
	flag.StringVar(&vimwikiPath, "vimwiki", wikihome, "Path to your vimwiki")
	flag.StringVar(&filePath, "file", "/index.md", "Path to your file")
	flag.StringVar(&geminiOutPath, "gmiout", defaultOut, "Path to your output folder")
}

func renderToFile(path string) {
	file := filepath.Join(vimwikiPath, path)
	mdcontent, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	gmirend := gemini.NewRenderer(
		gemini.RendererOptions{
			VimWikiRoot: vimwikiPath,
		},
	)
	gmi := vwtw.ToGEMINI([]byte(mdcontent), nil, gmirend)

	//fmt.Println(string(gmi))
	filePath = strings.Replace(path, ".md", ".gmi", -1)
	outPath := filepath.Join(geminiOutPath, filePath)
	os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
	err = ioutil.WriteFile(outPath, gmi, 0644)
	fmt.Println(path, " -> ", outPath)
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	target := filepath.Join(vimwikiPath, filePath)
	info, err := os.Stat(target)
	if err != nil {
		panic(err)
	}
	if info.IsDir() {
		err := filepath.Walk(target,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && strings.Contains(path, ".md") {
					renderToFile(strings.Replace(path, vimwikiPath, "", -1))
				}
				return nil
			})
		if err != nil {
			panic(err)
		}
	} else {
		renderToFile(filePath)
	}
}
