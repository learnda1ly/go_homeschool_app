/* Package school supports the homeschool app with datastructures. */
package school

type Learner struct {
	FamilyID    int32  `json:"family_id"`
	FirstName   string `json:"first_name"`
	Grade       string `json:"grade"`
	ID          int32  `json:"id"`
	ImageURL    string `json:"image_url"`
	LastName    string `json:"last_name"`
	LearnerType string `json:"learner_type"`
}

func GetLearnerList() []Learner {
	return []Learner{
		{
			ID:          1,
			FirstName:   "Stephen",
			LastName:    "Quinlan",
			LearnerType: "Self-Directed",
			Grade:       "na",
			ImageURL:    "test.learnda1ly.com",
			FamilyID:    1,
		},
		{
			ID:          2,
			FirstName:   "Caleb",
			LastName:    "Quinlan",
			LearnerType: "Grade School",
			Grade:       "1",
			ImageURL:    "test.learnda1ly.com",
			FamilyID:    1,
		},
	}
}
