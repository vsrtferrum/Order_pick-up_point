package imdb

import (
	"context"
	"sync"

	"github.com/opentracing/opentracing-go"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func NewRepository() *Repository {
	return &Repository{
		positionByName: make(map[uint32]*models.DataUnitJson),
	}
}

type Repository struct {
	lock           sync.RWMutex
	positionByName map[uint32]*models.DataUnitJson
}

func (r *Repository) UpsertPositions(ctx context.Context, positions []*models.DataUnitJson) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "imdbRepo.UpsertPositions")
	defer span.Finish()

	r.lock.Lock()
	for _, position := range positions {
		r.positionByName[position.IdOrder] = position
	}
	r.lock.Unlock()
	return nil
}

func (r *Repository) PositionList(ctx context.Context, filter *models.PositionRepoFilter) ([]*models.DataUnitJson, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "imdbRepo.PositionList")
	defer span.Finish()

	r.lock.RLock()
	defer r.lock.RUnlock()

	if filter.Empty() {
		var i int
		positions := make([]*models.DataUnitJson, len(r.positionByName))

		for _, position := range r.positionByName {
			positions[i] = position
			i++
		}

		return positions, nil
	}

	positions := make([]*models.DataUnitJson, 0, len(filter.IdOrders))

	for _, name := range filter.IdOrders {
		position, ok := r.positionByName[name]
		if !ok {
			continue
		}

		positions = append(positions, position)
	}

	return positions, nil
}
