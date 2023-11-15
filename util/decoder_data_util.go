package util

import (
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

const (
	Passphrase = "abcdefghijklmnopqrstuvwx"
)

func DecodeRequestAndValidate(writer http.ResponseWriter, req *http.Request, data interface{}) error {
	validate := validator.New()

	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return errors.New("failed to decode")
	}

	if err := validate.Struct(data); err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte(err.Error()))
		return errors.New(err.Error())
	}

	return nil
}

func ValidateImage(file multipart.File, header *multipart.FileHeader, writer http.ResponseWriter) error {
	if header.Size >= 1024*1024 {
		http.Error(writer, "Too big!!!", http.StatusBadRequest)
		return nil
	}

	defer file.Close()

	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		log.WithFields(log.Fields{
			"error": "server failed to process",
		}).Error("Fatal")
		return errors.New(err.Error())
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/jpg" && filetype != "image/webp" && filetype != "image/avif" && filetype != "image/png" {
		log.WithFields(log.Fields{
			"error": "invalid format",
		}).Error("Fatal")
		return errors.New("invalid format")
	}

	_, seekerr := file.Seek(0, io.SeekStart)
	if seekerr != nil {
		log.WithFields(log.Fields{
			"error": "server failed to process",
		}).Error("Fatal")
		return errors.New(seekerr.Error())
	}

	return nil
}

func ConvertStrInt64(data interface{}, base int, bitSize int) (int64, error) {
	format, formatErr := strconv.ParseInt(data.(string), base, bitSize)
	if formatErr != nil {
		return 0, formatErr
	}

	return format, nil
}

func ConvertStrInt(data interface{}, base int, bitSize int) (int, error) {
	format, formatErr := strconv.ParseInt(data.(string), base, bitSize)
	if formatErr != nil {
		return 0, formatErr
	}

	return int(format), nil
}
