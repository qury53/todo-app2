package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	_ "github.com/qury53/todo-app2/docs"
	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	core_pgx_pool "github.com/qury53/todo-app2/internal/core/repository/postgres"
	core_http_middleware "github.com/qury53/todo-app2/internal/core/transport/http/middleware"
	transport_http_server "github.com/qury53/todo-app2/internal/core/transport/http/server"
	auth_repostitory "github.com/qury53/todo-app2/internal/features/auth/repostitory"
	auth_service "github.com/qury53/todo-app2/internal/features/auth/service"
	auth_transport_http "github.com/qury53/todo-app2/internal/features/auth/transport/http"
	frontend_repository "github.com/qury53/todo-app2/internal/features/frontend/repository/file_system"
	frontend_service "github.com/qury53/todo-app2/internal/features/frontend/service"
	frontend_transport_http "github.com/qury53/todo-app2/internal/features/frontend/transport/http"
	statistics_repository "github.com/qury53/todo-app2/internal/features/statistics/repository"
	statistics_service "github.com/qury53/todo-app2/internal/features/statistics/service"
	statistics_transport "github.com/qury53/todo-app2/internal/features/statistics/transport/http"
	tasks_repository "github.com/qury53/todo-app2/internal/features/tasks/repository"
	tasks_service "github.com/qury53/todo-app2/internal/features/tasks/service"
	tasks_transport_http "github.com/qury53/todo-app2/internal/features/tasks/transport"
	users_repository "github.com/qury53/todo-app2/internal/features/users/repository"
	users_service "github.com/qury53/todo-app2/internal/features/users/service"
	users_transport_http "github.com/qury53/todo-app2/internal/features/users/transport/http"
)

// @title Golang TODOAPP API
// @version 1.0
// @description TODOAPP REST API scheme
// @contact.name qury53
// @contact.email lamerxdxd@gmail.com
// @BasePath /api/v1
func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGTERM,
	)

	defer cancel()

	//creating logger
	loggerConfig := core_logger.NewConfigMust()
	logger, err := core_logger.NewLogger(loggerConfig)
	if err != nil {
		log.Fatal("failed to create logger", err)
	}
	poolConfig := core_pgx_pool.NewConfigMust()
	pool, err := core_pgx_pool.NewPool(ctx, poolConfig)

	authConfig := auth_service.NewConfigMust()
	authRepository := auth_repostitory.NewAuthRepository(pool)
	authService := auth_service.NewAuthService(authConfig, authRepository)
	authTransport := auth_transport_http.NewAuthTransport(authService)

	//creating and running server
	serverConfig := transport_http_server.NewConfigMust()
	server := transport_http_server.NewHTTPServer(
		serverConfig,
		logger,
		core_http_middleware.CORS(),
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	//creating connection pool
	//poolConfig := core_pgx_pool.NewConfigMust()
	//pool, err := core_pgx_pool.NewPool(ctx, poolConfig)

	logger.Debug("creating feature users")
	//creating feature users
	usersRepository := users_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransport := users_transport_http.NewUserTransport(usersService)

	logger.Debug("creating feature tasks")
	//creating feature tasks
	tasksRepository := tasks_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	tasksTransport := tasks_transport_http.NewTasksTransport(tasksService)
	//creating feature statistics

	logger.Debug("creating feature statistics")
	statisticsRepository := statistics_repository.NewStatisticsRepository(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransport := statistics_transport.NewStatisticsTransport(statisticsService)

	logger.Debug("creating feature frontend")
	frontendRepository := frontend_repository.NewFrontendRepository()
	frontendService := frontend_service.NewFrontendService(frontendRepository)
	frontendTransport := frontend_transport_http.NewFrontendTransport(frontendService)

	//authConfig := auth_service.NewConfigMust()
	log.Println(authConfig)
	logger.Debug("creating feature auth")
	//authRepository := auth_repostitory.NewAuthRepository(pool)
	//authService := auth_service.NewAuthService(authConfig, authRepository)
	//authTransport := auth_transport_http.NewAuthTransport(authService)
	log.Printf("access TTL = %v, refresh TTL = %v", authService.AccessTokenTTL, authService.RefreshTokenTTL)
	// сейчас увидишь: access TTL = 0s, refresh TTL = 0s
	//creating apiVersionRouter
	apiVersionRouter := transport_http_server.NewAPIVersionRouter("v1")

	//registering routes
	apiVersionRouter.RegisterRoutes(usersTransport.Routes()...)
	apiVersionRouter.RegisterRoutes(tasksTransport.Routes()...)
	apiVersionRouter.RegisterRoutes(statisticsTransport.Routes()...)
	apiVersionRouter.RegisterRoutes(authTransport.Routes()...)
	server.RegisterRouters(apiVersionRouter)
	server.RegisterSwagger()
	server.RegisterRoutes(frontendTransport.Routes()...)
	err = server.Run(ctx)
	if err != nil {
		logger.Error("server error", "error", err)
	}
}
