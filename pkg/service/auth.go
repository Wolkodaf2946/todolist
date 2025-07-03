// реализация интерфейса
package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Wolkodaf2946/todolist"
	"github.com/Wolkodaf2946/todolist/pkg/repository"
)

const salt = "rdxcuahgsdbvihwtoivuwhbrovhnwlighblfkvwodfvbwfpv"

type AuthService struct { // структура, которую в конструкторе будет принимать репозиторий для работы с базой
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}

}

func (s *AuthService) CreateUser(user todolist.User) (int, error) { // метод, передающий структуру User на уровень ниже (repository)
	user.Password = s.generatePasswordHash(user.Password) // сначала хэшируем пароль, а только потом передаём структуру в слой репозитория

	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
