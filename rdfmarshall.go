package pgrdf

import (
	"fmt"
	"strings"

	"github.com/mrcook/pgrdf/_internal/marshaller"
	"github.com/mrcook/pgrdf/_internal/nodeid"
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
					NodeID:   nodeid.Generate(),
					Value:    &marshaller.Value{Data: e.BookType},
					MemberOf: &marshaller.MemberOf{Resource: "http://purl.org/dc/terms/DCMIType"},
				},
			},
			Issued: &marshaller.Issued{
				DataType: "http://www.w3.org/2001/XMLSchema#date",
				Value:    e.ReleaseDate,
			},
			Language: marshaller.Language{Description: marshaller.Description{
				NodeID: nodeid.Generate(),
				Value: &marshaller.Value{
					DataType: "http://purl.org/dc/terms/RFC4646",
					Data:     e.Language.Code,
				},
			}},
			LanguageDialect: e.Language.Dialect,
			LanguageNotes:   e.Language.Notes,
			License:         marshaller.License{Resource: "license"},
			Publisher:       e.Publisher,
			PublishedYear:   e.PublishedYear,
			Rights:          e.Copyright,
			Title:           strings.Join(e.Titles, "\n"),
			Alternative:     e.OtherTitles,
			Creators:        nil,
			Subjects:        nil,
			HasFormats:      nil,
			Bookshelves:     nil,
			Series:          e.Series,
			BookCover:       e.BookCoverFilename,
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
