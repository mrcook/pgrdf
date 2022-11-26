package marshaller_test

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"regexp"
	"testing"

	"github.com/mrcook/pgrdf/_internal/marshaller"
	"github.com/mrcook/pgrdf/_internal/unmarshaller"
)

// NOTE: there have been various manual changes to the test RDF (pg11.rdf):
//   - The order of tags in a Gutenberg RDF file is not fixed, but the output
//     of the XML marshall is, so the source file has been changed to match.
//   - Various characters are convert to HTML entities by the xml package
//     e.g. `'` -> `&#39;`, so these have been changed in the test file.
//   - Currently the xml package will not emit self-closing tags (`<tag />`)
//     so this test removes them before comparing the strings.
func TestRDF_FromUnmarshaller(t *testing.T) {
	file := openFile(t)
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		t.Fatalf("unable to read RDF bytes %s", err)
	}

	file = openFile(t)
	unmarshalledRdf := unmarshalRDF(t, file)
	rdf := marshaller.FromUnmarshaller(unmarshalledRdf)

	data, err := xml.MarshalIndent(rdf, "", "  ")
	if err != nil {
		t.Fatalf("error marshalling RDF %s", err)
	}

	// replace empty tags with self-closing tags
	r := regexp.MustCompile(`></[^>]+?>`)
	output := r.ReplaceAllString(string(data), "/>")

	// add the xml declaration
	dataXML := xml.Header + output

	if len(dataXML) != len(buf.String()) {
		t.Fatalf("strings of different length: got %d, want %d", len(dataXML), len(buf.String()))
	}
	if dataXML != buf.String() {
		println(dataXML)
		t.Fatalf("XML data does not match")
	}
}

func openFile(t *testing.T) *os.File {
	file, err := os.Open("../../samples/cache/epub/11/pg11.rdf")
	if err != nil {
		t.Fatalf("error opening test RDF file: %s", err)
	}
	return file
}

func unmarshalRDF(t *testing.T, reader io.Reader) *unmarshaller.RDF {
	t.Helper()

	rdf, err := unmarshaller.New(reader)
	if err != nil {
		t.Fatalf("unable to read RDF document: %s", err)
	}
	return rdf
}
