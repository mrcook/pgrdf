// Package marshaller contains a set of structs for generating a Project
// Gutenberg RDF XML document.
//
// NOTE: due to limitations in the Go xml package and the namespace complexity
// of the RDF documents, a separate set of marshaller and unmarshaller structs
// are required.
package marshaller

import (
	"encoding/xml"

	"github.com/mrcook/pgrdf/_internal/unmarshaller"
)

// RDF <rdf:RDF /> is the main document struct
type RDF struct {
	XMLName   xml.Name `xml:"rdf:RDF"`
	NsBase    string   `xml:"xml:base,attr,omitempty"`
	NsDcTerms string   `xml:"xmlns:dcterms,attr,omitempty"`
	NsPgTerms string   `xml:"xmlns:pgterms,attr,omitempty"`
	NsRdf     string   `xml:"xmlns:rdf,attr,omitempty"`
	NsRdfs    string   `xml:"xmlns:rdfs,attr,omitempty"`
	NsCC      string   `xml:"xmlns:cc,attr,omitempty"`
	NsMarcRel string   `xml:"xmlns:marcrel,attr,omitempty"`
	NsDCam    string   `xml:"xmlns:dcam,attr,omitempty"`

	Ebook        Ebook         `xml:"pgterms:ebook,omitempty"`
	Descriptions []Description `xml:"rdf:Description,omitempty"`
	Work         Work          `xml:"cc:Work,omitempty"`
}

// Ebook <pgterms:ebook /> holds the core metadata for this work.
type Ebook struct {
	About               string   `xml:"rdf:about,attr,omitempty"`
	Title               string   `xml:"dcterms:title,omitempty"`
	Alternative         []string `xml:"dcterms:alternative,omitempty"`
	Publisher           string   `xml:"dcterms:publisher,omitempty"`
	Issued              *Issued  `xml:"dcterms:issued,omitempty"`
	Summary             string   `xml:"pgterms:marc520,omitempty"`
	Series              string   `xml:"pgterms:marc440,omitempty"`
	Language            Language `xml:"dcterms:language,omitempty"`
	LanguageDialect     string   `xml:"pgterms:marc907,omitempty"`
	LanguageNotes       string   `xml:"pgterms:marc546,omitempty"`
	PublishedYear       int      `xml:"pgterms:marc906,omitempty"`
	OriginalPublication string   `xml:"pgterms:marc260,omitempty"`
	Edition             string   `xml:"pgterms:marc250,omitempty"`
	Credits             []string `xml:"pgterms:marc508,omitempty"`
	License             License  `xml:"dcterms:license,omitempty"`
	Rights              string   `xml:"dcterms:rights,omitempty"`
	PgDpClearance       string   `xml:"pgterms:marc905,omitempty"`
	Type                Type     `xml:"dcterms:type,omitempty"`
	Description         string   `xml:"dcterms:description,omitempty"`
	SourceDescription   string   `xml:"pgterms:marc300,omitempty"`
	SourceLink          string   `xml:"pgterms:marc904,omitempty"`
	LOC                 string   `xml:"pgterms:marc010,omitempty"`
	ISBN                string   `xml:"pgterms:marc020,omitempty"`
	BookCover           string   `xml:"pgterms:marc901,omitempty"`
	TitlePageImage      string   `xml:"pgterms:marc902,omitempty"`
	BackCover           string   `xml:"pgterms:marc903,omitempty"`

	Creators []Creator `xml:"dcterms:creator,omitempty"`

	// TODO: can these be marshalled programmatically?
	RelAdapters      []MarcRelator `xml:"marcrel:adp,omitempty"`
	RelAfterwords    []MarcRelator `xml:"marcrel:aft,omitempty"`
	RelAnnotators    []MarcRelator `xml:"marcrel:ann,omitempty"`
	RelArrangers     []MarcRelator `xml:"marcrel:arr,omitempty"`
	RelArtists       []MarcRelator `xml:"marcrel:art,omitempty"`
	RelIntroductions []MarcRelator `xml:"marcrel:aui,omitempty"`
	RelCommentators  []MarcRelator `xml:"marcrel:cmm,omitempty"`
	RelComposers     []MarcRelator `xml:"marcrel:cmp,omitempty"`
	RelConductors    []MarcRelator `xml:"marcrel:cnd,omitempty"`
	RelCompilers     []MarcRelator `xml:"marcrel:com,omitempty"`
	RelContributors  []MarcRelator `xml:"marcrel:ctb,omitempty"`
	RelDubious       []MarcRelator `xml:"marcrel:dub,omitempty"`
	RelEditors       []MarcRelator `xml:"marcrel:edt,omitempty"`
	RelEngravers     []MarcRelator `xml:"marcrel:egr,omitempty"`
	RelIllustrators  []MarcRelator `xml:"marcrel:ill,omitempty"`
	RelLibrettists   []MarcRelator `xml:"marcrel:lbt,omitempty"`
	RelOther         []MarcRelator `xml:"marcrel:oth,omitempty"`
	RelPublishers    []MarcRelator `xml:"marcrel:pbl,omitempty"`
	RelPhotographers []MarcRelator `xml:"marcrel:pht,omitempty"`
	RelPerformers    []MarcRelator `xml:"marcrel:prf,omitempty"`
	RelPrinters      []MarcRelator `xml:"marcrel:prt,omitempty"`
	RelResearchers   []MarcRelator `xml:"marcrel:res,omitempty"`
	RelTranscribers  []MarcRelator `xml:"marcrel:trc,omitempty"`
	RelTranslators   []MarcRelator `xml:"marcrel:trl,omitempty"`

	Subjects    []Subject   `xml:"dcterms:subject,omitempty"`
	HasFormats  []HasFormat `xml:"dcterms:hasFormat,omitempty"`
	Bookshelves []Bookshelf `xml:"pgterms:bookshelf,omitempty"`

	Downloads *Downloads `xml:"pgterms:downloads,omitempty"`
}

