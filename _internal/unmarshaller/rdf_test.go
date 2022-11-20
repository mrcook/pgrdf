package unmarshaller_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/mrcook/pgrdf/_internal/unmarshaller"
)

func TestNamespaces(t *testing.T) {
	r := openRDF(t)

	if r.NsBase != "http://www.gutenberg.org/" {
		t.Errorf("unexpected xml:base, got '%s'", r.NsBase)
	}
	if r.NsPgTerms != "http://www.gutenberg.org/2009/pgterms/" {
		t.Errorf("unexpected xmlns:pgterms, got '%s'", r.NsPgTerms)
	}
	if r.NsCC != "http://web.resource.org/cc/" {
		t.Errorf("unexpected xmlns:cc, got '%s'", r.NsCC)
	}
	if r.NsRdf != "http://www.w3.org/1999/02/22-rdf-syntax-ns#" {
		t.Errorf("unexpected xmlns:rdf, got '%s'", r.NsRdf)
	}
	if r.NsDcam != "http://purl.org/dc/dcam/" {
		t.Errorf("unexpected xmlns:dcam, got '%s'", r.NsDcam)
	}
	if r.NsRdfs != "http://www.w3.org/2000/01/rdf-schema#" {
		t.Errorf("unexpected xmlns:rdfs, got '%s'", r.NsRdfs)
	}
	if r.NsDcTerms != "http://purl.org/dc/terms/" {
		t.Errorf("unexpected xmlns:dcterms, got '%s'", r.NsDcTerms)
	}
	if r.NsMarcRel != "http://id.loc.gov/vocabulary/relators/" {
		t.Errorf("unexpected xmlns:marcrel, got '%s'", r.NsMarcRel)
	}
}

func TestWorkNode(t *testing.T) {
	r := openRDF(t)

	if r.Work.About != "" {
		t.Errorf("unexpected rdf:about, got '%s'", r.Work.About)
	}
	if r.Work.Comment != "Archives containing the RDF files for *all* our books can be downloaded at\n            https://www.gutenberg.org/wiki/Gutenberg:Feeds#The_Complete_Project_Gutenberg_Catalog" {
		t.Errorf("unexpected work comment, got '%s'", r.Work.Comment)
	}
	if r.Work.License.Resource != "https://creativecommons.org/publicdomain/zero/1.0/" {
		t.Errorf("unexpected license, got '%s'", r.Work.License.Resource)
	}
}

func TestDescriptionNodes(t *testing.T) {
	r := openRDF(t)

	if len(r.Descriptions) != 1 {
		t.Fatalf("expected 1 description, got %d", len(r.Descriptions))
	}
	d := r.Descriptions[0]

	if d.About != "https://en.wikipedia.org/wiki/Charles_Dickens" {
		t.Errorf("unexpected rdf:about, got '%s'", d.About)
	}
	if d.Description != "en.wikipedia" {
		t.Errorf("unexpected dcterms:description, got '%s'", d.Description)
	}
}

