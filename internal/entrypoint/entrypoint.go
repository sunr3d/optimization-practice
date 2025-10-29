package entrypoint

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wb-go/wbf/ginext"

	"github.com/sunr3d/gc-mem-stats/gcmemstats"

	"github.com/sunr3d/optimization-practice/internal/config"
	httphandlers "github.com/sunr3d/optimization-practice/internal/handlers"
	"github.com/sunr3d/optimization-practice/internal/server"
	"github.com/sunr3d/optimization-practice/internal/services/statssvc"
)

// Run - точка входа в приложение, сборка зависимостей, запуск сервера.
func Run(ctx context.Context, cfg *config.Config) error {
	// Инфраслой (БД и т.д.)
	// Для этого проекта не нужна

	// Сервисный слой (бизнес-логика)
	svc := statssvc.New()

	// Слой представления (ручки)
	engine := httphandlers.New(svc).RegisterHandlers()
	engine.GET("/metrics", gin.WrapH(gcmemstats.MetricsHandler()))
	registerPprof(engine)

	// Сервер
	srv := server.New(":"+cfg.HTTPPort, engine)

	// Запуск сервера
	return srv.Run(ctx)
}

// registerPprof - обертка над регистрацией pprof из gc-mem-stats в gin-engine.
func registerPprof(engine *ginext.Engine) {
	mux := http.NewServeMux()
	gcmemstats.RegisterPprof(mux)

	group := engine.Group("debug/pprof")
	group.Any("/*path", gin.WrapH(mux))
}
