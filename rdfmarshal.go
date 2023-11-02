package pgrdf

import (
	"fmt"

	"github.com/mrcook/pgrdf/internal/marshaler"
	"github.com/mrcook/pgrdf/internal/nodeid"
)

// rdfMarshal will serialise an Ebook object to a RDF object.
func rdfMarshal(e *Ebook) *marshaler.RDF {
	rdf := &marshaler.RDF{
		// TODO: only add them if they're needed.
		NsBase:    "http://www.gutenberg.org/",
		NsDcTerms: "http://purl.org/dc/terms/",
		NsPgTerms: "http://www.gutenberg.org/2009/pgterms/",
		NsRdf:     "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
		NsRdfs:    "http://www.w3.org/2000/01/rdf-schema#",
		NsCC:      "http://web.resource.org/cc/",
		NsDcDcam:  "http://purl.org/dc/dcam/",
		NsMarcRel: "http://id.loc.gov/vocabulary/relators/",

		Ebook: marshaler.Ebook{
			About:           fmt.Sprintf("ebooks/%d", e.ID),
			Titles:          e.Titles,
			Alternatives:    e.AlternateTitles,
			TableOfContents: e.TableOfContents,
			Publisher:       e.Publisher,
			PublishedYear:   e.PublishedYear,
			Issued: &marshaler.Issued{
				DataType: "http://www.w3.org/2001/XMLSchema#date",
				Value:    e.ReleaseDate,
			},
			Summary:         e.Summary,
			Series:          e.Series,
			Languages:       nil,
			LanguageDialect: e.LanguageDialect,
			LanguageNotes:   e.LanguageNotes,
			PublicationNote: e.PublicationNote,
			EditionNote:     e.EditionNote,
			ProductionNotes: e.ProductionNotes,
			License:         marshaler.License{Resource: "license"},
			Rights:          e.Copyright,
			DpClearanceCode: e.CopyrightClearanceCode,
			Type: marshaler.Type{
				Description: marshaler.Description{
					NodeID:   nodeid.Generate(),
					Value:    &marshaler.Value{Data: string(e.BookType)},
					MemberOf: &marshaler.MemberOf{Resource: "http://purl.org/dc/terms/DCMIType"},
				},
			},
			Descriptions:            e.Notes,
			PhysicalDescriptionNote: e.PhysicalDescriptionNote,
			SourceLinks:             e.SourceLinks,
			LCCN:                    e.LCCN,
			ISBN:                    e.ISBN,
			BookCoverImages:         e.BookCovers,
			TitlePageImage:          e.TitlePageImage,
			BackCoverImage:          e.BackCover,
			Creators:                nil,
			Subjects:                nil,
			HasFormats:              nil,
			Bookshelves:             nil,
			Downloads: &marshaler.Downloads{
				DataType: "http://www.w3.org/2001/XMLSchema#integer",
				Value:    e.Downloads,
			},
		},
		Descriptions: nil,
		Work: marshaler.Work{
			Comment: e.CCComment,
			License: marshaler.CCLicense{Resource: e.CCLicense},
		},
	}

	for _, lang := range e.Languages {
		rdf.Ebook.Languages = append(rdf.Ebook.Languages, marshaler.Language{
			Description: marshaler.Description{
				NodeID: nodeid.Generate(),
				Value: &marshaler.Value{
					DataType: "http://purl.org/dc/terms/RFC4646",
					Data:     lang,
				},
			},
		})
	}

	for _, c := range e.Creators {
		creator := marshaler.Creator{Agent: marshaler.Agent{
			About:   fmt.Sprintf("2009/agents/%d", c.ID),
			Name:    c.Name,
			Aliases: c.Aliases,
			BirthYear: &marshaler.Year{
				DataType: "http://www.w3.org/2001/XMLSchema#integer",
				Value:    c.Born,
			},
			DeathYear: &marshaler.Year{
				DataType: "http://www.w3.org/2001/XMLSchema#integer",
				Value:    c.Died,
			},
		}}
		for _, webpage := range c.WebPages {
			creator.Agent.Webpages = append(creator.Agent.Webpages, marshaler.Webpage{Resource: webpage})
		}
		rdf.Ebook.Creators = append(rdf.Ebook.Creators, creator)
	}

	for _, s := range e.Subjects {
		subject := marshaler.Subject{Description: marshaler.Description{
			NodeID:   nodeid.Generate(),
			Value:    &marshaler.Value{Data: s.Heading},
			MemberOf: &marshaler.MemberOf{Resource: s.Schema},
		}}
		rdf.Ebook.Subjects = append(rdf.Ebook.Subjects, subject)
	}

	for _, f := range e.Files {
		hasFormat := marshaler.HasFormat{
			File: marshaler.File{
				About: f.URL,
				Extent: marshaler.Extent{
					DataType: "http://www.w3.org/2001/XMLSchema#integer",
					Value:    f.Extent,
				},
				Modified: marshaler.Modified{
					DataType: "http://www.w3.org/2001/XMLSchema#dateTime",
					Value:    f.Modified,
				},
				IsFormatOf: marshaler.IsFormatOf{Resource: fmt.Sprintf("ebooks/%d", e.ID)},
				Formats:    nil,
			},
		}
		for _, enc := range f.Encodings {
			format := marshaler.Format{Description: marshaler.Description{
				NodeID:   nodeid.Generate(),
				Value:    &marshaler.Value{DataType: "http://purl.org/dc/terms/IMT", Data: enc},
				MemberOf: &marshaler.MemberOf{Resource: "http://purl.org/dc/terms/IMT"},
			}}
			hasFormat.File.Formats = append(hasFormat.File.Formats, format)
		}

		rdf.Ebook.HasFormats = append(rdf.Ebook.HasFormats, hasFormat)
	}

	for _, s := range e.Bookshelves {
		shelf := marshaler.Bookshelf{Description: marshaler.Description{
			NodeID:   nodeid.Generate(),
			Value:    &marshaler.Value{Data: s.Name},
			MemberOf: &marshaler.MemberOf{Resource: s.Resource},
		}}
		rdf.Ebook.Bookshelves = append(rdf.Ebook.Bookshelves, shelf)
	}

	for _, l := range e.AuthorLinks {
		link := marshaler.Description{
			About:       l.URL,
			Description: l.Description,
		}
		rdf.Descriptions = append(rdf.Descriptions, link)
	}

	return rdf
}
