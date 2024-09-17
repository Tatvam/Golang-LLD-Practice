package errors

type Error struct {
	statusCode   int
	errorCode    string
	errorMessage string
}

func NewError(statusCode int, errorCode, errorMessage string) *Error {
	return &Error{
		statusCode:   statusCode,
		errorCode:    errorCode,
		errorMessage: errorMessage,
	}
}

var (
	ErrGeneratingTicketFailure = &Error{
		statusCode:   500,
		errorMessage: "some error occured while generating ticket.",
		errorCode:    "PL-001-F",
	}

	ErrGeneratingTicketSuccess = &Error{
		statusCode:   200,
		errorMessage: "ticket generated successfully",
		errorCode:    "PL-001-S",
	}
	ErrNoTicketFound = &Error{
		statusCode:   500,
		errorMessage: "no ticket found",
		errorCode:    "PL-002-F",
	}
	ErrGeneratingBillFailure = &Error{
		statusCode:   500,
		errorMessage: "error generating bill",
		errorCode:    "PL-003-F",
	}
	ErrGeneratingBillSuccess = &Error{
		statusCode:   200,
		errorMessage: "bill generated successfully",
		errorCode:    "PL-003-S",
	}
	ErrParkingLotFull = &Error{
		statusCode:   400,
		errorMessage: "parking lot is full",
		errorCode:    "PL-004-F",
	}
)
