package modules

import (
	"errors"

	"github.com/blushft/strana"
	"github.com/blushft/strana/modules/broker/enhancer"
	"github.com/blushft/strana/modules/sink/loader"
	"github.com/blushft/strana/modules/sink/reporter"
	"github.com/blushft/strana/modules/source/collector"
	"github.com/blushft/strana/platform/config"
)

func New(conf config.Module) (strana.Module, error) {
	switch conf.Type {
	case "collector":
		return collector.New(conf)
	case "enhancer":
		return enhancer.New(conf)
	case "loader":
		return loader.New(conf)
	case "reporter":
		return reporter.New(conf)
	default:
		return nil, errors.New("invalid module: " + conf.Type)
	}
}
