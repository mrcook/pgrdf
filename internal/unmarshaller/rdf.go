// Package unmarshaller contains a set of structs for unmarshalling a Project
// Gutenberg RDF XML document.
//
// NOTE: due to limitations in the Go xml package and the namespace complexity
// of the RDF documents, a separate set of marshaller and unmarshaller structs
// are required.
package unmarshaller

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

	// convert the year string to an int after unmarshalling
	rdf.Ebook.PublishedYear, _ = strconv.Atoi(rdf.Ebook.PublishedYearString)

	return rdf, nil
}

type RDF struct {
	XMLName   xml.Name `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# RDF"`
	NsBase    string   `xml:"base,attr"`
	NsDcTerms string   `xml:"xmlns dcterms,attr"`
	NsPgTerms string   `xml:"xmlns pgterms,attr"`
	NsRdf     string   `xml:"xmlns rdf,attr"`
	NsRdfs    string   `xml:"xmlns rdfs,attr"`
	NsCC      string   `xml:"xmlns cc,attr"`
	NsMarcRel string   `xml:"xmlns marcrel,attr"`
	NsDcDcam  string   `xml:"xmlns dcam,attr"`

	Ebook        Ebook         `xml:"http://www.gutenberg.org/2009/pgterms/ ebook"`
	Descriptions []Description `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Description"`
	Work         Work          `xml:"http://web.resource.org/cc/ Work"`
}

type Ebook struct {
	About       string   `xml:"about,attr"`
	Title       string   `xml:"http://purl.org/dc/terms/ title"`
	Alternative []string `xml:"http://purl.org/dc/terms/ alternative"`
	Publisher   string   `xml:"http://purl.org/dc/terms/ publisher"`

	// NOTE: at least one RDF uses `Various` for marc906 instead of a year number, so this
	// must be parsed as a string, and then manually converted once the unmarshal is complete.
	PublishedYear       int
	PublishedYearString string `xml:"http://www.gutenberg.org/2009/pgterms/ marc906"`

	Issued              Issued     `xml:"http://purl.org/dc/terms/ issued"`
	Summary             string     `xml:"http://www.gutenberg.org/2009/pgterms/ marc520"`
	Series              []string   `xml:"http://www.gutenberg.org/2009/pgterms/ marc440"`
	Languages           []Language `xml:"http://purl.org/dc/terms/ language"`
	LanguageDialect     string     `xml:"http://www.gutenberg.org/2009/pgterms/ marc907"`
	LanguageNotes       []string   `xml:"http://www.gutenberg.org/2009/pgterms/ marc546"`
	OriginalPublication string     `xml:"http://www.gutenberg.org/2009/pgterms/ marc260"`
	Edition             string     `xml:"http://www.gutenberg.org/2009/pgterms/ marc250"`
	Credits             []string   `xml:"http://www.gutenberg.org/2009/pgterms/ marc508"`
	License             License    `xml:"http://purl.org/dc/terms/ license"`
	Rights              string     `xml:"http://purl.org/dc/terms/ rights"`
	PgDpClearance       string     `xml:"http://www.gutenberg.org/2009/pgterms/ marc905"`
	Type                Type       `xml:"http://purl.org/dc/terms/ type"`
	Description         string     `xml:"http://purl.org/dc/terms/ description"`
	SourceDescription   string     `xml:"http://www.gutenberg.org/2009/pgterms/ marc300"`
	SourceLinks         []string   `xml:"http://www.gutenberg.org/2009/pgterms/ marc904"`
	LOC                 string     `xml:"http://www.gutenberg.org/2009/pgterms/ marc010"`
	ISBN                string     `xml:"http://www.gutenberg.org/2009/pgterms/ marc020"`
	BookCovers          []string   `xml:"http://www.gutenberg.org/2009/pgterms/ marc901"`
	TitlePageImage      string     `xml:"http://www.gutenberg.org/2009/pgterms/ marc902"`
	BackCover           string     `xml:"http://www.gutenberg.org/2009/pgterms/ marc903"`
	Creators            []Creator  `xml:"http://purl.org/dc/terms/ creator"`

	// TODO: can these be unmarshalled programmatically?
	RelAdapters      []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ adp"`
	RelAfterwords    []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ aft"`
	RelAnnotators    []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ ann"`
	RelArrangers     []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ arr"`
	RelArtists       []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ art"`
	RelIntroductions []MarcRelator `xml:"http://id.loc.gov/vocabulary/relators/ aui"`
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

	Subjects    []Subject   `xml:"http://purl.org/dc/terms/ subject"`
	HasFormats  []HasFormat `xml:"http://purl.org/dc/terms/ hasFormat"`
	Bookshelves []Bookshelf `xml:"http://www.gutenberg.org/2009/pgterms/ bookshelf"`
	Downloads   Downloads   `xml:"http://www.gutenberg.org/2009/pgterms/ downloads"`
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

type Agent struct {
	About     string   `xml:"about,attr"`
	Name      string   `xml:"name"`
	Aliases   []string `xml:"alias"`
	BirthYear Year     `xml:"birthdate"`
	DeathYear Year     `xml:"deathdate"`
	Webpage   Webpage  `xml:"http://www.gutenberg.org/2009/pgterms/ webpage"`
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

type MarcRelator struct {
	Resource string `xml:"resource,attr"`
	Agent    Agent  `xml:"http://www.gutenberg.org/2009/pgterms/ agent"`
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

type File struct {
	About      string     `xml:"about,attr"`
	Extent     Extent     `xml:"http://purl.org/dc/terms/ extent"`
	Modified   Modified   `xml:"http://purl.org/dc/terms/ modified"`
	IsFormatOf IsFormatOf `xml:"http://purl.org/dc/terms/ isFormatOf"`
	Formats    []Format   `xml:"http://purl.org/dc/terms/ format"`
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

type Work struct {
	About   string    `xml:"about,attr"`
	Comment string    `xml:"http://www.w3.org/2000/01/rdf-schema# comment"`
	License CCLicense `xml:"http://web.resource.org/cc/ license"`
}

type CCLicense struct {
	Resource string `xml:"resource,attr"`
}

type Description struct {
	About       string   `xml:"about,attr"`
	NodeID      string   `xml:"nodeID,attr"`
	Value       Value    `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# value"`
	MemberOf    MemberOf `xml:"http://purl.org/dc/dcam/ memberOf"`
	Description string   `xml:"http://purl.org/dc/terms/ description"`
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
