package pgrdf

import "strings"

// File is a resource for the ebook such as .txt, .tei, .zip, etc.
type File struct {
	// URL of the file at gutenberg.org.
	// <pgterms:file rdf:about="...">
	URL string `json:"url"`

	// Extent is the size of the file in bytes.
	// `<dcterms:extent>`
	Extent int `json:"extent"`

	// Modified date for this resource.
	// `<dcterms:modified>`
	Modified string `json:"modified"`

	// Encodings for this resource, e.g. "image/jpeg"
	// `<dcterms:format>`
	Encodings []string `json:"encoding"`
}

func (f *File) AddEncoding(encoding string) {
	encoding = strings.TrimSpace(encoding)
	if len(encoding) == 0 {
		return
	}

	for _, enc := range f.Encodings {
		if enc == encoding {
			return
		}
	}

	f.Encodings = append(f.Encodings, encoding)
}
