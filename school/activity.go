package school

type activity struct {
	ID            int32     `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Purpose       string    `json:"purpose"`
	Participants  []learner `json:"participants"`
	StartDateTime int32     `json:"start_date_time"`
	EndDateTime   int32     `json:"end_date_time"`
}
