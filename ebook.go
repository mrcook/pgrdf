package pgrdf

import (
	"encoding/xml"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"

	"github.com/mrcook/pgrdf/_internal/marshaller"
	"github.com/mrcook/pgrdf/_internal/unmarshaller"
)

// MarcRelatorCode representing a creator role: `aut`, `edt`, etc.
type MarcRelatorCode string

const (
	RoleAut MarcRelatorCode = "aut"
	RoleCom MarcRelatorCode = "com"
	RoleCtb MarcRelatorCode = "ctb"
	RoleEdt MarcRelatorCode = "edt"
	RoleIll MarcRelatorCode = "ill"
	RoleTrl MarcRelatorCode = "trl"
)

// Ebook reads an RDF from the input stream and maps it onto this object, which
// hides the complexities of the source RDF for easier interaction. This type
// can also be un/marshalled to JSON.
type Ebook struct {
	ID                int          `json:"id"`           // PG eText ID.
	BookType          string       `json:"type"`         // Text, Sound, etc.
	ReleaseDate       string       `json:"released"`     // PG release date.
	Language          string       `json:"language"`     // Language of this work.
	Publisher         string       `json:"publisher"`    // Publisher of this work; usually Project Gutenberg.
	Copyright         string       `json:"copyright"`    // Rights for this work (e.g. PD).
	PublishedYear     int          `json:"published"`    // Year this work was published in.
	Titles            []string     `json:"titles"`       // Full title for this work (main, sub, etc.).
	OtherTitles       []string     `json:"other_titles"` // Alternative titles for this work.
	Creators          []Creator    `json:"creators"`     // List of creators for this work (author, illustration, etc.).
	Subjects          []Subject    `json:"subjects"`     // List of subjects for this work.
	Files             []File       `json:"files"`        // List of files for this work (txt, zip, images, etc.).
	Bookshelves       []Bookshelf  `json:"bookshelves"`  // List of bookshelves this work is available in.
	Series            string       `json:"series"`       // The series of work this belongs to.
	BookCoverFilename string       `json:"book_cover"`   // Book cover filename (found in the HTML ebook directory).
	Downloads         int          `json:"downloads"`    // Download count (previous 30 days).
	Note              string       `json:"note"`         // Additional notes about this eText.
	Comment           string       `json:"comment"`      // Usually just info on where to find the RDF files.
	CCLicense         string       `json:"cc_license"`   // Creative Commons license URL.
	AuthorLinks       []AuthorLink `json:"author_links"` // List of author links (typically Wikipedia links).

	// internal slice to prevent duplicate node IDs
	generatedNodeIDs []string
}

// NewEbook returns a new Ebook that reads from the provided io.Reader.
func NewEbook(r io.Reader) (*Ebook, error) {
	rdf, err := unmarshaller.New(r)
	if err != nil {
		return nil, err
	}
	ebook := mapUnmarshalled(rdf)

	return ebook, nil
}

// ToRDF writes RDF XML data to the provided io.Writer.
func (e *Ebook) ToRDF(w io.Writer) error {
	rdf := e.mapToMarshaller()

	data, err := xml.Marshal(rdf)
	if err != nil {
		return err
	}

	if _, err := w.Write(data); err != nil {
		return err
	}

	return nil
}

// Bookshelf is a Project Gutenberg bookshelf.
type Bookshelf struct {
	Resource string `json:"name"`    // Name of bookshelf at gutenberg.org. Usually "2009/pgterms/Bookshelf".
	Name     string `json:"subject"` // The bookshelf name.
}

// Creator is a person involved in the creation of the work, such as the author, illustrator, etc.
type Creator struct {
	ID      int             `json:"id"`                  // Unique Project Gutenberg ID.
	Name    string          `json:"name"`                // Name of the creator.
	Aliases []string        `json:"aliases,omitempty"`   // Any aliases for the creator.
	Born    int             `json:"born_year,omitempty"` // Date of Birth.
	Died    int             `json:"died_year,omitempty"` // Date of Death.
	Role    MarcRelatorCode `json:"role,omitempty"`      // Code indicating the role. e.g. `aut`, `edt`, `ill`, etc.
	WebPage string          `json:"webpage,omitempty"`   // URL for this creator (usually Wikipedia).
}

