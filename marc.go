package pgrdf

// MarcRelatorCode representing a creator role: `aut`, `edt`, etc.
type MarcRelatorCode string

const (
	RoleAut MarcRelatorCode = "aut"
	RoleCom MarcRelatorCode = "com"
	RoleCtb MarcRelatorCode = "ctb"
	RoleEdt MarcRelatorCode = "edt"
	RoleIll MarcRelatorCode = "ill"
	RoleTrl MarcRelatorCode = "trl"
)
