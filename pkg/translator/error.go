package translator

import (
	"encoding/json"
	"fmt"
	"io/ioutil" //nolint: staticcheck, SA1019
	"path/filepath"
	"strconv"
	"strings"
)

type (
	Error struct {
		translation map[string]map[string]ErrorTemplate // language: {error_code: template}
	}

	ErrorTemplate struct {
		Code           string `json:"code"`
		MessageTitle   string `json:"message_title"`
		Message        string `json:"message"`
		Severity       string `json:"severity"`
		HTTPStatusCode int    `json:"http_status_code"`
	}
)

const (
	defaultErrorTemplateKey = "default"
)

func NewError(path string) *Error {
	handler := Error{
		translation: map[string]map[string]ErrorTemplate{},
	}
	handler.readFiles(path)

	return &handler
}

func (e *Error) Translate(locale, code string, vars ...interface{}) ErrorTemplate {
	lang := e.getLanguage(locale)
	template := e.getTemplate(lang, code)
	template.Message = fillMessageVariables(template.Message, vars...)
	template.MessageTitle = fillMessageVariables(template.MessageTitle, vars...)

	cd, er := strconv.Atoi(code)
	if er != nil {
		cd = 400
	}

	template.HTTPStatusCode = cd
	return template
}

// getLanguage gets language from user locale.
// en_ID --> en
func (e *Error) getLanguage(locale string) string {
	localeParts := strings.Split(locale, "_")
	if len(localeParts) == 0 || len(localeParts[0]) == 0 {
		return "en"
	}
	if _, exists := e.translation[localeParts[0]]; !exists {
		return "en"
	}
	return localeParts[0]
}

func (e *Error) getTemplate(lang, code string) ErrorTemplate {
	template, exists := e.translation[lang][code]
	if !exists {
		return e.translation[lang][defaultErrorTemplateKey]
	}

	return template
}

func (e *Error) readFiles(dir string) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {

		return
	}

	translation := make(map[string]map[string]ErrorTemplate, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fullPath := filepath.Join(dir, file.Name())
		content, err := ioutil.ReadFile(fullPath)
		if err != nil {
			continue
		}

		var templates map[string]ErrorTemplate
		// Use the file name as the map key.
		// en.json --> {"en": {}}
		nameParts := strings.Split(file.Name(), ".")
		err = json.Unmarshal(content, &templates)
		if err != nil {

			continue
		}
		translation[nameParts[0]] = templates
	}

	e.translation = translation
}

func fillMessageVariables(message string, vars ...interface{}) string {
	// Count the number of format specifiers in the template
	specifierCount := strings.Count(message, "%v")

	// If there are fewer arguments than format specifiers, fill the extra specifiers with empty strings.
	if len(vars) < specifierCount {
		diff := specifierCount - len(vars)
		for i := 0; i < diff; i++ {
			vars = append(vars, "")
		}
	}

	// If there are more arguments than format specifiers, remove the extra specifiers.
	if len(vars) > specifierCount {
		vars = vars[:specifierCount]
	}

	return fmt.Sprintf(message, vars...)
}
