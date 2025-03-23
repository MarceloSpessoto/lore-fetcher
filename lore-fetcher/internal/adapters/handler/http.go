package handler

import (
  "lore-fetcher/internal/core/domain"
  "lore-fetcher/internal/core/services"
  "net/http"
  "github.com/gin-gonic/gin"
)

type HTTPHandler struct {
  service services.LFService
}

func NewHTTPHandler(LFService services.LFService) *HTTPHandler {
  return &HTTPHandler{service: LFService}
}

func (h *HTTPHandler) SavePatch(c *gin.Context) {
  var patch domain.Patch
  if err := c.ShouldBindJSON(&patch); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  if err := h.service.SavePatch(patch); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *HTTPHandler) ReadPatches(c *gin.Context) {
  patches, err := h.service.ReadPatches()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"data": patches})
}
