package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tensorchord/openmodelz/agent/api/types"
)

// @Summary     Scale the inferences.
// @Description Scale the inferences.
// @Tags        inference
// @Accept      json
// @Produce     json
// @Param       namespace query    string                    true "Namespace"  example("modelz-d3524a71-c17c-4c92-8faf-8603f02f4713")
// @Param       request   body     types.ScaleServiceRequest true "query params"
// @Success     202       {object} []types.ScaleServiceRequest
// @Router      /system/scale-inference [post]
func (s *Server) handleInferenceScale(c *gin.Context) error {
	var req types.ScaleServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return NewError(
			http.StatusBadRequest, err, "inference-scale")
	}

	namespace := c.Query("namespace")
	if namespace == "" {
		return NewError(
			http.StatusBadRequest, errors.New("namespace is required"), "inference-scale")
	}

	if err := s.runtime.InferenceScale(c.Request.Context(),
		namespace, req); err != nil {
		return errFromErrDefs(err, "inference-scale")
	}

	c.JSON(http.StatusAccepted, req)
	return nil
}