func TestEbook(t *testing.T) {
	r := openRDF(t)

	e := r.Ebook

	if e.About != "ebooks/1400" {
		t.Errorf("unexpected rdf:about, got '%s'", e.About)
	}
	if e.Description != "Message for the rdf_test.go" {
		t.Errorf("unexpected dcterms:description, got '%s'", e.Description)
	}
	if e.Type.Description.NodeID != "Nebb73a3dacde414382cc3a31ce400f17" {
		t.Errorf("unexpected dcterms:type//rdf:nodeID, got '%s'", e.Type.Description.NodeID)
	}
	if e.Type.Description.MemberOf.Resource != "http://purl.org/dc/terms/DCMIType" {
		t.Errorf("unexpected dcterms:type//dcam:memberOf, got '%s'", e.Type.Description.MemberOf.Resource)
	}
	if e.Type.Description.Value.Data != "Text" {
		t.Errorf("unexpected dcterms:type//rdf:value, got '%s'", e.Type.Description.Value.Data)
	}
	if e.Issued.DataType != "http://www.w3.org/2001/XMLSchema#date" {
		t.Errorf("unexpected dcterms:issued rdf:datatype, got '%s'", e.Issued.DataType)
	}
	if e.Issued.Value != "1998-07-01" {
		t.Errorf("unexpected dcterms:issued, got '%s'", e.Issued.Value)
	}
	if e.Language.Description.NodeID != "N73e956e8e5d049ac943dfe482ddd5802" {
		t.Errorf("unexpected dcterms:language//rdf:nodeID, got '%s'", e.Language.Description.NodeID)
	}
	if e.Language.Description.Value.DataType != "http://purl.org/dc/terms/RFC4646" {
		t.Errorf("unexpected dcterms:language//rdf:value.datatype, got '%s'", e.Language.Description.Value.DataType)
	}
	if e.Language.Description.Value.Data != "en" {
		t.Errorf("unexpected dcterms:language//rdf:value, got '%s'", e.Language.Description.Value.Data)
	}
	if e.LanguageSubCode != "GB" {
		t.Errorf("unexpected marc907 (language sub-code), got '%s'", e.LanguageSubCode)
	}
	if e.Publisher != "Project Gutenberg" {
		t.Errorf("unexpected dcterms:publisher, got '%s'", e.Publisher)
	}
	if e.PublishedYear != 1861 {
		t.Errorf("unexpected marc906 (published year), got '%d'", e.PublishedYear)
	}
	if e.License.Resource != "license" {
		t.Errorf("unexpected dcterms:license, got '%s'", e.License.Resource)
	}
	if e.Rights != "Public domain in the USA." {
		t.Errorf("unexpected dcterms:rights, got '%s'", e.Rights)
	}
	if e.Title != "Great Expectations" {
		t.Errorf("unexpected dcterms:title, got '%s'", e.Title)
	}
	if len(e.Alternative) != 1 {
		t.Errorf("expected 1 dcterms:title, got %d", len(e.Alternative))
	} else if e.Alternative[0] != "Alternative title for the rdf_test.go" {
		t.Errorf("unexpected dcterms:alternative, got '%s'", e.Alternative[0])
	}
	if e.Series != "Dickens Best Of" {
		t.Errorf("unexpected pgterms:marc440 (series), got '%s'", e.Series)
	}
	if e.BookCover != "file:///public/vhost/g/gutenberg/html/files/1400/1400-h/images/cover.jpg" {
		t.Errorf("unexpected pgterms:marc901 bookcover tag, got '%s'", e.BookCover)
	}
	if e.Downloads.DataType != "http://www.w3.org/2001/XMLSchema#integer" {
		t.Errorf("unexpected pgterms:downloads rdf:datatype, got '%s'", e.Downloads.DataType)
	}
	if e.Downloads.Value != 16579 {
		t.Errorf("unexpected pgterms:downloads count, got %d", e.Downloads.Value)
	}

	// Basic checking only. The contents of each slice are tested separately.
	if len(e.Creators) != 1 {
		t.Errorf("expected 1 dcterms:creator, got %d", len(e.Creators))
	}
	if len(e.Compilers) != 1 {
		t.Errorf("expected 1 marcrel:com, got %d", len(e.Compilers))
	}
	if len(e.Contributors) != 1 {
		t.Errorf("expected 1 marcrel:ctb, got %d", len(e.Contributors))
	}
	if len(e.Editors) != 1 {
		t.Errorf("expected 1 marcrel:edt, got %d", len(e.Editors))
	}
	if len(e.Illustrators) != 1 {
		t.Errorf("expected 1 marcrel:ill, got %d", len(e.Illustrators))
	}
	if len(e.Translators) != 1 {
		t.Errorf("expected 1 marcrel:trl, got %d", len(e.Translators))
	}
	if len(e.Subjects) != 9 {
		t.Errorf("expected 9 dcterms:subject, got %d", len(e.Subjects))
	}
	if len(e.HasFormats) != 15 {
		t.Errorf("expected 15 dcterms:hasFormat, got %d", len(e.HasFormats))
	}
	if len(e.Bookshelves) != 1 {
		t.Errorf("expected 1 pgterms:bookshelf, got %d", len(e.Bookshelves))
	}
}

