package pgrdf

// Language has the language and dialect details that the book is written in,
// along with any additional notes.
type Language struct {
	Code    string `json:"code"`              // ISO-639-1 two-letter language code. Example: `en`.
	Dialect string `json:"dialect,omitempty"` // ISO 3166-2 subdivision code. Example: `GB`.
	Notes   string `json:"notes,omitempty"`   // Additional details on the language used in the ebook. Example: `Uses 19th century spelling`.
}
