package handler

import (
	"errors"
	"net/http"
	"strconv"
	"tasker_go/internal/service"
	"tasker_go/internal/transport/http/responder"

	"github.com/gorilla/mux"
)

type JudgeHandler struct {
	judgeService service.JudgeService
}

func NewJudgeHandler(s service.JudgeService) *JudgeHandler {
	return &JudgeHandler{
		judgeService: s,
	}
}

func (h *JudgeHandler) GetJudge(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	userID := UserIDFromContext(ctx)

	taskId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "invalid id")
		return
	}

	judge, err := h.judgeService.GetByTaskID(ctx, userID, uint(taskId))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTaskNotFound):
			responder.RespondWithError(w, http.StatusNotFound, "judge not found")
		case errors.Is(err, service.ErrForbidden):
			responder.RespondWithError(w, http.StatusForbidden, "forbidden")
		default:
			responder.RespondWithError(w, http.StatusInternalServerError, "internal error")
		}
		return
	}

	responder.RespondWithJSON(w, http.StatusOK, dto.JudgeToResponse(judge))

}
