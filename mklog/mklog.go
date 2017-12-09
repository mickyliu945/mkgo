package mklog

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"mkgo/mkconfig"
)

var Logger *zap.Logger

func Init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	initLogger(mkconfig.Config.MLGO.Log.Path,
		mkconfig.Config.MLGO.Log.Level,
		mkconfig.Config.MLGO.Debug)
}

func initLogger(path string, level string, isDebug bool) {
	var js string
	if isDebug {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stdout"]
      }`, level)
	} else {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["%s"],
      "errorOutputPaths": ["%s"]
      }`, level, path, path)
	}

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
}