// File is a resource for the ebook, such as .txt, .tei, .zip, etc.
type File struct {
	URL       string   `json:"url"`      // URL of the file at gutenberg.org.
	Extent    int      `json:"extent"`   // Gutenberg extent number.
	Modified  string   `json:"modified"` // Modified date for this resource.
	Encodings []string `json:"encoding"` // Encoding(s) for this resource.
}

// Subject is a Dublin Core Vocabulary Encoding Scheme such as LCSH and LCC.
type Subject struct {
	Heading string `json:"heading"` // Heading or other label
	Schema  string `json:"schema"`  // Vocabulary Encoding Scheme (full URL).
}

// AuthorLink is an external resource about the author. Usually this is a link
// to Wikipedia.
type AuthorLink struct {
	URL         string `json:"url"`         // URL for this author.
	Description string `json:"description"` // Short description about this link.
}

// Maps the unmarshalled RDF to the exported Ebook type.
func mapUnmarshalled(rdf *unmarshaller.RDF) *Ebook {
	ebook := &Ebook{
		ID:                agentID(rdf.Ebook.About),
		BookType:          rdf.Ebook.Type.Description.Value.Data,
		ReleaseDate:       rdf.Ebook.Issued.Value,
		Language:          languageTag(rdf.Ebook),
		Publisher:         rdf.Ebook.Publisher,
		PublishedYear:     rdf.Ebook.PublishedYear,
		Copyright:         rdf.Ebook.Rights,
		Titles:            titles(rdf.Ebook.Title),
		OtherTitles:       rdf.Ebook.Alternative,
		BookCoverFilename: bookCoverFilename(rdf.Ebook.BookCover),
		Downloads:         rdf.Ebook.Downloads.Value,
		Series:            rdf.Ebook.Series,
		Note:              rdf.Ebook.Description,
		Comment:           rdf.Work.Comment,
		CCLicense:         rdf.Work.License.Resource,
	}

	for _, l := range rdf.Descriptions {
		ebook.addAuthorLink(l.Description, l.About)
	}
	for _, c := range rdf.Ebook.Creators {
		ebook.addCreator(&c.Agent, RoleAut)
	}
	for _, t := range rdf.Ebook.Compilers {
		ebook.addCreator(&t.Agent, RoleCom)
	}
	for _, t := range rdf.Ebook.Contributors {
		ebook.addCreator(&t.Agent, RoleCtb)
	}
	for _, e := range rdf.Ebook.Editors {
		ebook.addCreator(&e.Agent, RoleEdt)
	}
	for _, i := range rdf.Ebook.Illustrators {
		ebook.addCreator(&i.Agent, RoleIll)
	}
	for _, t := range rdf.Ebook.Translators {
		ebook.addCreator(&t.Agent, RoleTrl)
	}
	for _, s := range rdf.Ebook.Subjects {
		ebook.addSubject(s.Description.Value.Data, s.Description.MemberOf.Resource)
	}
	for _, f := range rdf.Ebook.HasFormats {
		ebook.addBookFile(&f)
	}
	for _, s := range rdf.Ebook.Bookshelves {
		ebook.addBookshelf(s.Description.Value.Data, s.Description.MemberOf.Resource)
	}

	return ebook
}

// addCreator appends an Agent to the creators list with the given role.
func (e *Ebook) addCreator(agent *unmarshaller.Agent, role MarcRelatorCode) {
	creator := Creator{
		ID:      agentID(agent.About),
		Name:    agent.Name,
		Aliases: agent.Aliases,
		Born:    agent.Birthdate.Value,
		Died:    agent.Deathdate.Value,
		Role:    role,
		WebPage: agent.Webpage.Resource,
	}
	e.Creators = append(e.Creators, creator)
}

