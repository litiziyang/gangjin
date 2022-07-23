package logger

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	sysLog "log"
	"os"
	"strings"
	"time"
)

var serverName string

type esHook struct {
	cmd string
	cli *elastic.Client
}

func (hook *esHook) Fire(entry *logrus.Entry) error {
	doc := newlog(entry)
	doc["cmd"] = hook.cmd
	go hook.sendEs(doc)
	return nil
}
func (hook *esHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

type config struct {
	LogLevel   string   // 日志级别
	EsAddress  []string //ES addr
	EsUser     string   //ES user
	EsPassword string   //ES password

}

func SetupLog(elasticsearch string, LogLevel string, srvName string) error {
	if LogLevel == "" {
		LogLevel = "debug"
	}
	c := config{
		LogLevel:   "error",
		EsAddress:  []string{fmt.Sprintf("http://%s/", elasticsearch)},
		EsUser:     "",
		EsPassword: "",
	}
	logLvl, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		return err
	}
	serverName = srvName
	logrus.SetLevel(logLvl)
	logrus.SetReportCaller(true)
	if elasticsearch != "" {
		Hook := sendEsclien(c)
		logrus.AddHook(Hook)
	}
	return nil
}

func sendEsclien(config config) *esHook {
	es, err := elastic.NewClient(
		elastic.SetURL(config.EsAddress...),
		elastic.SetBasicAuth(config.EsUser, config.EsPassword),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(15*time.Second),
		elastic.SetErrorLog(sysLog.New(os.Stderr, "ES:", sysLog.LstdFlags)),
		elastic.SetInfoLog(sysLog.New(os.Stdout, "ES:", sysLog.LstdFlags)),
	)
	if err != nil {
		sysLog.Fatal("es连接失败 ", err)
	}
	return &esHook{cli: es, cmd: strings.Join(os.Args, " ")}
}

//文件分割
func (m *appLogDocModel) indexName() string {
	return "jianlai.light-speak.com-" + time.Now().Local().Format("2006-01-02")
}

type appLogDocModel map[string]interface{}

//创建log
func newlog(entry *logrus.Entry) appLogDocModel {
	ins := map[string]interface{}{}
	for kk, vv := range entry.Data {
		ins[kk] = vv
	}
	ins["time"] = time.Now().Local()
	ins["level"] = entry.Level
	ins["message"] = entry.Message
	ins["caller"] = fmt.Sprintf("%s:%d  %#v", entry.Caller.File, entry.Caller.Line, entry.Caller.Func)
	ins["serverName"] = serverName
	return ins
}

func (hook *esHook) sendEs(doc appLogDocModel) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("发送日志到elasticsearch失败: ", r)
		}
	}()
	_, err := hook.cli.Index().Index(doc.indexName()).Type("_doc").BodyJson(doc).Do(context.Background())
	if err != nil {
		sysLog.Println(err)
	}
}
