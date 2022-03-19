package commonobjects

// The main api Parsing object:
type ParsingObject struct {
	DataIn string // Data type in e.g. tsv,csv,json.
	DataOut string // Data type out  e.g. tsv,csv,json.
	FileIn []byte // Original File bytes in.
	FileOut []byte // File bytes to be returned.
}