func (e *Ebook) addAuthorLink(description, url string) {
	wiki := AuthorLink{
		Description: description,
		URL:         url,
	}
	e.AuthorLinks = append(e.AuthorLinks, wiki)
}

func (e *Ebook) addSubject(heading, schema string) {
	sub := Subject{
		Heading: heading,
		Schema:  schema,
	}
	e.Subjects = append(e.Subjects, sub)
}

func (e *Ebook) addBookFile(format *unmarshaller.HasFormat) {
	file := File{
		URL:       format.File.About,
		Extent:    format.File.Extent.Value,
		Modified:  format.File.Modified.Value,
		Encodings: nil,
	}
	for _, f := range format.File.Formats {
		file.Encodings = append(file.Encodings, f.Description.Value.Data)
	}
	e.Files = append(e.Files, file)
}

func (e *Ebook) addBookshelf(name, resource string) {
	shelf := Bookshelf{
		Resource: resource,
		Name:     name,
	}
	e.Bookshelves = append(e.Bookshelves, shelf)
}

// Used for extracting the ebook ID and creator ID.
func agentID(about string) int {
	parts := strings.Split(about, "/")
	idString := parts[len(parts)-1]
	id, _ := strconv.Atoi(idString)
	return id
}

// Extract the book cover filename from the file path.
// marc901 tags contain a book cover filename from the HTML version of the ebook.
func bookCoverFilename(cover string) string {
	parts := strings.Split(cover, "-h")
	cover = parts[len(parts)-1]
	cover = strings.TrimPrefix(cover, "/")
	return cover
}

// Constructs a valid language localisation tag: e.g. `en`, `en-GB`, etc.
func languageTag(ebook unmarshaller.Ebook) string {
	var codes []string

	if len(ebook.Language.Description.Value.Data) > 0 {
		codes = append(codes, ebook.Language.Description.Value.Data)
	}
	if len(ebook.LanguageSubCode) > 0 {
		codes = append(codes, ebook.LanguageSubCode)
	}
	return strings.Join(codes, "-")
}

func titles(title string) []string {
	return strings.Split(title, "\n")
}

