package whitebit

type Error struct {
	Code    int
	Message string
	Errors  map[string][]string
}

func (err Error) Error() string {
	return err.Message
}
