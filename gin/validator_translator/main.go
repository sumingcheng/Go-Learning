package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
)

var trans ut.Translator

type Register struct {
	Username string `json:"username" binding:"required,max=10,min=5"`
	Password string `json:"password" binding:"required,max=10,min=5"`
}

func main() {
	r := gin.Default()

	if err := CreateTranslator("zh"); err != nil {
		fmt.Println("无法创建翻译器")
		return
	}

	r.POST("/register", func(c *gin.Context) {
		var register Register

		if err := c.ShouldBindJSON(&register); err != nil {
			e, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"code": 1,
					"msg":  err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code": 1001,
				"msg":  e.Translate(trans),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "success",
				"data": register,
			})
		}
	})

	r.Run(":8888")
}

//func removeStructName(msg string) string {
//	return msg
//
//}

func CreateTranslator(locale string) error {
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		fmt.Println("类型断言失败") // 添加了打印语句
		return fmt.Errorf("类型断言失败")
	}

	// 注册结构体的json标签名作为字段名
	validate.RegisterTagNameFunc(func(sf reflect.StructField) string {
		name := sf.Tag.Get("json")
		if name == "_" {
			return ""
		}
		return name
	})

	zhTranslation := zh.New()
	enTranslation := en.New()
	uniTranslations := ut.New(enTranslation, zhTranslation, enTranslation)

	translator, ok := uniTranslations.GetTranslator(locale)
	if !ok {
		fmt.Printf("找不到翻译器(%s)\n", locale) // 添加了打印语句
		return fmt.Errorf("找不到翻译器(%s)", locale)
	}

	switch locale {
	case "zh":
		if err := zhtranslations.RegisterDefaultTranslations(validate, translator); err != nil {
			fmt.Printf("注册中文翻译器失败: %v\n", err) // 添加了打印语句
			return fmt.Errorf("注册中文翻译器失败: %v", err)
		}
	case "en":
		if err := entranslations.RegisterDefaultTranslations(validate, translator); err != nil {
			fmt.Printf("注册英文翻译器失败: %v\n", err) // 添加了打印语句
			return fmt.Errorf("注册英文翻译器失败: %v", err)
		}
	default:
		fmt.Printf("不支持的语言类型(%s)\n", locale) // 添加了打印语句
		return fmt.Errorf("不支持的语言类型(%s)", locale)
	}

	trans = translator // 设置全局翻译器
	return nil
}
