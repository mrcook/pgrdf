# pgrdf - a Project Gutenberg RDF wrapper library

A small Go language library for interacting with the Project Gutenberg Catalog
Metadata (RDF XML), mapping them on to a more usable set of types.

`pgrdf` also provides a helper function for reading the RDF metadata directly
from a Project Gutenberg `tar` archive. See the usage section below for more
information.

A gutenberg RDF source file might look like [pg1400.rdf](samples/cache/epub/1400/pg1400.rdf)
from the sample directory:

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
```

Using the `pgrdf` library, the metadata can be accessed more easily via the `Ebook`
object, which can also be un/marshalled to JSON:

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
    "webpage": "https://en.wikipedia.org/wiki/Charles_Dickens"
  }]
}
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
	rdfFile, _ := os.Open("./pg1400.rdf")
	ebook, _ := pgrdf.ReadRDF(rdfFile)

	ebook.Titles = append(ebook.Titles, "In Three Volumes")

	w := bytes.NewBuffer([]byte{}) // create io.Writer
	_ = ebook.WriteRDF(w)             // marshall to RDF XML data
	fmt.Println(w.String())

	data, _ := json.Marshal(ebook) // marshall to JSON
	fmt.Println(string(data))
}
```

It is possible to read the metadata directly from the official Project Gutenberg
offline RDF catalog archive: http://www.gutenberg.org/cache/epub/feeds/.

There are currently two archives available:

    rdf-files.tar.bz2
    rdf-files.tar.zip

Reading from a bz2/zip is considerably slower than just the plain `tar` archive,
so it is recommended to first extract the tarball from the bz2/zip archive.
Example using Linux:

    $ bzip2 -dk rdf-files.tar.bz2

If this is not possible/desirable then the `.tar.bz2` must first be wrapped in a
`bzip2` reader:

    rdf, err := FromTarArchive(bzip2.NewReader(archiveFile), id)

When an archive is fully extracted to a local directory, there is a helper
function for looking up a RDF using the PG eText ID:

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

Copyright (c) 2018-2022 Michael R. Cook. All rights reserved.

This work is licensed under the terms of the MIT license.
For a copy, see <https://opensource.org/licenses/MIT>.
