// Package unmarshaler contains a set of structs for unmarshaling a Project
// Gutenberg RDF XML document.
//
// NOTE: due to limitations in the Go xml package and the namespace complexity
// of the RDF documents, a separate set of marshaler and unmarshaler structs
// are required.
package unmarshaler

import (
	"encoding/xml"
	"io"
	"strconv"
	"strings"
)

func New(r io.Reader) (*RDF, error) {
	rdf := &RDF{}

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(data, rdf); err != nil {
		return nil, err
	}

	// convert the year string to an int after unmarshaling
	rdf.Ebook.PublishedYear, _ = strconv.Atoi(rdf.Ebook.PublishedYearString)

	return rdf, nil
}

type RDF struct {
	XMLName xml.Name `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# RDF"`

	// all the namespace declarations contains in the RDF.
	NsBase    string `xml:"base,attr"`
	NsDcTerms string `xml:"xmlns dcterms,attr"`
	NsPgTerms string `xml:"xmlns pgterms,attr"`
	NsRdf     string `xml:"xmlns rdf,attr"`
	NsRdfs    string `xml:"xmlns rdfs,attr"`
	NsCC      string `xml:"xmlns cc,attr"`
	NsMarcRel string `xml:"xmlns marcrel,attr"`
	NsDcDcam  string `xml:"xmlns dcam,attr"`

	// The main ebook content.
	Ebook Ebook `xml:"http://www.gutenberg.org/2009/pgterms/ ebook"`

	// This description tag contains links to author biographies, usually to Wikipedia.
	Descriptions []Description `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Description"`

	// Creative Commons information about the work.
	Work Work `xml:"http://web.resource.org/cc/ Work"`
}

