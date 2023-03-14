// Synchronized package
package syncs

import (
	"fmt"
	"go-skeleton/config"
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/logs"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

// Go Routine Recovery. Execute Go Routine with panic handling to prevent system down due to panic
// and push panic log.
func GoRecover(f func()) {
	go func() {
		defer recoverPanic()
		f()
	}()
}

func recoverPanic() {
	if err := recover(); err != nil {
		if config.Config.System.Mode != gin.ReleaseMode {
			log.Println(fmt.Sprintf("[START-GORECOVER-PANIC]\n%s\n[END-GORECOVER-PANIC]", errors.GetStack(err)))
			return
		}

		fields := logs.Fields{
			"type_str":     "ERR-GORECOVER-PANIC",
			"mode":         config.Config.System.Mode,
			"error_type":   reflect.ValueOf(err).Type().String(),
			"error_string": errors.ToString(err),
			"error_stack":  errors.GetStack(err),
		}

		cl := logs.Log.WithFields(fields)
		logs.PushPanicLog(cl)
	}
}
