package pgrdf_test

import (
	"bytes"
	"testing"

	"github.com/mrcook/pgrdf"
)

func TestEbook_WriteRDF(t *testing.T) {
	ebook := generateEbook()

	w := bytes.NewBuffer([]byte{})
	err := ebook.WriteRDF(w)
	if err != nil {
		t.Fatalf("error marshalling ebook: %s", err)
	}

	data := w.String()
	if data != rdfMarshallExpected {
		t.Errorf("unexpected marshalled output, expected:\n'%s'\n\ngot:\n'%s'\n", rdfMarshallExpected, data)
	}
}

func generateEbook() *pgrdf.Ebook {
	return &pgrdf.Ebook{
		ID:          11,
		BookType:    "Text",
		ReleaseDate: "2008-06-27",
		Language: pgrdf.Language{
			Code:    "en",
			Dialect: "GB",
			Notes:   "Uses 19th century spelling.",
		},
		Publisher:     "Project Gutenberg",
		PublishedYear: 1909,
		Copyright:     "Public domain in the USA.",
		Titles:        []string{"Alice's Adventures in Wonderland"},
		OtherTitles:   []string{"Alice in Wonderland"},
		Creators: []pgrdf.Creator{{
			ID:      7,
			Name:    "Carroll, Lewis",
			Aliases: []string{"Dodgson, Charles Lutwidge"},
			Born:    1832,
			Died:    1898,
			WebPage: "https://en.wikipedia.org/wiki/Lewis_Carroll",
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
		Series:            "Best of Fantasy",
		BookCoverFilename: "images/cover.jpg",
		Downloads:         32144,
		Note:              "An improved version is available at #28885.",
		Comment:           "Archives containing the RDF files for *all* our books can be downloaded from our website.",
		CCLicense:         "https://creativecommons.org/publicdomain/zero/1.0/",
		AuthorLinks: []pgrdf.AuthorLink{{
			URL:         "https://en.wikipedia.org/wiki/Lewis_Carroll",
			Description: "en.wikipedia",
		}},
	}
}

var rdfMarshallExpected = `<rdf:RDF xml:base="http://www.gutenberg.org/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:cc="http://web.resource.org/cc/" xmlns:marcrel="http://id.loc.gov/vocabulary/relators/" xmlns:dcam="http://purl.org/dc/dcam/">
  <cc:Work>
    <rdfs:comment>Archives containing the RDF files for *all* our books can be downloaded from our website.</rdfs:comment>
    <cc:license rdf:resource="https://creativecommons.org/publicdomain/zero/1.0/"></cc:license>
  </cc:Work>
  <pgterms:ebook rdf:about="ebooks/11">
    <dcterms:description>An improved version is available at #28885.</dcterms:description>
    <dcterms:type>
      <rdf:Description rdf:nodeID="Nb915b0362e09cb245ffc942c959201f2">
        <rdf:value>Text</rdf:value>
        <dcam:memberOf rdf:resource="http://purl.org/dc/terms/DCMIType"></dcam:memberOf>
      </rdf:Description>
    </dcterms:type>
    <dcterms:issued rdf:datatype="http://www.w3.org/2001/XMLSchema#date">2008-06-27</dcterms:issued>
    <dcterms:language>
      <rdf:Description rdf:nodeID="N59f6317c7c4dbd8e93f3f12b2415d876">
        <rdf:value rdf:datatype="http://purl.org/dc/terms/RFC4646">en</rdf:value>
      </rdf:Description>
    </dcterms:language>
    <pgterms:marc907>GB</pgterms:marc907>
    <pgterms:marc546>Uses 19th century spelling.</pgterms:marc546>
    <dcterms:publisher>Project Gutenberg</dcterms:publisher>
    <pgterms:marc906>1909</pgterms:marc906>
    <dcterms:license rdf:resource="license"></dcterms:license>
    <dcterms:rights>Public domain in the USA.</dcterms:rights>
    <dcterms:title>Alice&#39;s Adventures in Wonderland</dcterms:title>
    <dcterms:alternative>Alice in Wonderland</dcterms:alternative>
    <dcterms:creator>
      <pgterms:agent rdf:about="2009/agents/7">
        <pgterms:name>Carroll, Lewis</pgterms:name>
        <pgterms:alias>Dodgson, Charles Lutwidge</pgterms:alias>
        <pgterms:birthdate rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">1832</pgterms:birthdate>
        <pgterms:deathdate rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">1898</pgterms:deathdate>
        <pgterms:webpage rdf:resource="https://en.wikipedia.org/wiki/Lewis_Carroll"></pgterms:webpage>
      </pgterms:agent>
    </dcterms:creator>
    <dcterms:subject>
      <rdf:Description rdf:nodeID="N4e3c9c524010316e93b7353ddc82cde1">
        <rdf:value>Fantasy fiction</rdf:value>
        <dcam:memberOf rdf:resource="http://purl.org/dc/terms/LCSH"></dcam:memberOf>
      </rdf:Description>
    </dcterms:subject>
    <dcterms:hasFormat>
      <pgterms:file rdf:about="https://www.gutenberg.org/files/11/11-0.txt">
        <dcterms:extent rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">174693</dcterms:extent>
        <dcterms:modified rdf:datatype="http://www.w3.org/2001/XMLSchema#dateTime">2020-10-12T03:45:53</dcterms:modified>
        <dcterms:isFormatOf rdf:resource="ebooks/11"></dcterms:isFormatOf>
        <dcterms:format>
          <rdf:Description rdf:nodeID="N2d78b15714cf8a43902bd3108479c078">
            <rdf:value rdf:datatype="http://purl.org/dc/terms/IMT">text/plain; charset=utf-8</rdf:value>
            <dcam:memberOf rdf:resource="http://purl.org/dc/terms/IMT"></dcam:memberOf>
          </rdf:Description>
        </dcterms:format>
      </pgterms:file>
    </dcterms:hasFormat>
    <pgterms:bookshelf>
      <rdf:Description rdf:nodeID="N5fe1f85f2ca92d66a964562166b9b4cc">
        <rdf:value>Children&#39;s Literature</rdf:value>
        <dcam:memberOf rdf:resource="2009/pgterms/Bookshelf"></dcam:memberOf>
      </rdf:Description>
    </pgterms:bookshelf>
    <pgterms:marc440>Best of Fantasy</pgterms:marc440>
    <pgterms:marc901>images/cover.jpg</pgterms:marc901>
    <pgterms:downloads rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">32144</pgterms:downloads>
  </pgterms:ebook>
  <rdf:Description rdf:about="https://en.wikipedia.org/wiki/Lewis_Carroll">
    <dcterms:description>en.wikipedia</dcterms:description>
  </rdf:Description>
</rdf:RDF>`
