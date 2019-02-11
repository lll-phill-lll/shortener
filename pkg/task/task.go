package task

type Task struct {
	URL       string
	Hash string
	HostURL string
}

func (t Task) GetHashedURL() string {
	return t.HostURL + "/" + t.Hash
}
