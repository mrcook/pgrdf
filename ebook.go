package pgrdf

import (
	"encoding/xml"
	"io"
	"regexp"
)

// Ebook reads an RDF from the input stream and maps it onto this object, which
// hides the complexities of the source RDF for easier interaction. This type
// can also be un/marshalled to JSON.
type Ebook struct {
	// PG eText ID.
	ID int `json:"id"`

	// The type of ebook: Text, Sound, etc.
	BookType BookType `json:"type"`

	// PG release date in ISO 8601 format. Example: 2006-01-02.
	ReleaseDate string `json:"released"`

	// Languages details of this book.
	Languages []Language `json:"languages"`

	// Publisher of this work; usually "Project Gutenberg".
	Publisher string `json:"publisher"`

	// Year this work was published in.
	PublishedYear int `json:"published"`

	// Rights for this work. Example: PD.
	Copyright string `json:"copyright"`

	// Full title for this work (main, sub, etc.).
	Titles []string `json:"titles"`

	// Alternative titles for this work.
	OtherTitles []string `json:"other_titles"`

	// List of creators for this work (author, illustrator, etc.).
	Creators []Creator `json:"creators"`

	// List of subjects for this work.
	Subjects []Subject `json:"subjects"`

	// List of files associated with this work (txt, zip, images, etc.).
	Files []File `json:"files"`

	// List of bookshelves this work is available in.
	Bookshelves []Bookshelf `json:"bookshelves"`

	// The series this work originally belonged to.
	Series string `json:"series"`

	// Book cover filename (found in the HTML ebook directory).
	BookCoverFilename string `json:"book_cover"`

	// Download count - from the previous 30 days, so can be more or less than
	// the last time the RDF was processed.
	Downloads int `json:"downloads"`

	// Additional notes about this eText.
	Note string `json:"note"`

	// Usually just info on where to find the RDF files.
	Comment string `json:"comment"`

	// Creative Commons license URL.
	CCLicense string `json:"cc_license"`

	// List of author links (typically to Wikipedia).
	AuthorLinks []AuthorLink `json:"author_links"`
}

// ReadRDF document from the given `io.Reader` and unmarshall to an Ebook.
func ReadRDF(r io.Reader) (*Ebook, error) {
	doc, err := rdfUnmarshall(r)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// WriteRDF marshals the Ebook to an RDF document and writes it to the provided `io.Writer`.
func (e *Ebook) WriteRDF(w io.Writer) error {
	rdf := rdfMarshall(e)

	data, err := xml.MarshalIndent(rdf, "", "  ")
	if err != nil {
		return err
	}

	// The xml package does not currently emit self-closing tags, e.g. `<tag />`.
	r := regexp.MustCompile(`></[^>]+?>`)
	data = r.ReplaceAll(data, []byte("/>"))

	// prepend the xml declaration
	data = append([]byte(xml.Header), data...)

	if _, err := w.Write(data); err != nil {
		return err
	}

	return nil
}

func (e *Ebook) AddSubject(heading, schema string) {
	sub := Subject{
		Heading: heading,
		Schema:  schema,
	}
	e.Subjects = append(e.Subjects, sub)
}

func (e *Ebook) AddBookshelf(name, resource string) {
	shelf := Bookshelf{
		Resource: resource,
		Name:     name,
	}
	e.Bookshelves = append(e.Bookshelves, shelf)
}

func (e *Ebook) AddAuthorLink(description, url string) {
	wiki := AuthorLink{
		Description: description,
		URL:         url,
	}
	e.AuthorLinks = append(e.AuthorLinks, wiki)
}

// AddCreator appends an Agent to the creators list with the given role.
func (e *Ebook) AddCreator(creator Creator) {
	e.Creators = append(e.Creators, creator)
}

func (e *Ebook) AddBookFile(file File) {
	e.Files = append(e.Files, file)
}

func (e *Ebook) SetBookType(value string) {
	switch value {
	case "collection":
		e.BookType = BookTypeCollection
	case "Dataset":
		e.BookType = BookTypeDataset
	case "Image":
		e.BookType = BookTypeImage
	case "MovingImage":
		e.BookType = BookTypeMovingImage
	case "Sound":
		e.BookType = BookTypeSound
	case "StillImage":
		e.BookType = BookTypeStillImage
	case "Text":
		e.BookType = BookTypeText
	default:
		e.BookType = BookTypeUnknown
	}
}

type BookType string

const (
	BookTypeUnknown     BookType = ""
	BookTypeCollection  BookType = "Collection"
	BookTypeDataset     BookType = "Dataset"
	BookTypeImage       BookType = "Image"
	BookTypeMovingImage BookType = "MovingImage"
	BookTypeSound       BookType = "Sound"
	BookTypeStillImage  BookType = "StillImage"
	BookTypeText        BookType = "Text"
)
