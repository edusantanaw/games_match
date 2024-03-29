package utils

type HttpResponse struct {
	StatusCode int
	Body       any
}

func Created[T comparable](data T) HttpResponse {
	httpResponse := HttpResponse{StatusCode: 201, Body: data}
	return httpResponse
}

func BadRequest[T comparable](data T) HttpResponse {
	httpResponse := HttpResponse{StatusCode: 400, Body: data}
	return httpResponse
}

func Ok[T comparable](data T) HttpResponse {
	httpResponse := HttpResponse{StatusCode: 200, Body: data}
	return httpResponse
}
