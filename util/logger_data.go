package util

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Its error ", entry.Level, entry.Message)
	return nil
}

func Log2File(message string) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	pathFile := fmt.Sprintf("/home/maulanazn/Devstack/blanja-api/log/blanjaapi-log-%q.log", time.Now())

	file, _ := os.OpenFile(pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	logger.SetLevel(logrus.TraceLevel)

	logger.Error(message)
}