// Type <dcterms:type /> the media type of this work: text, audio, etc.
type Type struct {
	Description Description
}

// Issued <dcterms:issued /> the date this eText was added to the Gutenberg collection.
type Issued struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Value    string `xml:",chardata"`
}

// Language <dcterms:language /> the language this work was written in.
type Language struct {
	Description Description
}

// License <dcterms:license />
type License struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
}

// Agent <pgterms:agent /> represents a contributor to this work; author, illustrator, etc.
type Agent struct {
	About     string   `xml:"rdf:about,attr,omitempty"`
	Name      string   `xml:"pgterms:name,omitempty"`
	Aliases   []string `xml:"pgterms:alias,omitempty"`
	BirthYear *Year    `xml:"pgterms:birthdate,omitempty"`
	DeathYear *Year    `xml:"pgterms:deathdate,omitempty"`
	Webpage   *Webpage `xml:"pgterms:webpage,omitempty"`
}

// Year is used for representing an `rdf:datatype` attribute for an Agent
// birth and death years.
type Year struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Value    int    `xml:",chardata"`
}

// Webpage <pgterms:webpage /> a webpage link.
type Webpage struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
}

// Creator <dcterms:creator /> is the creator (author) of this work.
type Creator struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
	Agent    Agent  `xml:"pgterms:agent"`
}

type MarcRelator struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
	Agent    *Agent `xml:"pgterms:agent,omitempty"`
}

// HasFormat <dcterms:hasFormat /> represents a file resource.
type HasFormat struct {
	XMLName xml.Name `xml:"dcterms:hasFormat"`
	File    File
}

// File <pgterms:file /> is a file resource.
type File struct {
	XMLName    xml.Name   `xml:"pgterms:file"`
	About      string     `xml:"rdf:about,attr,omitempty"`
	Extent     Extent     `xml:"dcterms:extent,omitempty"`
	Modified   Modified   `xml:"dcterms:modified,omitempty"`
	IsFormatOf IsFormatOf `xml:"dcterms:isFormatOf"`
	Formats    []Format   `xml:"dcterms:format,omitempty"`
}

