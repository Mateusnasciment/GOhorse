package eventhandlers

import (
	eventhandlers "github.com/mateusnasciment/APIGOLANG/internal/application/interfaces/event-handlers"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewBooksEventHandler,
		fx.As(new(eventhandlers.BooksEventHandler)),
	),
)
