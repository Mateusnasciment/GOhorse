package eventhandlers

import (
	events "github.com/mateusnasciment/APIGOLANG/internal/domain/events/books"
)

type BooksEventHandler interface {
	Handle(event events.BookEvent) 
}
