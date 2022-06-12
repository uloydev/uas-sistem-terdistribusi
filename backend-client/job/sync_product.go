package job

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sister-backend/app/model"
)

func SyncProduct() {
	var data model.WebResponse

	url := url.URL{
		Scheme: "http",
		Host:   "localhost:8693",
		Path:   "v1/product/sync",
	}
	resp, err := http.Get(url.String())

	if err != nil {
		fmt.Println("[Sync Error] : " + err.Error())
	}

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		fmt.Println("[Sync Error] : " + err.Error())
	}
	defer resp.Body.Close()

	if data.Code != http.StatusOK {
		fmt.Println("[Sync Error] : " + data.Data.(string))
	}

	fmt.Println("[Sync Success] : Product now synced with master")
}
