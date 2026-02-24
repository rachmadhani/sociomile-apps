package routes

import (
	"sociomile-apps/internal/cache"
	"sociomile-apps/internal/event"
	Handler "sociomile-apps/internal/handlers"
	"sociomile-apps/internal/middleware"
	repositories "sociomile-apps/internal/repositories"
	"sociomile-apps/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, dispatcher *event.Dispatcher, conversationCache *cache.ConversationCache, ticketCache *cache.TicketCache) *gin.Engine {
	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"}
	router.Use(cors.New(corsConfig))

	authService := services.NewAuthService(db)
	authHandler := Handler.NewAuthHandler(authService)

	channelService := services.NewChannelService(
		repositories.NewConversationRepository(db),
		repositories.NewMessageRepository(db),
	)

	channelHandler := Handler.NewChannelHandler(channelService)

	conversationService := services.NewConversationService(
		repositories.NewConversationRepository(db),
		repositories.NewMessageRepository(db),
		dispatcher,
	)

	conversationHandler := Handler.NewConversationHandler(conversationService)

	conversationQueryService := services.NewConversationQueryService(
		repositories.NewConversationRepository(db),
		conversationCache,
	)

	conversationQueryHandler := Handler.NewConversationQueryHandler(conversationQueryService)

	ticketService := services.NewTicketService(
		repositories.NewTicketRepository(db),
		repositories.NewConversationRepository(db),
		dispatcher,
		ticketCache,
	)

	ticketHandler := Handler.NewTicketHandler(ticketService)

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/list-agent", authHandler.GetListAgent)
			auth.POST("/logout", authHandler.Logout)
		}

		channel := api.Group("/channel")
		{
			channel.POST("/webhook", channelHandler.Webhook)
		}

		conversation := api.Group("/conversation")
		{
			conversation.POST("/:id/agent-reply", middleware.AuthMiddleware(), middleware.RequireRole("agent"), conversationHandler.AgentReply)
			conversation.GET("/:id", middleware.AuthMiddleware(), middleware.RequireRole("agent", "admin"), conversationQueryHandler.Detail)
			conversation.GET("/", middleware.AuthMiddleware(), middleware.RequireRole("agent", "admin"), conversationQueryHandler.List)

			conversation.POST("/:id/escalate", middleware.AuthMiddleware(), middleware.RequireRole("agent"), ticketHandler.EscalateTicket)
		}

		ticket := api.Group("/ticket")
		{
			ticket.GET("/", middleware.AuthMiddleware(), middleware.RequireRole("agent", "admin"), ticketHandler.List)
			ticket.POST("/:id/update-status", middleware.AuthMiddleware(), middleware.RequireRole("admin"), ticketHandler.UpdateStatus)
		}

	}
	return router
}
