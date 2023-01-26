package functions

import (
	"encoding/json"
	"maily/go-backend/src/structure"
	"net/http"
	"os"
)

func GetIpData(ip string) structure.IpLocationInfo {
	// Setup IP client and load .env
	//_ = godotenv.Load(".env")
	apiKey := os.Getenv("IP_ADDRESS_API_KEY")
	url := "https://api.ipdata.co/" + ip + "?api-key=" + apiKey

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	var locationInfo structure.IpLocationInfo
	json.NewDecoder(resp.Body).Decode(&locationInfo)

	return locationInfo
}
