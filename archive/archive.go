// Package archive contains a couple of helper functions for reading RDF files
// directly from an archive location such as a directory or .tar archive file.
package archive

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"

	"github.com/mrcook/pgrdf"
)

// FromDirectory will lookup the .rdf from the given base directory and PG eText number.
// If lookup speed is important, then this method is faster than using a .tar
// archive, at the cost of disk space.
//
// NOTE: the directory structure must be exactly as extracted from the archive:
//
//	rdf_files/
//	├─ cache/
//	│  └─ epub/
//	│     ├─ 1/
//	│     │  └─ pg1.rdf
//	│     ├─ 2/
//	│     ...
//
// The base dir "rdf_files" would be given with the eText id "1".
func FromDirectory(baseDir string, id int) (*pgrdf.Ebook, error) {
	filename := rdfFilename(baseDir, id)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return pgrdf.ReadRDF(file)
}

// FromTarArchive will read the .rdf file from the official Project Gutenberg archive,
// which contain the standard PG directory structure, i.e. `cache/epub/11/pg11.rdf`.
//
// These are currently available in two archive types:
//
//	rdf-files.tar.bz2
//	rdf-files.tar.zip
//
// Reading from bz2/zip is considerably slower than just from the plain tar,
// so it is recommended to first extract the tar from the bz2/zip archive.
//
//	$ bzip2 -dk rdf-files.tar.bz2
//
// If using the bz2/zip directly is required, then the `.tar.bz2` must first be
// wrapped in a bzip2 reader, before calling the function:
//
//	FromTarArchive(bzip2.NewReader(archiveFile), id)
func FromTarArchive(archiveFile io.Reader, id int) (*pgrdf.Ebook, error) {
	rdfFullPath := rdfFilename("", id)

	r := tar.NewReader(archiveFile)
	for {
		header, err := r.Next()
		if err == io.EOF || err != nil {
			err = errors.Wrapf(err, "eText ID '%d' not found in archive", id)
			return nil, err
		}
		if header.Name == rdfFullPath {
			break
		}
	}

	return pgrdf.ReadRDF(r)
}

func rdfFilename(directory string, id int) string {
	filename := fmt.Sprintf("pg%d.rdf", id)
	return filepath.Join(directory, "cache", "epub", strconv.Itoa(id), filename)
}
