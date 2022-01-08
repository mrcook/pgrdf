package unmarshaller

import (
	"encoding/xml"
	"io"
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
	NsDcam    string   `xml:"xmlns dcam,attr"`

	Work         Work
	Ebook        Ebook
	Descriptions []Description `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Description"`
}

type Work struct {
	XMLName xml.Name `xml:"http://web.resource.org/cc/ Work"`
	About   string   `xml:"about,attr"`
	Comment string   `xml:"http://www.w3.org/2000/01/rdf-schema# comment"`
	License CCLicense
}

type CCLicense struct {
	XMLName  xml.Name `xml:"http://web.resource.org/cc/ license"`
	Resource string   `xml:"resource,attr"`
}

type Ebook struct {
	XMLName       xml.Name `xml:"http://www.gutenberg.org/2009/pgterms/ ebook"`
	About         string   `xml:"about,attr"`
	Description   string   `xml:"http://purl.org/dc/terms/ description"`
	Type          Type
	Issued        Issued
	Language      Language
	Publisher     string `xml:"http://purl.org/dc/terms/ publisher"`
	PublishedYear int    `xml:"http://www.gutenberg.org/2009/pgterms/ marc906"`
	License       License
	Rights        string      `xml:"http://purl.org/dc/terms/ rights"`
	Title         string      `xml:"http://purl.org/dc/terms/ title"`
	Alternative   []string    `xml:"http://purl.org/dc/terms/ alternative"`
	Creators      []Creator   `xml:"http://purl.org/dc/terms/ creator"`
	Subjects      []Subject   `xml:"http://purl.org/dc/terms/ subject"`
	HasFormats    []HasFormat `xml:"http://purl.org/dc/terms/ hasFormat"`
	Bookshelves   []Bookshelf `xml:"http://www.gutenberg.org/2009/pgterms/ bookshelf"`
	Downloads     Downloads
}

type Agent struct {
	XMLName   xml.Name `xml:"http://www.gutenberg.org/2009/pgterms/ agent"`
	About     string   `xml:"about,attr"`
	Name      string   `xml:"name"`
	Aliases   []string `xml:"alias"`
	Birthdate Year     `xml:"birthdate"`
	Deathdate Year     `xml:"deathdate"`
	Webpage   Webpage
}

type Bookshelf struct {
	XMLName     xml.Name `xml:"http://www.gutenberg.org/2009/pgterms/ bookshelf"`
	Description Description
}

type Creator struct {
	XMLName xml.Name `xml:"http://purl.org/dc/terms/ creator"`
	Agent   Agent
}

type Description struct {
	XMLName     xml.Name `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Description"`
	About       string   `xml:"about,attr"`
	NodeID      string   `xml:"nodeID,attr"`
	Value       Value
	MemberOf    MemberOf
	Description string `xml:"http://purl.org/dc/terms/ description"`
}

type Downloads struct {
	XMLName  xml.Name `xml:"http://www.gutenberg.org/2009/pgterms/ downloads"`
	DataType string   `xml:"datatype,attr"`
	Value    int      `xml:",chardata"`
}

type Extent struct {
	XMLName  xml.Name `xml:"http://purl.org/dc/terms/ extent"`
	DataType string   `xml:"datatype,attr"`
	Value    int      `xml:",chardata"`
}

type File struct {
	XMLName    xml.Name `xml:"http://www.gutenberg.org/2009/pgterms/ file"`
	About      string   `xml:"about,attr"`
	Extent     Extent
	Modified   Modified
	IsFormatOf IsFormatOf
	Formats    []Format `xml:"http://purl.org/dc/terms/ format"`
}

type Format struct {
	XMLName     xml.Name `xml:"http://purl.org/dc/terms/ format"`
	Description Description
}

type HasFormat struct {
	XMLName xml.Name `xml:"http://purl.org/dc/terms/ hasFormat"`
	File    File
}

type IsFormatOf struct {
	XMLName  xml.Name `xml:"http://purl.org/dc/terms/ isFormatOf"`
	Resource string   `xml:"resource,attr"`
}

type Issued struct {
	XMLName  xml.Name `xml:"http://purl.org/dc/terms/ issued"`
	DataType string   `xml:"datatype,attr"`
	Value    string   `xml:",chardata"`
}

type Language struct {
	XMLName     xml.Name `xml:"http://purl.org/dc/terms/ language"`
	Description Description
}

type License struct {
	XMLName  xml.Name `xml:"http://purl.org/dc/terms/ license"`
	Resource string   `xml:"resource,attr"`
}

type MemberOf struct {
	XMLName  xml.Name `xml:"http://purl.org/dc/dcam/ memberOf"`
	Resource string   `xml:"resource,attr"`
}

type Modified struct {
	XMLName  xml.Name `xml:"http://purl.org/dc/terms/ modified"`
	DataType string   `xml:"datatype,attr"`
	Value    string   `xml:",chardata"`
}

type Subject struct {
	XMLName     xml.Name `xml:"http://purl.org/dc/terms/ subject"`
	Description Description
}

type Type struct {
	XMLName     xml.Name `xml:"http://purl.org/dc/terms/ type"`
	Description Description
}

type Value struct {
	XMLName  xml.Name `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# value"`
	DataType string   `xml:"datatype,attr"`
	Data     string   `xml:",chardata"`
}

type Webpage struct {
	XMLName  xml.Name `xml:"http://www.gutenberg.org/2009/pgterms/ webpage"`
	Resource string   `xml:"resource,attr"`
}

type Year struct {
	DataType string `xml:"datatype,attr"`
	Value    int    `xml:",chardata"`
}
