package validator

import (
	"fmt"
	"go-skeleton/pkg/utils/wording"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var cvtor *customValidator

type customValidator struct {
	inner *validator.Validate
}

func Validate(input interface{}) (string, error) {
	err := cvtor.inner.Struct(input)
	message := GetValidatorMessage(err)
	return message, err
}

func GetValidatorMessage(err error) string {
	message := ""
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {

				field := wording.ToSnakeCase(e.Field())
				params := e.Param()
				customValidator := findValidatorByTag(e.Tag())
				messageOverwriter := findMessageOverwriter(e.Tag(), field)
				validatorTranslation := findValidatorTranslation(e.Tag(), field, params)
				if customValidator != nil {
					message = customValidator.GetMessage(e, *customValidator)
				} else if validatorTranslation != nil {
					message = validatorTranslation.GetMessage(e, *validatorTranslation)
				} else if messageOverwriter != nil {
					message = messageOverwriter.GetMessage(e, *messageOverwriter)
				} else {

					baseFormat := "{0} does not meet {1}{2} criteria"
					param := e.Param()
					if len(param) == 0 {
						param = ""
					} else {
						param = "(" + param + ")"
					}
					message = baseFormat
					values := []string{field, e.Tag(), param}
					for i, v := range values {
						message = strings.Replace(message, "{"+fmt.Sprint(i)+"}", v, -1)
					}
				}
				break // just show first error
			}
		} else {
			message = "Failed input validation"
		}
	}

	return message
}

func init() {
	obj := validator.New()
	obj.RegisterTagNameFunc(func(fld reflect.StructField) string {
		splits := strings.SplitN(fld.Tag.Get("json"), ",", 2)
		if len(splits) == 0 {
			return ""
		}

		name := splits[0]
		if name == "-" {
			return ""
		}

		return name
	})

	for _, v := range ValidatorWrappers {
		obj.RegisterValidation(v.Key, func(fl validator.FieldLevel) bool {
			customValidator := findValidatorByTag(fl.GetTag())
			return customValidator.Validate(fl, *customValidator)
		})
	}

	cvtor = &customValidator{obj}
}
