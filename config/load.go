package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/spf13/viper"
)

/*
	获取配置文件目录: 优先级 1 > 2 > 3 > 4
	1. 指定目录.指定文件名
	2. 指定目录.默认文件名
	3. 默认目录.指定文件名
	4. 默认目录.默认文件名
*/

var (
	// 默认目录列表 config conf data
	configDir = []string{"config", "conf", "data"}
	// 默认文件名列表 custom:个人 test:测试 production:正式
	configNames = []string{"custom", "test", "production"}
)

/** LoadConfing 加载解析配置
  参数:
  *       conf            interface{}	反序列化对象,必须为非空指针
  *       dir             string     	配置文件目录
  *       fileName        string     	配置文件名
  返回值:
  *       error   error
*/
func LoadConfing(conf interface{}, dir, fileName string) error {
	if err := mustNotNilPtr(conf); err != nil {
		return err
	}

	if dir == "" {
		// 使用默认配置
		return nil
	}

	return nil
}

// 读取配置文件
func readConfFile(dir, fileName string) (*viper.Viper, error) {
	v := viper.New()
	// if fileName != "" {
	// 	// 有后缀去掉
	// }

	return v, nil
}

/** existFilePath 指定路径是否存在
  参数:
  *       dir     string	非空路径 (支持相对和绝对路径)
  返回值:
  *       string  string	存在则返回绝对路径
  *       error   error
*/
func existFilePath(dir string) (string, error) {
	// 是否为绝对路径
	fp, err := func(dir string) (string, error) {
		if !filepath.IsAbs(dir) {
			fp, err := filepath.Abs(dir)
			if err != nil {
				return "", errors.New("获取系统路径失败,构建绝对路径失败")
			}

			return fp, nil
		}

		return dir, nil
	}(dir)

	if err != nil {
		return "", err
	}

	info, err := os.Stat(fp)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return "", errors.New("指定路径不存在")
	}

	if !info.IsDir() {
		return "", fmt.Errorf("[%s] 非目录", dir)
	}

	return fp, nil
}

// getDefaultFilePath 获取默认路径 依次判定: config -> conf -> data
func getDefaultFilePath() (string, error) {
	for _, dir := range configDir {
		fp, err := existFilePath(dir)
		if err != nil {
			continue
		}

		return fp, nil
	}

	return "", fmt.Errorf("未找到符合的默认目录 %v", configDir)
}

// 对象不为空,且非空指针
func mustNotNilPtr(conf interface{}) error {
	if reflect.TypeOf(conf).Kind() != reflect.Ptr {
		return errors.New("conf参数必须是指针")
	}

	if reflect.ValueOf(conf).IsNil() {
		return errors.New("conf不能为nil")
	}

	return nil
}
