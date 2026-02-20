package routes

import (
	Handler "sociomile-apps/internal/handlers"
	repositories "sociomile-apps/internal/repositories"
	"sociomile-apps/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()

	authService := services.NewAuthService(db)
	authHandler := Handler.NewAuthHandler(authService)

	channelService := services.NewChannelService(
		repositories.NewConversationRepository(db),
		repositories.NewMessageRepository(db),
	)

	channelHandler := Handler.NewChannelHandler(channelService)

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", authHandler.Logout)
		}

		channel := api.Group("/channel")
		{
			channel.POST("/webhook", channelHandler.Webhook)
		}
	}
	return router
}
