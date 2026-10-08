package consoletoken

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidAlgorithm = errors.New("invalid jwt signing algorithm")

type Claims struct {
	ServerID  int    `json:"server_id"`
	Addr      string `json:"addr"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	SessionID string `json:"session_id"`

	jwt.RegisteredClaims
}

type Manager struct {
	jwtSecret []byte
	aesgcm    cipher.AEAD
}

func NewManager(
	jwtSecret []byte,
	encryptionKey []byte,
) (*Manager, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &Manager{
		jwtSecret: jwtSecret,
		aesgcm:    aesgcm,
	}, nil
}

func generateJTI() (string, error) {
	b := make([]byte, 16)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func (m *Manager) encrypt(value string) (string, error) {
	nonce := make([]byte, m.aesgcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := m.aesgcm.Seal(
		nonce,
		nonce,
		[]byte(value),
		nil,
	)

	return base64.RawURLEncoding.EncodeToString(ciphertext), nil
}

func (m *Manager) decrypt(value string) (string, error) {
	data, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	nonceSize := m.aesgcm.NonceSize()

	if len(data) < nonceSize {
		return "", errors.New("invalid encrypted value")
	}

	nonce := data[:nonceSize]
	ciphertext := data[nonceSize:]

	plaintext, err := m.aesgcm.Open(
		nil,
		nonce,
		ciphertext,
		nil,
	)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

type UserServer struct {
	ID       int
	Addr     string
	UserName string
	Password string
}

func (m *Manager) Create(
	server UserServer,
	sessionID string,
	expiresAt time.Time,
) (string, error) {
	jti, err := generateJTI()
	if err != nil {
		return "", err
	}

	encryptedPassword, err := m.encrypt(server.Password)
	if err != nil {
		return "", err
	}

	claims := Claims{
		ServerID:  server.ID,
		Addr:      server.Addr,
		UserName:  server.UserName,
		Password:  encryptedPassword,
		SessionID: sessionID,

		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Audience:  jwt.ClaimStrings{"micro-ssh"},
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(m.jwtSecret)
}

func (m *Manager) Parse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, ErrInvalidAlgorithm
			}

			return m.jwtSecret, nil
		},
		jwt.WithAudience("micro-ssh"),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func (m *Manager) Password(claims *Claims) (string, error) {
	return m.decrypt(claims.Password)
}