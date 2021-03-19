package stores

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type service struct {
	l *zap.Logger
}

func Build() fx.Option {
	return fx.Options(
		fx.Provide(func(l *zap.Logger) *service { return &service{l} }),
		fx.Invoke(func(r chi.Router, s *service) {
			rtr := chi.NewRouter()
			r.Mount("/stores", rtr)

			rtr.Get("/", s.get)
		}),
	)
}

func (s *service) get(w http.ResponseWriter, r *http.Request) {
	// etc...
}
