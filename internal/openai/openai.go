package openai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type OpenAIRequest struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	Temperature      int      `json:"temperature"`
	MaxTokens        int      `json:"max_tokens"`
	TopP             int      `json:"top_p"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	Stop             []string `json:"stop"`
}

func ChatAPI() {
	host := "https://api.openai.com/v1/completions"

	params := OpenAIRequest{
		Model:            "text-davinci-003",
		Prompt:           "Q:How open the acount about blockchain?\n A:",
		Temperature:      0,
		MaxTokens:        100,
		TopP:             1,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.0,
		Stop:             []string{"\n"},
	}
	jsons, _ := json.Marshal(params)
	bodyStr := strings.NewReader(string(jsons))

	// 发送 HTTP GET 请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", host, bodyStr)
	if err != nil {
		// handle error
		return
	}

	// 设置HTTP请求的header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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")
	req.Header.Set("Cookie", "__UONE__CSRF__TOKEN__=VNZFEiVPRq0TeRAq0A2DGGJS; xlly_s=1; cna=N6pvHAWGcXUCASp4Suq1/IAe; SSO_EMPID_HASH_V2=b8fe687b08ffe4e7d17af8c0a57db006; SSO_BU_HASH_V2=a619f8465211eef29f6cefe837ebc3c1; isg=BOXl1g-S6gzvuA7Q60yenxI79KcfIpm0Z3qrxefLoJwr_gFwrXAohqNfjGKIfrFs; l=fBMrAf9RT8gFzd-ZBO5ahurza779iIOfGsPzaNbMiIEGa6616FOSpOCeX8QDJdtjgTf0dexy58T6YdUeW4zT-PkDBeYQOAZPuDJ9-bpU-L5..; tfstk=c9kfBJcn7VHPQ1vGnrtybn3Vy5yCZ0C_vsZicEvBgmzva7ifiAWURsQQsieaXu1..; SSO_REFRESH_TOKEN=9cab76c16291466ba5ea3658e9f7ed000db26100; EGG_SESS=3WXCjuDrFY9xiv0wTq0vcAmjAw0nPPI6w054NrebxPNxF-bKoWdV0ofncositRnQ3ZlTL_pifwwp65paocEC_Dd2yoSjdfhAUBiT5kimU_S2SyN_fRjhc9C8b4lw6Er8EegGw6lAA4hEI14-FtMUIa-1v82r6zY-2MO2QkHyOpE9b6gF3cffwK5UX8wXDAzco-IJlldpFTkz2b0hqJPZYrfszdrft12jX0OlA_4krD6ghUYT_jjfvPHFXbk1Oh4LrMms5mAOrO78Wcls_1Gw8WxCQ_43UWgK1RT96KZ2jkpbQZXTZy3GDK09-QH4To1-jl9O-p9RCw6bd_sGJp1pgcWgTmP6q4rKMb9hjX7XBOKUt8dV43fRlz1X0iIYY_AGQmGJZ8tmEbobgbpgM1YuyLuWchi5H0Gq3B4FI-Oz-D1wm59_wcLuXMchSwTdqgC5tzp5uDPb8kQz5ntiVxsVn8GGWaNH-izigKzUFgag5FrCe1JHa9bK-wEYnFOsR0DsZA8zfYg1fGDs6eB9h3fcgtil3DLB1-N-s0W7f_75lUrR9iJc08U1C_2vWfLy8AYjKDF0AR49I77YBdZQKoedMLUwQTr_56AYgwNExCM8FjoCuWspKfUdHTCepIPkhIDriqhnmrJAoPcVgbjhXld-mcgDVvvnaifIU1j58FFLcqD8JFo8L2vCWfEOOyBCcjHE6x_fwJ-ODz5VnpitM00tgqfm3KqGaB0po6TK34oabuex_jwYDZK7rgZKJr1VG5Vk3vwH-kEpH7BD7BV_Gakol6PIYwtn0fdP-7tXMxH0UJNsPddqn77NVLnLxDGHxmHx6vFu4qZoJPgZ-c8c4WtbMCLasatsT9Gcm0lVZiH7en06VWuNSOiCfCENOpPP_TtCzvJUDmqwNYSf7lMS4zUqR5GdwiHA4dRdCMlbsca53X0-YtGKrij-AkdCqJLNXo6x-_ZUnqEgC-yw0mB8k30-3QHec4wrG4gzzccSUFS3PPSLXkjQP_ObV4ZTtLgCJ2DdRAbXu-yRhilZoKW8j9mTgHKuynEEyumtvSGZCdrvddMn5UxLUzpNdFC3qgIl4QT-Uoc6aRSEAaYuKHh9i8lY2gIkj9LTCuqZXeXL6j2rIJn93lyvfsOGCuR6gMHDNZkwzzZy-QqfIGDveRB-unolM0nYdMOQRy1RLEg26kJKzi_pP0lRogz7lqFvCoyTiUlaZEWEbi9Q1cZupmN16dH3kbgFMYsdQD6uPmvpmgdGdS1wYrVo7I0sohcmo9aitOkJZ-JALSvIoF1usuwzAiyGI8uaPEvNW-nzNaAIA0JK6nMEMgNRTayiV28jVL56m6eRmIorFlobsEwHhINFlFfwNHj7ATUVY02NsF-tQGYuJtpG82WDCJrxy8U9AW0sH6Oy")
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		return
	}

	defer resp.Body.Close()

	// 读取请求返回的数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
