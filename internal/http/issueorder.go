package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/logger"
	"gitlab.ozon.dev/berkinv/homework/internal/metrics"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

const IssueOrder = "Issued"

func (h *Handler) IssueCnt(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	span, ctx := opentracing.StartSpanFromContext(ctx, "http.AddPositions")
	defer span.Finish()

	statusCode, err := h.issueOrder(ctx, request)
	if err != nil {
		metrics.IncIssueByHandler(IssueOrder)
		http.Error(writer, err.Error(), statusCode)
		return
	}

	metrics.IncIssueByHandler(IssueOrder)
}

func (h *Handler) issueOrder(ctx context.Context, request *http.Request) (int, error) {

	if request.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, fmt.Errorf("invalid method: %s", request.Method)
	}

	decoder := json.NewDecoder(request.Body)

	var positions []*models.DataUnitJson
	err := decoder.Decode(&positions)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("failed to read wantPositions data: %s", err)
	}

	err = h.positionUpserter.UpsertPositions(ctx, positions)
	if err != nil {
		logger.Warnf(ctx, "failed to add wantPositions: %s", err)
		return http.StatusInternalServerError, errors.InternalServerError
	}

	return http.StatusOK, nil
}