type Ebook struct {
	// Contains the PG eText number.
	About string `xml:"about,attr"`

	// Title(s) of the work.
	Titles []string `xml:"http://purl.org/dc/terms/ title"`

	// An alternate set of titles for the work.
	Alternatives []string `xml:"http://purl.org/dc/terms/ alternative"`

	// The table of contents for a books.
	TableOfContents string `xml:"http://purl.org/dc/terms/ tableOfContents"`

	// Publisher of this work, which is always "Project Gutenberg".
	Publisher string `xml:"http://purl.org/dc/terms/ publisher"`

	// Original year of publication for this work.
	//
	// NOTE: at least one RDF uses `Various` for marc906 instead of a year integer.
	// This must be parsed as a string then manually converted once the unmarshal is complete.
	PublishedYear       int
	PublishedYearString string `xml:"http://www.gutenberg.org/2009/pgterms/ marc906"`

	// The date PG released this work.
	Issued *Issued `xml:"http://purl.org/dc/terms/ issued"`

	// A short summary of the work -- currently very few RDFs provide this.
	Summary string `xml:"http://www.gutenberg.org/2009/pgterms/ marc520"`

	// Series which this work belongs to, with many entries also including the volume number.
	Series []string `xml:"http://www.gutenberg.org/2009/pgterms/ marc440"`

	// Language the work contains - most use ISO 639-1, but some are ISO 639-3.
	// Usually the primary language is placed first.
	Languages []Language `xml:"http://purl.org/dc/terms/ language"`

	// Language dialect for the primary language, in ISO 3166-2 format.
	LanguageDialect string `xml:"http://www.gutenberg.org/2009/pgterms/ marc907"`

	// Notes on the language used in the work, e.g. "Uses 19th Century spelling".
	LanguageNotes []string `xml:"http://www.gutenberg.org/2009/pgterms/ marc546"`

	// Publication note of the source material, e.g. publisher, city, and year.
	PublicationNote string `xml:"http://www.gutenberg.org/2009/pgterms/ marc260"`

	// Edition note for this work, e.g. "2nd Edition".
	EditionNote string `xml:"http://www.gutenberg.org/2009/pgterms/ marc250"`

	// ProductionNotes for this release. Can also include "updated" dates,
	// either as a separate entry, or as part of the credit info.
	ProductionNotes []string `xml:"http://www.gutenberg.org/2009/pgterms/ marc508"`

	// License: tag is always empty and `resource` attr is always "license".
	License License `xml:"http://purl.org/dc/terms/ license"`

	// The copyright terms of the work. Usually "Public domain in the USA", or a copyrighted message.
	Rights string `xml:"http://purl.org/dc/terms/ rights"`

	// Distributed Proofreaders clearance code, e.g. "20050213050736stahl".
	DpClearanceCode string `xml:"http://www.gutenberg.org/2009/pgterms/ marc905"`

	// Type of this work, always one of: Collection, Dataset, Image, MovingImage, Sound, StillImage, Text.
	Type Type `xml:"http://purl.org/dc/terms/ type"`

	// A general description of the work.
	Descriptions []string `xml:"http://purl.org/dc/terms/ description"`

	// A description of the physical attributes of the source of this work,
	// e.g. "5 pages : illustrations, map, portraits.". Currently very few RDFs provide this.
	PhysicalDescriptionNote string `xml:"http://www.gutenberg.org/2009/pgterms/ marc300"`

	// Links to information about the source of this work.
	SourceLinks []string `xml:"http://www.gutenberg.org/2009/pgterms/ marc904"`

	// Library of Congress Control Number.
	LCCN string `xml:"http://www.gutenberg.org/2009/pgterms/ marc010"`

	// Original ISBN of this work. Currently only one RDF includes this.
	ISBN string `xml:"http://www.gutenberg.org/2009/pgterms/ marc020"`

	// Cover images for this work, including book covers and title pages.
	// Some are URLs and some are `file://` URIs.
	BookCoverImages []string `xml:"http://www.gutenberg.org/2009/pgterms/ marc901"`

	// Title page image URL. Currently very few RDFs provide this.
	TitlePageImage string `xml:"http://www.gutenberg.org/2009/pgterms/ marc902"`

	// Back cover image URL. Currently very few RDFs provide this.
	BackCoverImage string `xml:"http://www.gutenberg.org/2009/pgterms/ marc903"`

	// Creators of this work, i.e. the authors.
	Creators []Creator `xml:"http://purl.org/dc/terms/ creator"`

	// Contributors to the work, recorded as MARC Relator codes: aut, edt, ill, etc.
	// TODO: can these be unmarshaled programmatically?
	RelAdapters      []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ adp"`
	RelAfterwords    []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ aft"`
	RelAnnotators    []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ ann"`
	RelArrangers     []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ arr"`
	RelArtists       []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ art"`
	RelIntroductions []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ aui"`
	RelAuthors       []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ aut"`
	RelCommentators  []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ cmm"`
	RelComposers     []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ cmp"`
	RelConductors    []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ cnd"`
	RelCompilers     []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ com"`
	RelContributors  []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ ctb"`
	RelDubious       []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ dub"`
	RelEditors       []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ edt"`
	RelEngravers     []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ egr"`
	RelIllustrators  []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ ill"`
	RelLibrettists   []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ lbt"`
	RelOther         []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ oth"`
	RelPublishers    []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ pbl"`
	RelPhotographers []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ pht"`
	RelPerformers    []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ prf"`
	RelPrinters      []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ prt"`
	RelResearchers   []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ res"`
	RelTranscribers  []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ trc"`
	RelTranslators   []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ trl"`

	// NOTE: `clb` is a deprecated code (only pg6948.rdf uses this).
	RelCollaborators []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ clb"` // deprecated

	// NOTE: `unk` is not an official Realtor code (only 8 RDFs use this).
	RelUnknown []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ unk"`

	// Subjects, using LCSH and LCC codes.
	Subjects []Subject `xml:"http://purl.org/dc/terms/ subject"`

	// Metadata for all files associated with this work.
	HasFormats []HasFormat `xml:"http://purl.org/dc/terms/ hasFormat"`

	// The Project Gutenberg bookshelves this work belongs to.
	Bookshelves []Bookshelf `xml:"http://www.gutenberg.org/2009/pgterms/ bookshelf"`

	// The number of times this work has been downloaded in the 30 days prior
	// to the RDF being generated.
	Downloads *Downloads `xml:"http://www.gutenberg.org/2009/pgterms/ downloads"`
}

// Id taken from the about attribute.
func (e *Ebook) Id() int {
	return extractIdFromAttr(e.About)
}

type Type struct {
	Description Description `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Description"`
}

type Issued struct {
	DataType string `xml:"datatype,attr"`
	Value    string `xml:",chardata"`
}

type Language struct {
	Description Description `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Description"`
}

type License struct {
	Resource string `xml:"resource,attr"`
}

