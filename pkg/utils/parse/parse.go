package parse

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type Stashes struct {
	reqBody map[string]interface{}
}

func (s *Stashes) NewRequestBody(c *gin.Context) {
	ByteBody, _ := c.GetRawData()
	if len(ByteBody) != 0 {
		err := json.Unmarshal(ByteBody, &s.reqBody)
		if err != nil {
			log.Error(err)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))
	}
}

// unused, should try to find a case to be able to use this, probably outside this package
func (s *Stashes) GetRequestBody(c *gin.Context) map[string]interface{} {
	c.ShouldBind(s.reqBody)

	return s.reqBody
}

// Convert json interface into map string interface
func ToMapStringInterface(source *interface{}) *map[string]interface{} {
	if sourceParsed, err := json.Marshal(source); err == nil {
		resultMap := make(map[string]interface{})
		if err := json.Unmarshal(sourceParsed, &resultMap); err == nil {
			return &resultMap
		}
	}
	return nil
}
