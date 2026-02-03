package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"tasker_go/internal/analysis"
	"tasker_go/internal/config"
	"tasker_go/internal/llm/gemini"
	"tasker_go/internal/repository/gorm"
	"tasker_go/internal/service"
	"tasker_go/internal/transport/http/handler"
	"tasker_go/internal/transport/http/router"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func Run() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Println(".env не найден, используем системные переменные")
	}

	db := config.ConnectDB()
	taskRepo := gorm.NewTaskRepository(db)
	userRepo := gorm.NewUserRepository(db)
	judgeRepo := gorm.NewJudgeRepository(db)

	taskService := service.NewTaskService(taskRepo)
	authService := service.NewAuthService(userRepo)

	llm, err := gemini.NewClient(ctx, "")
	if err != nil {
		log.Fatal(err)
	}
	judgeAnalyzer := analysis.NewJudgeAnalyzer(llm)
	judgeService := service.NewJudgeService(taskRepo, judgeRepo, judgeAnalyzer)

	taskHandler := handler.NewTaskHandler(taskService)
	authHandler := handler.NewAuthHandler(authService)
	judgeHandler := handler.NewJudgeHandler(judgeService)

	r := router.New(taskHandler, authHandler, judgeHandler)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:*",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
		},
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = http.ListenAndServe(":"+port, c.Handler(r))
	if err != nil {
		log.Printf("Server error: %v", err)
	}

}
