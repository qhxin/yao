package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/share"
)

var appPath string
var envFile string

var lang = os.Getenv("YAO_LANG")
var langs = map[string]string{
	"Start Engine":                          "启动象传应用引擎",
	"One or more arguments are not correct": "参数错误",
	"Application directory":                 "指定应用路径",
	"Environment file":                      "指定环境变量文件",
	"Help for yao":                          "显示命令帮助文档",
	"Show app configure":                    "显示应用配置信息",
	"Update database schema":                "更新数据库结构",
	"Execute process":                       "更新数据库结构",
	"Show version":                          "显示当前版本号",
}

// L 多语言切换
func L(words string) string {
	if lang == "" {
		return words
	}

	if trans, has := langs[words]; has {
		return trans
	}
	return words
}

var rootCmd = &cobra.Command{
	Use:   share.BUILDNAME,
	Short: "Yao App Engine",
	Long:  `Yao App Engine`,
	Args:  cobra.MinimumNArgs(1),
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(os.Stderr, L("One or more arguments are not correct"), args)
		os.Exit(1)
	},
}

// 加载命令
func init() {
	rootCmd.AddCommand(
		versionCmd,
		migrateCmd,
		inspectCmd,
		startCmd,
		runCmd,
	)
	rootCmd.SetHelpCommand(helpCmd)
	rootCmd.PersistentFlags().StringVarP(&appPath, "app", "a", "", L("Application directory"))
	rootCmd.PersistentFlags().StringVarP(&envFile, "env", "e", "", L("Environment file"))
}

// Execute 运行Root
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Boot 设定配置
func Boot() {

	if envFile == "" && appPath != "" {
		config.Conf = config.LoadFrom(filepath.Join(appPath, ".env"))
		config.Conf.Root = appPath
		return
	}

	if envFile != "" && appPath == "" { // 指定环境变量文件
		// config.SetEnvFile(envFile)
		config.Conf = config.LoadFrom(envFile)
	}

	if envFile != "" && appPath != "" {
		config.Conf = config.LoadFrom(envFile)
		config.Conf.Root = appPath
	}
}
