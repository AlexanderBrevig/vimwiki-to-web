package pkg

import (
	markdown "github.com/gomarkdown/markdown"
	parser "github.com/gomarkdown/markdown/parser"
)

func ToGEMINI(md []byte, p *parser.Parser, renderer markdown.Renderer) []byte {
	doc := markdown.Parse(md, p)
	if renderer == nil {
		//Set options
	}
	return markdown.Render(doc, renderer)
}