func TestCreators(t *testing.T) {
	r := openRDF(t)

	if len(r.Ebook.Creators) != 1 {
		t.Fatalf("expected 1 dcterms:creator, got %d", len(r.Ebook.Creators))
	}
	a := r.Ebook.Creators[0].Agent

	if a.About != "2009/agents/37" {
		t.Errorf("unexpected dcterms:creator/agent.about, got '%s'", a.About)
	}
	if a.Name != "Dickens, Charles" {
		t.Errorf("unexpected dcterms:creator/agent/name, got '%s'", a.Name)
	}
	if len(a.Aliases) != 2 {
		t.Fatalf("expected 2 dcterms:creator/agent/alias, got %d", len(a.Aliases))
	}
	if a.Aliases[1] != "Boz" {
		t.Errorf("unexpected dcterms:creator/agent/, got '%s'", a.Aliases[1])
	}
	if a.Birthdate.DataType != "http://www.w3.org/2001/XMLSchema#integer" {
		t.Errorf("unexpected dcterms:creator/agent/birthdate.datatype, got '%s'", a.Birthdate.DataType)
	}
	if a.Birthdate.Value != 1812 {
		t.Errorf("unexpected dcterms:creator/agent/birthdate, got %d", a.Birthdate.Value)
	}
	if a.Deathdate.DataType != "http://www.w3.org/2001/XMLSchema#integer" {
		t.Errorf("unexpected dcterms:creator/agent/deathdate.datatype, got '%s'", a.Deathdate.DataType)
	}
	if a.Deathdate.Value != 1870 {
		t.Errorf("unexpected dcterms:creator/agent/deathdate, got %d", a.Deathdate.Value)
	}
	if a.Webpage.Resource != "https://en.wikipedia.org/wiki/Charles_Dickens" {
		t.Errorf("unexpected dcterms:creator/agent/webpage, got '%s'", a.Webpage.Resource)
	}
}

func TestCompilers(t *testing.T) {
	r := openRDF(t)

	if len(r.Ebook.Compilers) != 1 {
		t.Fatalf("expected 1 marcrel:trl, got %d", len(r.Ebook.Compilers))
	}
	a := r.Ebook.Compilers[0].Agent

	if a.About != "2009/agents/54317" {
		t.Errorf("unexpected compiler dcterms:creator/agent.about, got '%s'", a.About)
	}
	if a.Name != "Paz, M." {
		t.Errorf("unexpected compiler dcterms:creator/agent/name, got '%s'", a.Name)
	}
}

func TestContributors(t *testing.T) {
	r := openRDF(t)

	if len(r.Ebook.Contributors) != 1 {
		t.Fatalf("expected 1 marcrel:ctb, got %d", len(r.Ebook.Contributors))
	}
	a := r.Ebook.Contributors[0].Agent

	if a.About != "2009/agents/6198" {
		t.Errorf("unexpected contributor dcterms:creator/agent.about, got '%s'", a.About)
	}
	if a.Name != "Robert, Clémence" {
		t.Errorf("unexpected contributor dcterms:creator/agent/name, got '%s'", a.Name)
	}
}

func TestEditors(t *testing.T) {
	r := openRDF(t)

	if len(r.Ebook.Editors) != 1 {
		t.Fatalf("expected 1 marcrel:edt, got %d", len(r.Ebook.Creators))
	}
	a := r.Ebook.Editors[0].Agent

	if a.About != "2009/agents/8397" {
		t.Errorf("unexpected editor dcterms:creator/agent.about, got '%s'", a.About)
	}
	if a.Name != "Snell, F. J. (Frederick John)" {
		t.Errorf("unexpected editor dcterms:creator/agent/name, got '%s'", a.Name)
	}
}

func TestIllustrators(t *testing.T) {
	r := openRDF(t)

	if len(r.Ebook.Illustrators) != 1 {
		t.Fatalf("expected 1 marcrel:ill, got %d", len(r.Ebook.Illustrators))
	}
	a := r.Ebook.Illustrators[0].Agent

	if a.About != "2009/agents/9473" {
		t.Errorf("unexpected illustrator dcterms:creator/agent.about, got '%s'", a.About)
	}
	if a.Name != "Leech, John" {
		t.Errorf("unexpected illustrator dcterms:creator/agent/name, got '%s'", a.Name)
	}
}

func TestTranslators(t *testing.T) {
	r := openRDF(t)

	if len(r.Ebook.Translators) != 1 {
		t.Fatalf("expected 1 marcrel:trl, got %d", len(r.Ebook.Translators))
	}
	a := r.Ebook.Translators[0].Agent

	if a.About != "2009/agents/1736" {
		t.Errorf("unexpected translator dcterms:creator/agent.about, got '%s'", a.About)
	}
	if a.Name != "Wyllie, David" {
		t.Errorf("unexpected translator dcterms:creator/agent/name, got '%s'", a.Name)
	}
}

func TestSubjects(t *testing.T) {
	r := openRDF(t)

	if len(r.Ebook.Subjects) != 9 {
		t.Fatalf("expected 9 dcterms:subject, got %d", len(r.Ebook.Subjects))
	}
	d := r.Ebook.Subjects[5].Description

	if d.NodeID != "Nb6ba2be5822749bd8470f99ddf722bb3" {
		t.Errorf("unexpected dcterms:subject rdf:nodeID, got '%s'", d.NodeID)
	}
	if d.Value.Data != "Orphans -- Fiction" {
		t.Errorf("unexpected dcterms:subject, got '%s'", d.Value.Data)
	}
	if d.MemberOf.Resource != "http://purl.org/dc/terms/LCSH" {
		t.Errorf("unexpected dcterms:subject, got '%s'", d.MemberOf.Resource)
	}
}

