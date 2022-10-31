package rest

const (
	MsgUserCreated        = "User created successfully"
	MsgEmailConflict      = "Email is already taken"
	MsgInternalSeverErr   = "Internal server error"
	MsgBadRequest         = "Bad request"
	MsgNotFound           = "Not found"
	MsgTimeOut            = "Connection timeout"
	MSgBadURL             = "Error while parsing URL"
	handlerTimeoutSeconds = 30
	readTimeoutSeconds    = 2
	writeTimeoutSeconds   = 5
	statusValAccepted     = "accepted"
	statusValDeclined     = "declined"
	payValCash            = "cash"
	payValCard            = "card"
	keyTaID               = "transaction_id"
	keyTmID               = "terminal_id"
	keyStatus             = "status"
	keyPayType            = "payment_type"
	keyDatePost           = "date_post"
	keyPayNar             = "payment_narrative"
	timeLayout            = "2006-01-02 15:04:05"
	dateLayout            = "2006-01-02"
	payNarMinLen          = 4
)

// const (
// 	offset    = "offset"
// 	limit     = "limit"
// 	sort      = "sort"
// 	firstName = "first_name"
// 	lastName  = "last_name"
// 	createdAt = "created_at"
// )
