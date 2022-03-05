# pgrdf changelog

## HEAD


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
