package main

import "net/http"

// listProductsHandler godoc
// @Summary      List products
// @Description  List products
// @Tags products
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500 {object} error
// @Security     ApiKeyAuth
// @Router       /products [get]
func (app *application) listProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
