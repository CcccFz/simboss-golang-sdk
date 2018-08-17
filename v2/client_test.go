package simboss

import (
	"testing"
	"strconv"
	"time"
	"net/url"
	"encoding/json"
)

func TestClient_sign(t *testing.T) {
	const appId string = "1111"
	const  appSecret = "2222"
	client := NewClient(appId, appSecret)
	data := url.Values{}
	data.Set("iccid", "1001")
	data.Set("type", "cmcc")
	data.Set("appid", appId)
	data.Set("timestamp", strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	sign := client.sign(data)
	t.Log(sign)
}

func TestResponse_Unmarshal(t *testing.T) {
	resp := Response{
		Code:    "0",
		Message: "",
		Data:    Pool{Id: 1001},
		Detail:  "",
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}

	nResp := Response{}

	if err := json.Unmarshal(respBytes, &nResp); err != nil {
		t.Fatal(err)
	}

	dataBytes, err := json.Marshal(nResp.Data)
	if err != nil {
		t.Fatal(err)
	}

	detail := Pool{}

	if err := json.Unmarshal(dataBytes, &detail); err != nil {
		t.Fatal(err)
	}

	if detail.Id != 1001 {
		t.Error("id should be 1001")
	}
}