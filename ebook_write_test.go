package pgrdf_test

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/mrcook/pgrdf"
)

var nodeIdRE = regexp.MustCompile(`nodeID="N[^"]+"`)

func TestEbook_WriteRDF(t *testing.T) {
	ebook := generateEbook()

	w := bytes.NewBuffer([]byte{})
	err := ebook.WriteRDF(w)
	if err != nil {
		t.Fatalf("error marshaling ebook: %s", err)
	}

	data := w.String()

	// nodeID are auto-generated so replace them before checking
	data = nodeIdRE.ReplaceAllString(data, `nodeID="N"`)
	rdfMarshalExpected = nodeIdRE.ReplaceAllString(rdfMarshalExpected, `nodeID="N"`)
	if data == rdfMarshalExpected {
		return
	}

	if len(data) != len(rdfMarshalExpected) {
		t.Fatalf("expected the length to be %d, got %d", len(data), len(rdfMarshalExpected))
	}

	// show where the diversion happens
	var index int
	for i := 0; i < len(data); i++ {
		if data[i] != rdfMarshalExpected[i] {
			index = i
			break
		}
	}
	t.Errorf("unexpected marshaled output at position %d\n%s\n", index, data[0:index])
}

func generateEbook() *pgrdf.Ebook {
	return &pgrdf.Ebook{
		ID:                      11,
		Titles:                  []string{"Alice's Adventures in Wonderland"},
		AlternateTitles:         []string{"Alice in Wonderland"},
		TableOfContents:         "Prefatory Note -- Chapter 1 -- Chapter 2 -- Chapter 3 -- Conclusion",
		Publisher:               "Project Gutenberg",
		PublishedYear:           1909,
		ReleaseDate:             "2008-06-27",
		Summary:                 "A short story about short summaries.",
		Series:                  []string{"Best of Fantasy"},
		Languages:               []string{"en"},
		LanguageDialect:         "GB",
		LanguageNotes:           []string{"Uses 19th century spelling."},
		PublicationNote:         "United Kingdom: J. Johnson, 1794.",
		EditionNote:             "2nd Edition",
		ProductionNotes:         []string{"D. Smith", "Updated: 2022-07-14"},
		Copyright:               "Public domain in the USA.",
		CopyrightClearanceCode:  "20001231235959randomthing",
		BookType:                pgrdf.BookTypeText,
		Notes:                   []string{"An improved version is available at #28885."},
		PhysicalDescriptionNote: "5 pages : illustrations, map, portraits.",
		SourceLinks:             []string{"https://example.com/ebooks/11/something"},
		LCCN:                    "77177892",
		ISBN:                    "978-0-919366-14-5",
		BookCovers:              []string{"https://www.gutenberg.org/cache/epub/11/pg11.cover.medium.jpg"},
		TitlePageImage:          "https://example.org/ebook11/title.jpg",
		BackCover:               "https://example.org/ebook11/back.jpg",
		Creators: []pgrdf.Creator{{
			ID:       7,
			Name:     "Carroll, Lewis",
			Aliases:  []string{"Dodgson, Charles Lutwidge"},
			Born:     1832,
			Died:     1898,
			WebPages: []string{"https://en.wikipedia.org/wiki/Lewis_Carroll"},
		}},
		Subjects: []pgrdf.Subject{{
			Heading: "Fantasy fiction",
			Schema:  "http://purl.org/dc/terms/LCSH",
		}},
		Files: []pgrdf.File{{
			URL:       "https://www.gutenberg.org/files/11/11-0.txt",
			Extent:    174693,
			Modified:  "2020-10-12T03:45:53",
			Encodings: []string{"text/plain; charset=utf-8"},
		}},
		Bookshelves: []pgrdf.Bookshelf{{
			Resource: "2009/pgterms/Bookshelf",
			Name:     "Children's Literature",
		}},
		Downloads: 32144,
		AuthorLinks: []pgrdf.AuthorLink{{
			URL:         "https://en.wikipedia.org/wiki/Lewis_Carroll",
			Description: "en.wikipedia",
		}},
		CCComment: "Archives containing the RDF files for *all* our books can be downloaded from our website.",
		CCLicense: "https://creativecommons.org/publicdomain/zero/1.0/",
	}
}

