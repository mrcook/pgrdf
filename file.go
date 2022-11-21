package pgrdf

// File is a resource for the ebook, such as .txt, .tei, .zip, etc.
type File struct {
	URL       string   `json:"url"`      // URL of the file at gutenberg.org.
	Extent    int      `json:"extent"`   // Extent is the size of the file in bytes.
	Modified  string   `json:"modified"` // Modified date for this resource.
	Encodings []string `json:"encoding"` // Encoding(s) for this resource.
}

func (f *File) AddEncoding(encoding string) {
	// TODO: prevent duplicates?
	f.Encodings = append(f.Encodings, encoding)
}
