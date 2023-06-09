package parse

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewRequestBody(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = &http.Request{
		Body: ioutil.NopCloser(bytes.NewBufferString(`{"name":"test"}`)),
	}
	s := &Stashes{}
	s.NewRequestBody(c)

	body, _ := c.GetRawData()

	assert.Equal(t, `{"name":"test"}`, string(body))
	assert.NoError(t, c.Err())
	assert.NoError(t, c.Err())
}
