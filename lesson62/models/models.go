package models

type Book struct {
    ID       string  `json:"id"`
    Title    string  `json:"title"`
    Author   string  `json:"author"`
    Quantity int32   `json:"quantity"`
    Price    float64 `json:"price"`
}

type CreateBookRequest struct {
    Title    string  `json:"title"`
    Author   string  `json:"author"`
    Quantity int32   `json:"quantity"`
    Price    float64 `json:"price"`
}

type CreateBookResponse struct {
    Message string `json:"message"`
}

type GetBookRequest struct {
    ID string `json:"id"`
}

type GetBookResponse struct {
    Book Book `json:"book"`
}

type UpdateBookRequest struct {
    ID       string  `json:"id"`
    Title    string  `json:"title"`
    Author   string  `json:"author"`
    Quantity int32   `json:"quantity"`
    Price    float64 `json:"price"`
}

type UpdateBookResponse struct {
    Message string `json:"message"`
}

type DeleteBookRequest struct {
    ID string `json:"id"`
}

type DeleteBookResponse struct {
    Message string `json:"message"`
}

type BuyBookRequest struct {
    ID       string `json:"id"`
    Quantity int32  `json:"quantity"`
}

type BuyBookResponse struct {
    Message string `json:"message"`
}

type SellBookRequest struct {
    ID       string `json:"id"`
    Quantity int32  `json:"quantity"`
}

type SellBookResponse struct {
    Message string `json:"message"`
}
