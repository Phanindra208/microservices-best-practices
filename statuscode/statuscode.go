package statuscode

const (
	//SuccessCodeOK ... Success
	SuccessCodeOK = 200
	//SuccessMultipleCandidates ... Success with multiple candidates
	SuccessMultipleCandidates = 300
	//ErrorCodeUnauthorized ... Unauthorized
	ErrorCodeUnauthorized = 401
	//ErrorCodeBadRequest ... Bad request
	ErrorCodeBadRequest = 400
	//ErrorCodeItemNotFound ... Item not found
	ErrorCodeItemNotFound = 404
	//ErrorCodeOutOfStock ... Item outofstock
	ErrorCodeOutOfStock = 1000
	//ErrorMaxToOrderExceeded ... Item Max to order Exceeded
	ErrorMaxToOrderExceeded = 1001
	//ErrorServiceNotAvailable ... Service not available
	ErrorServiceNotAvailable = 1002
	//ErrorCodeItemRemoved ... Item removed
	ErrorCodeItemRemoved = 1003
	//ErrorCodeUnitNotAllowed ... Unit Not allowed
	ErrorCodeUnitNotAllowed = 1004
	// ErrorCodeEmptyList is used when a list is empty.
	ErrorCodeEmptyList = 1005

	//SuccessIncreaseProductAllowed ... Increase product success
	SuccessIncreaseProductAllowed = 2000
	//ErrorCodeIncreaseMoreThanAllowed ... Increase request more than what is allowed for product
	ErrorCodeIncreaseMoreThanAllowed = 2001
	//ErrorCodeIncreaseUnitNotAllowed ... Increase request unit not allowed
	ErrorCodeIncreaseUnitNotAllowed = 2002
	//ErrorCodeIncreaseGramsToKG ... Increase request grams to kg
	ErrorCodeIncreaseGramsToKG = 2003

	//SuccessDecreaseProductAllowed ... Decrease product success
	SuccessDecreaseProductAllowed = 3000
	//ErrorCodeDecreaseMoreThanAllowed ... Decrease request more than what is allowed for product
	ErrorCodeDecreaseMoreThanAllowed = 3001
	//ErrorCodeDecreaseUnitNotAllowed ... Decrease request quantity unit is not allowed
	ErrorCodeDecreaseUnitNotAllowed = 3002
	//ErrorCodeDecreaseGramsToKG ... Decrease request grams to kg
	ErrorCodeDecreaseGramsToKG = 3003

	// ErrorCodeInvalidArea is returned when the area is invalid.
	ErrorCodeInvalidArea = 4001

	// ErrorCodeDeliverySlotExpired is returned when the delivery slot expires.
	ErrorCodeDeliverySlotExpired = 5001
	// ErrorCodeInvalidDeliverySlot is returned when the delivery slot is invalid.
	ErrorCodeInvalidDeliverySlot = 5002
	// ErrorCodeInvalidCartCannotProcess is returned when the cart is invalid and cannot create an order.
	ErrorCodeInvalidCartCannotProcess = 5003
	// ErrorCodeOrderItemOutOfStock is returned when an item is out of stock.
	ErrorCodeOrderItemOutOfStock = 5004
	// ErrorCodeInvalidAmountOfPoints is returned when invalid amount of points are attempted to be redeemed.
	ErrorCodeInvalidAmountOfPoints = 5005

	// ErrorCodePointsExceedCartTotal is returned when used points exceed total cart amount.
	ErrorCodePointsExceedCartTotal = 5101
	// ErrorCodePointsExceedAvailable is returned when used points exceed user's available points.
	ErrorCodePointsExceedAvailable = 5102
)
