package value_object

type PullRequestEvent struct {
	Action      string      `json:"action"`
	Repository  Repository  `json:"repository"`
	PullRequest PullRequest `json:"pull_request"`
}

type Repository struct {
	Name string `json:"name"`
}

type PullRequest struct {
	ID   int `json:"id"`
	User struct {
		Login string `json:"login"`
	}
	Title string `json:"title"`
	Base  struct {
		Ref string `json:"base"`
	}
	Head struct {
		Ref string `json:"head"`
	}
	URL string `json:"url"`
}