// Maps Ebook to the marshaller RDF.
func (e *Ebook) mapToMarshaller() *marshaller.RDF {
	rdf := &marshaller.RDF{
		// TODO: only add them if they're needed.
		NsBase:    "http://www.gutenberg.org/",
		NsDcTerms: "http://purl.org/dc/terms/",
		NsPgTerms: "http://www.gutenberg.org/2009/pgterms/",
		NsRdf:     "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
		NsRdfs:    "http://www.w3.org/2000/01/rdf-schema#",
		NsCC:      "http://web.resource.org/cc/",
		NsDCam:    "http://purl.org/dc/dcam/",
		NsMarcRel: "http://id.loc.gov/vocabulary/relators/",

		Work: marshaller.Work{
			Comment: e.Comment,
			License: marshaller.CCLicense{Resource: e.CCLicense},
		},
		Ebook: marshaller.Ebook{
			About:       fmt.Sprintf("ebooks/%d", e.ID),
			Description: e.Note,
			Type: marshaller.Type{
				Description: marshaller.Description{
					NodeID:   e.nodeIdGenerator(),
					Value:    &marshaller.Value{Data: e.BookType},
					MemberOf: &marshaller.MemberOf{Resource: "http://purl.org/dc/terms/DCMIType"},
				},
			},
			Issued: &marshaller.Issued{
				DataType: "http://www.w3.org/2001/XMLSchema#date",
				Value:    e.ReleaseDate,
			},
			Language: marshaller.Language{Description: marshaller.Description{
				NodeID: e.nodeIdGenerator(),
				Value: &marshaller.Value{
					DataType: "http://purl.org/dc/terms/RFC4646",
					Data:     e.Language,
				},
			}},
			License:     marshaller.License{Resource: "license"},
			Publisher:   e.Publisher,
			Rights:      e.Copyright,
			Title:       strings.Join(e.Titles, "\n"),
			Alternative: e.OtherTitles,
			Creators:    nil,
			Subjects:    nil,
			HasFormats:  nil,
			Bookshelves: nil,
			Downloads: &marshaller.Downloads{
				DataType: "http://www.w3.org/2001/XMLSchema#integer",
				Value:    e.Downloads,
			},
		},
		Descriptions: nil,
	}

	for _, c := range e.Creators {
		creator := marshaller.Creator{Agent: marshaller.Agent{
			About:   fmt.Sprintf("2009/agents/%d", c.ID),
			Name:    c.Name,
			Aliases: c.Aliases,
			Birthdate: &marshaller.Year{
				DataType: "http://www.w3.org/2001/XMLSchema#integer",
				Value:    c.Born,
			},
			Deathdate: &marshaller.Year{
				DataType: "http://www.w3.org/2001/XMLSchema#integer",
				Value:    c.Died,
			},
			Webpage: &marshaller.Webpage{Resource: c.WebPage},
		}}
		rdf.Ebook.Creators = append(rdf.Ebook.Creators, creator)
	}

	for _, s := range e.Subjects {
		subject := marshaller.Subject{Description: marshaller.Description{
			NodeID:   e.nodeIdGenerator(),
			Value:    &marshaller.Value{Data: s.Heading},
			MemberOf: &marshaller.MemberOf{Resource: s.Schema},
		}}
		rdf.Ebook.Subjects = append(rdf.Ebook.Subjects, subject)
	}

	for _, f := range e.Files {
		hasFormat := marshaller.HasFormat{
			File: marshaller.File{
				About: f.URL,
				Extent: marshaller.Extent{
					DataType: "http://www.w3.org/2001/XMLSchema#integer",
					Value:    f.Extent,
				},
				Modified: marshaller.Modified{
					DataType: "http://www.w3.org/2001/XMLSchema#dateTime",
					Value:    f.Modified,
				},
				IsFormatOf: marshaller.IsFormatOf{Resource: fmt.Sprintf("ebooks/%d", e.ID)},
				Formats:    nil,
			},
		}
		for _, enc := range f.Encodings {
			format := marshaller.Format{Description: marshaller.Description{
				NodeID:   e.nodeIdGenerator(),
				Value:    &marshaller.Value{DataType: "http://purl.org/dc/terms/IMT", Data: enc},
				MemberOf: &marshaller.MemberOf{Resource: "http://purl.org/dc/terms/IMT"},
			}}
			hasFormat.File.Formats = append(hasFormat.File.Formats, format)
		}

		rdf.Ebook.HasFormats = append(rdf.Ebook.HasFormats, hasFormat)
	}

	for _, s := range e.Bookshelves {
		shelf := marshaller.Bookshelf{Description: marshaller.Description{
			NodeID:   e.nodeIdGenerator(),
			Value:    &marshaller.Value{Data: s.Name},
			MemberOf: &marshaller.MemberOf{Resource: s.Resource},
		}}
		rdf.Ebook.Bookshelves = append(rdf.Ebook.Bookshelves, shelf)
	}

	for _, l := range e.AuthorLinks {
		link := marshaller.Description{
			About:       l.URL,
			Description: l.Description,
		}
		rdf.Descriptions = append(rdf.Descriptions, link)
	}

	return rdf
}

func (e *Ebook) nodeIdGenerator() string {
	const letters = "abcdef0123456789"

	var id string
	for {
		b := make([]byte, 32)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		id = "N" + string(b)

		// check for duplicate ID
		uniq := true
		for _, i := range e.generatedNodeIDs {
			if id == i {
				uniq = false
				break
			}
		}
		if uniq {
			break
		}
	}
	e.generatedNodeIDs = append(e.generatedNodeIDs, id)

	return id
}
