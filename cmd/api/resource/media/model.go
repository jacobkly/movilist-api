package media

type DTO struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
	Director      string `json:"director"`
}

type Form struct {
	Title         string `json:"title"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
	Director      string `json:"director"`
}
