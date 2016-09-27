package logging

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

type apiLogHook struct {
	logEndpoint string
	httpClient  *http.Client
	host        string
}

func NewAPILogHook(logEndpoint, host string) logrus.Hook {
	return &apiLogHook{
		logEndpoint: logEndpoint,
		httpClient:  &http.Client{},
		host:        host,
	}
}

func (h *apiLogHook) Fire(entry *logrus.Entry) error {
	level := strings.ToUpper(entry.Level.String())

	if level == "PANIC" {
		level = "FATAL"
	}

	logData := make(map[string]interface{})
	logData["ts"] = time.Now().UTC().Format(time.RFC3339)
	logData["level"] = level
	logData["message"] = entry.Message
	logData["host"] = h.host

	for k, v := range entry.Data {
		logData[k] = v
	}

	data, err := json.Marshal(logData)

	if err != nil {
		return err
	}

	dataReader := bytes.NewReader(data)
	req, err := http.NewRequest("POST", h.logEndpoint, dataReader)
	if err != nil {
		return err
	}

	// Need to set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

// Levels returns the levels that this hook is listening to
func (h *apiLogHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
