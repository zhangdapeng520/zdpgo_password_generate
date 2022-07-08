package zdpgo_password_generate

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

/*
@Time : 2022/6/21 17:14
@Author : 张大鹏
@File : generator.go
@Software: Goland2021.3.1
@Description: 用于生成密码
*/

const (
	LengthWeak                uint = 6                                 // 弱密码长度
	LengthOK                  uint = 12                                // 一般密码长度
	LengthStrong              uint = 24                                // 强密码长度
	LengthVeryStrong          uint = 36                                // 高强密码长度
	DefaultLetterSet               = "abcdefghijklmnopqrstuvwxyz"      // 默认小写字母集合
	DefaultLetterAmbiguousSet      = "ijlo"                            // 默认相似字母集合
	DefaultNumberSet               = "0123456789"                      // 默认数字集合
	DefaultNumberAmbiguousSet      = "01"                              // 默认相似数字集合
	DefaultSymbolSet               = "!$%^&*()_+{}:@[];'#<>?,./|\\-=?" // 默认特殊符号集合
	DefaultSymbolAmbiguousSet      = "<>[](){}:;'/|\\,"                //默认转义符号集合
)

var (
	// DefaultConfig 默认密码配置
	DefaultConfig = Config{
		Length:                     LengthStrong,
		IncludeSymbols:             true,
		IncludeNumbers:             true,
		IncludeLowercaseLetters:    true,
		IncludeUppercaseLetters:    true,
		ExcludeSimilarCharacters:   true,
		ExcludeAmbiguousCharacters: true,
	}

	// ErrConfigIsEmpty 密码配置为空错误
	ErrConfigIsEmpty = errors.New("config is empty")

	// 默认密码生成器
	DefaultGenerator = New()
)

// Generator 密码生成器
type Generator struct {
	Config *Config
}

// Config 生成密码的配置
type Config struct {
	Length                     uint   `yaml:"length" json:"length"`                                             // 密码长度
	CharacterSet               string `yaml:"character_set" json:"character_set"`                               // 字符池子，从池子里面取密码字符
	IncludeSymbols             bool   `yaml:"include_symbols" json:"include_symbols"`                           // 是否包含标志，比如!"£*
	IncludeNumbers             bool   `yaml:"include_numbers" json:"include_numbers"`                           // 是否包含数字
	IncludeLowercaseLetters    bool   `yaml:"include_lowercase_letters" json:"include_lowercase_letters"`       // 是否包含小写字母
	IncludeUppercaseLetters    bool   `yaml:"include_uppercase_letters" json:"include_uppercase_letters"`       // 是否包含大写字母
	ExcludeSimilarCharacters   bool   `yaml:"exclude_similar_characters" json:"exclude_similar_characters"`     // 是否包含相似的字符，比如i1jIo0
	ExcludeAmbiguousCharacters bool   `yaml:"exclude_ambiguous_characters" json:"exclude_ambiguous_characters"` // 是否包含特殊转义符号<>{}[]()/|\
}

func New() *Generator {
	return NewWithConfig(&DefaultConfig)
}

// NewWithConfig 返回密码生成器
func NewWithConfig(config *Config) *Generator {
	// 配置
	g := &Generator{}
	if config.Length == 0 {
		config.Length = LengthStrong
	}
	if config.CharacterSet == "" {
		config.CharacterSet = buildCharacterSet(config)
	}
	g.Config = config

	// 返回
	return g
}

func buildCharacterSet(config *Config) string {
	var characterSet string
	if config.IncludeLowercaseLetters {
		characterSet += DefaultLetterSet
		if config.ExcludeSimilarCharacters {
			characterSet = removeCharacters(characterSet, DefaultLetterAmbiguousSet)
		}
	}

	if config.IncludeUppercaseLetters {
		characterSet += strings.ToUpper(DefaultLetterSet)
		if config.ExcludeSimilarCharacters {
			characterSet = removeCharacters(characterSet, strings.ToUpper(DefaultLetterAmbiguousSet))
		}
	}

	if config.IncludeNumbers {
		characterSet += DefaultNumberSet
		if config.ExcludeSimilarCharacters {
			characterSet = removeCharacters(characterSet, DefaultNumberAmbiguousSet)
		}
	}

	if config.IncludeSymbols {
		characterSet += DefaultSymbolSet
		if config.ExcludeAmbiguousCharacters {
			characterSet = removeCharacters(characterSet, DefaultSymbolAmbiguousSet)
		}
	}

	return characterSet
}

func removeCharacters(str, characters string) string {
	return strings.Map(func(r rune) rune {
		if !strings.ContainsRune(characters, r) {
			return r
		}
		return -1
	}, str)
}

// Generate generates one password with length set in the
// config
func (g Generator) Generate() (*string, error) {
	var generated string
	characterSet := strings.Split(g.Config.CharacterSet, "")
	max := big.NewInt(int64(len(characterSet)))

	for i := uint(0); i < g.Config.Length; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		generated += characterSet[val.Int64()]
	}
	return &generated, nil
}

// GenerateByLength 生成指定长度的命名
func (g Generator) GenerateByLength(length uint) string {
	g.Config.Length = length
	password, err := g.Generate()
	if err != nil {
		fmt.Println("生成密码失败", err)
		return ""
	}
	return *password
}

// GenerateByWeak 生成弱密码
func (g Generator) GenerateByWeak() string {
	return g.GenerateByLength(LengthWeak)
}

// GenerateByOK 生成普通密码
func (g Generator) GenerateByOK() string {
	return g.GenerateByLength(LengthOK)
}

// GenerateByStrong 生成强密码
func (g Generator) GenerateByStrong() string {
	return g.GenerateByLength(LengthStrong)
}

// GenerateByVeryStrong 生成超强密码
func (g Generator) GenerateByVeryStrong() string {
	return g.GenerateByLength(LengthVeryStrong)
}

// GenerateMany 批量生产密码
func (g Generator) GenerateMany(amount uint) ([]string, error) {
	var generated []string
	for i := uint(0); i < amount; i++ {
		str, err := g.Generate()
		if err != nil {
			return nil, err
		}

		generated = append(generated, *str)
	}
	return generated, nil
}

// GenerateWithLength generate one password with set length
func (g Generator) GenerateWithLength(length uint) (*string, error) {
	var generated string
	characterSet := strings.Split(g.Config.CharacterSet, "")
	max := big.NewInt(int64(len(characterSet)))
	for i := uint(0); i < length; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		generated += characterSet[val.Int64()]
	}
	return &generated, nil
}

// GenerateManyWithLength generates multiple passwords with set length
func (g Generator) GenerateManyWithLength(amount, length uint) ([]string, error) {
	var generated []string
	for i := uint(0); i < amount; i++ {
		str, err := g.GenerateWithLength(length)
		if err != nil {
			return nil, err
		}
		generated = append(generated, *str)
	}
	return generated, nil
}
