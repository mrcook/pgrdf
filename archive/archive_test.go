package archive_test

import (
	"os"
	"testing"

	"github.com/mrcook/pgrdf/archive"
)

func TestDirectoryLookup(t *testing.T) {
	rdf, err := archive.FromDirectory("../samples", 999991234)
	if err != nil {
		t.Fatalf("unpexcted error reading RDF file: %s", err)
	}

	if len(rdf.Titles) != 1 {
		t.Fatalf("expected 1 title, got %d", len(rdf.Titles))
	}
	if rdf.Titles[0] != "Great Expectations" {
		t.Errorf("unexpected title found, got '%s'", rdf.Titles[0])
	}
}

func TestArchiveLookup(t *testing.T) {
	file, err := os.Open("../samples/rdf-files-test.tar")
	if err != nil {
		t.Fatalf("Unable to open RDF tar archive: %s", err)
	}

	rdf, err := archive.FromTarArchive(file, 1400)
	if err != nil {
		t.Fatalf("unpexcted error reading RDF tar archive: %s", err)
	}

	if len(rdf.Titles) != 1 {
		t.Fatalf("expected 1 title, got %d", len(rdf.Titles))
	}
	if rdf.Titles[0] != "Great Expectations" {
		t.Errorf("unexpected title found, got '%s'", rdf.Titles[0])
	}
}
