package pgrdf_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/mrcook/pgrdf"
)

func TestEbook(t *testing.T) {
	file, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

	e, err := pgrdf.NewEbook(file)
	if err != nil {
		t.Fatalf("error processing sample file: %s", err)
	}

	if e.ID != 1400 {
		t.Errorf("unexpected ebook ID, got %d", e.ID)
	}
	if e.Note != "Message for the rdf_test.go" {
		t.Errorf("unexpected ebook description, got '%s'", e.Note)
	}
	if e.BookType != "Text" {
		t.Errorf("unexpected ebook book type, got '%s'", e.BookType)
	}
	if e.ReleaseDate != "1998-07-01" {
		t.Errorf("unexpected ebook book type, got '%s'", e.ReleaseDate)
	}
	if e.Language != "en-GB" {
		t.Errorf("unexpected ebook language, got '%s'", e.Language)
	}
	if e.Publisher != "Project Gutenberg" {
		t.Errorf("unexpected ebook publisher, got '%s'", e.Publisher)
	}
	if e.PublishedYear != 1861 {
		t.Errorf("unexpected ebook published date, got '%d'", e.PublishedYear)
	}
	if e.Copyright != "Public domain in the USA." {
		t.Errorf("unexpected ebook copyright, got '%s'", e.Copyright)
	}
	if e.Downloads != 16579 {
		t.Errorf("unexpected ebook downloads, got %d", e.Downloads)
	}
	if e.Comment != "Archives containing the RDF files for *all* our books can be downloaded at\n            https://www.gutenberg.org/wiki/Gutenberg:Feeds#The_Complete_Project_Gutenberg_Catalog" {
		t.Errorf("unexpected work comment, got '%s'", e.Comment)
	}
	if e.CCLicense != "https://creativecommons.org/publicdomain/zero/1.0/" {
		t.Errorf("unexpected license, got '%s'", e.CCLicense)
	}

	if len(e.AuthorLinks) != 1 {
		t.Fatalf("expected 1 wikipedia authors, got %d\n", len(e.AuthorLinks))
	}
	wiki := e.AuthorLinks[0]

	if wiki.Description != "en.wikipedia" {
		t.Errorf("unexpected Wikipedia language, got '%s'", wiki.Description)
	}
	if wiki.URL != "https://en.wikipedia.org/wiki/Charles_Dickens" {
		t.Errorf("unexpected author URL, got '%s'", wiki.URL)
	}

	if len(e.Titles) != 1 {
		t.Errorf("expected 1 ebook title, got %d\n", len(e.Titles))
	}
	if len(e.Creators) != 4 {
		t.Errorf("expected 4 ebook creators, got %d\n", len(e.Creators))
	}
	if len(e.Subjects) != 9 {
		t.Errorf("expected 9 ebook subjects, got %d\n", len(e.Subjects))
	}
	if len(e.Files) != 15 {
		t.Errorf("expected 15 ebook book formats, got %d\n", len(e.Files))
	}
	if len(e.Bookshelves) != 1 {
		t.Errorf("expected 1 ebook bookshelves, got %d\n", len(e.Bookshelves))
	}
}

