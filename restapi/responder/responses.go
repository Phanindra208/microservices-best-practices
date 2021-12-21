package responder

// OK sets response to be an "OK" response with no errors.
func (resp *Responder) OK(body interface{}) *Responder {
	resp.status = 200
	resp.body = body
	return resp
}

// Created sets response to be a "created" response with no errors.
func (resp *Responder) Created(body interface{}) *Responder {
	resp.status = 201
	resp.body = body
	return resp
}

// NoContent sets response to be a "no content" response with no errors.
func (resp *Responder) NoContent() *Responder {
	resp.status = 204
	return resp
}
