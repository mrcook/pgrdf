# pgrdf changelog

## HEAD


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
