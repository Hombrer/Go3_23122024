package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	db "Bankstore/db/sqlc"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// TODO: add routes to router
	router.POST("/accounts", server.CreateAccount)

	server.router = router
	return server
}

type CreateAccountRequest struct {
	Owner string `json:"owner" binding:"required"`
	Currency db.Currency `json:"currency" binding:"required,oneof=USR EUR"`

}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
	}

	arg := db.CreateAccountParams {
		Owner: req.Owner,
		Currency: req.Currency,
		Balance: 0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// errorResponce return gin.H -> map[string]interface{}
func errorResponce(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start server method
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}