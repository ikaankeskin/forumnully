package main

func filterByCategories(posts []Post, filterCategories []string) []Post {
	if len(filterCategories) == 0 {
		return posts
	}
	result := []Post{}
	for _, post := range posts {
		categoties := post.Categories
		for _, cat := range categoties {
			if contains(filterCategories, cat) {
				result = append(result, post)
				break
			}
		}
	}
	return result
}
