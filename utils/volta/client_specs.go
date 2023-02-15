package volta

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/saucesteals/mimic"
)

type BrowserClient struct {
	UserAgent string
	SecChUA   string
	M         *mimic.ClientSpec
}

var (
	browserClients = []BrowserClient{}
)

func init() {
	for i := 100; i <= 110; i++ {
		m, _ := mimic.Chromium(mimic.BrandChrome, strconv.Itoa(i))

		browserClients = append(browserClients, BrowserClient{
			UserAgent: fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", m.Version()),
			SecChUA:   m.ClientHintUA(),
			M:         m,
		})
	}
}

func RandomClient() BrowserClient {
	return browserClients[rand.Intn(len(browserClients)-1)]
}
