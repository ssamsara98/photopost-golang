package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssamsara98/photopost-golang/src/constants"
	"github.com/ssamsara98/photopost-golang/src/infrastructure"
	"github.com/ssamsara98/photopost-golang/src/lib"
	"github.com/ssamsara98/photopost-golang/src/utils"
)

// DBTransactionMiddleware -> struct for transaction
type DBTransactionMiddleware struct {
	logger *lib.Logger
	db     *infrastructure.Database
}

// NewDBTransactionMiddleware -> new instance of transaction
func NewDBTransactionMiddleware(
	logger *lib.Logger,
	db *infrastructure.Database,
) *DBTransactionMiddleware {
	return &DBTransactionMiddleware{
		logger: logger,
		db:     db,
	}
}

// Handle -> It setup the database transaction middleware
func (m DBTransactionMiddleware) Handle() gin.HandlerFunc {
	m.logger.Debug("setting up database transaction middleware")

	return func(c *gin.Context) {
		txHandle := m.db.DB.Begin()
		m.logger.Debug("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set(constants.DBTransaction, txHandle)
		c.Next()

		if utils.StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			m.logger.Debug("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				m.logger.Error("trx commit error: ", err)
			}
		} else {
			m.logger.Debug("rolling back transaction due to status code: ", c.Writer.Status())
			txHandle.Rollback()
		}
	}
}
