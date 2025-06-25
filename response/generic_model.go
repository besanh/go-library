package response

type PaginationBodyResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Limit   int64  `json:"limit"`
	Offset  int64  `json:"offset"`
	Total   int64  `json:"total"`
}

type PaginationResponse[T any] struct {
	Body PaginationBodyResponse[T]
}

type GenericResponse[T any] struct {
	// Status int
	Body BodyResponse[T]
}

type BodyResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type MediaResponse struct {
	ContentType             string `header:"Content-Type"`
	ContentLength           int    `header:"Content-Length"`
	AcceptRanges            string `header:"Accept-Ranges"`
	AllowControlAllowOrigin string `header:"Access-Control-Allow-Origin"`
	CacheControl            string `header:"Cache-Control"`
	Body                    []byte
}

type IdResponse struct {
	Id string `json:"id"`
}
