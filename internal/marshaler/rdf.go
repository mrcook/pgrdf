// Package marshaler contains a set of structs for generating a Project
// Gutenberg RDF XML document.
//
// NOTE: due to limitations in the Go xml package and the namespace complexity
// of the RDF documents, a separate set of marshaler and unmarshaler structs
// are required.
package marshaler

import (
	"encoding/xml"

	"github.com/mrcook/pgrdf/internal/unmarshaler"
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
	NsDcDcam  string   `xml:"xmlns:dcam,attr,omitempty"`

	Ebook        Ebook         `xml:"pgterms:ebook,omitempty"`
	Descriptions []Description `xml:"rdf:Description,omitempty"`
	Work         Work          `xml:"cc:Work,omitempty"`
}

// Ebook <pgterms:ebook /> holds the core metadata for this work.
type Ebook struct {
	About                   string     `xml:"rdf:about,attr,omitempty"`
	Titles                  []string   `xml:"dcterms:title,omitempty"`
	Alternatives            []string   `xml:"dcterms:alternative,omitempty"`
	TableOfContents         string     `xml:"dcterms:tableOfContents,omitempty"`
	Publisher               string     `xml:"dcterms:publisher,omitempty"`
	PublishedYear           int        `xml:"pgterms:marc906,omitempty"`
	Issued                  *Issued    `xml:"dcterms:issued,omitempty"`
	Summary                 string     `xml:"pgterms:marc520,omitempty"`
	Series                  []string   `xml:"pgterms:marc440,omitempty"`
	Languages               []Language `xml:"dcterms:language,omitempty"`
	LanguageDialect         string     `xml:"pgterms:marc907,omitempty"`
	LanguageNotes           []string   `xml:"pgterms:marc546,omitempty"`
	PublicationNote         string     `xml:"pgterms:marc260,omitempty"`
	EditionNote             string     `xml:"pgterms:marc250,omitempty"`
	ProductionNotes         []string   `xml:"pgterms:marc508,omitempty"`
	License                 License    `xml:"dcterms:license,omitempty"`
	Rights                  string     `xml:"dcterms:rights,omitempty"`
	DpClearanceCode         string     `xml:"pgterms:marc905,omitempty"`
	Type                    Type       `xml:"dcterms:type,omitempty"`
	Descriptions            []string   `xml:"dcterms:description,omitempty"`
	PhysicalDescriptionNote string     `xml:"pgterms:marc300,omitempty"`
	SourceLinks             []string   `xml:"pgterms:marc904,omitempty"`
	LCCN                    string     `xml:"pgterms:marc010,omitempty"`
	ISBN                    string     `xml:"pgterms:marc020,omitempty"`
	BookCoverImages         []string   `xml:"pgterms:marc901,omitempty"`
	TitlePageImage          string     `xml:"pgterms:marc902,omitempty"`
	BackCoverImage          string     `xml:"pgterms:marc903,omitempty"`
	Creators                []Creator  `xml:"dcterms:creator,omitempty"`

	// TODO: can these be marshaled programmatically?
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

	RelCollaborators []MarcRelator `xml:"marcrel:clb,omitempty"`
	RelUnknown       []MarcRelator `xml:"marcrel:unk,omitempty"`

	Subjects    []Subject   `xml:"dcterms:subject,omitempty"`
	HasFormats  []HasFormat `xml:"dcterms:hasFormat,omitempty"`
	Bookshelves []Bookshelf `xml:"pgterms:bookshelf,omitempty"`
	Downloads   *Downloads  `xml:"pgterms:downloads,omitempty"`
}

// Type <dcterms:type /> the media type of this work: text, audio, etc.
type Type struct {
	Description Description `xml:"rdf:Description"`
}

// Issued <dcterms:issued /> the date this eText was added to the Gutenberg collection.
type Issued struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Value    string `xml:",chardata"`
}

// Language <dcterms:language /> the language this work was written in.
type Language struct {
	Description Description `xml:"rdf:Description"`
}

// License <dcterms:license />
type License struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
}

// Agent <pgterms:agent /> represents a contributor to this work; author, illustrator, etc.
type Agent struct {
	About     string    `xml:"rdf:about,attr,omitempty"`
	Name      string    `xml:"pgterms:name,omitempty"`
	Aliases   []string  `xml:"pgterms:alias,omitempty"`
	BirthYear *Year     `xml:"pgterms:birthdate,omitempty"`
	DeathYear *Year     `xml:"pgterms:deathdate,omitempty"`
	Webpages  []Webpage `xml:"pgterms:webpage,omitempty"`
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
	File File `xml:"pgterms:file"`
}

// File <pgterms:file /> is a file resource.
type File struct {
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
	Description Description `xml:"rdf:Description"`
}

// Bookshelf <pgterms:bookshelf /> is a link to a Project Gutenberg bookshelf.
// that this work is part of.
type Bookshelf struct {
	Description Description `xml:"rdf:Description"`
}

// Subject <dcterms:subject /> and object representing a subject/genre.
type Subject struct {
	Description Description `xml:"rdf:Description"`
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
	About  string `xml:"rdf:about,attr,omitempty"`
	NodeID string `xml:"rdf:nodeID,attr,omitempty"`

	// depending on context only one of these is included
	Value       *Value    `xml:"rdf:value,omitempty"`
	MemberOf    *MemberOf `xml:"dcam:memberOf,omitempty"`
	Description string    `xml:"dcterms:description,omitempty"`
}

