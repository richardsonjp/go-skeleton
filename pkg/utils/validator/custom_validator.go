// custom_validator.go
package validator

import (
	"go-skeleton/pkg/utils/lang"
	timeutil "go-skeleton/pkg/utils/time"

	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

type ValidatorWrapper struct {
	// Validator key
	Key string
	// Return NIL to make it valid
	ValidateField func(value interface{}) interface{}
	// Return validation message
	GetMessage func(e validator.FieldError, vw ValidatorWrapper) string
}

var ValidatorWrappers = []ValidatorWrapper{
	*whitelistIpValidator,
	*alphaSpaceValidator,
	*dateYYYYMMDDValidator,
	*passwordValidator,
	*nameValidator,
}

func findValidatorByTag(tag string) *ValidatorWrapper {
	for _, v := range ValidatorWrappers {
		if v.Key == tag {
			return &v
		}
	}
	return nil
}

var whitelistIpValidator = &ValidatorWrapper{
	Key: "whitelist_ip_validator",
	ValidateField: func(value interface{}) interface{} {
		isIPAllowed := func(IP string) bool {
			forbiddenIPs := []string{"0.0.0.0"}
			for _, v := range forbiddenIPs {
				if strings.TrimSpace(IP) == v {
					return false
				}
			}
			return true
		}

		isValidIP4 := func(ipAddress string) bool {
			ipAddress = strings.TrimSpace(ipAddress)
			re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
			if re.MatchString(ipAddress) {
				return true
			}
			return false
		}

		whitelistIps := value.([]string)
		for _, ip := range whitelistIps {
			if !isIPAllowed(ip) || !isValidIP4(ip) {
				return ip
			}
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_WHITELIST", templateData)
		return msg
	},
}

var alphaSpaceValidator = &ValidatorWrapper{
	Key: "alphaspace_validator",
	ValidateField: func(value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return value
		}
		re, _ := regexp.Compile(`^[a-zA-Z ]+$`)
		if !re.MatchString(str) {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_ALPHASPACE", templateData)
		return msg
	},
}

var nameValidator = &ValidatorWrapper{
	Key: "name_validator",
	ValidateField: func(value interface{}) interface{} {
		if str, ok := value.(string); !ok {
			return value
		} else {
			str = strings.TrimSpace(str)
			if len(str) == 0 {
				return str
			}
			if re, err := regexp.Compile(`^[a-zA-Z',. ]+$`); err != nil {
				return str
			} else if !re.MatchString(str) {
				return str
			}
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_VALUE", templateData)
		return msg
	},
}

var dateYYYYMMDDValidator = &ValidatorWrapper{
	Key: "date_yyyymmdd_validator",
	ValidateField: func(value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return value
		}
		_, err := timeutil.Parse(str, timeutil.ISO8601TimeDate)
		if err != nil {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_DATE", templateData)
		return msg
	},
}

var passwordValidator = &ValidatorWrapper{
	Key: "password_validator",
	ValidateField: func(value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return value
		}
		if !regexp.MustCompile(`^((?:.*)(?:.*\d)(?:.*[a-zA-Z])|(?:.*[a-zA-Z])(?:.*\d)).{0,}$`).MatchString(str) {
			return value
		}
		if len(str) < 8 || len(str) > 64 {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		msg, _ := lang.CurrentTranslation.Translate("ERR_PASSWORD")
		return msg
	},
}

func (o *ValidatorWrapper) Validate(fl validator.FieldLevel, vw ValidatorWrapper) bool {
	return vw.ValidateField(fl.Field().Interface()) == nil
}
