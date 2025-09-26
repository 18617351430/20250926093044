package services

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"anti-fake-system/utils"
)

// RuleConfig 防伪码规则配置
type RuleConfig struct {
	Prefix       *PrefixConfig    `json:"prefix,omitempty"`     // 前置位配置
	MerchantCode string           `json:"merchant_code"`        // 商户标识码
	RandomNum    *RandomNumConfig `json:"random_num,omitempty"` // 随机数配置
	BatchCode    string           `json:"batch_code"`           // 批次标识
	FixedNum     *FixedNumConfig  `json:"fixed_num,omitempty"`  // 指定数字配置
	Sequence     *SequenceConfig  `json:"sequence"`             // 序号配置
	Separator    string           `json:"separator,omitempty"`  // 分隔符
	TotalLength  int              `json:"total_length"`         // 总长度限制
}

type PrefixConfig struct {
	Length  int    `json:"length"`  // 长度(1-5位)
	Content string `json:"content"` // 内容(支持数字、字母、特定符号)
}

type RandomNumConfig struct {
	Length   int    `json:"length"`   // 长度(1-8位)
	Position string `json:"position"` // 位置: before_merchant, after_merchant, before_batch, after_batch
}

type FixedNumConfig struct {
	Length  int    `json:"length"`  // 长度(1-6位)
	Content string `json:"content"` // 固定序列
}

type SequenceConfig struct {
	Length int `json:"length"` // 长度(最大9位)
	Start  int `json:"start"`  // 起始序号
}

// CodeGenerator 防伪码生成器
type CodeGenerator struct {
	ruleConfig    *RuleConfig
	encryptionKey []byte
}

// NewCodeGenerator 创建防伪码生成器
func NewCodeGenerator(ruleConfig *RuleConfig, encryptionKey []byte) *CodeGenerator {
	return &CodeGenerator{
		ruleConfig:    ruleConfig,
		encryptionKey: encryptionKey,
	}
}

// GenerateSingle 生成单个防伪码
func (g *CodeGenerator) GenerateSingle(sequence int) (string, error) {
	var parts []string

	// 生成前置位
	if g.ruleConfig.Prefix != nil {
		parts = append(parts, g.generatePrefix())
	}

	// 添加商户标识码
	parts = append(parts, g.ruleConfig.MerchantCode)

	// 生成随机数（根据位置）
	if g.ruleConfig.RandomNum != nil {
		randomNum, err := g.generateRandomNumber(g.ruleConfig.RandomNum.Length)
		if err != nil {
			return "", err
		}

		// 根据位置插入随机数
		switch g.ruleConfig.RandomNum.Position {
		case "before_merchant":
			parts = append([]string{randomNum}, parts...)
		case "after_merchant":
			if len(parts) > 1 {
				parts = append(parts[:1], append([]string{randomNum}, parts[1:]...)...)
			} else {
				parts = append(parts, randomNum)
			}
		case "before_batch":
			// 在批次标识前插入
			parts = append(parts, randomNum)
		case "after_batch":
			// 在批次标识后插入
			parts = append(parts, randomNum)
		}
	}

	// 添加批次标识
	parts = append(parts, g.ruleConfig.BatchCode)

	// 添加指定数字
	if g.ruleConfig.FixedNum != nil {
		parts = append(parts, g.ruleConfig.FixedNum.Content)
	}

	// 生成序号
	sequenceStr := fmt.Sprintf("%0"+strconv.Itoa(g.ruleConfig.Sequence.Length)+"d", sequence)
	parts = append(parts, sequenceStr)

	// 组合所有部分
	code := strings.Join(parts, g.ruleConfig.Separator)

	// 检查总长度
	if len(code) > g.ruleConfig.TotalLength {
		return "", fmt.Errorf("生成的防伪码长度超过限制: %d > %d", len(code), g.ruleConfig.TotalLength)
	}

	return code, nil
}

// GenerateBatch 批量生成防伪码
func (g *CodeGenerator) GenerateBatch(startSeq, count int) ([]string, error) {
	codes := make([]string, 0, count)

	for i := 0; i < count; i++ {
		code, err := g.GenerateSingle(startSeq + i)
		if err != nil {
			return nil, err
		}
		codes = append(codes, code)
	}

	return codes, nil
}

// EncryptRuleConfig 加密规则配置
func (g *CodeGenerator) EncryptRuleConfig() (string, error) {
	configJSON, err := json.Marshal(g.ruleConfig)
	if err != nil {
		return "", err
	}

	return utils.Encrypt(configJSON, g.encryptionKey)
}

// DecryptRuleConfig 解密规则配置
func (g *CodeGenerator) DecryptRuleConfig(encryptedConfig string) error {
	decrypted, err := utils.Decrypt(encryptedConfig, g.encryptionKey)
	if err != nil {
		return err
	}

	return json.Unmarshal(decrypted, &g.ruleConfig)
}

// generatePrefix 生成前置位
func (g *CodeGenerator) generatePrefix() string {
	if g.ruleConfig.Prefix == nil {
		return ""
	}

	// 如果内容长度不足，用0填充
	prefix := g.ruleConfig.Prefix.Content
	if len(prefix) < g.ruleConfig.Prefix.Length {
		prefix = strings.Repeat("0", g.ruleConfig.Prefix.Length-len(prefix)) + prefix
	}

	return prefix[:g.ruleConfig.Prefix.Length]
}

// generateRandomNumber 生成随机数字
func (g *CodeGenerator) generateRandomNumber(length int) (string, error) {
	if length <= 0 {
		return "", nil
	}

	max := big.NewInt(1)
	for i := 0; i < length; i++ {
		max.Mul(max, big.NewInt(10))
	}

	randomNum, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	// 格式化为固定长度
	format := "%0" + strconv.Itoa(length) + "d"
	return fmt.Sprintf(format, randomNum), nil
}

// ValidateCode 验证防伪码格式
func (g *CodeGenerator) ValidateCode(code string) bool {
	// 检查总长度
	if len(code) != g.ruleConfig.TotalLength {
		return false
	}

	// 检查各部分是否符合规则
	// 这里可以根据具体规则实现更详细的验证逻辑
	return true
}
