package models

// Book represents the book structure used in the API.
type Book struct {
    ID       string  `json:"id"`
    Title    string  `json:"title"`
    Author   string  `json:"author"`
    Quantity int32   `json:"quantity"`
    Price    float64 `json:"price"`
}

// CreateBookRequest represents the request payload for creating a book.
type CreateBookRequest struct {
    Title    string  `json:"title"`
    Author   string  `json:"author"`
    Quantity int32   `json:"quantity"`
    Price    float64 `json:"price"`
}

// CreateBookResponse represents the response for the create book request.
type CreateBookResponse struct {
    Message string `json:"message"`
}

// GetBookRequest represents the request payload for getting a book.
type GetBookRequest struct {
    ID string `json:"id"`
}

// GetBookResponse represents the response for the get book request.
type GetBookResponse struct {
    Book Book `json:"book"`
}

// UpdateBookRequest represents the request payload for updating a book.
type UpdateBookRequest struct {
    ID       string  `json:"id"`
    Title    string  `json:"title"`
    Author   string  `json:"author"`
    Quantity int32   `json:"quantity"`
    Price    float64 `json:"price"`
}

// UpdateBookResponse represents the response for the update book request.
type UpdateBookResponse struct {
    Message string `json:"message"`
}

// DeleteBookRequest represents the request payload for deleting a book.
type DeleteBookRequest struct {
    ID string `json:"id"`
}

// DeleteBookResponse represents the response for the delete book request.
type DeleteBookResponse struct {
    Message string `json:"message"`
}

// BuyBookRequest represents the request payload for buying a book.
type BuyBookRequest struct {
    ID       string `json:"id"`
    Quantity int32  `json:"quantity"`
}

// BuyBookResponse represents the response for the buy book request.
type BuyBookResponse struct {
    Message string `json:"message"`
}

// SellBookRequest represents the request payload for selling a book.
type SellBookRequest struct {
    ID       string `json:"id"`
    Quantity int32  `json:"quantity"`
}

// SellBookResponse represents the response for the sell book request.
type SellBookResponse struct {
    Message string `json:"message"`
}
