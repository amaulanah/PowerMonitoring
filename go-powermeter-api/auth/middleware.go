package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah fungsi middleware untuk memvalidasi token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header dibutuhkan"})
			return
		}

		// Token biasanya dikirim dengan format "Bearer <token>"
		// Kita perlu memisahkan "Bearer" dari tokennya
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Format Authorization header salah"})
			return
		}

		tokenString := parts[1]

		claims, err := ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			return
		}

		// Simpan username di dalam konteks Gin untuk digunakan di handler selanjutnya
		c.Set("username", claims.Username)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}
