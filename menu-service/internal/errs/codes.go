package errs

type ErrorCode string

const (
	CodeRestaurantNotFound ErrorCode = "RESTAURANT_NOT_FOUND"
	CodeDishNotFound       ErrorCode = "DISH_NOT_FOUND"
	CodeInternalError      ErrorCode = "INTERNAL_ERROR"
)
