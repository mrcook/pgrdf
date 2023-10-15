package pgrdf

// AuthorLink is an external resource about the author, usually a Wikipedia link.
// `<rdf:Description>`
type AuthorLink struct {
	// URL for this author.
	// `<rdf:Description rdf:about="">`
	URL string `json:"url"`

	// A short description about this link.
	// `<dcterms:description>`
	Description string `json:"description"`
}
