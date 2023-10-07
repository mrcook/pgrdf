package pgrdf_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mrcook/pgrdf"
)

func TestReadRDF(t *testing.T) {
	ebook := getEbookFromSampleRdf(t)

	if ebook.ID != 999991234 {
		t.Errorf("unexpected ebook ID, got %d", ebook.ID)
	}
	if ebook.Note != "A description for this RDF" {
		t.Errorf("unexpected ebook description, got '%s'", ebook.Note)
	}
	if ebook.BookType != pgrdf.BookTypeText {
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
	if len(ebook.Series) != 2 {
		t.Errorf("expected 2 series, got %d", len(ebook.Series))
	} else if ebook.Series[0] != "Dickens Best Of" {
		t.Errorf("unexpected series, got '%s'", ebook.Series[0])
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
	if len(ebook.OtherTitles) != 2 {
		t.Errorf("expected 2 other titles, got %d\n", len(ebook.OtherTitles))
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

func TestEbookReadLanguages(t *testing.T) {
	ebook := getEbookFromSampleRdf(t)

	if len(ebook.Languages) != 2 {
		t.Fatalf("expected 2 languages, got %d", len(ebook.Languages))
	}

	if ebook.Languages[0].Code != "en" {
		t.Errorf("unexpected language #1, got '%s'", ebook.Languages[0].Code)
	}
	if ebook.Languages[0].Dialect != "GB" {
		t.Errorf("unexpected language #1 dielect, got '%s'", ebook.Languages[0].Dialect)
	}
	if len(ebook.Languages[0].Notes) != 2 {
		t.Errorf("expected two language notes, got %d", len(ebook.Languages[0].Notes))
	} else {
		if ebook.Languages[0].Notes[0] != "Uses 19th century spelling." {
			t.Errorf("unexpected language #1 notes #1, got '%s'", ebook.Languages[0].Notes[0])
		}
		if ebook.Languages[0].Notes[1] != "This ebook uses a beginning of the 20th century spelling." {
			t.Errorf("unexpected language #1 notes #2, got '%s'", ebook.Languages[0].Notes[1])
		}
	}

	if ebook.Languages[1].Code != "de" {
		t.Errorf("unexpected language #2, got '%s'", ebook.Languages[1].Code)
	}
	if ebook.Languages[1].Dialect != "" {
		t.Errorf("expected language #2 dielect to be blank, got '%s'", ebook.Languages[1].Dialect)
	}
	if len(ebook.Languages[1].Notes) != 2 {
		t.Errorf("expected 2 notes for language #2, got %d", len(ebook.Languages[1].Notes))
	}
}

func TestEbookReadAuthor(t *testing.T) {
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

func TestEbookReadCreators(t *testing.T) {
	ebook := getEbookFromSampleRdf(t)

	cases := []struct {
		id   int
		role pgrdf.MarcRelator
		name string
	}{
		{id: 37, role: pgrdf.RoleAut, name: "Dickens, Charles"},
		{id: 1, role: pgrdf.RoleAdp, name: ""},
		{id: 2, role: pgrdf.RoleAft, name: ""},
		{id: 3, role: pgrdf.RoleArr, name: ""},
		{id: 4, role: pgrdf.RoleAnn, name: ""},
		{id: 5, role: pgrdf.RoleArt, name: ""},
		{id: 6, role: pgrdf.RoleAui, name: ""},
		{id: 7, role: pgrdf.RoleCmm, name: ""},
		{id: 8, role: pgrdf.RoleCmp, name: ""},
		{id: 9, role: pgrdf.RoleCnd, name: ""},
		{id: 10, role: pgrdf.RoleCom, name: ""},
		{id: 11, role: pgrdf.RoleCtb, name: ""},
		{id: 12, role: pgrdf.RoleDub, name: ""},
		{id: 13, role: pgrdf.RoleEdt, name: ""},
		{id: 8397, role: pgrdf.RoleEdt, name: "Snell, F. J. (Frederick John)"},
		{id: 14, role: pgrdf.RoleEgr, name: ""},
		{id: 15, role: pgrdf.RoleIll, name: ""},
		{id: 9473, role: pgrdf.RoleIll, name: "Leech, John"},
		{id: 16, role: pgrdf.RoleLbt, name: ""},
		{id: 17, role: pgrdf.RoleOth, name: ""},
		{id: 18, role: pgrdf.RolePbl, name: ""},
		{id: 19, role: pgrdf.RolePht, name: ""},
		{id: 53417, role: pgrdf.RolePht, name: "Richardson, John A."},
		{id: 20, role: pgrdf.RolePrf, name: ""},
		{id: 21, role: pgrdf.RolePrt, name: ""},
		{id: 22, role: pgrdf.RoleRes, name: ""},
		{id: 23, role: pgrdf.RoleTrc, name: ""},
		{id: 8397, role: pgrdf.RoleTrl, name: ""},
		{id: 1736, role: pgrdf.RoleTrl, name: "Wyllie, David"},
	}

	for _, data := range cases {
		t.Run(fmt.Sprintf("validate ID %d is present", data.id), func(t *testing.T) {
			var creator *pgrdf.Creator
			for i, _ := range ebook.Creators {
				if ebook.Creators[i].ID == data.id && ebook.Creators[i].Role == data.role {
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

func TestEbookReadSubjects(t *testing.T) {
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

func TestEbookReadFiles(t *testing.T) {
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

func TestEbookReadBookshelves(t *testing.T) {
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

	ebook, err := pgrdf.ReadRDF(file)
	if err != nil {
		t.Fatalf("error processing test RDF file: %s", err)
	}
	return ebook
}
