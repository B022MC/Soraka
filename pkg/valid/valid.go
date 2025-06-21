package valid

import (
	"reflect"
	"regexp"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type (
	BoolStr          = string
	DefaultValidator struct {
		once     sync.Once
		validate *validator.Validate
	}
)

const (
	BoolStrTrue    BoolStr = "true"
	BoolStrFalse   BoolStr = "false"
	defaultTagName         = "binding"
)

var (
	phoneReg = regexp.MustCompile(`^1[23456789]\d{9}$`)
)

var _ binding.StructValidator = &DefaultValidator{}

func validPhone(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	return phoneReg.MatchString(val)
}
func validPhoneOrEmpty(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	return val == "" || phoneReg.MatchString(val)
}
func validBoolStr(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	return val == BoolStrTrue || val == BoolStrFalse
}
func (v *DefaultValidator) ValidateStruct(obj any) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyInit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}
func (v *DefaultValidator) Engine() any {
	v.lazyInit()
	return v.validate
}
func (v *DefaultValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName(defaultTagName)
		reg := v.validate.RegisterValidation
		_ = reg("validBoolStr", validBoolStr)
		_ = reg("phone", validPhone)
		_ = reg("phoneOrEmpty", validPhoneOrEmpty)
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
