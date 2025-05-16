// handler/transfer.go
package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "transaction/usecase"
)

type TransferHandler struct {
    Usecase usecase.TransferUsecase
}

func (h *TransferHandler) TransferMoney(c *gin.Context) {
    token := c.GetHeader("Authorization") // Bearer ...
    if token == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
        return
    }

    var req struct {
        ToUserID string `json:"to_user_id"`
        Amount   int64  `json:"amount"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    resp, err := h.Usecase.TransferMoney(c.Request.Context(), token, req.ToUserID, req.Amount)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}
