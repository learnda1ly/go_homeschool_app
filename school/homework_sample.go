package school

type homeworkSample struct {
	ID          int32   `json:"id"`
	SubjectID   int32   `json:"subject_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	FileURL     string  `json:"file_url"`
	Thumbnail   string  `json:"thumbnail"`
	Grade       float32 `json:"grade"`
	LetterGrade rune    `json:"letter_grade"`
}
