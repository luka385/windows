package application

import "github.com/luka385/my-pet-app/domain"

type PetUsecase struct {
	repo domain.PetRepository
}

func NewPetUsecase(repo domain.PetRepository) *PetUsecase {
	return &PetUsecase{
		repo: repo,
	}
}

func (u *PetUsecase) CreatePet(pet *domain.Pet) error {
	return u.repo.Create(pet)
}

func (u *PetUsecase) GetPetByID(id string) (*domain.Pet, error) {
	return u.repo.GetByID(id)
}

func (u *PetUsecase) UpdatePet(id string, UpdatePet *domain.Pet) error {
	return u.repo.Update(id, UpdatePet)
}

func (u *PetUsecase) DeletePet(id string) error {
	return u.repo.Delete(id)
}
