package logger

import (
	"fmt"
	"github.com/heirko/go-contrib/logrusHelper"
	mate "github.com/heralight/logrus_mate"
	//_ "github.com/heralight/logrus_mate/hooks/file"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/txvier/dcoin/logger/hook"
	"io"
	"os"
	"strings"
	"time"
)

type RotatelogsConf struct {
	LinkName      string `json:"link-name"`
	Clock         string `json:"clock"`
	RotationTime  string `json:"rotation-time"`
	RotationCount int64  `json:"rotation-count"`
	MaxAge        string `json:"max-age"`
}

func (c RotatelogsConf) GetRotationTime() (d time.Duration) {
	if d, err := time.ParseDuration(c.RotationTime); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		return d
	}
	return
}

func (c RotatelogsConf) GetRotationCount() uint {
	return uint(c.RotationCount)

}

func (c RotatelogsConf) GetMaxAge() (d time.Duration) {
	if d, err := time.ParseDuration(c.MaxAge); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		return d
	}
	return
}

func (c RotatelogsConf) GetLinkName() string {
	return c.LinkName
}

func (c RotatelogsConf) GetClock() rotatelogs.Clock {
	if "UTC" == strings.ToUpper(c.Clock) {
		return rotatelogs.UTC
	}
	return rotatelogs.Local
}

func init() {

	// ########## Init Viper
	var viper = viper.New()

	viper.SetConfigName("logger") // name of config file (without extension), here we use some logrus_mate sample
	//viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// ########### End Init Viper

	logrus.SetReportCaller(true)

	// Read configuration
	//mate.RegisterWriter("rotatelogs", NewRotatelogsWriter)
	var c = logrusHelper.UnmarshalConfiguration(viper) // Unmarshal configuration from Viper
	logrusHelper.SetConfig(logrus.StandardLogger(), c) // for e.g. apply it to logrus default instance
	// ### End Read Configuration

}

func NewRotatelogsWriter(options mate.Options) (writer io.Writer, err error) {
	fmt.Println(options)
	var conf RotatelogsConf
	if err = options.ToObject(&conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return rotatelogs.New(conf.GetLinkName()+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(conf.GetLinkName()),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(conf.GetMaxAge()),             // 文件最大保存时间
		rotatelogs.WithRotationTime(conf.GetRotationTime()), // 日志切割时间间隔
		rotatelogs.WithClock(conf.GetClock()),
		rotatelogs.WithRotationCount(conf.GetRotationCount()),
	)
}
