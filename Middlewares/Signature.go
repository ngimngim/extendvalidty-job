package Middlewares

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func SignRequest(context *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: context.Writer}
		context.Writer = w

		hash := MD5(w.body.String())
		context.Header("Signature", hash)

		context.Next()
		fmt.Println("Response body: \"" + w.body.String()+"\"")

	}

