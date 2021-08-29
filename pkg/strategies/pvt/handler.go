package pvt

import (
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func Handler(w http.ResponseWriter, r *http.Request) {
	Logger.Info("request received", zap.Any("url", r.URL))
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	Logger.Info("request data", zap.Any("data", string(b)))
}
