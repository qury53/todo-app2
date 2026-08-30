package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	core_pgx_pool "github.com/qury53/todo-app2/internal/core/repository/postgres"
	core_http_middleware "github.com/qury53/todo-app2/internal/core/transport/http/middleware"
	transport_http_server "github.com/qury53/todo-app2/internal/core/transport/http/server"
	tasks_repository "github.com/qury53/todo-app2/internal/features/tasks/repository"
	tasks_service "github.com/qury53/todo-app2/internal/features/tasks/service"
	tasks_transport_http "github.com/qury53/todo-app2/internal/features/tasks/transport"
	users_repository "github.com/qury53/todo-app2/internal/features/users/repository"
	users_service "github.com/qury53/todo-app2/internal/features/users/service"
	users_transport_http "github.com/qury53/todo-app2/internal/features/users/transport/http"
)

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

	//creating and running server
	serverConfig := transport_http_server.NewConfigMust()
	server := transport_http_server.NewHTTPServer(
		serverConfig,
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	//creating connection pool
	poolConfig := core_pgx_pool.NewConfigMust()
	pool, err := core_pgx_pool.NewPool(ctx, poolConfig)

	//creating feature users
	usersRepository := users_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransport := users_transport_http.NewUserTransport(usersService)

	//creating feature tasks
	tasksRepository := tasks_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	tasksTransport := tasks_transport_http.NewTasksTransport(tasksService)

	//creating apiVersionRouter
	apiVersionRouter := transport_http_server.NewAPIVersionRouter("v1")

	//registering routes
	apiVersionRouter.RegisterRoutes(usersTransport.Routes()...)
	apiVersionRouter.RegisterRoutes(tasksTransport.Routes()...)
	server.RegisterRouters(apiVersionRouter)
	err = server.Run(ctx)
	if err != nil {
		logger.Error("server error", "error", err)
	}
}
