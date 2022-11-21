package pgrdf_test

import (
	"testing"

	"github.com/mrcook/pgrdf"
)

func TestEbook_AddSubject(t *testing.T) {
	ebook := pgrdf.Ebook{}
	ebook.AddSubject("Fiction", "http://purl.org/dc/terms/LCSH")

	if len(ebook.Subjects) != 1 {
		t.Fatalf("expected 1 subject, got: %d", len(ebook.Subjects))
	}

	subject := ebook.Subjects[0]

	if subject.Heading != "Fiction" {
		t.Errorf("unexpected heading: '%s'", subject.Heading)
	}
	if subject.Schema != "http://purl.org/dc/terms/LCSH" {
		t.Errorf("unexpected schema: '%s'", subject.Schema)
	}
}

func TestEbook_AddBookshelf(t *testing.T) {
	ebook := pgrdf.Ebook{}
	ebook.AddBookshelf("My Bookshelf", "2009/pgterms/Bookshelf")

	if len(ebook.Bookshelves) != 1 {
		t.Fatalf("expected 1 bookshelf, got: %d", len(ebook.Bookshelves))
	}

	shelf := ebook.Bookshelves[0]

	if shelf.Name != "My Bookshelf" {
		t.Errorf("unexpected name: '%s'", shelf.Name)
	}
	if shelf.Resource != "2009/pgterms/Bookshelf" {
		t.Errorf("unexpected resource: '%s'", shelf.Resource)
	}
}

func TestEbook_AddAuthorLink(t *testing.T) {
	ebook := pgrdf.Ebook{}
	ebook.AddAuthorLink("en.wikipedia", "https://en.wikipedia.org/wiki/Charles_Dickens")

	if len(ebook.AuthorLinks) != 1 {
		t.Fatalf("expected 1 link, got: %d", len(ebook.AuthorLinks))
	}

	link := ebook.AuthorLinks[0]

	if link.Description != "en.wikipedia" {
		t.Errorf("unexpected description: '%s'", link.Description)
	}
	if link.URL != "https://en.wikipedia.org/wiki/Charles_Dickens" {
		t.Errorf("unexpected url: '%s'", link.URL)
	}
}
