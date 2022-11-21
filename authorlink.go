package pgrdf

// AuthorLink is an external resource about the author.
// Usually this is a link to Wikipedia.
type AuthorLink struct {
	// URL for this author.
	URL string `json:"url"`

	// A short description about this link.
	Description string `json:"description"`
}
