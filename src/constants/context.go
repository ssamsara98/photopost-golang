package constants

const (
	// DBTransaction is database transaction handle set at router context
	DBTransaction = "db_trx"

	// Claims -> authentication claims
	Claims = "Claims"

	// UID -> authenticated user's id
	UID = "UID"

	// File uploaded file from file upload middleware
	File = "@uploaded_file"

	// Limit for get all api
	Limit = "Limit"
	// Page
	Page = "Page"
	// Page
	Cursor = "Cursor"

	// Rate Limit
	RateLimit = "RateLimit"

	// Token -> bearer token
	Token        = "Token"
	TokenPrefix  = "Bearer"
	TokenAccess  = "Access"
	TokenRefresh = "Refresh"

	// User -> user
	User = "User"
)
