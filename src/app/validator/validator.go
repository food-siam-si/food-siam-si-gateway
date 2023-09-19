package validator

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/en"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	v     *validator.Validate
	trans ut.Translator
}

type IValidator interface {
	Validate(i interface{}) *[]string
}

func NewValidator() IValidator {
	v := validator.New()

	english := en.New()
	uni := ut.New(english, english)

	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(v, trans)

	return &Validator{v, trans}
}

func (v *Validator) Validate(i interface{}) *[]string {
	err := v.v.Struct(i)

	if err != nil {
		validatorErrs := err.(validator.ValidationErrors)
		errs := []string{}

		for _, e := range validatorErrs {
			translatedErr := fmt.Errorf(e.Translate(v.trans))
			errs = append(errs, translatedErr.Error())
		}

		return &errs
	}

	return nil
}
