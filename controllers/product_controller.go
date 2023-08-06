package controllers

import (
	"gocommerce/models"
	"gocommerce/services"
	"gocommerce/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		productService: services.NewProductService(db),
	}
}

func NewProductControllerTest(productService *services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newProduct, err := c.productService.CreateProduct(&product)
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	utils.SendSuccessResponse(ctx, http.StatusCreated, "User created successfully", newProduct)
}

func (c *ProductController) GetProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	product, err := c.productService.GetProductByID(productID)
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusNotFound, "Product Not Found")
		return
	}

	utils.SendSuccessResponse(ctx, http.StatusFound, "User created successfully", product)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	var updatedProduct models.Product
	if err := ctx.ShouldBindJSON(&updatedProduct); err != nil {
		utils.SendErrorResponse(ctx, 400, "Invalid request payload")
		return
	}

	// Convert the productID (string) to a uint
	parsedProductID, err := strconv.ParseUint(productID, 10, 64)
	if err != nil {
		utils.SendErrorResponse(ctx, 400, "Invalid product ID")
		return
	}

	// Assign the productID to the ID field of updatedProduct
	updatedProduct.ID = uint(parsedProductID)

	// Call the UpdateProduct function from the service
	err = c.productService.UpdateProduct(&updatedProduct)
	if err != nil {
		utils.SendErrorResponse(ctx, 500, "Failed to update product")
		return
	}

	utils.SendSuccessResponse(ctx, 200, "Product updated successfully", updatedProduct)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	err := c.productService.DeleteProduct(productID)
	if err != nil {
		utils.SendErrorResponse(ctx, 500, "Failed to delete product")
		return
	}

	utils.SendSuccessResponse(ctx, 200, "Product deleted successfully", nil)
}

func (c *ProductController) ListProducts(ctx *gin.Context) {
	products, err := c.productService.ListProducts()
	if err != nil {
		utils.SendErrorResponse(ctx, 500, "Failed to fetch products")
		return
	}

	utils.SendSuccessResponse(ctx, 200, "Products fetched successfully", products)
}