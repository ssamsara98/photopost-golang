package middlewares

import (
	"time"

	"github.com/ssamsara98/photopost-golang/src/constants"
	"github.com/ssamsara98/photopost-golang/src/lib"
)

// // Global store
// // using in-memory store with goroutine which clears expired keys.
// var store = memory.NewStore()

type RateLimitOption struct {
	period time.Duration
	limit  int64
}

type Option func(*RateLimitOption)

type RateLimitMiddleware struct {
	logger *lib.Logger
	option RateLimitOption
}

func NewRateLimitMiddleware(logger *lib.Logger) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		logger: logger,
		option: RateLimitOption{
			period: constants.RateLimitPeriod,
			limit:  constants.RateLimitRequests,
		},
	}
}

// func (lm RateLimitMiddleware) Handle(options ...Option) gin.HandlerFunc {
// 	// func (lm RateLimitMiddleware) Handle(options ...Option) fiber.Handler {
// 	lm.logger.Debug("setting up rate limit middleware")

// 	return func(c *gin.Context) {
// 		// return func(c *fiber.Ctx) (err error) {
// 		key := c.ClientIP() // Gets cient IP Address
// 		// key := c.IP() // Gets cient IP Address

// 		// Setting up rate limit
// 		// Limit -> # of API Calls
// 		// Period -> in a given time frame
// 		// setting default values
// 		opt := RateLimitOption{
// 			period: lm.option.period,
// 			limit:  lm.option.limit,
// 		}

// 		for _, o := range options {
// 			o(&opt)
// 		}

// 		rate := limiter.Rate{
// 			Limit:  opt.limit,
// 			Period: opt.period,
// 		}

// 		// Limiter instance
// 		instance := limiter.New(store, rate)

// 		// Returns the rate limit details for given identifier.
// 		// FullPath is appended with IP address. `/api/users&&127.0.0.1` as key
// 		context, err := instance.Get(c, c.FullPath()+"&&"+key)
// 		if err != nil {
// 			lm.logger.Panic(err.Error())
// 		}

// 		c.Set(constants.RateLimit, instance)

// 		// Setting custom headers
// 		c.Header("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
// 		c.Header("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
// 		c.Header("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))

// 		// Limit exceeded
// 		if context.Reached {
// 			errNew := errors.New("rate limit exceed")
// 			utils.ErrorJSON(c, http.StatusTooManyRequests, errNew)
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }

// func WithOptions(period time.Duration, limit int64) Option {
// 	return func(o *RateLimitOption) {
// 		o.period = period
// 		o.limit = limit
// 	}
// }
