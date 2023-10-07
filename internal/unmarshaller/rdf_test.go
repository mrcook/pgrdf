package unmarshaller_test

import (
	"io"
	"os"
	"testing"

	"github.com/mrcook/pgrdf/internal/unmarshaller"
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
	if r.Work.Comment != "Archives containing the RDF files for *all* our books can be downloaded from our website." {
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

func TestEbookGeneral(t *testing.T) {
	r := openRDF(t)

	e := r.Ebook

	if e.About != "ebooks/999991234" {
		t.Errorf("unexpected rdf:about, got '%s'", e.About)
	}
	if e.Description != "A description for this RDF" {
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
		t.Errorf("expected 2 dcterms:alternative, got %d", len(e.Alternative))
	} else if e.Alternative[0] != "Alternate Title\r\nWith a newline separation" {
		t.Errorf("unexpected dcterms:alternative, got '%s'", e.Alternative[0])
	}
	if len(e.Series) != 2 {
		t.Errorf("expected 2 pgterms:marc440 (series), got %d", len(e.Series))
	} else if e.Series[0] != "Dickens Best Of" {
		t.Errorf("unexpected pgterms:marc440 (series), got '%s'", e.Series[0])
	}
	if e.BookCover != "file:///files/999991234/999991234-h/images/cover.jpg" {
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

func TestLanguage(t *testing.T) {
	r := openRDF(t)
	e := r.Ebook

	if len(e.Languages) != 2 {
		t.Fatalf("expected 2 languages, got %d", len(e.Languages))
	}

	if e.Languages[0].Description.NodeID != "N73e956e8e5d049ac943dfe482ddd5802" {
		t.Errorf("unexpected dcterms:language//rdf:nodeID, got '%s'", e.Languages[0].Description.NodeID)
	}
	if e.Languages[0].Description.Value.DataType != "http://purl.org/dc/terms/RFC4646" {
		t.Errorf("unexpected dcterms:language//rdf:value.datatype, got '%s'", e.Languages[0].Description.Value.DataType)
	}
	if e.Languages[0].Description.Value.Data != "en" {
		t.Errorf("unexpected dcterms:language//rdf:value, got '%s'", e.Languages[0].Description.Value.Data)
	}
	if e.LanguageDialect != "GB" {
		t.Errorf("unexpected marc907 (language sub-code), got '%s'", e.LanguageDialect)
	}

	if e.Languages[1].Description.NodeID != "N9bd0e8afb25241038817304e8e0ff2a9" {
		t.Errorf("unexpected dcterms:language//rdf:nodeID, got '%s'", e.Languages[1].Description.NodeID)
	}
	if e.Languages[1].Description.Value.Data != "de" {
		t.Errorf("unexpected dcterms:language//rdf:value, got '%s'", e.Languages[1].Description.Value.Data)
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
	if a.BirthYear.DataType != "http://www.w3.org/2001/XMLSchema#integer" {
		t.Errorf("unexpected dcterms:creator/agent/birthdate.datatype, got '%s'", a.BirthYear.DataType)
	}
	if a.BirthYear.Value != 1812 {
		t.Errorf("unexpected dcterms:creator/agent/birthdate, got %d", a.BirthYear.Value)
	}
	if a.DeathYear.DataType != "http://www.w3.org/2001/XMLSchema#integer" {
		t.Errorf("unexpected dcterms:creator/agent/deathdate.datatype, got '%s'", a.DeathYear.DataType)
	}
	if a.DeathYear.Value != 1870 {
		t.Errorf("unexpected dcterms:creator/agent/deathdate, got %d", a.DeathYear.Value)
	}
	if a.Webpage.Resource != "https://en.wikipedia.org/wiki/Charles_Dickens" {
		t.Errorf("unexpected dcterms:creator/agent/webpage, got '%s'", a.Webpage.Resource)
	}
}

func TestMarcRelators(t *testing.T) {
	r := openRDF(t)

	if len(r.Ebook.RelAdapters) != 1 {
		t.Errorf("expected 1 adp, got %d", len(r.Ebook.RelAdapters))
	} else if r.Ebook.RelAdapters[0].AgentId() != 1 {
		t.Errorf("unexpected adp agent ID %d", r.Ebook.RelAdapters[0].AgentId())
	}
	if len(r.Ebook.RelAfterwords) != 1 {
		t.Errorf("expected 1 aft, got %d", len(r.Ebook.RelAfterwords))
	} else if r.Ebook.RelAfterwords[0].AgentId() != 2 {
		t.Errorf("unexpected aft agent ID %d", r.Ebook.RelAfterwords[0].AgentId())
	}
	if len(r.Ebook.RelArrangers) != 1 {
		t.Errorf("expected 1 arr, got %d", len(r.Ebook.RelArrangers))
	} else if r.Ebook.RelArrangers[0].AgentId() != 3 {
		t.Errorf("unexpected arr agent ID %d", r.Ebook.RelArrangers[0].AgentId())
	}
	if len(r.Ebook.RelAnnotators) != 1 {
		t.Errorf("expected 1 ann, got %d", len(r.Ebook.RelAnnotators))
	} else if r.Ebook.RelAnnotators[0].AgentId() != 4 {
		t.Errorf("unexpected ann agent ID %d", r.Ebook.RelAnnotators[0].AgentId())
	}
	if len(r.Ebook.RelArtists) != 1 {
		t.Errorf("expected 1 art, got %d", len(r.Ebook.RelArtists))
	} else if r.Ebook.RelArtists[0].AgentId() != 5 {
		t.Errorf("unexpected art agent ID %d", r.Ebook.RelArtists[0].AgentId())
	}
	if len(r.Ebook.RelIntroductions) != 1 {
		t.Errorf("expected 1 aui, got %d", len(r.Ebook.RelIntroductions))
	} else if r.Ebook.RelIntroductions[0].AgentId() != 6 {
		t.Errorf("unexpected aui agent ID %d", r.Ebook.RelIntroductions[0].AgentId())
	}
	if len(r.Ebook.RelCommentators) != 1 {
		t.Errorf("expected 1 cmm, got %d", len(r.Ebook.RelCommentators))
	} else if r.Ebook.RelCommentators[0].AgentId() != 7 {
		t.Errorf("unexpected cmm agent ID %d", r.Ebook.RelCommentators[0].AgentId())
	}
	if len(r.Ebook.RelComposers) != 1 {
		t.Errorf("expected 1 cmm, got %d", len(r.Ebook.RelComposers))
	} else if r.Ebook.RelComposers[0].AgentId() != 8 {
		t.Errorf("unexpected cmm agent ID %d", r.Ebook.RelComposers[0].AgentId())
	}
	if len(r.Ebook.RelConductors) != 1 {
		t.Errorf("expected 1 cnd, got %d", len(r.Ebook.RelConductors))
	} else if r.Ebook.RelConductors[0].AgentId() != 9 {
		t.Errorf("unexpected cnd agent ID %d", r.Ebook.RelConductors[0].AgentId())
	}
	if len(r.Ebook.RelCompilers) != 1 {
		t.Errorf("expected 1 com, got %d", len(r.Ebook.RelCompilers))
	} else if r.Ebook.RelCompilers[0].AgentId() != 10 {
		t.Errorf("unexpected com agent ID %d", r.Ebook.RelCompilers[0].AgentId())
	}
	if len(r.Ebook.RelContributors) != 1 {
		t.Errorf("expected 1 ctb, got %d", len(r.Ebook.RelContributors))
	} else if r.Ebook.RelContributors[0].AgentId() != 11 {
		t.Errorf("unexpected ctb agent ID %d", r.Ebook.RelContributors[0].AgentId())
	}
	if len(r.Ebook.RelDubious) != 1 {
		t.Errorf("expected 1 dub, got %d", len(r.Ebook.RelDubious))
	} else if r.Ebook.RelDubious[0].AgentId() != 12 {
		t.Errorf("unexpected dub agent ID %d", r.Ebook.RelDubious[0].AgentId())
	}
	if len(r.Ebook.RelEditors) != 2 {
		t.Errorf("expected 1 edt, got %d", len(r.Ebook.RelEditors))
	} else {
		if r.Ebook.RelEditors[0].AgentId() != 13 {
			t.Errorf("unexpected edt agent ID %d", r.Ebook.RelEditors[0].AgentId())
		}
		if r.Ebook.RelEditors[1].AgentId() != 8397 {
			t.Errorf("unexpected edt agent ID %d", r.Ebook.RelEditors[1].AgentId())
		} else if r.Ebook.RelEditors[1].Agent.Name != "Snell, F. J. (Frederick John)" {
			t.Errorf("unexpected edt agent name %s", r.Ebook.RelEditors[1].Agent.Name)
		}
	}
	if len(r.Ebook.RelEngravers) != 1 {
		t.Errorf("expected 1 egr, got %d", len(r.Ebook.RelEngravers))
	} else if r.Ebook.RelEngravers[0].AgentId() != 14 {
		t.Errorf("unexpected egr agent ID %d", r.Ebook.RelEngravers[0].AgentId())
	}
	if len(r.Ebook.RelIllustrators) != 2 {
		t.Errorf("expected 1 ill, got %d", len(r.Ebook.RelIllustrators))
	} else {
		if r.Ebook.RelIllustrators[0].AgentId() != 15 {
			t.Errorf("unexpected ill agent ID %d", r.Ebook.RelIllustrators[0].AgentId())
		}
		if r.Ebook.RelIllustrators[1].AgentId() != 9473 {
			t.Errorf("unexpected ill agent ID %d", r.Ebook.RelIllustrators[1].AgentId())
		} else if r.Ebook.RelIllustrators[1].Agent.Name != "Leech, John" {
			t.Errorf("unexpected ill agent name %s", r.Ebook.RelIllustrators[1].Agent.Name)
		}
	}
	if len(r.Ebook.RelLibrettists) != 1 {
		t.Errorf("expected 1 lbt, got %d", len(r.Ebook.RelLibrettists))
	} else if r.Ebook.RelLibrettists[0].AgentId() != 16 {
		t.Errorf("unexpected lbt agent ID %d", r.Ebook.RelLibrettists[0].AgentId())
	}
	if len(r.Ebook.RelOther) != 1 {
		t.Errorf("expected 1 oth, got %d", len(r.Ebook.RelOther))
	} else if r.Ebook.RelOther[0].AgentId() != 17 {
		t.Errorf("unexpected oth agent ID %d", r.Ebook.RelOther[0].AgentId())
	}
	if len(r.Ebook.RelPublishers) != 1 {
		t.Errorf("expected 1 pbl, got %d", len(r.Ebook.RelPublishers))
	} else if r.Ebook.RelPublishers[0].AgentId() != 18 {
		t.Errorf("unexpected pbl agent ID %d", r.Ebook.RelPublishers[0].AgentId())
	}
	if len(r.Ebook.RelPhotographers) != 2 {
		t.Errorf("expected 1 pht, got %d", len(r.Ebook.RelPhotographers))
	} else {
		if r.Ebook.RelPhotographers[0].AgentId() != 19 {
			t.Errorf("unexpected pht agent ID %d", r.Ebook.RelPhotographers[0].AgentId())
		}
		if r.Ebook.RelPhotographers[1].AgentId() != 53417 {
			t.Errorf("unexpected pht agent ID %d", r.Ebook.RelPhotographers[1].AgentId())
		} else if r.Ebook.RelPhotographers[1].Agent.Name != "Richardson, John A." {
			t.Errorf("unexpected pht agent name %s", r.Ebook.RelPhotographers[1].Agent.Name)
		}
	}
	if len(r.Ebook.RelPerformers) != 1 {
		t.Errorf("expected 1 prf, got %d", len(r.Ebook.RelPerformers))
	} else if r.Ebook.RelPerformers[0].AgentId() != 20 {
		t.Errorf("unexpected prf agent ID %d", r.Ebook.RelPerformers[0].AgentId())
	}
	if len(r.Ebook.RelPrinters) != 1 {
		t.Errorf("expected 1 prt, got %d", len(r.Ebook.RelPrinters))
	} else if r.Ebook.RelPrinters[0].AgentId() != 21 {
		t.Errorf("unexpected prt agent ID %d", r.Ebook.RelPrinters[0].AgentId())
	}
	if len(r.Ebook.RelResearchers) != 1 {
		t.Errorf("expected 1 res, got %d", len(r.Ebook.RelResearchers))
	} else if r.Ebook.RelResearchers[0].AgentId() != 22 {
		t.Errorf("unexpected res agent ID %d", r.Ebook.RelResearchers[0].AgentId())
	}
	if len(r.Ebook.RelTranscribers) != 1 {
		t.Errorf("expected 1 trc, got %d", len(r.Ebook.RelTranscribers))
	} else if r.Ebook.RelTranscribers[0].AgentId() != 23 {
		t.Errorf("unexpected trc agent ID %d", r.Ebook.RelTranscribers[0].AgentId())
	}
	if len(r.Ebook.RelTranslators) != 2 {
		t.Errorf("expected 1 trl, got %d", len(r.Ebook.RelTranslators))
	} else {
		if r.Ebook.RelTranslators[0].AgentId() != 8397 {
			t.Errorf("unexpected trl agent ID %d", r.Ebook.RelTranslators[0].AgentId())
		}
		if r.Ebook.RelTranslators[1].AgentId() != 1736 {
			t.Errorf("unexpected trl agent ID %d", r.Ebook.RelTranslators[1].AgentId())
		} else if r.Ebook.RelTranslators[1].Agent.Name != "Wyllie, David" {
			t.Errorf("unexpected trl agent name %s", r.Ebook.RelTranslators[1].Agent.Name)
		}
	}
}

// NOTE: test for when an RDF uses `Various` for marc906 instead of a year number.
func TestMarcRelators_InvalidMarc906(t *testing.T) {
	file, err := os.Open("../../samples/marc906-error.rdf")
	if err != nil {
		t.Fatalf("error opening test RDF file: %s", err)
	}
	r, err := unmarshaller.New(file)
	if err != nil {
		t.Fatalf("unexpected RDF read error: %s", err)
	}
	if r.Ebook.PublishedYear != 0 {
		t.Errorf("expected PublishedYear to not have been set, got %d", r.Ebook.PublishedYear)
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

	if f.About != "https://www.example.org/files/999991234/999991234-h.zip" {
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
	if f.IsFormatOf.Resource != "ebooks/999991234" {
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
	rdf := openRDF(t)
	if rdf.Ebook.LOC != "77177891" {
		t.Errorf("unexpected marc010 LoC Number '%s'", rdf.Ebook.LOC)
	}
	if rdf.Ebook.ISBN != "0-397-00033-2" {
		t.Errorf("unexpected marc020 ISBN '%s'", rdf.Ebook.ISBN)
	}
	if rdf.Ebook.Edition != "The Charles Dickens Edition" {
		t.Errorf("unexpected marc250 edition '%s'", rdf.Ebook.Edition)
	}
	if rdf.Ebook.OriginalPublication != "United Kingdom: J. Johnson, 1794." {
		t.Errorf("unexpected marc260 original publication '%s'", rdf.Ebook.OriginalPublication)
	}
	if rdf.Ebook.SourceDescription != "Musical score" {
		t.Errorf("unexpected marc300 source description '%s'", rdf.Ebook.SourceDescription)
	}
	if len(rdf.Ebook.Series) != 2 {
		t.Errorf("expected 2 marc440 series, got %d", len(rdf.Ebook.Series))
	} else if rdf.Ebook.Series[0] != "Dickens Best Of" {
		t.Errorf("unexpected marc440 series, got '%s'", rdf.Ebook.Series[0])
	}
	if len(rdf.Ebook.Credits) != 2 {
		t.Errorf("expected 2 marc508 credit entries, got %d", len(rdf.Ebook.Credits))
	}
	if rdf.Ebook.Credits[0] != "Updated: 2022-07-14" {
		t.Errorf("unexpected marc508 credit[0] '%s'", rdf.Ebook.Credits[0])
	}
	if rdf.Ebook.Credits[1] != "Produced by Anon." {
		t.Errorf("unexpected marc508 credit[1] '%s'", rdf.Ebook.Credits[1])
	}
	if rdf.Ebook.Summary != "A fun version of A Christmas Carol." {
		t.Errorf("unexpected marc520 summary '%s'", rdf.Ebook.Summary)
	}
	if rdf.Ebook.LanguageDialect != "GB" {
		t.Errorf("unexpected marc907 summary '%s'", rdf.Ebook.LanguageDialect)
	}
	if len(rdf.Ebook.LanguagesNotes) != 2 {
		t.Errorf("expected 1 language note, got %d", len(rdf.Ebook.LanguagesNotes))
	} else {
		if rdf.Ebook.LanguagesNotes[0] != "Uses 19th century spelling." {
			t.Errorf("unexpected marc546 language note #1 '%s'", rdf.Ebook.LanguagesNotes[0])
		}
		if rdf.Ebook.LanguagesNotes[1] != "This ebook uses a beginning of the 20th century spelling." {
			t.Errorf("unexpected marc546 language note #2 '%s'", rdf.Ebook.LanguagesNotes[1])
		}
	}
	if rdf.Ebook.BookCover != "file:///files/999991234/999991234-h/images/cover.jpg" {
		t.Errorf("unexpected marc901 book cover link '%s'", rdf.Ebook.BookCover)
	}
	if rdf.Ebook.TitlePageImage != "https://example.org/ebook1/title.jpg" {
		t.Errorf("unexpected marc902 title page link '%s'", rdf.Ebook.TitlePageImage)
	}
	if rdf.Ebook.BackCover != "https://example.org/ebook1/back.jpg" {
		t.Errorf("unexpected marc903 back cover link '%s'", rdf.Ebook.BackCover)
	}
	if rdf.Ebook.SourceLink != "https://example.com/ebooks/1/something" {
		t.Errorf("unexpected marc904 source link '%s'", rdf.Ebook.SourceLink)
	}
	if rdf.Ebook.PgDpClearance != "19991231235959randomthing" {
		t.Errorf("unexpected marc905 PGDP clearance code '%s'", rdf.Ebook.PgDpClearance)
	}
	if rdf.Ebook.PublishedYear != 1861 {
		t.Errorf("unexpected marc906 published year '%d'", rdf.Ebook.PublishedYear)
	}
	if rdf.Ebook.LanguageDialect != "GB" {
		t.Errorf("expected marc907 language dielect code '%s'", rdf.Ebook.LanguageDialect)
	}
}

func openRDF(t *testing.T) *unmarshaller.RDF {
	t.Helper()

	file, err := os.Open("../../samples/cache/epub/999991234/pg999991234.rdf")
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
