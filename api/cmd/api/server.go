package api

import (
	"PopcornMovie/cmd/middleware"
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/gateway/payment"
	"PopcornMovie/gateway/rest"
	"PopcornMovie/resolver"
	"PopcornMovie/service"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

func NewServerCmd(configs *config.Configurations, logger *zap.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "run api server",
		Long:  "run api server with graphql",
		Run: func(cmd *cobra.Command, args []string) {
			// connect db to postgres database
			db, err := ent.Open("postgres", configs.Postgres.ConnectionString, ent.Debug())
			if err != nil {
				logger.Error("Getting error connect to postgresql database", zap.Error(err))
				os.Exit(1)
			}
			// Create validator
			validator := validator.New()
			// Add translator for validator
			en := en.New()
			uni := ut.New(en, en)
			validationTranslator, _ := uni.GetTranslator("en")
			// Register default translation for validator
			err = en_translations.RegisterDefaultTranslations(validator, validationTranslator)
			if err != nil {
				logger.Error("Getting error from register default translation", zap.Error(err))
				os.Exit(1)
			}

			err = payment.NewPaymentService(configs.Payos)
			if err != nil {
				logger.Error("Getting error from create payment service", zap.Error(err))
				os.Exit(1)
			}
			// GraphQL schema resolver handler.
			service := service.New(db, logger, *configs)
			resolverHandler := handler.NewDefaultServer(resolver.NewExecutableSchema(service, logger))
			restCall := rest.New(service)
			// Create a Gin router instance
			app := gin.Default()

			// middleware
			app.Use(
				middleware.CorsMiddleware(),
				middleware.RequestCtxMiddleware(),
				middleware.AuthMiddleware(configs.AppConfig.JWTSecret),
			)

			// Define the GraphQL endpoint
			app.POST("/query", gin.WrapH(resolverHandler))

			app.POST("/verify-hook", restCall.WebHooktype)
			// Define the GraphQL Playground endpoint
			app.GET("/", func(c *gin.Context) {
				playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
			})

			// Listen on port 8000
			logger.Info("Listening on port: 8000")
			if err := app.Run(":8000"); err != nil {
				logger.Error("Get error from run server", zap.Error(err))
			}
		},
	}
}
