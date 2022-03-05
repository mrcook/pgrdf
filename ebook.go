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
	RoleEdt MarcRelatorCode = "edt"
	RoleIll MarcRelatorCode = "ill"
	RoleTrl MarcRelatorCode = "trl"
)

// Ebook reads an RDF from the input stream and maps it onto this object, which
// hides the complexities of the source RDF for easier interaction. This type
// can also be un/marshalled to JSON.
type Ebook struct {
	ID            int          `json:"id"`           // PG eText ID.
	BookType      string       `json:"type"`         // Text, Sound, etc.
	ReleaseDate   string       `json:"released"`     // PG release date.
	Language      string       `json:"language"`     // Language of this work.
	Publisher     string       `json:"publisher"`    // Publisher of this work; usually Project Gutenberg.
	Copyright     string       `json:"copyright"`    // Rights for this work (e.g. PD).
	PublishedYear int          `json:"published"`    // Year this work was published in
	Titles        []string     `json:"titles"`       // Full title for this work (main, sub, etc.).
	OtherTitles   []string     `json:"other_titles"` // Alternative titles for this work.
	Creators      []Creator    `json:"creators"`     // List of creators for this work (author, illustration, etc.).
	Subjects      []Subject    `json:"subjects"`     // List of subjects for this work.
	Files         []File       `json:"files"`        // List of files for this work (txt, zip, images, etc.).
	Bookshelves   []Bookshelf  `json:"bookshelves"`  // List of bookshelves this work is available in.
	Downloads     int          `json:"downloads"`    // Download count (previous 30 days).
	Note          string       `json:"note"`         // Additional notes about this eText.
	Comment       string       `json:"comment"`      // Usually just info on where to find the RDF files.
	CCLicense     string       `json:"cc_license"`   // Creative Commons license URL.
	AuthorLinks   []AuthorLink `json:"author_links"` // List of author links (typically Wikipedia links).

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
func mapUnmarshalled(r *unmarshaller.RDF) *Ebook {
	ebook := &Ebook{
		ID:            extractAgentID(r.Ebook.About),
		BookType:      r.Ebook.Type.Description.Value.Data,
		ReleaseDate:   r.Ebook.Issued.Value,
		Language:      constructLanguageTag(r.Ebook),
		Publisher:     r.Ebook.Publisher,
		PublishedYear: r.Ebook.PublishedYear,
		Copyright:     r.Ebook.Rights,
		Titles:        titles(r.Ebook.Title),
		OtherTitles:   r.Ebook.Alternative,
		Creators:      nil,
		Subjects:      nil,
		Files:         nil,
		Bookshelves:   nil,
		Downloads:     r.Ebook.Downloads.Value,
		Note:          r.Ebook.Description,

		Comment:     r.Work.Comment,
		CCLicense:   r.Work.License.Resource,
		AuthorLinks: nil,
	}

	for _, l := range r.Descriptions {
		wiki := AuthorLink{
			Description: l.Description,
			URL:         l.About,
		}
		ebook.AuthorLinks = append(ebook.AuthorLinks, wiki)
	}

	for _, c := range r.Ebook.Creators {
		ebook.addCreator(&c.Agent, RoleAut)
	}
	for _, c := range r.Ebook.Editors {
		ebook.addCreator(&c.Agent, RoleEdt)
	}
	for _, c := range r.Ebook.Illustrators {
		ebook.addCreator(&c.Agent, RoleIll)
	}
	for _, c := range r.Ebook.Translators {
		ebook.addCreator(&c.Agent, RoleTrl)
	}

	for _, s := range r.Ebook.Subjects {
		sub := Subject{
			Heading: s.Description.Value.Data,
			Schema:  s.Description.MemberOf.Resource,
		}
		ebook.Subjects = append(ebook.Subjects, sub)
	}

	for _, s := range r.Ebook.HasFormats {
		file := File{
			URL:       s.File.About,
			Extent:    s.File.Extent.Value,
			Modified:  s.File.Modified.Value,
			Encodings: nil,
		}
		for _, f := range s.File.Formats {
			file.Encodings = append(file.Encodings, f.Description.Value.Data)
		}

		ebook.Files = append(ebook.Files, file)
	}

	for _, s := range r.Ebook.Bookshelves {
		shelf := Bookshelf{
			Resource: s.Description.MemberOf.Resource,
			Name:     s.Description.Value.Data,
		}
		ebook.Bookshelves = append(ebook.Bookshelves, shelf)
	}

	return ebook
}

// Constructs a valid language localisation tag: e.g. `en`, `en-GB`, etc.
func constructLanguageTag(ebook unmarshaller.Ebook) string {
	var codes []string

	if len(ebook.Language.Description.Value.Data) > 0 {
		codes = append(codes, ebook.Language.Description.Value.Data)
	}
	if len(ebook.LanguageSubCode) > 0 {
		codes = append(codes, ebook.LanguageSubCode)
	}
	return strings.Join(codes, "-")
}

// Used for extracting the ebook ID and creator ID.
func extractAgentID(about string) int {
	parts := strings.Split(about, "/")
	idString := parts[len(parts)-1]
	id, _ := strconv.Atoi(idString)
	return id
}

func titles(title string) []string {
	return strings.Split(title, "\n")
}

// addCreator appends an Agent to the creators list with the given role.
func (e *Ebook) addCreator(agent *unmarshaller.Agent, role MarcRelatorCode) {
	creator := Creator{
		ID:      extractAgentID(agent.About),
		Name:    agent.Name,
		Aliases: agent.Aliases,
		Born:    agent.Birthdate.Value,
		Died:    agent.Deathdate.Value,
		Role:    role,
		WebPage: agent.Webpage.Resource,
	}
	e.Creators = append(e.Creators, creator)
}

// Maps Ebook to the marshaller RDF.
func (e *Ebook) mapToMarshaller() *marshaller.RDF {
	rdf := &marshaller.RDF{
		NsBase:    "http://www.gutenberg.org/",
		NsDcTerms: "http://purl.org/dc/terms/",
		NsPgTerms: "http://www.gutenberg.org/2009/pgterms/",
		NsRdf:     "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
		NsRdfs:    "http://www.w3.org/2000/01/rdf-schema#",
		NsCC:      "http://web.resource.org/cc/",
		NsDCam:    "http://purl.org/dc/dcam/",

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
