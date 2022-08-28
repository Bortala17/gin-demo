package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}
