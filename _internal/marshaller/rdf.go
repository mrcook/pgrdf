package marshaller

import (
	"encoding/xml"

	"github.com/mrcook/pgrdf/_internal/unmarshaller"
)

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

	Work         Work          `xml:"cc:Work,omitempty"`
	Ebook        Ebook         `xml:"pgterms:ebook,omitempty"`
	Descriptions []Description `xml:"rdf:Description,omitempty"`
}

type Work struct {
	About   string    `xml:"rdf:about,attr,omitempty"`
	Comment string    `xml:"rdfs:comment,omitempty"`
	License CCLicense `xml:"cc:license,omitempty"`
}

type CCLicense struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
}

type Ebook struct {
	About               string        `xml:"rdf:about,attr,omitempty"`
	Summary             string        `xml:"pgterms:marc520,omitempty"`
	Description         string        `xml:"dcterms:description,omitempty"`
	Type                Type          `xml:"dcterms:type,omitempty"`
	Issued              *Issued       `xml:"dcterms:issued,omitempty"`
	Language            Language      `xml:"dcterms:language,omitempty"`
	LanguageDialect     string        `xml:"pgterms:marc907,omitempty"`
	LanguageNotes       string        `xml:"pgterms:marc546,omitempty"`
	Publisher           string        `xml:"dcterms:publisher,omitempty"`
	PublishedYear       int           `xml:"pgterms:marc906,omitempty"`
	License             License       `xml:"dcterms:license,omitempty"`
	Rights              string        `xml:"dcterms:rights,omitempty"`
	PgDpClearance       string        `xml:"pgterms:marc905,omitempty"`
	Title               string        `xml:"dcterms:title,omitempty"`
	Alternative         []string      `xml:"dcterms:alternative,omitempty"`
	Creators            []Creator     `xml:"dcterms:creator,omitempty"`
	Editors             []Editor      `xml:"marcrel:edt,omitempty"`
	Illustrators        []Illustrator `xml:"marcrel:ill,omitempty"`
	Translators         []Translator  `xml:"marcrel:trl,omitempty"`
	Subjects            []Subject     `xml:"dcterms:subject,omitempty"`
	HasFormats          []HasFormat   `xml:"dcterms:hasFormat,omitempty"`
	Bookshelves         []Bookshelf   `xml:"pgterms:bookshelf,omitempty"`
	Series              string        `xml:"pgterms:marc440,omitempty"`
	BookCover           string        `xml:"pgterms:marc901,omitempty"`
	TitlePageImage      string        `xml:"pgterms:marc902,omitempty"`
	BackCover           string        `xml:"pgterms:marc903,omitempty"`
	Edition             string        `xml:"pgterms:marc250,omitempty"`
	OriginalPublication string        `xml:"pgterms:marc260,omitempty"`
	SourceDescription   string        `xml:"pgterms:marc300,omitempty"`
	SourceLink          string        `xml:"pgterms:marc904,omitempty"`
	Credits             []string      `xml:"pgterms:marc508,omitempty"`
	LOC                 string        `xml:"pgterms:marc010,omitempty"`
	ISBN                string        `xml:"pgterms:marc020,omitempty"`
	Downloads           *Downloads    `xml:"pgterms:downloads,omitempty"`
}

type Agent struct {
	About     string   `xml:"rdf:about,attr,omitempty"`
	Name      string   `xml:"pgterms:name,omitempty"`
	Aliases   []string `xml:"pgterms:alias,omitempty"`
	Birthdate *Year    `xml:"pgterms:birthdate,omitempty"`
	Deathdate *Year    `xml:"pgterms:deathdate,omitempty"`
	Webpage   *Webpage `xml:"pgterms:webpage,omitempty"`
}

type Bookshelf struct {
	Description Description
}

type Creator struct {
	Agent Agent `xml:"pgterms:agent"`
}

type Editor struct {
	Agent Agent `xml:"pgterms:agent"`
}

type Illustrator struct {
	Agent Agent `xml:"pgterms:agent"`
}

type Translator struct {
	Agent Agent `xml:"pgterms:agent"`
}

type Description struct {
	XMLName     xml.Name  `xml:"rdf:Description"`
	About       string    `xml:"rdf:about,attr,omitempty"`
	NodeID      string    `xml:"rdf:nodeID,attr,omitempty"`
	Value       *Value    `xml:"rdf:value,omitempty"`
	MemberOf    *MemberOf `xml:"dcam:memberOf,omitempty"`
	Description string    `xml:"dcterms:description,omitempty"`
}

type Downloads struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Value    int    `xml:",chardata"`
}

type Extent struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Value    int    `xml:",chardata"`
}

