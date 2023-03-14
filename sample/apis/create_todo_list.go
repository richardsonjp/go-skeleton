package apis

import (
	"fmt"

	"go-skeleton/pkg/utils/json"

	"github.com/go-resty/resty/v2"
)

func CreateTodoItem() {
	client := resty.New()
	url := "http://127.0.0.1:7000/v1/todos"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Content-Type":    "application/json",
			"Accept-Language": "id",
		}).
		SetBody(map[string]interface{}{
			"task":   "todo-context",
			"status": 1,
		}).
		Post(url)
	boom(err)

	body := map[string]interface{}{}
	err = json.Unmarshal(resp.Body(), &body)
	boom(err)

	str, err := json.MarshalIndent(body, "", " ")
	fmt.Printf("Body: %s\n", str)
	boom(err)
}
