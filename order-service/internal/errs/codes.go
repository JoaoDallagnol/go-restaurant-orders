package errs

type ErrorCode string

const (
	CodeOrderNotFound               ErrorCode = "ORDER_NOT_FOUND"
	CodeOrderItemNotFound           ErrorCode = "ORDER_ITEM_NOT_FOUND"
	CodeInternalError               ErrorCode = "INTERNAL_ERROR"
	CodeMenuServiceIntegrationError ErrorCode = "MENU_SERVICE_INTEGRATION_ERROR"
)
