package handler

import (
	"strconv"

	"github.com/agusheryanto182/go-online-store-mvp/domain/product"
	"github.com/agusheryanto182/go-online-store-mvp/domain/product/dto"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/agusheryanto182/go-online-store-mvp/helper/response"
	"github.com/agusheryanto182/go-online-store-mvp/helper/validator"
	"github.com/gofiber/fiber/v2"
)

type ProductHandlerImpl struct {
	productService product.ProductServiceInterface
}

func NewProductService(productService product.ProductServiceInterface) product.ProductHandlerInterface {
	return &ProductHandlerImpl{productService: productService}
}

func (h *ProductHandlerImpl) CreateProduct(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Unauthorized")
	}

	var input dto.CreateProductRequest
	if err := c.BodyParser(&input); err != nil {
		return response.SendStatusBadRequest(c, "invalid input : "+err.Error())
	}

	if err := validator.ValidateStruct(input); err != nil {
		return response.SendStatusBadRequest(c, "error validating input : "+err.Error())
	}

	_, err := h.productService.CreateProduct(&input)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to create product : "+err.Error())
	}

	return response.SendStatusCreatedResponse(c, "success to create product")
}

func (h *ProductHandlerImpl) UpdateProduct(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Unauthorized")
	}

	ID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id")
	}

	var input dto.UpdateProductRequest
	if err := c.BodyParser(&input); err != nil {
		return response.SendStatusBadRequest(c, "invalid input : "+err.Error())
	}

	if err := validator.ValidateStruct(input); err != nil {
		return response.SendStatusBadRequest(c, "error validating input : "+err.Error())
	}

	updatedProduct, err := h.productService.UpdateProduct(ID, &input)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to update product : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success to update product", dto.GetProductByID(updatedProduct))
}

func (h *ProductHandlerImpl) DeleteProduct(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Unauthorized")
	}

	ID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id")
	}

	if err := h.productService.DeleteProduct(ID); err != nil {
		return response.SendStatusBadRequest(c, "failed to delete product : "+err.Error())
	}

	return response.SendStatusOkResponse(c, "success to delete product")
}

func (h *ProductHandlerImpl) GetProductByCategoryID(c *fiber.Ctx) error {
	ID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id : "+err.Error())
	}

	products, err := h.productService.GetProductByCategoryID(ID)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to get products by category : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", dto.GetProducts(products))
}
