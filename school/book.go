package school

type book struct {
	ID          int32    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Isbn        string   `json:"isbn"`
	Author      []author `json:"author"`
	ImageURL    string   `json:"image_url"`
}

type author struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ImageURL  string `json:"image_url"`
}
