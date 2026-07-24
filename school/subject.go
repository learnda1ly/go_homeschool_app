package school

type subject struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Duration    int32  `json:"duration"`
	LogoURL     string `json:"logo_url"`
}
