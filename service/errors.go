package service

var (
	// user api errors

	ErrInvalidEmail             = &Error{Code: -1000, Message: "invalid email"}
	ErrInvalidPassword          = &Error{Code: -1001, Message: "invalid password"}
	ErrInvalidUserType          = &Error{Code: -1002, Message: "invalid user type"}
	ErrPasswordMismatch         = &Error{Code: -1003, Message: "password and confirm password must match"}
	ErrEmailNotExists           = &Error{Code: -1004, Message: "email doesn't exist"}
	ErrEmailAlreadyExists       = &Error{Code: -1005, Message: "email already exists"}
	ErrInvalidCredentials       = &Error{Code: -1006, Message: "invalid credentials"}
	ErrInvalidVerificationToken = &Error{Code: -1007, Message: "invalid verification token"}
	ErrInactiveAccount          = &Error{Code: -1008, Message: "your account is inactive"}
	ErrMissingPubKey            = &Error{Code: -1009, Message: "public key is required for lender user"}

	// portal api errors

	ErrBorrowNotFound     = &Error{Code: -2000, Message: "borrow not found"}
	ErrInvalidBorrowState = &Error{Code: -2001, Message: "invalid borrow state"}

	// exchange api errors

	ErrInvalidOrderType = &Error{Code: -3000, Message: "invalid order type"}
	ErrInvalidOrderSide = &Error{Code: -3001, Message: "invalid order side"}
	ErrInvalidSymbol    = &Error{Code: -3002, Message: "invalid symbol"}

	// general api errors

	ErrInvalidArgument     = &Error{Code: -9000, Message: "invalid argument"}
	ErrInternalServerError = &Error{Code: -9001, Message: "internal server error"}
	ErrInvalidLimit        = &Error{Code: -9002, Message: "invalid limit"}
	ErrInvalidPage         = &Error{Code: -9003, Message: "invalid page"}
)

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return e.Message
}
