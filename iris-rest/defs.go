package main

type irisTestResponse struct {
	Status string `json:"status"`
}

type BookCreateRequest struct {
	Name string `json:"name"`
}

type BookUpdateRequest struct {
	Name string `json:"name"`
}

type BookCreateSuccessResponse struct {
	Status string `json:"status"`
	Name   string `json:"name"`
	Id     int    `json:"id"`
}

type BookCreateFailureResponse struct {
	Status   string `json:"status"`
	ErrorMsg string `json:"error_msg"`
}

type BooksListResponse struct {
	Status string `json:"status"`
	Books  []Book `json:"books"`
}

type BookItemSuccessResponse struct {
	Status string `json:"status"`
	Item   Book   `json:"book"`
}

type BookItemFailureResponse struct {
	Status   string `json:"status"`
	ErrorMsg string `json:"error_msg"`
}

type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
