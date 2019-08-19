package controllers

import (
	"mime/multipart"
)

type Context interface {
	MustGet(string) interface{}
	Get(string) (interface{}, bool)
	Param(string) string
	Bind(interface{}) error
	ShouldBind(interface{}) error
	Status(int)
	JSON(int, interface{})
	FormFile(string) (*multipart.FileHeader, error)
	MultipartForm() (*multipart.Form, error)
	SaveUploadedFile(*multipart.FileHeader, string) error
	String(int, string, ...interface{})
}
