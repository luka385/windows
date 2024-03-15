package domain

type Pet struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Age      int    `json:"age"`
	OwnerID  string `json:"owner_id"`
	ImageURL string `json:"image_url"`
}

type PetRepository interface {
	Create(pet *Pet) error
	GetByID(id string) (*Pet, error)
	Update(id string, pet *Pet) error
	Delete(id string) error
}
