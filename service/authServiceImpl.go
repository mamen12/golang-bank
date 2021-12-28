package service

import (
	"bank/dto"
	"bank/entity"
	"bank/repo"
	"time"

	"github.com/golang-jwt/jwt"
)

func NewAuthService(authRepo *repo.AuthRepo) AuthService {
	historyRepo := repo.NewHistoryRepo()
	historyService := NewHistoryService(&historyRepo)

	return &AuthServiceImpl{
		HistoryService: historyService,
		AuthRepo:       *authRepo,
	}
}

type AuthServiceImpl struct {
	HistoryService HistoryService
	AuthRepo       repo.AuthRepo
}

func (service *AuthServiceImpl) GenerateToken(customer entity.Customer) string {

	credential := dto.UserCredential{
		Id:       customer.GetId(),
		Username: customer.GetUsername(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, credential)
	tokenString, _ := token.SignedString([]byte("supersecret"))
	return tokenString
}

func (service *AuthServiceImpl) Login(username, password string) (string, error) {
	customer, err := service.AuthRepo.Login(username, password)
	if err != nil {
		return "", err
	}
	token := service.GenerateToken(customer)

	err = service.AuthRepo.SaveToken(token)

	if err != nil {
		return "", err
	}
	//save To History
	var history entity.History
	currentTime := time.Now()
	history.SetId("login-" + username)
	history.SetWhen(currentTime.Format("2006-01-02 15:04:05"))
	history.SetName("Login")
	service.HistoryService.Create(history)

	return token, nil

}
func (service *AuthServiceImpl) Logout(token string) error {
	//save To history
	var history entity.History
	currentTime := time.Now()
	history.SetId("Logout" + token)
	history.SetWhen(currentTime.Format("2006-01-02 15:04:05"))
	history.SetName("Logout Customer")
	service.HistoryService.Create(history)

	return service.AuthRepo.Logout(token)
}
