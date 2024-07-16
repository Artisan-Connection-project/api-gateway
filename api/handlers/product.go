package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	pro "api_gateway/genproto/product_service"
	"api_gateway/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler interface {
	AddProduct(c *gin.Context)
	EditProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	GetProduct(c *gin.Context)
	SearchProducts(c *gin.Context)
	AddRating(c *gin.Context)
	GetRatings(c *gin.Context)
	PlaceOrder(c *gin.Context)
	CancelOrder(c *gin.Context)
	UpdateOrderStatus(c *gin.Context)
	GetOrders(c *gin.Context)
	GetOrder(c *gin.Context)
	PayOrder(c *gin.Context)
	CheckPaymentStatus(c *gin.Context)
	UpdateShippingInfo(c *gin.Context)
	AddArtisanCategory(c *gin.Context)
	AddProductCategory(c *gin.Context)
	GetStatistics(c *gin.Context)
	GetUserActivity(c *gin.Context)
	GetRecommendations(c *gin.Context)
	GetArtisanRankings(c *gin.Context)
}

type productHandler struct {
	ProductService pro.ProductServiceClient
}

func NewProductHandler(productClient pro.ProductServiceClient) ProductHandler {
	return &productHandler{ProductService: productClient}
}

// @Summary Create Product
// @Description Creating a new product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param Add_Product_Request body product_service.AddProductRequest true "Create"
// @Success 201 {object} product_service.AddProductResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/create	 [post]
func (h *productHandler) AddProduct(c *gin.Context) {
	reqPro := pro.AddProductRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.AddProduct(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resPro)
}

// @Summary Update Product
// @Description Update a new product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param Product_id path string true "product_id"
// @Param UpdateProduct body product_service.EditProductRequest true "Update"
// @Success 200 {object} product_service.EditProductResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/{id}	 [put]
func (h *productHandler) EditProduct(c *gin.Context) {
	id := c.Param("id")
	reqPro := pro.EditProductRequest{Id: id}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.EditProduct(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

// @Summary Delete Product
// @Description Delete a product by id
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" format(uuid)
// @Success 200 {object} product_service.DeleteProductResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/{id} [delete]
func (h *productHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	reqPro := pro.DeleteProductRequest{Id: id}
	_, err := h.ProductService.DeleteProduct(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully ID:" + reqPro.Id})
}

// @Summary Get All Products
// @Description get all products in page by limit
// @Tags Product
// @Accept  json
// @Produce  json
// @Param limit query string true "limit"
// @Param page query string true "page"
// @Success 200 {object} []product_service.Product
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/all	 [get]
func (h *productHandler) GetProducts(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	reqPro := pro.GetProductsRequest{
		Limit: limit,
		Page:  page,
	}
	log.Println(page, limit)
	resPro, err := h.ProductService.GetProducts(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resPro.Products)
}

// @Summary Get a Product
// @Description Get a Product by its ID
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" format(uuid)
// @Success 200 {object} product_service.Product
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/{id} [get]
func (h *productHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	reqPro := pro.GetProductRequest{Id: id}

	resPro, err := h.ProductService.GetProduct(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resPro == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, resPro.Product)
}

// @Summary Get a Product
// @Description Get a Product by its ID
// @Tags Product
// @Accept  json
// @Produce  json
// @Param query query string false "product name"
// @Param category query string false "category name"
// @Param min_price query string false "minimum price"
// @Param max_price query string false "maximum price"
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Success 200 {object} []product_service.Product
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/search [post]
func (h *productHandler) SearchProducts(c *gin.Context) {
	query := c.Query("query")
	category := c.Query("category")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")
	limit := c.Query("limit")
	page := c.Query("page")

	minPricef, err := strconv.ParseFloat(minPrice, 32)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	maxPricef, err := strconv.ParseFloat(maxPrice, 32)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	reqPro := pro.SearchProductsRequest{
		Query:    query,
		Category: category,
		MinPrice: float32(minPricef),
		MaxPrice: float32(maxPricef),
		Limit:    limit,
		Page:     page,
	}
	resPro, err := h.ProductService.SearchProducts(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro.Products)
}

// @Summary Rate A Product
// @Description Rate a product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param product_id path string true "Product ID" format(uuid)
// @Param AddRatingRequest body models.AddRatingRequest true "rating a product"
// @Success 201 {object} product_service.AddRatingResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/{product_id}/rate [post]
func (h *productHandler) AddRating(c *gin.Context) {
	id := c.Param("product_id")
	reqPro := pro.AddRatingRequest{ProductId: id}

	reqModel := models.AddRatingRequest{}

	err := c.ShouldBindJSON(&reqModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//converting request body to grpc struct
	reqPro.UserId = reqModel.UserId
	reqPro.Comment = reqModel.Comment
	reqPro.Rating = reqModel.Rating

	resPro, err := h.ProductService.AddRating(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resPro)
}

// @Summary Get all ratings for a product
// @Description Get all ratings for a product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" format(uuid)
// @Success 200 {object} []product_service.Rating
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/{id}/ratings [get]
func (h *productHandler) GetRatings(c *gin.Context) {
	id := c.Param("id")

	reqPro := pro.GetRatingsRequest{ProductId: id}

	resPro, err := h.ProductService.GetRatings(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro.Ratings)
}

func (h *productHandler) PlaceOrder(c *gin.Context) {
	reqPro := pro.PlaceOrderRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.PlaceOrder(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) CancelOrder(c *gin.Context) {
	reqPro := pro.CancelOrderRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.CancelOrder(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) UpdateOrderStatus(c *gin.Context) {
	reqPro := pro.UpdateOrderStatusRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.UpdateOrderStatus(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) GetOrders(c *gin.Context) {
	reqPro := pro.GetOrdersRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.GetOrders(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) GetOrder(c *gin.Context) {
	reqPro := pro.GetOrderRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.GetOrder(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) PayOrder(c *gin.Context) {
	reqPro := pro.PayOrderRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.PayOrder(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) CheckPaymentStatus(c *gin.Context) {
	reqPro := pro.CheckPaymentStatusRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.CheckPaymentStatus(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) UpdateShippingInfo(c *gin.Context) {
	reqPro := pro.UpdateShippingInfoRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.UpdateShippingInfo(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) AddArtisanCategory(c *gin.Context) {
	reqPro := pro.AddArtisanCategoryRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.AddArtisanCategory(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

// @Summary Create New Category for Product
// @Description Create new category for Product
// @Tags Product Category
// @Accept  json
// @Produce  json
// @Param Product_Category body product_service.AddProductCategoryRequest true "product category"
// @Success 200 {object} product_service.AddProductCategoryResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/categories/product [post]
func (h *productHandler) AddProductCategory(c *gin.Context) {
	reqPro := pro.AddProductCategoryRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.AddProductCategory(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) GetStatistics(c *gin.Context) {
	reqPro := pro.GetStatisticsRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.GetStatistics(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) GetUserActivity(c *gin.Context) {
	reqPro := pro.GetUserActivityRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.GetUserActivity(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) GetRecommendations(c *gin.Context) {
	reqPro := pro.GetRecommendationsRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.GetRecommendations(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) GetArtisanRankings(c *gin.Context) {
	reqPro := pro.GetArtisanRankingsRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.GetArtisanRankings(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}
