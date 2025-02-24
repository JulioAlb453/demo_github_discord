package value_object

type PullRequestEvent struct {
	Action       string      `json:"action"`
	PullResquest PullRequest `json:"pull_request"`
}

type PullRequest struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Head   Branch `json: head"`
	Base  Branch `json:"base"`
	URL   string `json:"url"`
}

type Branch struct {
	Ref string `json:"ref"`
	SHA string `json:"sha"`
}
