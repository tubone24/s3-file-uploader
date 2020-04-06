package utils

import (
	"path/filepath"
	"strings"
)

func GetContentType(key string) string {
	//pos := strings.LastIndex(key, ".")

	if strings.HasSuffix(key, ".tar.gz") {
		return "application/x-tar"
	}

	ext := filepath.Ext(key)

	switch ext {
	case ".txt":
		return "text/plain"
	case ".tsv":
		return "text/tab-separated-values"
	case ".csv":
		return "text/csv"
	case ".json":
		return "application/json"
	case ".pdf":
		return "application/pdf"
	case ".xls":
		return "application/vnd.ms-excel"
	case ".xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case ".ppt":
		return "application/vnd.ms-powerpoint"
	case ".pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case ".doc":
		return "application/msword"
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".zip":
		return "application/zip"
	case ".lzh":
		return "application/x-lzh"
	case ".tar.gz":
		return "application/x-tar"
	case ".tgz":
		return "application/x-tar"
	case ".tar":
		return "application/x-tar"
	case ".bz":
		return "application/x-bzip"
	case ".bz2":
		return "application/x-bzip2"
	case ".gz":
		return "application/gzip"
	case ".rar":
		return "application/vnd.rar"
	case ".7z":
		return "application/x-7z-compressed"
	case ".xml":
		return "application/xml"
	case ".bin":
		return "application/octet-stream"
	default:
		return "application/x-www-form-urlencoded"
	}
}
