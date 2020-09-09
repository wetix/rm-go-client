package rm

// Error :
type Error struct {
	raw []byte
}

// Error :
func (e Error) Error() string {
	return "rm: " + string(e.raw)
}

// Raw :
func (e Error) Raw() []byte {
	return e.raw
}
