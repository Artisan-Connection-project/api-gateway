package handlers

import (
	"net/http"

	pro "api_gateway/genproto/product_service"

	"github.com/gin-gonic/gin"
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
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/product [post]
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

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) EditProduct(c *gin.Context) {
	reqPro := pro.EditProductRequest{}
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

func (h *productHandler) DeleteProduct(c *gin.Context) {
	reqPro := pro.DeleteProductRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.DeleteProduct(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) GetProducts(c *gin.Context) {
	reqPro := pro.GetProductsRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.GetProducts(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) GetProduct(c *gin.Context) {
	reqPro := pro.GetProductRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.GetProduct(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) SearchProducts(c *gin.Context) {
	reqPro := pro.SearchProductsRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.SearchProducts(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) AddRating(c *gin.Context) {
	reqPro := pro.AddRatingRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.AddRating(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
}

func (h *productHandler) GetRatings(c *gin.Context) {
	reqPro := pro.GetRatingsRequest{}
	err := c.ShouldBindJSON(&reqPro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resPro, err := h.ProductService.GetRatings(c, &reqPro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resPro)
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
