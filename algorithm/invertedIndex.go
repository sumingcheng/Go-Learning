package algorithm

type Doc struct {
	Id       int
	Keywords []string
}

func BuildInvertedIndex(docs []Doc) map[string][]int {
	index := make(map[string][]int)
	for _, doc := range docs {
		for _, keyword := range doc.Keywords {
			index[keyword] = append(index[keyword], doc.Id)
		}
	}
	return index
}
