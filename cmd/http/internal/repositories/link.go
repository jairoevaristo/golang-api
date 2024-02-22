package repositories

var links = make(map[string]string)

func Insert(hash string, url string) {
	links[hash] = url
}

func GetAll() map[string]string {
	return links
}

func FindLinkByHash(hash string) string {
	return links[hash]
}