// Value <rdf:value /> for a Description object.
type Value struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Data     string `xml:",chardata"`
}

// MemberOf <dcam:memberOf /> the schema for a given subject/genre.
type MemberOf struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
}

// FromUnmarshaler is a helper function for mapping an unmarshaler.RDF to
// this marshaler.RDF object so that we have both marshal and unmarshal
// functionality for an RDF XML document.
// This is required due to the Go xml package limitations, as described at the
// top of this file.
func FromUnmarshaler(in *unmarshaler.RDF) *RDF {
	out := &RDF{
		NsBase:    in.NsBase,
		NsDcTerms: in.NsDcTerms,
		NsPgTerms: in.NsPgTerms,
		NsRdf:     in.NsRdf,
		NsRdfs:    in.NsRdfs,
		NsCC:      in.NsCC,
		NsMarcRel: in.NsMarcRel,
		NsDcDcam:  in.NsDcDcam,
		Ebook: Ebook{
			About:                   in.Ebook.About,
			Titles:                  in.Ebook.Titles,
			Alternatives:            in.Ebook.Alternatives,
			TableOfContents:         in.Ebook.TableOfContents,
			Publisher:               in.Ebook.Publisher,
			PublishedYear:           in.Ebook.PublishedYear,
			Summary:                 in.Ebook.Summary,
			Series:                  in.Ebook.Series,
			LanguageDialect:         in.Ebook.LanguageDialect,
			LanguageNotes:           in.Ebook.LanguageNotes,
			PublicationNote:         in.Ebook.PublicationNote,
			EditionNote:             in.Ebook.EditionNote,
			ProductionNotes:         in.Ebook.ProductionNotes,
			License:                 License{Resource: in.Ebook.License.Resource},
			Rights:                  in.Ebook.Rights,
			DpClearanceCode:         in.Ebook.DpClearanceCode,
			Type:                    Type{Description: description(&in.Ebook.Type.Description)},
			Descriptions:            in.Ebook.Descriptions,
			PhysicalDescriptionNote: in.Ebook.PhysicalDescriptionNote,
			SourceLinks:             in.Ebook.SourceLinks,
			LCCN:                    in.Ebook.LCCN,
			ISBN:                    in.Ebook.ISBN,
			BookCoverImages:         in.Ebook.BookCoverImages,
			TitlePageImage:          in.Ebook.TitlePageImage,
			BackCoverImage:          in.Ebook.BackCoverImage,
		},
		Work: Work{
			About:   in.Work.About,
			Comment: in.Work.Comment,
			License: CCLicense{Resource: in.Work.License.Resource},
		},
	}

	if in.Ebook.Issued != nil {
		out.Ebook.Issued = &Issued{
			DataType: in.Ebook.Issued.DataType,
			Value:    in.Ebook.Issued.Value,
		}
	}

	for _, lang := range in.Ebook.Languages {
		out.Ebook.Languages = append(out.Ebook.Languages, Language{Description: description(&lang.Description)})
	}

	for _, c := range in.Ebook.Creators {
		creator := Creator{
			Resource: c.Resource,
			Agent:    createAgent(&c.Agent),
		}
		out.Ebook.Creators = append(out.Ebook.Creators, creator)
	}

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

	for _, s := range in.Ebook.Bookshelves {
		shelf := Bookshelf{
			Description: description(&s.Description),
		}
		out.Ebook.Bookshelves = append(out.Ebook.Bookshelves, shelf)
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

	return out
}

// maps the unmarshaler description
func description(d *unmarshaler.Description) Description {
	desc := Description{
		About:       d.About,
		NodeID:      d.NodeID,
		Description: d.Description,
	}
	if d.Value != nil {
		desc.Value = &Value{
			DataType: d.Value.DataType,
			Data:     d.Value.Data,
		}
	}
	if d.MemberOf != nil {
		desc.MemberOf = &MemberOf{
			Resource: d.MemberOf.Resource,
		}
	}
	return desc
}

func createAgent(in *unmarshaler.Agent) Agent {
	if in == nil {
		return Agent{}
	}

	out := Agent{
		About:   in.About,
		Name:    in.Name,
		Aliases: in.Aliases,
	}
	if in.BirthYear != nil {
		out.BirthYear = &Year{
			DataType: in.BirthYear.DataType,
			Value:    in.BirthYear.Value,
		}
	}
	if in.DeathYear != nil {
		out.DeathYear = &Year{
			DataType: in.DeathYear.DataType,
			Value:    in.DeathYear.Value,
		}
	}
	for _, webpage := range in.Webpages {
		if len(webpage.Resource) > 0 {
			out.Webpages = append(out.Webpages, Webpage{Resource: webpage.Resource})
		}
	}
	return out
}

// NOTE: the assumption here is that a Resource will only ever exist on a
// MarcRelator when there is no Agent info!
func generateMarcRelator(c *unmarshaler.MarcRelator) MarcRelator {
	if c.Agent == nil {
		return MarcRelator{Resource: c.Resource}
	}
	agent := createAgent(c.Agent)
	return MarcRelator{Agent: &agent}
}
