// реализация интерфейса
package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Wolkodaf2946/todolist"
	"github.com/Wolkodaf2946/todolist/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "rdxcuahgsdbvihwtoivuwhbrovhnwlighblfkvwodfvbwfpv"
	signingKey = "qrkjk#4#5FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

// tokenClaims определяет вашу структуру утверждений для JWT.
// Теперь мы встраиваем jwt.RegisteredClaims вместо jwt.StandardClaims.
type tokenClaims struct {
	jwt.RegisteredClaims     // Используем RegisteredClaims для стандартных утверждений JWT
	UserId               int `json:"user_id"`
}

type AuthService struct { // структура, которую в конструкторе будет принимать репозиторий для работы с базой
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}

}

func (s *AuthService) CreateUser(user todolist.User) (int, error) { // метод, передающий структуру User на уровень ниже (repository)
	user.Password = generatePasswordHash(user.Password) // сначала хэшируем пароль, а только потом передаём структуру в слой репозитория

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// Получаем пользователя из репозитория, используя хэш пароля.
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		// Важно возвращать более специфичные ошибки, например, "пользователь не найден" или "неверный пароль"
		// Сейчас это просто общая ошибка из репозитория.
		return "", fmt.Errorf("Error getting user: %w", err)
	}

	// Вычисляем время истечения срока действия токена.
	expirationTime := time.Now().Add(tokenTTL)

	// Создаем объект утверждений.
	claims := tokenClaims{
		// Инициализируем RegisteredClaims с помощью jwt.NewNumericDate() для полей времени.
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Токен перестанет быть валидным через tokenTTL
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // Время выдачи токена
		},
		UserId: user.Id, // Приватное утверждение: ID пользователя
	}

	// Создаем новый JWT с указанными методом подписи и утверждениями.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с помощью секретного ключа.
	signedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", fmt.Errorf("token signature error: %w", err)
	}

	return signedToken, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
