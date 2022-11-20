package pgrdf_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/mrcook/pgrdf"
)

func TestEbook(t *testing.T) {
	ebook := getEbookFromSampleRdf(t)

	if ebook.ID != 999991234 {
		t.Errorf("unexpected ebook ID, got %d", ebook.ID)
	}
	if ebook.Note != "A description for this RDF" {
		t.Errorf("unexpected ebook description, got '%s'", ebook.Note)
	}
	if ebook.BookType != "Text" {
		t.Errorf("unexpected ebook book type, got '%s'", ebook.BookType)
	}
	if ebook.ReleaseDate != "1998-07-01" {
		t.Errorf("unexpected ebook book type, got '%s'", ebook.ReleaseDate)
	}
	if ebook.Publisher != "Project Gutenberg" {
		t.Errorf("unexpected ebook publisher, got '%s'", ebook.Publisher)
	}
	if ebook.PublishedYear != 1861 {
		t.Errorf("unexpected ebook published date, got '%d'", ebook.PublishedYear)
	}
	if ebook.Copyright != "Public domain in the USA." {
		t.Errorf("unexpected ebook copyright, got '%s'", ebook.Copyright)
	}
	if ebook.Series != "Dickens Best Of" {
		t.Errorf("unexpected series, got '%s'", ebook.Series)
	}
	if ebook.BookCoverFilename != "images/cover.jpg" {
		t.Errorf("unexpected book cover filename, got '%s'", ebook.BookCoverFilename)
	}
	if ebook.Downloads != 16579 {
		t.Errorf("unexpected ebook downloads, got %d", ebook.Downloads)
	}
	if ebook.Comment != "Archives containing the RDF files for *all* our books can be downloaded from our website." {
		t.Errorf("unexpected work comment, got '%s'", ebook.Comment)
	}
	if ebook.CCLicense != "https://creativecommons.org/publicdomain/zero/1.0/" {
		t.Errorf("unexpected license, got '%s'", ebook.CCLicense)
	}

	if len(ebook.AuthorLinks) != 1 {
		t.Fatalf("expected 1 wikipedia authors, got %d\n", len(ebook.AuthorLinks))
	}
	wiki := ebook.AuthorLinks[0]

	if wiki.Description != "en.wikipedia" {
		t.Errorf("unexpected Wikipedia language, got '%s'", wiki.Description)
	}
	if wiki.URL != "https://en.wikipedia.org/wiki/Charles_Dickens" {
		t.Errorf("unexpected author URL, got '%s'", wiki.URL)
	}

	if len(ebook.Titles) != 1 {
		t.Errorf("expected 1 ebook title, got %d\n", len(ebook.Titles))
	}
	if len(ebook.Creators) == 0 {
		t.Error("expected one or more ebook creators, got none")
	}
	if len(ebook.Subjects) != 9 {
		t.Errorf("expected 9 ebook subjects, got %d\n", len(ebook.Subjects))
	}
	if len(ebook.Files) != 15 {
		t.Errorf("expected 15 ebook book formats, got %d\n", len(ebook.Files))
	}
	if len(ebook.Bookshelves) != 1 {
		t.Errorf("expected 1 ebook bookshelves, got %d\n", len(ebook.Bookshelves))
	}
}

func TestEbookLanguage(t *testing.T) {
	ebooks := getEbookFromSampleRdf(t)

	if ebooks.Language.Code != "en" {
		t.Errorf("unexpected ebook language, got '%s'", ebooks.Language.Code)
	}
	if ebooks.Language.Dialect != "GB" {
		t.Errorf("unexpected ebook language dielect, got '%s'", ebooks.Language.Dialect)
	}
	if ebooks.Language.Notes != "Uses 19th century spelling." {
		t.Errorf("unexpected ebook language notes, got '%s'", ebooks.Language.Notes)
	}
}

func TestEbookAuthor(t *testing.T) {
	ebook := getEbookFromSampleRdf(t)

	if len(ebook.Creators) == 0 {
		t.Fatal("expected at least one ebook creator, got none")
	}

	t.Run("validates author data", func(t *testing.T) {
		a := ebook.Creators[0]

		if a.ID != 37 {
			t.Errorf("unexpected author ID, got %d", a.ID)
		}
		if a.Name != "Dickens, Charles" {
			t.Errorf("unexpected author name, got '%s'", a.Name)
		}
		if a.Born != 1812 {
			t.Errorf("unexpected author birthdate, got %d", a.Born)
		}
		if a.Died != 1870 {
			t.Errorf("unexpected author deathdate, got %d", a.Died)
		}
		if a.Role != pgrdf.RoleAut {
			t.Errorf("unexpected creator role, got '%s'", a.Role)
		}
		if len(a.Aliases) != 2 {
			t.Errorf("expected 2 ebook author aliases, got %d\n", len(a.Aliases))
		} else if a.Aliases[1] != "Boz" {
			t.Errorf("unexpected author name, got '%s'", a.Aliases[1])
		}
		if a.WebPage != "https://en.wikipedia.org/wiki/Charles_Dickens" {
			t.Errorf("unexpected author webpage, got '%s'", a.WebPage)
		}
	})
}

