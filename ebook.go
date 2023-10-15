package pgrdf

import (
	"encoding/xml"
	"io"
	"regexp"
)

// Ebook reads an RDF from the input stream and maps it onto this object, which
// hides the complexities of the source RDF for easier interaction. This type
// can also be un/marshaled to JSON.
type Ebook struct {
	// PG eText ID.
	// `<pgterms:ebook rdf:about="...">`.
	ID int `json:"id"`

	// Full title for this work (main, sub, etc.).
	// `<dcterms:title>`
	Titles []string `json:"titles"`

	// Alternate titles for this work.
	// `<dcterms:alternative>`
	AlternateTitles []string `json:"alternate_titles,omitempty"`

	// Publisher of this work; always "Project Gutenberg".
	// `<dcterms:publisher>`
	Publisher string `json:"publisher"`

	// Year this work was published in.
	// `<pgterms:marc906>`
	PublishedYear int `json:"published"`

	// PG release/issued date in ISO 8601 format. Example: 2006-01-02.
	// `<dcterms:issued>`
	ReleaseDate string `json:"released"`

	// A short summary of the work.
	// `<pgterms:marc520>`
	Summary string `json:"summary,omitempty"`

	// The series this work originally belonged to.
	// `<pgterms:marc440>`
	Series []string `json:"series,omitempty"`

	// Languages used in this book.
	// `<dcterms:language>`
	Languages []string `json:"languages,omitempty"`

	// Language dialect (ISO 3166-2), _probably_ for the primary language only.
	// `<pgterms:marc907>`
	LanguageDialect string `json:"language_dialect,omitempty"`

	// Notes about the language of the work, e.g. "Uses 19th century spelling."
	// `<pgterms:marc546>`
	LanguageNotes []string `json:"language_notes,omitempty"`

	// Source publication information: publisher, city, year, etc.
	// `<pgterms:marc260>`
	SrcPublicationInfo string `json:"src_publication_info,omitempty"`

	// Edition of this work, e.g. "2nd Edition", "A new edition with eleven new poems.", etc.
	// `<pgterms:marc250>:`
	Edition string `json:"edition,omitempty"`

	// Credits for this ebook. This can also include "updated" dates, either as a
	// separate entry or as part of the credit, e.g. "J. Smith\nUpdated: 2022-07-14".
	// `<pgterms:marc508>`
	Credits []string `json:"credits,omitempty"`

	// Rights for this work. Most are "Public domain in the USA."
	// `<dcterms:rights>`
	Copyright string `json:"copyright"`

	// Distributed Proofreaders clearance code, e.g. "20050213050736stahl".
	// `<pgterms:marc905>`
	CopyrightClearanceCode string `json:"pg_dp_clearance"`

	// Type of this work, one of:
	//   Collection, Dataset, Image, MovingImage, Sound, StillImage, Text
	// `<dcterms:type>`
	BookType BookType `json:"type"`

	// Additional notes about this eText.
	// `<dcterms:description>`
	Notes []string `json:"note"`

	// Misc information about the source of this work, e.g. "5 pages : illustrations, map, portraits.".
	// `<pgterms:marc300>`
	SourceDescription string `json:"source_description"`

	// URLs to information about the source of this work, e.g. image scans on Internet Archive website.
	// `<pgterms:marc904>`
	SourceLinks []string `json:"source_link"`

	// LOC (Library of Congress) code.
	// `<pgterms:marc010>`
	LOC string `json:"loc"`

	// ISBN of this work - possibly of the work used for the OCR.
	// `<pgterms:marc020>`
	ISBN string `json:"isbn"`

	// Book covers, or images acting as a book cover, i.e. this could be a title page.
	// A URL or file path in the HTML ebook directory.
	// `<pgterms:marc901>`
	BookCovers []string `json:"book_covers"`

	// URL of an image file representing the title page of a book.
	// `<pgterms:marc902>`
	TitlePageImage string `json:"title_page_image"`

	// URL of an image file representing the back cover of a book.
	// `<pgterms:marc903>`
	BackCover string `json:"back_cover"`

	// List of creators involved in the creation of this work, such as the
	// authors, editors, and illustrators.
	// `<dcterms:creator>`, `<marcrel:*>`, `<pgterms:agent>`
	Creators []Creator `json:"creators"`

	// List of subjects for this work.
	// `<dcterms:subject>`
	Subjects []Subject `json:"subjects"`

	// List of files associated with this work (txt, zip, images, etc.).
	// `<dcterms:hasFormat>`, `<pgterms:file>`
	Files []File `json:"files"`

	// List of PG bookshelves this work is available in.
	// `<pgterms:bookshelf>`
	Bookshelves []Bookshelf `json:"bookshelves"`

	// Download count - only from the previous 30 days at the time the RDF was generated.
	// <pgterms:downloads>
	Downloads int `json:"downloads"`

	// List of author links (typically to Wikipedia).
	// `<rdf:Description>`
	AuthorLinks []AuthorLink `json:"author_links"`

	// A Creative Commons comment, usually just info on where to find the RDF files.
	// `<cc:Work><rdfs:comment>`
	CCComment string `json:"cc_comment"`

	// A Creative Commons license URL.
	// `<cc:Work><cc:license>`
	CCLicense string `json:"cc_license"`
}

// ReadRDF document from the given `io.Reader` and unmarshal to an Ebook.
func ReadRDF(r io.Reader) (*Ebook, error) {
	doc, err := rdfUnmarshal(r)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// WriteRDF marshals the Ebook to an RDF document and writes it to the provided `io.Writer`.
func (e *Ebook) WriteRDF(w io.Writer) error {
	rdf := rdfMarshal(e)

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
