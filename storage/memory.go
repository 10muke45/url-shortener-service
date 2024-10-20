package storage

var urlStore = make(map[string]string)

func Store(shortURL, originalURL string) {
	urlStore[shortURL] = originalURL
}

func Retrieve(shortURL string) string {
	return urlStore[shortURL]
}
