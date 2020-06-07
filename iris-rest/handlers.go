package main

import (
	"fmt"
	"strconv"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/mgo.v2/bson"
)

func irisTest(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, irisTestResponse{SUCCESS})
}

func getBooks(ctx *iris.Context) {
	books := make([]Book, 128)
	booksTableConnector.Find(bson.M{}).Iter().All(&books)
	ctx.JSON(iris.StatusOK, BooksListResponse{Status: SUCCESS, Books: books})
}

func createBook(ctx *iris.Context) {
	request := &BookCreateRequest{}
	if err := ctx.ReadJSON(request); err != nil {
		ctx.JSON(iris.StatusBadRequest, fmt.Sprintf("INVALID-INPUT-JSON error: %v", err))
		return
	}

	errorMsg, flag := validate(request.Name)
	if !flag {
		ctx.JSON(iris.StatusBadRequest, BookCreateFailureResponse{Status: FAILURE, ErrorMsg: errorMsg})
		return
	}

	iter := cassandraSession.Query(`SELECT * FROM booksid`).Iter()
	row := make(map[string]interface{})
	for {
		if !iter.MapScan(row) {
			break
		}
	}

	fmt.Println("Current Id value : ", row["value"])
	new_id := row["value"].(int)
	new_id += 1
	query := "update booksid set  value=" + strconv.Itoa(new_id) + " where id=1"

	err := cassandraSession.Query(query).Exec()
	if err != nil {
		errorMsg = "Cannot Update the index in Cassandra "
		fmt.Println("Cannot Update the index in Cassandra ", err)
		ctx.JSON(iris.StatusBadRequest, BookCreateFailureResponse{Status: FAILURE, ErrorMsg: errorMsg})
		return
	}

	err = booksTableConnector.Insert(Book{Name: request.Name, Id: new_id})
	if err != nil {
		errorMsg = "Cannot insert the record in mongo"
		fmt.Println("Cannot insert the record in mongo : ", err)
		ctx.JSON(iris.StatusBadRequest, BookCreateFailureResponse{Status: FAILURE, ErrorMsg: errorMsg})
		return
	}

	ctx.JSON(iris.StatusOK, BookCreateSuccessResponse{Status: SUCCESS, Name: request.Name, Id: new_id})
}

func validate(name string) (string, bool) {
	if name == "" {
		return "Empty Name Not Allowed", false
	}
	books := make([]Book, 128)
	booksTableConnector.Find(bson.M{}).Iter().All(&books)

	for _, book := range books {
		if name == book.Name {
			return "Book with same name already exists!", false
		}
	}
	return "", true

}

func deleteBook(ctx *iris.Context) {
	name := ctx.Param("name")
	fmt.Println("name to be deleted : ", name)
	booksTableConnector.Remove(bson.M{"name": name})
	ctx.JSON(iris.StatusOK, irisTestResponse{"SUCCESS"})
}

func getBook(ctx *iris.Context) {
	name := ctx.Param("name")
	books := make([]Book, 128)
	booksTableConnector.Find(bson.M{"name": name}).All(&books)
	var item Book
	for _, book := range books {
		if name == book.Name {
			item = book
			break
		}
	}
	if item.Name == "" {
		ctx.JSON(iris.StatusOK, BookItemFailureResponse{Status: FAILURE, ErrorMsg: "BooK Not Found!"})
	} else {
		ctx.JSON(iris.StatusOK, BookItemSuccessResponse{Status: SUCCESS, Item: item})
	}

}
func updateBook(ctx *iris.Context) {
	name := ctx.Param("name")

	books := make([]Book, 128)
	booksTableConnector.Find(bson.M{"name": name}).All(&books)
	var item Book
	for _, book := range books {
		if name == book.Name {
			item = book
			break
		}
	}
	if item.Name == "" {
		ctx.JSON(iris.StatusOK, BookItemFailureResponse{Status: FAILURE, ErrorMsg: "BooK Not Found!"})
		return
	}

	request := &BookUpdateRequest{}
	if err := ctx.ReadJSON(request); err != nil {
		ctx.JSON(iris.StatusBadRequest, fmt.Sprintf("INVALID-INPUT-JSON error: %v", err))
		return
	}

	item.Name = request.Name
	booksTableConnector.Update(bson.M{"name": name}, bson.M{"$set": bson.M{"name": request.Name}})
	ctx.JSON(iris.StatusOK, BookItemSuccessResponse{Status: SUCCESS, Item: item})
}
