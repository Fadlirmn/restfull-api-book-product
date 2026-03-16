package services

import(
	"go-roadmap/models"
	"go-roadmap/repository"
	"go-roadmap/utils"
	
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{
	userRepo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService  {
	return &AuthService{userRepo: repo}
}

func (s *AuthService) Register(user models.User) error {
	
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	user.Role = "user"

	s.userRepo.Save(user)

	return nil
}
func (s *AuthService) Login(username string, password string) (string,error)  {
	
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err != nil {
		return "", err
	}

	return utils.GenerateToken(user.UserID,user.Role)
}