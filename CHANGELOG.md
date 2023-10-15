# pgrdf changelog

## HEAD


## v1.7.0 (2023-10-15)

Many of the previously added `marcrel` tags have now been made available in
the `Ebook` struct, for example, MARC 546 (language notes), MARC 440 (series),
MARC 250 (edition), and MARC 904 (source links).

### BREAKING CHANGES

Various tags in the RDFs may have multiple entries but they were previously
handled as single tags. The un/marshalers have been updated to handle these
and the `Ebook` struct updated to support slices where needed. Some fields
have also been renamed.

* `OtherTitles` has been renamed to `AlternateTitles`.
* `Language` is now a slice named `Languages` and includes only an alpha 2 or
  alpha 3 language code string, meaning the dialect and language notes have
  been moved to their own struct fields.
* `BookCoverFilename` is now a slice named `BookCovers`.
* `Note` is now a slice named `Notes`.
* `WebPage` in `Creator` is now a slice named `WebPages`.


## v1.6.1 (2023-09-30)

* Handle the case when an RDF uses `Various` in `marc906` instead of a year number.


## v1.6.0 (2023-09-05)

* Use a custom type for the `BookType` for all know types in the PG collection.
* Improve unmarshaller to handle both CR and LF for `dcterms:title`/`dcterms:alternative`.


## v1.5.0 (2022-11-27)

* Renamed `MarcRelatorCode` to `MarcRelator`.
* Added the remaining `marcrel` tags:
  * `marcrel:pbl`, `marcrel:adp`, `marcrel:pht`, etc.
  * Matching what are found in the current PG collection.
  * Which also means supporting empty tags, e.g. `<marcrel:adp rdf:resource="2009/agents/1" />`.
* The order of the marshalled RDF XML tags has been change to make it a little
  easier for humans to find information about the work.
* The RDF marshaller structs now have some useful comments.
* General cleanup and improvements related to the above topics.


## v1.4.0 (2022-11-26)

IMPORTANT: this release contains breaking changes! 

* `pgrdf.ToRDF()` has been renamed to `WriteRDF()`
* `pgrdf.NewEbook()` has been renamed to `ReadRDF()`
* The `Language` field on `Ebook` has changed from a `string` type to `pgrdf.Language`.

Additional changes:

* RDF unmarshalling now processes all MARC codes used by PG
  - that's all codes found in the 202-11-05 `rdf_files.tar.bz2` archive
* RDF marshalling now includes:
  - all missing tags, such as the contributors, and the new marc tags.
  - the generated XML is now tested against the `pg11.rdf` sample file.
* WriteRDF function now includes the XML declaration header and fixes the self-closing tags.
* Updated the sample RDF with more fake data
  - its number was also changed to a value PG will never use


## v1.3.0 (2022-03-05)

* Extract MARC Relators data for compilers (`marcrel:com`)
* Extract MARC Relators data for contributors (`marcrel:ctb`)
* Extract book Series from `pgterms:marc440`


## v1.2.0 (2022-03-05)

* Extract MARC Relators (`marcrel`) data:
  - editors (`edt`), illustrators (`ill`), and translators (`trl`).
  - adding them as `creators` in the JSON output.
  - this adds a new `role` field to the `Creator` object.
* Extract `marc901` book cover filename.
* Extract `marc907` language locale.
  - appending the locale to the language field, e.g. `"language": "en-GB"`.
* Refactoring of `ebook.go`, in particular `mapUnmarshalled()`.


## v1.1.0 (2022-01-08)

* Extract published year from the `marc906` tag.


## v1.0.1 (2021-06-12)

* Add missing LICENSE file


## v1.0.0 (2021-04-17)

* First release.
