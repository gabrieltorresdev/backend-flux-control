package controller

import (
	"net/http"
	"strconv"

	"github.com/gabrieltorresdev/backend-flux-control/internal/application/interfaces"
	"github.com/gabrieltorresdev/backend-flux-control/internal/infrastructure/http/v1/rest/gin/request/transaction"
	transactionResponse "github.com/gabrieltorresdev/backend-flux-control/internal/infrastructure/http/v1/rest/gin/response/transaction"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionController struct {
	transactionService interfaces.TransactionServiceInterface
}

func NewTransactionController(transactionService interfaces.TransactionServiceInterface) *TransactionController {
	return &TransactionController{
		transactionService: transactionService,
	}
}

func (c *TransactionController) GetTransactions(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	transactions, pagination, err := c.transactionService.FindAllPaginated(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := transactionResponse.BuildTransactionsResponse(
		ctx,
		transactions,
		pagination.Page,
		pagination.PageSize,
		http.StatusOK,
	)

	if response.PageInfo != nil {
		response.PageInfo.TotalItems = int(pagination.TotalItems)
		response.PageInfo.TotalPages = pagination.TotalPages
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var createTransactionRequest transaction.CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&createTransactionRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: get user id from auth
	userId := uuid.New()
	createTransactionDTO := createTransactionRequest.ToCreateTransactionDTO(userId)

	transaction, err := c.transactionService.Create(createTransactionDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := transactionResponse.BuildTransactionResponse(ctx, *transaction, http.StatusCreated)
	ctx.JSON(http.StatusCreated, response)
}
