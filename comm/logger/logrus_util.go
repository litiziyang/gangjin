package logger

import "github.com/sirupsen/logrus"

func Info(logger *logrus.Logger, msg string, err error, data interface{}) {
	log(logger, msg, data, nil, "info")
}
func Panic(logger *logrus.Logger, msg string, err error, data interface{}) {
	log(logger, msg, data, nil, "Panic")
}
func Error(logger *logrus.Logger, msg string, err error, data interface{}) {
	log(logger, msg, data, nil, "Error")
}
func Debug(logger *logrus.Logger, msg string, err error, data interface{}) {
	log(logger, msg, data, nil, "Debug")
}

func log(log *logrus.Logger, msg string, data interface{}, err error, t string) {
	fields := logrus.Fields{"data": data}
	if err != nil {
		fields = logrus.Fields{"data": data, "err": err}
	}
	field := log.WithFields(fields)
	switch t {
	case "debug":
		field.Debug(msg)
	case "info":
		field.Info(msg)
	case "Error":
		field.Error(msg)
	case "panic":
		field.Panic(msg)

	}
}
