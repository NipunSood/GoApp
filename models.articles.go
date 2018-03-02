package main

type article struct {
	ID      int    `json:"Id"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

var articlesList = []article{
	article{
		ID:      4,
		Title:   "Hello",
		Content: "This is great",
	},

	article{
		ID:      5,
		Title:   "Byee",
		Content: "This is unstopable",
	},
}

func getAllArticle() []article {
	return articlesList
}
