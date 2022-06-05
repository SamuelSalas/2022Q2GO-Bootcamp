package entity

type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Gender  string `json:"gender"`
	Image   string `json:"image"`
	Url     string `json:"url"`
	Created string `json:"created"`
}
