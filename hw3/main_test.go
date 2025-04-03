package main

import (
	"testing"
)

// Unit Tests
func TestTextDocument(t *testing.T) {
	textDoc := &TextDocument{}
	if got := textDoc.PrepareData(); got != "This is the raw text data." {
		t.Errorf("PrepareData() = %v, want %v", got, "This is the raw text data.")
	}

	formatted := textDoc.FormatContent("This is the raw text data.")
	if formatted != "Formatted Text: This is the raw text data." {
		t.Errorf("FormatContent() = %v, want %v", formatted, "Formatted Text: This is the raw text data.")
	}

	saved := textDoc.Save(formatted)
	if saved != "Saving text document: Formatted Text: This is the raw text data." {
		t.Errorf("Save() = %v, want %v", saved, "Saving text document: Formatted Text: This is the raw text data.")
	}
}

func TestHTMLDocument(t *testing.T) {
	htmlDoc := &HTMLDocument{}
	if got := htmlDoc.PrepareData(); got != "<html><body>This is raw HTML data.</body></html>" {
		t.Errorf("PrepareData() = %v, want %v", got, "<html><body>This is raw HTML data.</body></html>")
	}

	formatted := htmlDoc.FormatContent("<html><body>This is raw HTML data.</body></html>")
	if formatted != "<div><html><body>This is raw HTML data.</body></html></div>" {
		t.Errorf("FormatContent() = %v, want %v", formatted, "<div><html><body>This is raw HTML data.</body></html></div>")
	}

	saved := htmlDoc.Save(formatted)
	if saved != "Saving HTML document: <div><html><body>This is raw HTML data.</body></html></div>" {
		t.Errorf("Save() = %v, want %v", saved, "Saving HTML document: <div><html><body>This is raw HTML data.</body></html></div>")
	}
}

func TestBaseGenerator(t *testing.T) {
	textDoc := &TextDocument{}
	textGenerator := BaseGenerator{generator: textDoc}
	if got := textGenerator.Generate(); got != "Saving text document: Formatted Text: This is the raw text data." {
		t.Errorf("Generate() = %v, want %v", got, "Saving text document: Formatted Text: This is the raw text data.")
	}

	htmlDoc := &HTMLDocument{}
	htmlGenerator := BaseGenerator{generator: htmlDoc}
	if got := htmlGenerator.Generate(); got != "Saving HTML document: <div><html><body>This is raw HTML data.</body></html></div>" {
		t.Errorf("Generate() = %v, want %v", got, "Saving HTML document: <div><html><body>This is raw HTML data.</body></html></div>")
	}
}
