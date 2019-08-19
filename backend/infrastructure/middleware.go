package infrastructure

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"treasure-app/backend/domain"
	"treasure-app/backend/interfaces/database"
	"treasure-app/backend/usecase"

	gin "github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
)

type Auth struct {
	// client     *auth.Client
	Interactor usecase.UserInteractor
}

func NewAuth(sqlHandler database.SqlHandler) *Auth {
	return &Auth{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (auth *Auth) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Firebase SDK のセットアップ
		// opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
		app, err := firebase.NewApp(context.Background(), nil) //, opt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			c.JSON(500, err)
			return
		}
		client, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			c.JSON(500, err)
			return
		}

		// end init

		// クライアントから送られてきた JWT 取得
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT の検証
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			// JWT が無効なら Handler に進まず別処理
			fmt.Printf("error verifying ID token: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		// TODO: 後で消す
		// Set admin privilege on the user corresponding to uid.
		claims := map[string]interface{}{"admin": true}
		err = client.SetCustomUserClaims(context.Background(), token.UID, claims)
		if err != nil {
			log.Fatalf("error setting custom claims %v\n", err)
		}

		claims = token.Claims
		if admin, ok := claims["admin"]; ok {
			if admin.(bool) {
				log.Printf("ts")
				//Allow access to requested admin resource.
			}
		}

		userRecord, err := client.GetUser(context.Background(), token.UID)

		if err != nil {
			fmt.Printf("error: %v\n", err)
			c.JSON(500, err)
			return
		}

		u := domain.User{
			FirebaseUID: userRecord.UID,
			Name:        userRecord.DisplayName,
			Email:       userRecord.Email,
			Image:       userRecord.PhotoURL,
		}

		user, err := auth.Interactor.UserByFirebaseId(u.FirebaseUID)
		if err != nil {
			c.JSON(500, err)
			return
		}

		err = auth.Interactor.Update(u)
		if err != nil {
			c.JSON(500, err)
			return
		}

		c.Set("AuthorizedUser", user)

		log.Printf("Verified ID token: %v\n", token)
	}
}