func TestEbookCreators(t *testing.T) {
	ebook := getEbookFromSampleRdf(t)

	cases := []struct {
		id   int
		role pgrdf.MarcRelatorCode
		name string
	}{
		{id: 37, role: pgrdf.RoleAut, name: "Dickens, Charles"},
		{id: 6198, role: pgrdf.RoleCtb, name: "Robert, Clémence"},
		{id: 8397, role: pgrdf.RoleEdt, name: "Snell, F. J. (Frederick John)"},
		{id: 54317, role: pgrdf.RoleCom, name: "Paz, M."},
		{id: 9473, role: pgrdf.RoleIll, name: "Leech, John"},
		{id: 1736, role: pgrdf.RoleTrl, name: "Wyllie, David"},
	}

	for _, data := range cases {
		t.Run(fmt.Sprintf("validate ID %d is present", data.id), func(t *testing.T) {
			var creator *pgrdf.Creator
			for i, _ := range ebook.Creators {
				if ebook.Creators[i].ID == data.id {
					creator = &ebook.Creators[i]
					break
				}
			}
			if creator == nil {
				t.Errorf("expected to find creator ID '%d', none found", data.id)
			} else {
				if creator.ID != data.id {
					t.Errorf("expected creator ID %d, got %d", data.id, creator.ID)
				}
				if creator.Role != data.role {
					t.Errorf("expected creator role '%s', got '%s'", data.role, creator.Role)
				}
				if creator.Name != data.name {
					t.Errorf("unexpected creator name '%s', got '%s'", data.name, creator.Name)
				}
			}
		})
	}
}

func TestEbookSubjects(t *testing.T) {
	ebook := getEbookFromSampleRdf(t)

	if len(ebook.Subjects) != 9 {
		t.Fatalf("expected 9 ebook subjects, got %d\n", len(ebook.Subjects))
	}
	s := ebook.Subjects[7]

	if s.Heading != "Revenge -- Fiction" {
		t.Errorf("unexpected subject heading, got '%s'", s.Heading)
	}
	if s.Schema != "http://purl.org/dc/terms/LCSH" {
		t.Errorf("unexpected subject schema, got '%s'", s.Schema)
	}
}

func TestEbookFiles(t *testing.T) {
	ebook := getEbookFromSampleRdf(t)

	if len(ebook.Files) != 15 {
		t.Fatalf("expected 15 ebook files, got %d\n", len(ebook.Files))
	}
	f := ebook.Files[4]

	if f.Extent != 393579 {
		t.Errorf("unexpected file extent, got %d", f.Extent)
	}
	if f.Modified != "2015-11-06T09:50:04" {
		t.Errorf("unexpected file modified timestamp, got '%s'", f.Modified)
	}
	if f.URL != "https://www.example.org/files/999991234/999991234-8.zip" {
		t.Errorf("unexpected file URI, got '%s'", f.URL)
	}
	if len(f.Encodings) != 2 {
		t.Fatalf("expected 12 ebook hasFormat, got %d\n", len(f.Encodings))
	} else if f.Encodings[1] != "text/plain; charset=iso-8859-1" {
		t.Errorf("unexpected file URI, got '%s'", f.Encodings[1])
	}
}

func TestEbookBookshelves(t *testing.T) {
	ebook := getEbookFromSampleRdf(t)

	if len(ebook.Bookshelves) != 1 {
		t.Fatalf("expected 1 ebook bookshelves, got %d\n", len(ebook.Bookshelves))
	}
	s := ebook.Bookshelves[0]

	if s.Resource != "2009/pgterms/Bookshelf" {
		t.Errorf("unexpected bookshelf name, got '%s'", s.Resource)
	}
	if s.Name != "Best Books Ever Listings" {
		t.Errorf("unexpected bookshelf subject, got '%s'", s.Name)
	}
}

func getEbookFromSampleRdf(t *testing.T) *pgrdf.Ebook {
	t.Helper()

	file, err := os.Open("samples/cache/epub/999991234/pg999991234.rdf")
	if err != nil {
		t.Fatalf("error opening test RDF file: %s", err)
	}

	ebook, err := pgrdf.NewEbook(file)
	if err != nil {
		t.Fatalf("error processing test RDF file: %s", err)
	}
	return ebook
}

func TestToRDF(t *testing.T) {
	ebook := pgrdf.Ebook{
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

	w := bytes.NewBuffer([]byte{})
	err := ebook.ToRDF(w)
	if err != nil {
		t.Fatalf("error marshalling ebook: %s", err)
	}

	data := w.String()
	if data != rdfMarshallExpected {
		t.Errorf("unexpected marshalled output, expected:\n'%s'\n\ngot:\n'%s'\n", rdfMarshallExpected, data)
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
