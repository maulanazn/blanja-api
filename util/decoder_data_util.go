package util

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
)

const (
	Passphrase = "abcdefghijklmnopqrstuvwx"
)

func DecodeRequestAndValidate(writer http.ResponseWriter, req *http.Request, data interface{}) {
	validate := validator.New()

	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		Log2File(err.Error())
	}

	if err := validate.Struct(data); err != nil {
		Log2File(err.Error())
		writer.WriteHeader(400)
		writer.Write([]byte(err.Error()))
	}
}

func ValidateImage(file multipart.File, header *multipart.FileHeader, writer http.ResponseWriter) {
	if header.Size >= 1024*1024 {
		http.Error(writer, "Too big!!!", http.StatusBadRequest)
		return
	}

	defer file.Close()

	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		Log2File(err.Error())
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/jpg" && filetype != "image/webp" && filetype != "image/avif" && filetype != "image/png" {
		Log2File(err.Error())
		return
	}

	_, seekerr := file.Seek(0, io.SeekStart)
	if seekerr != nil {
		Log2File(err.Error())
		return
	}
}

func ConvertStrInt64(data interface{}, base int, bitSize int) int64 {
	format, formatErr := strconv.ParseInt(data.(string), base, bitSize)
	if formatErr != nil {
		return 0
	}

	return format
}

func ConvertStrInt(data interface{}, base int, bitSize int) int {
	format, formatErr := strconv.ParseInt(data.(string), base, bitSize)
	if formatErr != nil {
		return 0
	}

	return int(format)
}
