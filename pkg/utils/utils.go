package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/InstayPMS/backend/internal/application/dto"
	"github.com/InstayPMS/backend/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/mssola/useragent"
	"golang.org/x/crypto/bcrypt"
)

func APIResponse(c *gin.Context, status, code int, message string, data any) {
	c.JSON(status, dto.APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func ISEResponse(c *gin.Context) {
	APIResponse(c, http.StatusInternalServerError, constants.CodeInternalError, "Internal server error", nil)
}

func BadRequestResponse(c *gin.Context) {
	APIResponse(c, http.StatusBadRequest, constants.CodeBadRequest, "Invalid data", nil)
}

func OKResponse(c *gin.Context, data any) {
	APIResponse(c, http.StatusOK, constants.CodeSuccess, "Operation successful", data)
}

func GenerateSlug(str string) string {
	return slug.Make(str)
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateRefreshToken() (string, string, error) {
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", "", err
	}

	rawToken := base64.RawURLEncoding.EncodeToString(randomBytes)

	hash := sha256.New()
	hash.Write([]byte(rawToken))
	hashToken := hex.EncodeToString(hash.Sum(nil))

	return rawToken, hashToken, nil
}

func ValidateRefreshToken(rawToken string, storedHash string) bool {
	hash := sha256.New()
	hash.Write([]byte(rawToken))
	computedHash := hex.EncodeToString(hash.Sum(nil))

	return computedHash == storedHash
}

func ConvertUserAgent(uaReq string) string {
	ua := useragent.New(uaReq)
	browser, _ := ua.Browser()
	os := ua.OS()

	if os == "" {
		os = "my computer"
	}

	return fmt.Sprintf("%s on %s", browser, os)
}

func ExtractRootDomain(host string) string {
	if idx := strings.Index(host, ":"); idx != -1 {
		host = host[:idx]
	}

	if host == "localhost" || !strings.Contains(host, ".") {
		return host
	}

	parts := strings.Split(host, ".")
	if len(parts) == 4 {
		isIP := true
		for _, part := range parts {
			for _, ch := range part {
				if ch < '0' || ch > '9' {
					isIP = false
					break
				}
			}
		}
		if isIP {
			return host
		}
	}

	if len(parts) >= 2 {
		return "." + parts[len(parts)-2] + "." + parts[len(parts)-1]
	}

	return host
}