var rdfMarshalExpected = `<?xml version="1.0" encoding="UTF-8"?>
<rdf:RDF xml:base="http://www.gutenberg.org/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:cc="http://web.resource.org/cc/" xmlns:marcrel="http://id.loc.gov/vocabulary/relators/" xmlns:dcam="http://purl.org/dc/dcam/">
  <pgterms:ebook rdf:about="ebooks/11">
    <dcterms:title>Alice&#39;s Adventures in Wonderland</dcterms:title>
    <dcterms:alternative>Alice in Wonderland</dcterms:alternative>
    <dcterms:tableOfContents>Prefatory Note -- Chapter 1 -- Chapter 2 -- Chapter 3 -- Conclusion</dcterms:tableOfContents>
    <dcterms:publisher>Project Gutenberg</dcterms:publisher>
    <pgterms:marc906>1909</pgterms:marc906>
    <dcterms:issued rdf:datatype="http://www.w3.org/2001/XMLSchema#date">2008-06-27</dcterms:issued>
    <pgterms:marc520>A short story about short summaries.</pgterms:marc520>
    <pgterms:marc440>Best of Fantasy</pgterms:marc440>
    <dcterms:language>
      <rdf:Description rdf:nodeID="N4311cfb96de1a76167c27d1af1026fa7">
        <rdf:value rdf:datatype="http://purl.org/dc/terms/RFC4646">en</rdf:value>
      </rdf:Description>
    </dcterms:language>
    <pgterms:marc907>GB</pgterms:marc907>
    <pgterms:marc546>Uses 19th century spelling.</pgterms:marc546>
    <pgterms:marc260>United Kingdom: J. Johnson, 1794.</pgterms:marc260>
    <pgterms:marc250>2nd Edition</pgterms:marc250>
    <pgterms:marc508>D. Smith</pgterms:marc508>
    <pgterms:marc508>Updated: 2022-07-14</pgterms:marc508>
    <dcterms:license rdf:resource="license"/>
    <dcterms:rights>Public domain in the USA.</dcterms:rights>
    <pgterms:marc905>20001231235959randomthing</pgterms:marc905>
    <dcterms:type>
      <rdf:Description rdf:nodeID="N59f6317c7c4dbd8e93f3f12b2415d876">
        <rdf:value>Text</rdf:value>
        <dcam:memberOf rdf:resource="http://purl.org/dc/terms/DCMIType"/>
      </rdf:Description>
    </dcterms:type>
    <dcterms:description>An improved version is available at #28885.</dcterms:description>
    <pgterms:marc300>5 pages : illustrations, map, portraits.</pgterms:marc300>
    <pgterms:marc904>https://example.com/ebooks/11/something</pgterms:marc904>
    <pgterms:marc010>77177892</pgterms:marc010>
    <pgterms:marc020>978-0-919366-14-5</pgterms:marc020>
    <pgterms:marc901>https://www.gutenberg.org/cache/epub/11/pg11.cover.medium.jpg</pgterms:marc901>
    <pgterms:marc902>https://example.org/ebook11/title.jpg</pgterms:marc902>
    <pgterms:marc903>https://example.org/ebook11/back.jpg</pgterms:marc903>
    <dcterms:creator>
      <pgterms:agent rdf:about="2009/agents/7">
        <pgterms:name>Carroll, Lewis</pgterms:name>
        <pgterms:alias>Dodgson, Charles Lutwidge</pgterms:alias>
        <pgterms:birthdate rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">1832</pgterms:birthdate>
        <pgterms:deathdate rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">1898</pgterms:deathdate>
        <pgterms:webpage rdf:resource="https://en.wikipedia.org/wiki/Lewis_Carroll"/>
      </pgterms:agent>
    </dcterms:creator>
    <dcterms:subject>
      <rdf:Description rdf:nodeID="N4e3c9c524010316e93b7353ddc82cde1">
        <rdf:value>Fantasy fiction</rdf:value>
        <dcam:memberOf rdf:resource="http://purl.org/dc/terms/LCSH"/>
      </rdf:Description>
    </dcterms:subject>
    <dcterms:hasFormat>
      <pgterms:file rdf:about="https://www.gutenberg.org/files/11/11-0.txt">
        <dcterms:extent rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">174693</dcterms:extent>
        <dcterms:modified rdf:datatype="http://www.w3.org/2001/XMLSchema#dateTime">2020-10-12T03:45:53</dcterms:modified>
        <dcterms:isFormatOf rdf:resource="ebooks/11"/>
        <dcterms:format>
          <rdf:Description rdf:nodeID="N2d78b15714cf8a43902bd3108479c078">
            <rdf:value rdf:datatype="http://purl.org/dc/terms/IMT">text/plain; charset=utf-8</rdf:value>
            <dcam:memberOf rdf:resource="http://purl.org/dc/terms/IMT"/>
          </rdf:Description>
        </dcterms:format>
      </pgterms:file>
    </dcterms:hasFormat>
    <pgterms:bookshelf>
      <rdf:Description rdf:nodeID="N5fe1f85f2ca92d66a964562166b9b4cc">
        <rdf:value>Children&#39;s Literature</rdf:value>
        <dcam:memberOf rdf:resource="2009/pgterms/Bookshelf"/>
      </rdf:Description>
    </pgterms:bookshelf>
    <pgterms:downloads rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">32144</pgterms:downloads>
  </pgterms:ebook>
  <rdf:Description rdf:about="https://en.wikipedia.org/wiki/Lewis_Carroll">
    <dcterms:description>en.wikipedia</dcterms:description>
  </rdf:Description>
  <cc:Work>
    <rdfs:comment>Archives containing the RDF files for *all* our books can be downloaded from our website.</rdfs:comment>
    <cc:license rdf:resource="https://creativecommons.org/publicdomain/zero/1.0/"/>
  </cc:Work>
</rdf:RDF>`
