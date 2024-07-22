package middlewares

import (
	"strconv"

	"github.com/ssamsara98/photopost-golang/src/constants"
	"github.com/ssamsara98/photopost-golang/src/lib"

	"github.com/gin-gonic/gin"
)

type PaginationMiddleware struct {
	logger *lib.Logger
}

func NewPaginationMiddleware(logger *lib.Logger) *PaginationMiddleware {
	return &PaginationMiddleware{logger: logger}
}

func (p PaginationMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		p.logger.Debug("setting up pagination middleware")

		limit, err := strconv.ParseInt(c.Query("limit"), 10, 0)
		if err != nil || limit < 5 {
			limit = 10
		}

		page, err := strconv.ParseInt(c.Query("page"), 10, 0)
		if err != nil || limit < 1 {
			page = 1
		}

		c.Set(constants.Limit, limit)
		c.Set(constants.Page, page)

		c.Next()
	}
}

func (p PaginationMiddleware) HandleCursor() gin.HandlerFunc {
	p.logger.Debug("setting up cursor pagination middleware")

	return func(c *gin.Context) {
		limit, err := strconv.ParseInt(c.Query("limit"), 10, 0)
		if err != nil || limit < 5 {
			limit = 10
		}

		cursor, err := strconv.ParseInt(c.Query("cursor"), 10, 0)
		if err != nil {
			cursor = 0
		}

		c.Set(constants.Limit, limit)
		c.Set(constants.Cursor, cursor)

		c.Next()
	}
}
