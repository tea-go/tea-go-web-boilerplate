package app

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func getCallStack(skip int) string {
	buf := new(bytes.Buffer)
	for i := skip; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "\n%s:%d", file, line)
	}
	return buf.String()
}

// Transaction transaction middleware
func Transaction(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tx := db.Begin()

		rs := GetRequestScope(ctx)
		rs.SetTx(tx)

		defer func() {
			e := recover()

			if e != nil || rs.Rollback() {
				tx.Rollback()
			} else {
				tx.Commit()
			}

			if e != nil {
				stack := getCallStack(4)
				httprequest, _ := httputil.DumpRequest(ctx.Request, false)
				rs.Infof("[Recovery] %s panic recovered:\n%s\n%s\n%s%s", time.Now(), string(httprequest), e, stack)

				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": "Internal Server Error",
				})
			}
		}()

		ctx.Next()
	}
}