func TestEbookCreators(t *testing.T) {
	file, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

	e, err := pgrdf.NewEbook(file)
	if err != nil {
		t.Fatalf("error processing sample file: %s", err)
	}

	if len(e.Creators) != 4 {
		t.Fatalf("expected 4 ebook creators, got %d\n", len(e.Creators))
	}

	t.Run("first creator should be the author", func(t *testing.T) {
		a := e.Creators[0]

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

	t.Run("second creator should be the editor", func(t *testing.T) {
		a := e.Creators[1]

		if a.ID != 8397 {
			t.Errorf("unexpected editor ID, got %d", a.ID)
		}
		if a.Name != "Snell, F. J. (Frederick John)" {
			t.Errorf("unexpected editor name, got '%s'", a.Name)
		}
		if a.Role != pgrdf.RoleEdt {
			t.Errorf("unexpected editor role, got '%s'", a.Role)
		}
	})

	t.Run("third creator should be the illustrator", func(t *testing.T) {
		a := e.Creators[2]

		if a.ID != 9473 {
			t.Errorf("unexpected editor ID, got %d", a.ID)
		}
		if a.Name != "Leech, John" {
			t.Errorf("unexpected illustrator name, got '%s'", a.Name)
		}
		if a.Role != pgrdf.RoleIll {
			t.Errorf("unexpected illustrator role, got '%s'", a.Role)
		}
	})

	t.Run("fourth creator should be the translator", func(t *testing.T) {
		a := e.Creators[3]

		if a.ID != 1736 {
			t.Errorf("unexpected editor ID, got %d", a.ID)
		}
		if a.Name != "Wyllie, David" {
			t.Errorf("unexpected translator name, got '%s'", a.Name)
		}
		if a.Role != pgrdf.RoleTrl {
			t.Errorf("unexpected translator role, got '%s'", a.Role)
		}
	})
}

func TestEbookSubjects(t *testing.T) {
	file, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

	e, err := pgrdf.NewEbook(file)
	if err != nil {
		t.Fatalf("error processing sample file: %s", err)
	}

	if len(e.Subjects) != 9 {
		t.Fatalf("expected 9 ebook subjects, got %d\n", len(e.Subjects))
	}
	s := e.Subjects[7]

	if s.Heading != "Revenge -- Fiction" {
		t.Errorf("unexpected subject heading, got '%s'", s.Heading)
	}
	if s.Schema != "http://purl.org/dc/terms/LCSH" {
		t.Errorf("unexpected subject schema, got '%s'", s.Schema)
	}
}

func TestEbookFiles(t *testing.T) {
	file, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

	e, err := pgrdf.NewEbook(file)
	if err != nil {
		t.Fatalf("error processing sample file: %s", err)
	}

	if len(e.Files) != 15 {
		t.Fatalf("expected 15 ebook files, got %d\n", len(e.Files))
	}
	f := e.Files[4]

	if f.Extent != 393579 {
		t.Errorf("unexpected file extent, got %d", f.Extent)
	}
	if f.Modified != "2015-11-06T09:50:04" {
		t.Errorf("unexpected file modified timestamp, got '%s'", f.Modified)
	}
	if f.URL != "https://www.gutenberg.org/files/1400/1400-8.zip" {
		t.Errorf("unexpected file URI, got '%s'", f.URL)
	}
	if len(f.Encodings) != 2 {
		t.Fatalf("expected 12 ebook hasFormat, got %d\n", len(f.Encodings))
	} else if f.Encodings[1] != "text/plain; charset=iso-8859-1" {
		t.Errorf("unexpected file URI, got '%s'", f.Encodings[1])
	}
}

func TestEbookBookshelves(t *testing.T) {
	file, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

	e, err := pgrdf.NewEbook(file)
	if err != nil {
		t.Fatalf("error processing sample file: %s", err)
	}

	if len(e.Bookshelves) != 1 {
		t.Fatalf("expected 1 ebook bookshelves, got %d\n", len(e.Bookshelves))
	}
	s := e.Bookshelves[0]

	if s.Resource != "2009/pgterms/Bookshelf" {
		t.Errorf("unexpected bookshelf name, got '%s'", s.Resource)
	}
	if s.Name != "Best Books Ever Listings" {
		t.Errorf("unexpected bookshelf subject, got '%s'", s.Name)
	}
}

func TestToRDF(t *testing.T) {
	ebook := pgrdf.Ebook{
		ID:          11,
		BookType:    "Text",
		ReleaseDate: "2008-06-27",
		Language:    "en",
		Publisher:   "Project Gutenberg",
		Copyright:   "Public domain in the USA.",
		Titles:      []string{"Alice's Adventures in Wonderland"},
		OtherTitles: []string{"Alice in Wonderland"},
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
		Downloads: 32144,
		Note:      "An improved version is available at #28885.",
		Comment:   "Archives containing the RDF files for *all* our books can be downloaded at\n            https://www.gutenberg.org/wiki/Gutenberg:Feeds#The_Complete_Project_Gutenberg_Catalog",
		CCLicense: "https://creativecommons.org/publicdomain/zero/1.0/",
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
	if data != `<rdf:RDF xml:base="http://www.gutenberg.org/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:cc="http://web.resource.org/cc/" xmlns:dcam="http://purl.org/dc/dcam/"><cc:Work><rdfs:comment>Archives containing the RDF files for *all* our books can be downloaded at&#xA;            https://www.gutenberg.org/wiki/Gutenberg:Feeds#The_Complete_Project_Gutenberg_Catalog</rdfs:comment><cc:license rdf:resource="https://creativecommons.org/publicdomain/zero/1.0/"></cc:license></cc:Work><pgterms:ebook rdf:about="ebooks/11"><dcterms:description>An improved version is available at #28885.</dcterms:description><dcterms:type><rdf:Description rdf:nodeID="Nb915b0362e09cb245ffc942c959201f2"><rdf:value>Text</rdf:value><dcam:memberOf rdf:resource="http://purl.org/dc/terms/DCMIType"></dcam:memberOf></rdf:Description></dcterms:type><dcterms:issued rdf:datatype="http://www.w3.org/2001/XMLSchema#date">2008-06-27</dcterms:issued><dcterms:language><rdf:Description rdf:nodeID="N59f6317c7c4dbd8e93f3f12b2415d876"><rdf:value rdf:datatype="http://purl.org/dc/terms/RFC4646">en</rdf:value></rdf:Description></dcterms:language><dcterms:publisher>Project Gutenberg</dcterms:publisher><dcterms:license rdf:resource="license"></dcterms:license><dcterms:rights>Public domain in the USA.</dcterms:rights><dcterms:title>Alice&#39;s Adventures in Wonderland</dcterms:title><dcterms:alternative>Alice in Wonderland</dcterms:alternative><dcterms:creator><pgterms:agent rdf:about="2009/agents/7"><pgterms:name>Carroll, Lewis</pgterms:name><pgterms:alias>Dodgson, Charles Lutwidge</pgterms:alias><pgterms:birthdate rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">1832</pgterms:birthdate><pgterms:deathdate rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">1898</pgterms:deathdate><pgterms:webpage rdf:resource="https://en.wikipedia.org/wiki/Lewis_Carroll"></pgterms:webpage></pgterms:agent></dcterms:creator><dcterms:subject><rdf:Description rdf:nodeID="N4e3c9c524010316e93b7353ddc82cde1"><rdf:value>Fantasy fiction</rdf:value><dcam:memberOf rdf:resource="http://purl.org/dc/terms/LCSH"></dcam:memberOf></rdf:Description></dcterms:subject><dcterms:hasFormat><pgterms:file rdf:about="https://www.gutenberg.org/files/11/11-0.txt"><dcterms:extent rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">174693</dcterms:extent><dcterms:modified rdf:datatype="http://www.w3.org/2001/XMLSchema#dateTime">2020-10-12T03:45:53</dcterms:modified><dcterms:isFormatOf rdf:resource="ebooks/11"></dcterms:isFormatOf><dcterms:format><rdf:Description rdf:nodeID="N2d78b15714cf8a43902bd3108479c078"><rdf:value rdf:datatype="http://purl.org/dc/terms/IMT">text/plain; charset=utf-8</rdf:value><dcam:memberOf rdf:resource="http://purl.org/dc/terms/IMT"></dcam:memberOf></rdf:Description></dcterms:format></pgterms:file></dcterms:hasFormat><pgterms:bookshelf><rdf:Description rdf:nodeID="N5fe1f85f2ca92d66a964562166b9b4cc"><rdf:value>Children&#39;s Literature</rdf:value><dcam:memberOf rdf:resource="2009/pgterms/Bookshelf"></dcam:memberOf></rdf:Description></pgterms:bookshelf><pgterms:downloads rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">32144</pgterms:downloads></pgterms:ebook><rdf:Description rdf:about="https://en.wikipedia.org/wiki/Lewis_Carroll"><dcterms:description>en.wikipedia</dcterms:description></rdf:Description></rdf:RDF>` {
		t.Errorf("unexpected marshalled output, got: %s", data)
	}
}

func openRDF() (*os.File, error) {
	file, err := os.Open("samples/cache/epub/1400/pg1400.rdf")
	if err != nil {
		return nil, err
	}
	return file, nil
}
