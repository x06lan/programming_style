package main

import (
	"fmt"
)

type DocumentGenerator interface {
	PrepareData() string
	FormatContent(data string) string
	Save(content string) string
}

type BaseGenerator struct {
	generator DocumentGenerator
}

func (g *BaseGenerator) Generate() string {
	data := g.generator.PrepareData()
	formatted := g.generator.FormatContent(data)
	return g.generator.Save(formatted)
}

type TextDocument struct{}

func (t *TextDocument) PrepareData() string {
	return "This is the raw text data."
}
func (t *TextDocument) FormatContent(data string) string {
	return fmt.Sprintf("Formatted Text: %s", data)
}
func (t *TextDocument) Save(content string) string {
	return fmt.Sprintf("Saving text document: %s", content)
}

type HTMLDocument struct{}

func (h *HTMLDocument) PrepareData() string {
	data := "This is raw HTML data."
	return fmt.Sprintf("<html><body>%s</body></html>", data)
}
func (h *HTMLDocument) FormatContent(data string) string {
	return fmt.Sprintf("<div>%s</div>", data)
}
func (h *HTMLDocument) Save(content string) string {
	return fmt.Sprintf("Saving HTML document: %s", content)
}
