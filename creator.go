package pgrdf

// Creator is a person involved in the creation of the work, such as the author, illustrator, etc.
type Creator struct {
	ID      int             `json:"id"`                  // Unique Project Gutenberg ID.
	Name    string          `json:"name"`                // Name of the creator.
	Aliases []string        `json:"aliases,omitempty"`   // Any aliases for the creator.
	Born    int             `json:"born_year,omitempty"` // Date of Birth.
	Died    int             `json:"died_year,omitempty"` // Date of Death.
	Role    MarcRelatorCode `json:"role,omitempty"`      // Code indicating the role. e.g. `aut`, `edt`, `ill`, etc.
	WebPage string          `json:"webpage,omitempty"`   // URL for this creator (usually Wikipedia).
}
