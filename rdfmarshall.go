package pgrdf

import (
	"fmt"
	"strings"

	"github.com/mrcook/pgrdf/internal/marshaller"
	"github.com/mrcook/pgrdf/internal/nodeid"
)

// rdfMarshall marshalls an Ebook to an RDF document.
func rdfMarshall(e *Ebook) *marshaller.RDF {
	rdf := &marshaller.RDF{
		// TODO: only add them if they're needed.
		NsBase:    "http://www.gutenberg.org/",
		NsDcTerms: "http://purl.org/dc/terms/",
		NsPgTerms: "http://www.gutenberg.org/2009/pgterms/",
		NsRdf:     "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
		NsRdfs:    "http://www.w3.org/2000/01/rdf-schema#",
		NsCC:      "http://web.resource.org/cc/",
		NsDcDcam:  "http://purl.org/dc/dcam/",
		NsMarcRel: "http://id.loc.gov/vocabulary/relators/",

		Ebook: marshaller.Ebook{
			About:       fmt.Sprintf("ebooks/%d", e.ID),
			Title:       strings.Join(e.Titles, "\n"),
			Alternative: e.OtherTitles,
			Publisher:   e.Publisher,
			Issued: &marshaller.Issued{
				DataType: "http://www.w3.org/2001/XMLSchema#date",
				Value:    e.ReleaseDate,
			},
			Series:        e.Series,
			Languages:     nil,
			LanguageNotes: nil,
			PublishedYear: e.PublishedYear,
			License:       marshaller.License{Resource: "license"},
			Rights:        e.Copyright,
			Type: marshaller.Type{
				Description: marshaller.Description{
					NodeID:   nodeid.Generate(),
					Value:    &marshaller.Value{Data: string(e.BookType)},
					MemberOf: &marshaller.MemberOf{Resource: "http://purl.org/dc/terms/DCMIType"},
				},
			},
			Description: e.Note,
			BookCovers:  []string{e.BookCoverFilename},
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
		Work: marshaller.Work{
			Comment: e.Comment,
			License: marshaller.CCLicense{Resource: e.CCLicense},
		},
	}

	for i, lang := range e.Languages {
		if i == 0 {
			rdf.Ebook.LanguageDialect = lang.Dialect // apply only to first language
		}

		rdf.Ebook.Languages = append(rdf.Ebook.Languages, marshaller.Language{
			Description: marshaller.Description{
				NodeID: nodeid.Generate(),
				Value: &marshaller.Value{
					DataType: "http://purl.org/dc/terms/RFC4646",
					Data:     lang.Code,
				},
			},
		})
		rdf.Ebook.LanguageNotes = lang.Notes
	}

	for _, c := range e.Creators {
		creator := marshaller.Creator{Agent: marshaller.Agent{
			About:   fmt.Sprintf("2009/agents/%d", c.ID),
			Name:    c.Name,
			Aliases: c.Aliases,
			BirthYear: &marshaller.Year{
				DataType: "http://www.w3.org/2001/XMLSchema#integer",
				Value:    c.Born,
			},
			DeathYear: &marshaller.Year{
				DataType: "http://www.w3.org/2001/XMLSchema#integer",
				Value:    c.Died,
			},
			Webpage: &marshaller.Webpage{Resource: c.WebPage},
		}}
		rdf.Ebook.Creators = append(rdf.Ebook.Creators, creator)
	}

	for _, s := range e.Subjects {
		subject := marshaller.Subject{Description: marshaller.Description{
			NodeID:   nodeid.Generate(),
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
				NodeID:   nodeid.Generate(),
				Value:    &marshaller.Value{DataType: "http://purl.org/dc/terms/IMT", Data: enc},
				MemberOf: &marshaller.MemberOf{Resource: "http://purl.org/dc/terms/IMT"},
			}}
			hasFormat.File.Formats = append(hasFormat.File.Formats, format)
		}

		rdf.Ebook.HasFormats = append(rdf.Ebook.HasFormats, hasFormat)
	}

	for _, s := range e.Bookshelves {
		shelf := marshaller.Bookshelf{Description: marshaller.Description{
			NodeID:   nodeid.Generate(),
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
