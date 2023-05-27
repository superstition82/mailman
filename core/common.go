package core

type response struct {
	Data any `json:"data"`
}

func composeResponse(data any) response {
	return response{
		Data: data,
	}
}
