package volta

import (
	"net/http"
	"time"
)

type VoltaClient struct {
	Client *http.Client
}

func NewClient() VoltaClient {
	return VoltaClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (V VoltaClient) ReqDo(req http.Request, additionalHeaders map[string]string) (*http.Response, error) {
	brclient := RandomClient()
	req.Header = http.Header{
		"sec-ch-ua":          {brclient.SecChUA},
		"sec-ch-ua-mobile":   {"?0"},
		"user-agent":         {brclient.UserAgent},
		"content-type":       {"application/json"},
		"accept":             {"*/*"},
		"sec-ch-ua-platform": {`"Windows"`},
		"sec-fetch-site":     {"same-site"},
		"sec-fetch-mode":     {"cors"},
		"sec-fetch-dest":     {"empty"},
		"accept-encoding":    {"gzip, deflate, br"},
		"accept-language":    {"en-US,en;q=0.9"},
		"x-api-key":          {"u74w38X44fa7m3calbsu69blJVcC739z8NWJggVv"},
	}
	for k, v := range additionalHeaders {
		req.Header.Set(k, v)
	}
	return V.Client.Do(&req)
}
