package clients

import (
	"context"
	"net/http"
	"time"
)

func sendToSDL() {
	ctx := context.WithTimeout(context.Background(), 5*time.Second)
	req, _ := http.NewRequest("GET", "https://mota.cf", nil)
	req = req.WithContext(ctx)

	resp, _ := http.DefaultClient.Do(req)
}