func TestHasFormats(t *testing.T) {
	r := openRDF(t)

	if len(r.Ebook.HasFormats) != 15 {
		t.Fatalf("expected 15 dcterms:hasFormat, got %d", len(r.Ebook.HasFormats))
	}
	f := r.Ebook.HasFormats[0].File // using #0 because it has 2 formats

	if f.About != "https://www.gutenberg.org/files/1400/1400-h.zip" {
		t.Errorf("unexpected dcterms:hasFormat/file rdf:about, got '%s'", f.About)
	}
	if f.Extent.DataType != "http://www.w3.org/2001/XMLSchema#integer" {
		t.Errorf("unexpected dcterms:hasFormat/file/extent.datatype, got '%s'", f.Extent.DataType)
	}
	if f.Extent.Value != 31020490 {
		t.Errorf("unexpected dcterms:hasFormat/file/extent, got %d", f.Extent.Value)
	}
	if f.Modified.DataType != "http://www.w3.org/2001/XMLSchema#dateTime" {
		t.Errorf("unexpected dcterms:hasFormat/file/modified.datatype, got '%s'", f.Modified.DataType)
	}
	if f.Modified.Value != "2020-04-27T16:52:30" {
		t.Errorf("unexpected dcterms:hasFormat/file/modified, got '%s'", f.Modified.Value)
	}
	if f.IsFormatOf.Resource != "ebooks/1400" {
		t.Errorf("unexpected dcterms:hasFormat/file, got '%s'", f.IsFormatOf.Resource)
	}

	if len(f.Formats) != 2 {
		t.Fatalf("expected 12 dcterms:hasFormat/file/format, got %d", len(f.Formats))
	}
	d := f.Formats[0].Description

	if d.NodeID != "Nff24e040b39d4164ae80fcc149363ff5" {
		t.Errorf("unexpected dcterms:hasFormat//format.nodeID, got '%s'", d.NodeID)
	}
	if d.MemberOf.Resource != "http://purl.org/dc/terms/IMT" {
		t.Errorf("unexpected dcterms:hasFormat//format/memberOf, got '%s'", d.MemberOf.Resource)
	}
	if d.Value.DataType != "http://purl.org/dc/terms/IMT" {
		t.Errorf("unexpected dcterms:hasFormat//format/value.datatype, got '%s'", d.Value.DataType)
	}
	if d.Value.Data != "application/zip" {
		t.Errorf("unexpected dcterms:hasFormat//format/value, got '%s'", d.Value.Data)
	}
}

func TestBookshelves(t *testing.T) {
	r := openRDF(t)

	if len(r.Ebook.Bookshelves) != 1 {
		t.Fatalf("expected 1 pgterms:bookshelf, got %d", len(r.Ebook.Bookshelves))
	}
	s := r.Ebook.Bookshelves[0].Description

	if s.NodeID != "N546300c3a4394e77b56ded5d234ca5fd" {
		t.Errorf("unexpected pgterms:bookshelf rdf:nodeID, got '%s'", s.NodeID)
	}
	if s.Value.Data != "Best Books Ever Listings" {
		t.Errorf("unexpected pgterms:bookshelf/value, got '%s'", s.Value.Data)
	}
	if s.MemberOf.Resource != "2009/pgterms/Bookshelf" {
		t.Errorf("unexpected pgterms:bookshelf/memberOf, got '%s'", s.MemberOf.Resource)
	}
}