// Extent <dcterms:extent /> indicates the size of a file, in bytes.
type Extent struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Value    int    `xml:",chardata"`
}

// Modified <dcterms:modified /> date a file resource was last modified.
type Modified struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Value    string `xml:",chardata"`
}

// IsFormatOf <dcterms:isFormatOf /> contains the ebook ID.
type IsFormatOf struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
}

// Format <dcterms:format /> contains file type information.
type Format struct {
	Description Description
}

// Bookshelf <pgterms:bookshelf /> is a link to a Project Gutenberg bookshelf.
// that this work is part of.
type Bookshelf struct {
	Description Description
}

// Subject <dcterms:subject /> and object representing a subject/genre.
type Subject struct {
	Description Description
}

// Downloads <pgterms:downloads /> contains the number of times this work has
// been downloaded over the last 30 days.
type Downloads struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Value    int    `xml:",chardata"`
}

// Work <cc:Work /> is the Creative Commons license information for this work.
type Work struct {
	About   string    `xml:"rdf:about,attr,omitempty"`
	Comment string    `xml:"rdfs:comment,omitempty"`
	License CCLicense `xml:"cc:license,omitempty"`
}

// CCLicense <cc:license /> is a link for the Creative Commons license.
type CCLicense struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
}

// Description <rdf:Description /> is a generic struct for describing the
// node it is included in.
type Description struct {
	XMLName     xml.Name  `xml:"rdf:Description"`
	About       string    `xml:"rdf:about,attr,omitempty"`
	NodeID      string    `xml:"rdf:nodeID,attr,omitempty"`
	Value       *Value    `xml:"rdf:value,omitempty"`
	MemberOf    *MemberOf `xml:"dcam:memberOf,omitempty"`
	Description string    `xml:"dcterms:description,omitempty"`
}

// Value <rdf:value /> for a Description object.
type Value struct {
	XMLName  xml.Name `xml:"rdf:value"`
	DataType string   `xml:"rdf:datatype,attr,omitempty"`
	Data     string   `xml:",chardata"`
}

// MemberOf <dcam:memberOf /> the schema for a given subject/genre.
type MemberOf struct {
	XMLName  xml.Name `xml:"dcam:memberOf"`
	Resource string   `xml:"rdf:resource,attr,omitempty"`
}

