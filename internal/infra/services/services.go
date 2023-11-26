package services

import (
	"github.com/mateusnasciment/APIGOLANG/internal/application/interfaces"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewPlaygroundJsonValidator,
		fx.As(new(interfaces.JsonValidator)),
	),
)
