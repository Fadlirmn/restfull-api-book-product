package services

import (
	"errors"
	"go-roadmap/models"
	"go-roadmap/repository"
	"go-roadmap/utils"
	"time"

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
func (s *AuthService) Login(username string, password string) (string,string,error)  {
	
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "","", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err != nil {
		return "","", err
	}

	
	accessToken, _ := utils.GenerateToken(user.UserID,user.Role)
	rfTokenStr, _ := utils.GenerateRandomString(32)

	expiresAt := time.Now().Add(24*time.Hour)

	err = s.userRepo.SaveRefreshToken(user.UserID,rfTokenStr,expiresAt)
	if err != nil {
		return "","", err
	}

	return accessToken, rfTokenStr, nil
}

func (s *AuthService)RefreshToken(oldToken string)(string, error)  {
	storedToken, err := s.userRepo.FindRefreshToken(oldToken)
	if err != nil {
		return "", errors.New("refresh token invalid")
	}

	user ,err := s.userRepo.FindByID(storedToken.UserID)
	if err != nil {
		return "", err
	}

	newAccessToken,err := utils.GenerateToken(user.UserID,user.Role)
	if err != nil {
		return "", err
	}
	return newAccessToken,err
}
