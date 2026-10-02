package persist

func Recover(path string) ([]Entry, error) {
	log := New(path)
	return log.ReadAll()
}
