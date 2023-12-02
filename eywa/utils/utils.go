package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// IsExist returns true if path exist
func IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

func GetFiles(path, ext string) ([]string, error) {
	if !IsExist(path) {
		return nil, errors.New("path dose not exist:" + path)
	}

	var files []string
	err := filepath.Walk(path, func(file_path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(file_path) == ext {
			files = append(files, file_path)
		}

		return nil
	})

	if err != nil {
		return nil, errors.New("filepath.Walk() failed:" + err.Error())
	}

	if len(files) == 0 {
		return nil, errors.New("file not found")
	}

	return files, nil
}

func Write(path, text string) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("cannot create file:%v\n", err)
		return err
	}
	defer f.Close()

	f.WriteString(text)

	return nil
}

func IsIPv6Addr(ip string) bool {
	return !strings.Contains(ip, ".")
}

func ExtractPort(addr string) string {
	var terms []string
	if IsIPv6Addr(addr) {
		terms = strings.Split(addr, "]:")
	} else {
		terms = strings.Split(addr, ":")
	}

	if len(terms) > 1 {
		return terms[1]
	}

	return ""
}

func FindKeys(str string, keys []string, sep, prefix string, end byte) []string {
	if len(str) == 0 {
		return nil
	}
	var texts []string

	tmp := strings.ReplaceAll(str, " ", "")
	for _, k := range keys {
		key := k + sep
		if prefix != "" {
			key = k + prefix + sep
		}
		texts = append(texts, FindKey(tmp, key, end))
	}

	return texts
}

func FindKey(str, key string, end byte) string {
	x := strings.Index(str, key)
	if x < 0 {
		return ""
	}

	m := []byte(str)
	for i := x + len(key); i < len(str); i++ {
		if m[i] == end {
			return string(m[x+len(key) : i])
		}
	}

	return ""
}

func StrAdapter(value, condition, newValue string) string {
	if value == condition {
		return newValue
	}

	return value
}

func IntAdapter(value, condition, newValue int) int {
	if value == condition {
		return newValue
	}

	return value
}

func BoolAdapter(value, condition, newValue bool) bool {
	if value == condition {
		return newValue
	}

	return value
}
