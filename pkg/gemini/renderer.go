package gemini

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	ast "github.com/gomarkdown/markdown/ast"
)

type RendererOptions struct {
	VimWikiRoot string
	Debug       bool
}

type Renderer struct {
	opts RendererOptions
}

func NewRenderer(opts RendererOptions) *Renderer {
	return &Renderer{
		opts: opts,
	}
}

func (r *Renderer) debugNode(node ast.Node) {
	if r.opts.Debug {
		cont := node.AsContainer()
		if cont != nil && len(cont.Literal) > 0 {
			fmt.Printf("%T\t%s\n", node, cont.Literal)
		}
		leaf := node.AsLeaf()
		if leaf != nil && len(leaf.Literal) > 0 {
			fmt.Printf("%T\t%s\n", node, leaf.Literal)
		}
	}
}

func (r *Renderer) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	switch node := node.(type) {
	case *ast.Heading:
		r.debugNode(node)
		if entering {
			level := node.Level
			if level > 3 {
				level = 3
			}
			w.Write([]byte(strings.Repeat("#", level)))
			w.Write([]byte(" "))
			w.Write(node.Literal)
		} else {
			w.Write([]byte("\n\n"))
		}
	case *ast.List:
		r.debugNode(node)
		if !entering {
			w.Write([]byte("\n"))
		}
	case *ast.Paragraph:
		r.debugNode(node)
		if !entering {
			w.Write([]byte("\n\n"))
		}
	case *ast.CodeBlock:
		r.debugNode(node)
		if entering {
			w.Write([]byte("\n```\n"))
			w.Write(node.Literal)
			w.Write([]byte("\n```\n\n"))
		}
	case *ast.Code:
		r.debugNode(node)
		if entering {
			w.Write([]byte("`"))
			w.Write(node.Literal)
			w.Write([]byte("`"))
		}

	case *ast.BlockQuote:
		r.debugNode(node)
		if entering {
			w.Write([]byte("> "))
		} else {
			w.Write([]byte("\n"))
		}
	case *ast.Link:
		r.debugNode(node)
		if entering {
			w.Write([]byte("=> "))
			destination := string(node.Destination)
			destination = strings.Replace(destination, ".md", ".gmi", -1)
			w.Write([]byte(destination))
			w.Write([]byte(" "))
		}
	case *ast.Image:
		r.debugNode(node)
		if entering {
			w.Write([]byte("=> "))
			w.Write(node.Destination)
			w.Write([]byte(" "))
		}
	case *ast.Text:
		r.debugNode(node)
		if entering {
			_, isLink := node.Parent.(*ast.Link)
			_, isList := node.Parent.GetParent().(*ast.ListItem)
			_, isInclude := node.Parent.(*ast.Del)
			if !isLink && isList && len(node.Literal) > 0 {
				// Handle output of wikitasks
				if !strings.Contains(string(node.Literal), "[ ]") {
					w.Write([]byte("* "))
				}
				w.Write([]byte(node.Literal))
			} else if isInclude {
				files, err := filepath.Glob(r.opts.VimWikiRoot + string(node.Literal))
				if err != nil {
					fmt.Println(err)
				}
				for _, file := range files {
					dir := filepath.Dir(string(node.Literal))
					link := filepath.Base(file)
					link = strings.Replace(link, ".md", ".gmi", -1)
					output := "=> " + dir + "/" + link + " " + link
					w.Write([]byte(output))
				}
			} else {
				w.Write([]byte(strings.Replace(string(node.Literal), "\n", " ", -1)))
			}
		}
	case *ast.Del:
		r.debugNode(node)
	case *ast.ListItem:
		r.debugNode(node)
	case *ast.Emph:
		r.debugNode(node)
	case *ast.Document:
		r.debugNode(node)
	default:
		panic(fmt.Sprintf("%T is not supported for Gemini.gmi files", node))
	}
	return ast.GoToNext
}
func (r *Renderer) RenderHeader(w io.Writer, ast ast.Node) {
	banner := filepath.Join(r.opts.VimWikiRoot, "banner.txt")
	_, err := os.Stat(banner)
	if err == nil {
		banner, err := ioutil.ReadFile(banner)
		if err != nil {
			panic(err)
		}
		w.Write([]byte("```\n"))
		w.Write(banner)
		w.Write([]byte("\n```\n"))
	}
}
func (r *Renderer) RenderFooter(w io.Writer, ast ast.Node) {
}
