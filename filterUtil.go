package main

func filterByCategories(posts []Post, filterCategories []string) []Post {
	// if a user didn't chose to filter the categories, we return posts as they are
	if len(filterCategories) == 0 {
		return posts
	}

	sortedPosts := []Post{}
	for _, post := range posts {
		categoties := post.Categories
		if containsArr(filterCategories, categoties) {
			sortedPosts = append(sortedPosts, post)
		}
	}
	return sortedPosts
}
