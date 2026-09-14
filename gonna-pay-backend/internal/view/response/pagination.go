package response

type PaginatedResponse[T any] struct {
	Data       []T   `json:"data"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"totalPages"`
}

func NewPaginatedResponse[T any](data []T, page, limit int, total int64) PaginatedResponse[T] {
	totalPages := int(total) / limit

	if int(total)%limit != 0 {
		totalPages++
	}

	return PaginatedResponse[T]{
		Data: data,
		Page: page,
		Limit: limit,
		Total: total,
		TotalPages: totalPages,
	}
}
