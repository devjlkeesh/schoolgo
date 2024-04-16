package utils

import (
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
	"school23/internal/models"
	"strconv"
	"time"
)

func RequireNonNilOrElseGet[T any](a T, defaultValue T) T {
	if !reflect.ValueOf(a).IsZero() {
		return a
	}
	return defaultValue
}

func UploadFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	generatedFileName := getFilename(header.Filename)
	dst, err := os.Create(filepath.Join("uploads", generatedFileName))
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}
	return dst.Name(), nil
}

func IsNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func ParseTOIntSafe(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return value
}

func GetPagination(totalCount int) []models.Pagination {
	var pagination []models.Pagination
	var pages int = int(math.Ceil(float64(totalCount) / 5))
	for i := 0; i < pages; i++ {
		pagination = append(pagination, models.Pagination{Index: i, Value: i + 1})
	}
	return pagination
}

func getFilename(filename string) string {
	ext := filepath.Ext(filename)
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}
