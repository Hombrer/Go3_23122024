package api

import (
	db "Bankstore/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)


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

type getAccountRequest struct {
	// URL: /accounts/1 
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == pgx.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponce(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type listAccountsRequest struct {
	// URL: /accounts/?page_id=1&page_size=10
	// Обработка query parameters в Path
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=15"`
}

func (server *Server) ListAccounts(ctx *gin.Context) {
	var req listAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}
	arg := db.ListAccountsParams {
		Limit: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}