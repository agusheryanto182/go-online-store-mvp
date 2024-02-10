package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/agusheryanto182/go-online-store-mvp/domain/category"
	"github.com/agusheryanto182/go-online-store-mvp/domain/category/dto"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/agusheryanto182/go-online-store-mvp/helper/response"
	"github.com/agusheryanto182/go-online-store-mvp/helper/validator"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandlerImpl struct {
	categoryService category.CategoryServiceInterface
}

func NewCategoryHandler(categoryService category.CategoryServiceInterface) category.CategoryHandlerInterface {
	return &CategoryHandlerImpl{categoryService: categoryService}
}

func (h *CategoryHandlerImpl) CreateCategory(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Unauthorized")
	}

	input := dto.CreateCategoryRequest{}
	if err := c.BodyParser(&input); err != nil {
		return response.SendStatusBadRequest(c, "invalid input : "+err.Error())
	}

	if err := validator.ValidateStruct(input); err != nil {
		return response.SendStatusBadRequest(c, "error validating input : "+err.Error())
	}

	file, err := c.FormFile("image")
	if err != nil {
		return response.SendStatusBadRequest(c, "failed get image"+err.Error())
	}

	uniqueID := time.Now().Format("20060102150405")

	path := fmt.Sprintf("assets/images/%s-%v", uniqueID, file.Filename)

	if err := c.SaveFile(file, path); err != nil {
		return response.SendStatusBadRequest(c, "failed to save file images : "+err.Error())
	}

	input.Image = path

	createdCategory, err := h.categoryService.CreateCategory(&input)
	if err != nil {
		return response.SendStatusBadRequest(c, "error to create category : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success to create category", dto.CreateCategoryResponseFormatter(createdCategory))
}

func (h *CategoryHandlerImpl) UpdateCategory(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Unauthorized")
	}

	ID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id : "+err.Error())
	}

	input := dto.UpdateCategoryRequest{}
	if err := c.BodyParser(&input); err != nil {
		return response.SendStatusBadRequest(c, "invalid input : "+err.Error())
	}

	if err := validator.ValidateStruct(input); err != nil {
		return response.SendStatusBadRequest(c, "error validating input : "+err.Error())
	}

	updatedCategory, err := h.categoryService.UpdateCategory(ID, &input)
	if err != nil {
		return response.SendStatusBadRequest(c, "error to create category : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success to create category", updatedCategory)
}

func (h *CategoryHandlerImpl) DeleteCategory(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Unauthorized")
	}

	ID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id : "+err.Error())
	}

	if err := h.categoryService.DeleteCategory(ID); err != nil {
		return response.SendStatusBadRequest(c, "failed to delete category : "+err.Error())
	}
	return nil
}

func (h *CategoryHandlerImpl) GetAllCategory(c *fiber.Ctx) error {
	ID, _ := strconv.Atoi(c.Params("id"))

	if ID > 0 {
		category, err := h.categoryService.GetCategoryByID(ID)
		if err != nil {
			return response.SendStatusBadRequest(c, "failed get category : "+err.Error())
		}
		return response.SendStatusOkWithDataResponse(c, "success", dto.GetCategoryResponseFormatter(category))
	} else {
		categories, err := h.categoryService.GetAllCategory()
		if err != nil {
			return response.SendStatusBadRequest(c, "failed get all category : "+err.Error())
		}
		return response.SendStatusOkWithDataResponse(c, "success", categories)
	}
}
