package controller

import (
	"belajar-go-rest-api/helper"
	"belajar-go-rest-api/model/web"
	"belajar-go-rest-api/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err3 := strconv.Atoi(categoryId)
	helper.PanicIfError(err3)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err3 := strconv.Atoi(categoryId)
	helper.PanicIfError(err3)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindyById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err3 := strconv.Atoi(categoryId)
	helper.PanicIfError(err3)

	categoryRespnse := controller.CategoryService.FindyById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryRespnse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryRespnses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryRespnses,
	}

	helper.WriteResponseBody(writer, webResponse)
}
