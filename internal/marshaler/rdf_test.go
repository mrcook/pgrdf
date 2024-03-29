package marshaler_test

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"regexp"
	"testing"

	"github.com/mrcook/pgrdf/internal/marshaler"
	"github.com/mrcook/pgrdf/internal/unmarshaler"
)

// NOTE: there have been various manual changes to the test RDF (pg999991234.rdf):
//   - The order of tags in a Gutenberg RDF file is not fixed, but the output
//     of the XML marshal is, so the source file has been changed to match.
//   - Various characters are convert to HTML entities by the xml package
//     e.g. `'` -> `&#39;`, so these have been changed in the test file.
//   - Currently the xml package will not emit self-closing tags (`<tag />`)
//     so this test removes them before comparing the strings.
func TestRDF_FromUnmarshaler(t *testing.T) {
	file := openFile(t)
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		t.Fatalf("unable to read RDF bytes %s", err)
	}

	file = openFile(t)
	unmarshaledRdf := unmarshalRDF(t, file)
	rdf := marshaler.FromUnmarshaler(unmarshaledRdf)

	data, err := xml.MarshalIndent(rdf, "", "  ")
	if err != nil {
		t.Fatalf("error marshaling RDF %s", err)
	}

	// replace empty tags with self-closing tags
	r := regexp.MustCompile(`></[^>]+?>`)
	output := r.ReplaceAllString(string(data), "/>")

	// add the xml declaration
	dataXML := xml.Header + output

	// show where the diversion happens
	sourceBytes := buf.Bytes()
	sourceBytes = bytes.ReplaceAll(sourceBytes, []byte("Alternate Title&#13;\nWith a newline separation"), []byte("Alternate Title&#xD;&#xA;With a newline separation"))
	index := -1
	for i := 0; i < len(dataXML); i++ {
		if dataXML[i] != sourceBytes[i] {
			index = i
			break
		}
	}
	if index >= 0 {
		endIndex := index
		if len(dataXML) > index {
			endIndex += 1
		}
		t.Errorf("unexpected marshaled output at position %d\n%s\n", index, dataXML[0:endIndex])
	}
}

func openFile(t *testing.T) *os.File {
	file, err := os.Open("../../samples/cache/epub/999991234/pg999991234.rdf")
	if err != nil {
		t.Fatalf("error opening test RDF file: %s", err)
	}
	return file
}

func unmarshalRDF(t *testing.T, reader io.Reader) *unmarshaler.RDF {
	t.Helper()

	rdf, err := unmarshaler.New(reader)
	if err != nil {
		t.Fatalf("unable to read RDF document: %s", err)
	}
	return rdf
}
