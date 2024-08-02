package infrastructure

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	lm "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ssamsara98/photopost-golang/src/lib"
	"github.com/ssamsara98/photopost-golang/src/utils"
)

const idleTimeout = 5 * time.Second

// Router -> Fiber Router
type Router struct {
	*fiber.App
}

// NewRouter : all the routes are defined here
func NewRouter(
	env *lib.Env,
	logger *lib.Logger,
) *Router {

	app := fiber.New(fiber.Config{
		IdleTimeout:  idleTimeout,
		ErrorHandler: fiberErrorHandler,
	})

	/* MaxMultipartMemory */

	app.Use(idempotency.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 10 * time.Second,
	}))
	app.Use(lm.New(lm.Config{
		Output: logger.GetFiberLogger(),
		Format: fmt.Sprintf(
			"pid:${%s} | ${%s} | ${%s} | ${%s} | ${%s} | ${%s} | ${%s}",
			lm.TagPid, lm.TagStatus, lm.TagLatency, lm.TagIP, lm.TagMethod, lm.TagPath, lm.TagError,
		),
		DisableColors: true,
	}))
	app.Use(etag.New())
	app.Use(healthcheck.New())

	router := &Router{app}
	return router
}

func fiberErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return utils.ErrorJSON(c, err, code)
}
