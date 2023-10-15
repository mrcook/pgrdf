# pgrdf - a Project Gutenberg RDF library

A library written in the Go language for reading and writing Project Gutenberg
RDF documents using a simpler set of intermediary data types, which can also be
marshaled to JSON for a more compact representation of the metadata.

Helper functions are provided for reading RDF files directly from their `tar`
archive. See the usage section below for more information.

The `Ebook` struct is used as an intermediary representation of the metadata,
which provides a much easier set of data types than needing to handle RDF
directly, and can also be marshaled to JSON.

The following is a (truncated) JSON example:

```json
{
  "id": 1400,
  "released": "1998-07-01",
  "titles": ["Great Expectations"],
  "creators": [{
    "id": 37,
    "name": "Dickens, Charles",
    "aliases": [
      "Dickens, Charles John Huffam",
      "Boz"
    ],
    "born_year": 1812,
    "died_year": 1870,
    "webpages": ["https://en.wikipedia.org/wiki/Charles_Dickens"]
  }]
}
```

And here's the corresponding RDF snippet for [Great Expectations](https://gutenberg.org/ebooks/1400.rdf):

```xml
<?xml version="1.0" encoding="utf-8"?>
<rdf:RDF xml:base="http://www.gutenberg.org/"
         xmlns:dcam="http://purl.org/dc/dcam/"
         xmlns:dcterms="http://purl.org/dc/terms/"
         xmlns:pgterms="http://www.gutenberg.org/2009/pgterms/"
         xmlns:cc="http://web.resource.org/cc/"
         xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
         xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#">
    <pgterms:ebook rdf:about="ebooks/1400">
        <dcterms:title>Great Expectations</dcterms:title>
        <dcterms:creator>
            <pgterms:agent rdf:about="2009/agents/37">
                <pgterms:name>Dickens, Charles</pgterms:name>
                <pgterms:alias>Dickens, Charles John Huffam</pgterms:alias>
                <pgterms:alias>Boz</pgterms:alias>
                <pgterms:webpage rdf:resource="https://en.wikipedia.org/wiki/Charles_Dickens"/>
                <pgterms:birthdate rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">1812</pgterms:birthdate>
                <pgterms:deathdate rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">1870</pgterms:deathdate>
            </pgterms:agent>
        </dcterms:creator>
        <dcterms:issued rdf:datatype="http://www.w3.org/2001/XMLSchema#date">1998-07-01</dcterms:issued>
        <!-- ... -->
    </pgterms:ebook>
</rdf:RDF>
```


## Usage

A basic example might be:

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mrcook/pgrdf"
)

func main() {
	rdfFile, _ := os.Open("/path/to/pg1400.rdf")
	ebook, _ := pgrdf.ReadRDF(rdfFile)

	ebook.Titles = append(ebook.Titles, "In Three Volumes")

	w := bytes.NewBuffer([]byte{}) // create an io.Writer
	_ = ebook.WriteRDF(w)          // write the RDF data

	data, _ := json.Marshal(ebook) // marshal to JSON
	fmt.Println(string(data))
}
```

It is possible to read an RDF directly from the official Project Gutenberg
offline catalog archive: http://www.gutenberg.org/cache/epub/feeds/.

There are currently two archives available:

    rdf-files.tar.bz2
    rdf-files.tar.zip

Reading from a `bz2` is considerably slower than just the plain `tar` archive,
so it is recommended to first extract the tarball from the `bz2` archive.
Example on Linux:

    $ bzip2 -dk rdf-files.tar.bz2

If this is not possible/desirable then the `.tar.bz2` must first be wrapped in a
`bzip2` reader:

    rdf, err := archive.FromTarArchive(bzip2.NewReader(archiveFile), id)

When an archive is fully extracted to a local directory, the `FromDirectory`
function can be used:

    rdf, err := archive.FromDirectory("/rdf_files_dir", 1400)

It is important to note that the directory structure must be exactly as
extracted from the archive:

    rdf_files_dir/
    ├─ cache/
    │  └─ epub/
    │     ├─ 1/
    │     │  └─ pg1.rdf
    │     ├─ 2/
    │     ...


## LICENSE

Copyright (c) 2018-2023 Michael R. Cook. All rights reserved.

This work is licensed under the terms of the MIT license.
For a copy, see <https://opensource.org/licenses/MIT>.