func TestMarcCodes(t *testing.T) {
	t.Run("marc010 LoC Number", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc010>77177891</pgterms:marc010></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.LOC != "77177891" {
			t.Fatalf("unexpected LoC '%s'", rdf.Ebook.LOC)
		}
	})

	t.Run("marc020 ISBN", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc020>0-397-00033-2</pgterms:marc020></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.ISBN != "0-397-00033-2" {
			t.Fatalf("unexpected ISBN '%s'", rdf.Ebook.ISBN)
		}
	})

	t.Run("marc250 Edition", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc250>The Charles Dickens Edition</pgterms:marc250></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.Edition != "The Charles Dickens Edition" {
			t.Fatalf("unexpected edition '%s'", rdf.Ebook.Edition)
		}
	})

	t.Run("marc260 Original Publication Details", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc260>United Kingdom: J. Johnson, 1794.</pgterms:marc260></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.OriginalPublication != "United Kingdom: J. Johnson, 1794." {
			t.Fatalf("unexpected original publication details '%s'", rdf.Ebook.OriginalPublication)
		}
	})

	t.Run("marc300 Source Description", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc300>Musical score</pgterms:marc300></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.SourceDescription != "Musical score" {
			t.Fatalf("unexpected summary '%s'", rdf.Ebook.SourceDescription)
		}
	})

	t.Run("marc440 Series Title", func(t *testing.T) {
		example := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc440>Dickens Best Of</pgterms:marc440></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(example))
		if rdf.Ebook.Series != "Dickens Best Of" {
			t.Fatalf("unexpected series '%s'", rdf.Ebook.Series)
		}
	})

	t.Run("marc508 Credits", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc508>Updated: 2022-07-14</pgterms:marc508><pgterms:marc508>Produced by Anon.</pgterms:marc508></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if len(rdf.Ebook.Credits) != 2 {
			t.Fatalf("expected 2 credit entries, got %d", len(rdf.Ebook.Credits))
		}
		if rdf.Ebook.Credits[0] != "Updated: 2022-07-14" {
			t.Errorf("nnexpected credit '%s'", rdf.Ebook.Credits[0])
		}
		if rdf.Ebook.Credits[1] != "Produced by Anon." {
			t.Errorf("unexpected credit '%s'", rdf.Ebook.Credits[1])
		}
	})

	t.Run("marc520 Summary", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc520>A fun version of Night Before Christmas.</pgterms:marc520></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.Summary != "A fun version of Night Before Christmas." {
			t.Fatalf("unexpected summary '%s'", rdf.Ebook.Summary)
		}
	})

	t.Run("marc546 Language Note", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc546>Uses 19th century spelling</pgterms:marc546></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.LanguageNote != "Uses 19th century spelling" {
			t.Fatalf("unexpected book cover link '%s'", rdf.Ebook.LanguageNote)
		}
	})

	t.Run("marc901 Bookcover link", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc901>file:///images/cover.jpg</pgterms:marc901></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.BookCover != "file:///images/cover.jpg" {
			t.Fatalf("unexpected book cover link '%s'", rdf.Ebook.BookCover)
		}
	})

	t.Run("marc902 Title page image link", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc902>https://example.org/ebook1/title.jpg</pgterms:marc902></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.TitlePageImage != "https://example.org/ebook1/title.jpg" {
			t.Fatalf("unexpected title page link '%s'", rdf.Ebook.TitlePageImage)
		}
	})

	t.Run("marc903 Back cover link", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc903>https://example.org/ebook1/back.jpg</pgterms:marc903></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.BackCover != "https://example.org/ebook1/back.jpg" {
			t.Fatalf("unexpected back cover link '%s'", rdf.Ebook.BackCover)
		}
	})

	t.Run("marc904 Source Link", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc904>https://example.org/ebook1</pgterms:marc904></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.SourceLink != "https://example.org/ebook1" {
			t.Fatalf("unexpected source link '%s'", rdf.Ebook.SourceLink)
		}
	})

	t.Run("marc905 PGDP Copyright Clearance Code", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc905>20051128110452chambers</pgterms:marc905></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.PgDpClearance != "20051128110452chambers" {
			t.Fatalf("unexpected clearance code value '%s'", rdf.Ebook.PgDpClearance)
		}
	})

	t.Run("marc906 Published Year", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc906>1911</pgterms:marc906></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.PublishedYear != 1911 {
			t.Fatalf("expected series '1911', got '%d'", rdf.Ebook.PublishedYear)
		}
	})

	t.Run("marc907 Language Code", func(t *testing.T) {
		rdfXml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"><pgterms:ebook><pgterms:marc907>US</pgterms:marc907></pgterms:ebook></rdf:RDF>`
		rdf := unmarshalRDF(t, strings.NewReader(rdfXml))
		if rdf.Ebook.LanguageSubCode != "US" {
			t.Fatalf("expected series 'US', got '%s'", rdf.Ebook.LanguageSubCode)
		}
	})
}

func openRDF(t *testing.T) *unmarshaller.RDF {
	t.Helper()

	file, err := os.Open("../../samples/cache/epub/1400/pg1400.rdf")
	if err != nil {
		t.Fatalf("error opening test RDF file: %s", err)
	}
	return unmarshalRDF(t, file)
}

func unmarshalRDF(t *testing.T, reader io.Reader) *unmarshaller.RDF {
	t.Helper()

	rdf, err := unmarshaller.New(reader)
	if err != nil {
		t.Fatalf("unable to read RDF document: %s", err)
	}
	return rdf
}
