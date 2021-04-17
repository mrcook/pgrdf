package unmarshaller_test

import (
	"os"
	"testing"

	"github.com/mrcook/pgrdf/_internal/unmarshaller"
)

func TestNamespaces(t *testing.T) {
	r, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

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
}

func TestWorkNode(t *testing.T) {
	r, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

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
	r, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

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
	r, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}
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
	if e.Publisher != "Project Gutenberg" {
		t.Errorf("unexpected dcterms:publisher, got '%s'", e.Publisher)
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
	r, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

	if len(r.Ebook.Creators) != 1 {
		t.Errorf("expected 1 dcterms:creator, got %d", len(r.Ebook.Creators))
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

func TestSubjects(t *testing.T) {
	r, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

	if len(r.Ebook.Subjects) != 9 {
		t.Errorf("expected 9 dcterms:subject, got %d", len(r.Ebook.Subjects))
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
	r, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

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
	r, err := openRDF()
	if err != nil {
		t.Fatalf("unable to read sample file: %s", err)
	}

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

func openRDF() (*unmarshaller.RDF, error) {
	r := &unmarshaller.RDF{}

	file, err := os.Open("../../samples/cache/epub/1400/pg1400.rdf")
	if err != nil {
		return r, err
	}

	return unmarshaller.New(file)
}
