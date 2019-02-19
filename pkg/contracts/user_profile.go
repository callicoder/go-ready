package contracts

type UserProfile struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ImageUrl  string `json:"image_url"`
}