// Agent metadata used in creator and marcrel tags.
type Agent struct {
	// About holds the PG agent ID, e.g. `2009/agents/10`.
	About string `xml:"about,attr"`

	// Primary name of the author/agent.
	Name string `xml:"http://www.gutenberg.org/2009/pgterms/ name"`

	// A list of aliases for the author.
	Aliases []string `xml:"http://www.gutenberg.org/2009/pgterms/ alias"`

	// Year of Birth of the person.
	BirthYear *Year `xml:"http://www.gutenberg.org/2009/pgterms/ birthdate"`

	// Year of death of the person.
	DeathYear *Year `xml:"http://www.gutenberg.org/2009/pgterms/ deathdate"`

	// URLs linking to biographies for the person, usually Wikipedia.
	Webpages []Webpage `xml:"http://www.gutenberg.org/2009/pgterms/ webpage"`
}

// Id taken from the about attribute.
func (a *Agent) Id() int {
	return extractIdFromAttr(a.About)
}

type Year struct {
	DataType string `xml:"datatype,attr"`
	Value    int    `xml:",chardata"`
}

type Webpage struct {
	Resource string `xml:"resource,attr"`
}

type Creator struct {
	Resource string `xml:"resource,attr"`
	Agent    Agent  `xml:"http://www.gutenberg.org/2009/pgterms/ agent"`
}

// AgentId is the Gutenberg ID for this agent. Usually this is taken from the
// Agent `About` field, however, in rare situations when this is blank, the
// Creator `Resource` needs to be used.
func (c Creator) AgentId() int {
	if len(c.Agent.About) > 0 {
		return c.Agent.Id()
	}
	return extractIdFromAttr(c.Resource)
}

type MarcRelator struct {
	Resource string `xml:"resource,attr"`
	Agent    *Agent `xml:"http://www.gutenberg.org/2009/pgterms/ agent"`
}

func (m MarcRelator) AgentId() int {
	if len(m.Resource) > 0 {
		return extractIdFromAttr(m.Resource)
	}
	return m.Agent.Id()
}

type HasFormat struct {
	File File `xml:"http://www.gutenberg.org/2009/pgterms/ file"`
}

// File metadata
type File struct {
	// A URL where the file can be downloaded.
	About string `xml:"about,attr"`

	// Size of the file in bytes.
	Extent Extent `xml:"http://purl.org/dc/terms/ extent"`

	// Date the file was last modified.
	Modified Modified `xml:"http://purl.org/dc/terms/ modified"`

	// Contains only the ID of the work, e.g. `ebooks/11`
	IsFormatOf IsFormatOf `xml:"http://purl.org/dc/terms/ isFormatOf"`

	// The mimetype of the file, e.g. `image/jpeg`
	Formats []Format `xml:"http://purl.org/dc/terms/ format"`
}

type Extent struct {
	DataType string `xml:"datatype,attr"`
	Value    int    `xml:",chardata"`
}

type Modified struct {
	DataType string `xml:"datatype,attr"`
	Value    string `xml:",chardata"`
}

type IsFormatOf struct {
	Resource string `xml:"resource,attr"`
}

type Format struct {
	Description Description `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Description"`
}

type Bookshelf struct {
	Description Description `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Description"`
}

type Subject struct {
	Description Description `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Description"`
}

type Downloads struct {
	DataType string `xml:"datatype,attr"`
	Value    int    `xml:",chardata"`
}

// Work holds Creative Commons information about this work.
type Work struct {
	// Currently no RDF contains a value, but the attribute is always present.
	About string `xml:"about,attr"`

	// Comment from PG about where to find all the RDFs.
	Comment string `xml:"http://www.w3.org/2000/01/rdf-schema# comment"`

	// Creative Commons license URL.
	License CCLicense `xml:"http://web.resource.org/cc/ license"`
}

type CCLicense struct {
	Resource string `xml:"resource,attr"`
}

// Description is used in various different tags, e.g. languages, type, subjects tags.
// Depending on context the included child tags (Value, MemberOf, and Description)
// can change, i.e. <dcterms:language> only uses the Value tag.
type Description struct {
	About  string `xml:"about,attr"`
	NodeID string `xml:"nodeID,attr"`

	// depending on context only one of these is included
	Value       *Value    `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# value"`
	MemberOf    *MemberOf `xml:"http://purl.org/dc/dcam/ memberOf"`
	Description string    `xml:"http://purl.org/dc/terms/ description"`
}

type Value struct {
	DataType string `xml:"datatype,attr"`
	Data     string `xml:",chardata"`
}

type MemberOf struct {
	Resource string `xml:"resource,attr"`
}

// Extracts the ID from an attribute, e.g. the string `2009/agents/7`
func extractIdFromAttr(attr string) int {
	parts := strings.Split(attr, "/")
	idString := parts[len(parts)-1]
	id, _ := strconv.Atoi(idString)
	return id
}