// FromUnmarshaller is a helper function for mapping an unmarshaller.RDF to
// this marshaller.RDF object so that we have both marshall and unmarshall
// functionality for an RDF XML document.
// This is required due to the Go xml package limitations, as described at the
// top of this file.
func FromUnmarshaller(in *unmarshaller.RDF) *RDF {
	out := &RDF{
		NsBase:    in.NsBase,
		NsDcTerms: in.NsDcTerms,
		NsPgTerms: in.NsPgTerms,
		NsRdf:     in.NsRdf,
		NsRdfs:    in.NsRdfs,
		NsCC:      in.NsCC,
		NsMarcRel: in.NsMarcRel,
		NsDCam:    in.NsDcam,
		Ebook: Ebook{
			About:       in.Ebook.About,
			Description: in.Ebook.Description,
			Type: Type{
				Description: description(&in.Ebook.Type.Description),
			},
			Issued: nil,
			Language: Language{
				Description: description(&in.Ebook.Language.Description),
			},
			Publisher:     in.Ebook.Publisher,
			PublishedYear: in.Ebook.PublishedYear,
			License:       License{Resource: in.Ebook.License.Resource},
			Rights:        in.Ebook.Rights,
			Title:         in.Ebook.Title,
			Alternative:   in.Ebook.Alternative,
			Creators:      nil,

			RelAdapters:      nil,
			RelAfterwords:    nil,
			RelAnnotators:    nil,
			RelArrangers:     nil,
			RelArtists:       nil,
			RelIntroductions: nil,
			RelCommentators:  nil,
			RelComposers:     nil,
			RelConductors:    nil,
			RelCompilers:     nil,
			RelContributors:  nil,
			RelDubious:       nil,
			RelEditors:       nil,
			RelEngravers:     nil,
			RelIllustrators:  nil,
			RelLibrettists:   nil,
			RelOther:         nil,
			RelPublishers:    nil,
			RelPhotographers: nil,
			RelPerformers:    nil,
			RelPrinters:      nil,
			RelResearchers:   nil,
			RelTranscribers:  nil,
			RelTranslators:   nil,

			Subjects:    nil,
			HasFormats:  nil,
			Bookshelves: nil,

			LOC:                 in.Ebook.LOC,
			ISBN:                in.Ebook.ISBN,
			Edition:             in.Ebook.Edition,
			OriginalPublication: in.Ebook.OriginalPublication,
			SourceDescription:   in.Ebook.SourceDescription,
			Series:              in.Ebook.Series,
			Credits:             in.Ebook.Credits,
			Summary:             in.Ebook.Summary,
			LanguageNotes:       in.Ebook.LanguageNotes,
			BookCover:           in.Ebook.BookCover,
			TitlePageImage:      in.Ebook.TitlePageImage,
			BackCover:           in.Ebook.BackCover,
			SourceLink:          in.Ebook.SourceLink,
			PgDpClearance:       in.Ebook.PgDpClearance,
			LanguageDialect:     in.Ebook.LanguageDialect,

			Downloads: nil,
		},
		Descriptions: nil,
		Work: Work{
			About:   in.Work.About,
			Comment: in.Work.Comment,
			License: CCLicense{Resource: in.Work.License.Resource},
		},
	}

	if len(in.Ebook.Issued.DataType) > 0 || len(in.Ebook.Issued.Value) > 0 {
		out.Ebook.Issued = &Issued{
			DataType: in.Ebook.Issued.DataType,
			Value:    in.Ebook.Issued.Value,
		}
	}

	if len(in.Ebook.Downloads.DataType) > 0 || in.Ebook.Downloads.Value > 0 {
		out.Ebook.Downloads = &Downloads{
			DataType: in.Ebook.Downloads.DataType,
			Value:    in.Ebook.Downloads.Value,
		}
	}

	for _, s := range in.Descriptions {
		out.Descriptions = append(out.Descriptions, description(&s))
	}

	for _, s := range in.Ebook.Bookshelves {
		shelf := Bookshelf{
			Description: description(&s.Description),
		}
		out.Ebook.Bookshelves = append(out.Ebook.Bookshelves, shelf)
	}

	for _, c := range in.Ebook.Creators {
		creator := Creator{
			Resource: c.Resource,
			Agent:    *createAgent(&c.Agent),
		}
		out.Ebook.Creators = append(out.Ebook.Creators, creator)
	}

	//
	// Now we add all the MARC relators!
	//
	for _, c := range in.Ebook.RelAdapters {
		out.Ebook.RelAdapters = append(out.Ebook.RelAdapters, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelAfterwords {
		out.Ebook.RelAfterwords = append(out.Ebook.RelAfterwords, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelAnnotators {
		out.Ebook.RelAnnotators = append(out.Ebook.RelAnnotators, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelArrangers {
		out.Ebook.RelArrangers = append(out.Ebook.RelArrangers, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelArtists {
		out.Ebook.RelArtists = append(out.Ebook.RelArtists, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelIntroductions {
		out.Ebook.RelIntroductions = append(out.Ebook.RelIntroductions, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelCommentators {
		out.Ebook.RelCommentators = append(out.Ebook.RelCommentators, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelComposers {
		out.Ebook.RelComposers = append(out.Ebook.RelComposers, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelConductors {
		out.Ebook.RelConductors = append(out.Ebook.RelConductors, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelCompilers {
		out.Ebook.RelCompilers = append(out.Ebook.RelCompilers, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelContributors {
		out.Ebook.RelContributors = append(out.Ebook.RelContributors, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelDubious {
		out.Ebook.RelDubious = append(out.Ebook.RelDubious, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelEditors {
		out.Ebook.RelEditors = append(out.Ebook.RelEditors, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelEngravers {
		out.Ebook.RelEngravers = append(out.Ebook.RelEngravers, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelIllustrators {
		out.Ebook.RelIllustrators = append(out.Ebook.RelIllustrators, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelLibrettists {
		out.Ebook.RelLibrettists = append(out.Ebook.RelLibrettists, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelOther {
		out.Ebook.RelOther = append(out.Ebook.RelOther, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelPublishers {
		out.Ebook.RelPublishers = append(out.Ebook.RelPublishers, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelPhotographers {
		out.Ebook.RelPhotographers = append(out.Ebook.RelPhotographers, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelPerformers {
		out.Ebook.RelPerformers = append(out.Ebook.RelPerformers, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelPrinters {
		out.Ebook.RelPrinters = append(out.Ebook.RelPrinters, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelResearchers {
		out.Ebook.RelResearchers = append(out.Ebook.RelResearchers, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelTranscribers {
		out.Ebook.RelTranscribers = append(out.Ebook.RelTranscribers, generateMarcRelator(&c))
	}
	for _, c := range in.Ebook.RelTranslators {
		out.Ebook.RelTranslators = append(out.Ebook.RelTranslators, generateMarcRelator(&c))
	}

	for _, s := range in.Ebook.Subjects {
		subject := Subject{
			Description: description(&s.Description),
		}
		out.Ebook.Subjects = append(out.Ebook.Subjects, subject)
	}

	for _, s := range in.Ebook.HasFormats {
		format := HasFormat{
			File: File{
				About: s.File.About,
				Extent: Extent{
					DataType: s.File.Extent.DataType,
					Value:    s.File.Extent.Value,
				},
				Modified: Modified{
					DataType: s.File.Modified.DataType,
					Value:    s.File.Modified.Value,
				},
				Formats: nil,
				IsFormatOf: IsFormatOf{
					Resource: s.File.IsFormatOf.Resource,
				},
			},
		}

		for _, f := range s.File.Formats {
			fileFormat := Format{
				Description: description(&f.Description),
			}
			format.File.Formats = append(format.File.Formats, fileFormat)
		}

		out.Ebook.HasFormats = append(out.Ebook.HasFormats, format)
	}

	return out
}

// maps the unmarshaller description
func description(d *unmarshaller.Description) Description {
	desc := Description{
		About:       d.About,
		NodeID:      d.NodeID,
		MemberOf:    nil,
		Value:       nil,
		Description: d.Description,
	}

	if len(d.Value.DataType) > 0 || len(d.Value.Data) > 0 {
		desc.Value = &Value{
			DataType: d.Value.DataType,
			Data:     d.Value.Data,
		}
	}

	if len(d.MemberOf.Resource) > 0 {
		desc.MemberOf = &MemberOf{
			Resource: d.MemberOf.Resource,
		}
	}

	return desc
}

func createAgent(in *unmarshaller.Agent) *Agent {
	out := Agent{
		About:     in.About,
		Name:      in.Name,
		Aliases:   in.Aliases,
		BirthYear: nil, // added below
		DeathYear: nil, // added below
		Webpage:   nil, // added below
	}
	if in.BirthYear.Value != 0 {
		out.BirthYear = &Year{
			DataType: in.BirthYear.DataType,
			Value:    in.BirthYear.Value,
		}
	}
	if in.DeathYear.Value != 0 {
		out.DeathYear = &Year{
			DataType: in.DeathYear.DataType,
			Value:    in.DeathYear.Value,
		}
	}
	if len(in.Webpage.Resource) > 0 {
		out.Webpage = &Webpage{
			Resource: in.Webpage.Resource,
		}
	}
	return &out
}

// NOTE: the assumption here is that a Resource will only ever exist on a
// MarcRelator when there is no Agent info!
func generateMarcRelator(c *unmarshaller.MarcRelator) MarcRelator {
	if len(c.Agent.About) == 0 {
		return MarcRelator{Resource: c.Resource}
	}
	return MarcRelator{Agent: createAgent(&c.Agent)}
}