type File struct {
	XMLName    xml.Name   `xml:"pgterms:file"`
	About      string     `xml:"rdf:about,attr,omitempty"`
	Extent     Extent     `xml:"dcterms:extent,omitempty"`
	Modified   Modified   `xml:"dcterms:modified,omitempty"`
	IsFormatOf IsFormatOf `xml:"dcterms:isFormatOf"`
	Formats    []Format   `xml:"dcterms:format,omitempty"`
}

type Format struct {
	Description Description
}

type HasFormat struct {
	XMLName xml.Name `xml:"dcterms:hasFormat"`
	File    File
}

type IsFormatOf struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
}

type Issued struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Value    string `xml:",chardata"`
}

type Language struct {
	Description Description
}

type License struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
}

type MemberOf struct {
	XMLName  xml.Name `xml:"dcam:memberOf"`
	Resource string   `xml:"rdf:resource,attr,omitempty"`
}

type Modified struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Value    string `xml:",chardata"`
}

type Subject struct {
	Description Description
}

type Type struct {
	Description Description
}

type Value struct {
	XMLName  xml.Name `xml:"rdf:value"`
	DataType string   `xml:"rdf:datatype,attr,omitempty"`
	Data     string   `xml:",chardata"`
}

type Webpage struct {
	Resource string `xml:"rdf:resource,attr,omitempty"`
}

type Year struct {
	DataType string `xml:"rdf:datatype,attr,omitempty"`
	Value    int    `xml:",chardata"`
}

func FromUnmarshaller(r *unmarshaller.RDF) *RDF {
	m := &RDF{
		NsBase:    r.NsBase,
		NsPgTerms: r.NsPgTerms,
		NsCC:      r.NsCC,
		NsRdf:     r.NsRdf,
		NsDCam:    r.NsDcam,
		NsRdfs:    r.NsRdfs,
		NsDcTerms: r.NsDcTerms,
		Work: Work{
			About:   r.Work.About,
			Comment: r.Work.Comment,
			License: CCLicense{Resource: r.Work.License.Resource},
		},
		Ebook: Ebook{
			About:       r.Ebook.About,
			Description: r.Ebook.Description,
			Bookshelves: nil,
			Title:       r.Ebook.Title,
			Alternative: r.Ebook.Alternative,
			Creators:    nil,
			Subjects:    nil,
			Publisher:   r.Ebook.Publisher,
			Issued:      nil,
			Rights:      r.Ebook.Rights,
			License:     License{Resource: r.Ebook.License.Resource},
			Downloads:   nil,
			HasFormats:  nil,
			Type: Type{
				Description: description(&r.Ebook.Type.Description),
			},
			Language: Language{
				Description: description(&r.Ebook.Language.Description),
			},
		},
		Descriptions: nil,
	}

	if len(r.Ebook.Issued.DataType) > 0 || len(r.Ebook.Issued.Value) > 0 {
		m.Ebook.Issued = &Issued{
			DataType: r.Ebook.Issued.DataType,
			Value:    r.Ebook.Issued.Value,
		}
	}

	if len(r.Ebook.Downloads.DataType) > 0 || r.Ebook.Downloads.Value > 0 {
		m.Ebook.Downloads = &Downloads{
			DataType: r.Ebook.Downloads.DataType,
			Value:    r.Ebook.Downloads.Value,
		}
	}

	for _, s := range r.Descriptions {
		m.Descriptions = append(m.Descriptions, description(&s))
	}

	for _, s := range r.Ebook.Bookshelves {
		shelf := Bookshelf{
			Description: description(&s.Description),
		}
		m.Ebook.Bookshelves = append(m.Ebook.Bookshelves, shelf)
	}

	for _, c := range r.Ebook.Creators {
		creator := Creator{
			Agent: Agent{
				About:     c.Agent.About,
				Webpage:   nil,
				Birthdate: nil,
				Aliases:   c.Agent.Aliases,
				Deathdate: nil,
				Name:      c.Agent.Name,
			},
		}

		if len(c.Agent.Webpage.Resource) > 0 {
			creator.Agent.Webpage = &Webpage{
				Resource: c.Agent.Webpage.Resource,
			}
		}

		if c.Agent.Birthdate.Value != 0 {
			creator.Agent.Birthdate = &Year{
				DataType: c.Agent.Birthdate.DataType,
				Value:    c.Agent.Birthdate.Value,
			}
		}
		if c.Agent.Deathdate.Value != 0 {
			creator.Agent.Deathdate = &Year{
				DataType: c.Agent.Deathdate.DataType,
				Value:    c.Agent.Deathdate.Value,
			}
		}

		m.Ebook.Creators = append(m.Ebook.Creators, creator)
	}

	for _, s := range r.Ebook.Subjects {
		subject := Subject{
			Description: description(&s.Description),
		}
		m.Ebook.Subjects = append(m.Ebook.Subjects, subject)
	}

	for _, s := range r.Ebook.HasFormats {
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

		m.Ebook.HasFormats = append(m.Ebook.HasFormats, format)
	}

	return m
}

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
