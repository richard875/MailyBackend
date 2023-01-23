package structure

type IpLocationInfo struct {
	IP           string  `json:"ip"`
	IsEU         bool    `json:"is_eu"`
	City         string  `json:"city"`
	Region       string  `json:"region"`
	RegionCode   string  `json:"region_code"`
	RegionType   string  `json:"region_type"`
	CountryName  string  `json:"country_name"`
	CountryCode  string  `json:"country_code"`
	ContinetName string  `json:"continent_name"`
	ContinetCode string  `json:"continent_code"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Postal       string  `json:"postal"`
	CallingCode  string  `json:"calling_code"`
	Flag         string  `json:"flag"`
	EmojiFlag    string  `json:"emoji_flag"`
	EmojiUnicode string  `json:"emoji_unicode"`
	ASN          struct {
		ASN    string `json:"asn"`
		Name   string `json:"name"`
		Domain string `json:"domain"`
		Route  string `json:"route"`
		Type   string `json:"type"`
	} `json:"asn"`
	Languages []struct {
		Name   string `json:"name"`
		Native string `json:"native"`
		Code   string `json:"code"`
	} `json:"languages"`
	Currency struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Symbol string `json:"symbol"`
		Native string `json:"native"`
		Plural string `json:"plural"`
	} `json:"currency"`
	TimeZone struct {
		Name        string `json:"name"`
		Abbr        string `json:"abbr"`
		Offset      string `json:"offset"`
		IsDST       bool   `json:"is_dst"`
		CurrentTime string `json:"current_time"`
	} `json:"time_zone"`
	Threat struct {
		IsTor           bool          `json:"is_tor"`
		IsICloudRelay   bool          `json:"is_icloud_relay"`
		IsProxy         bool          `json:"is_proxy"`
		IsDatacenter    bool          `json:"is_datacenter"`
		IsAnonymous     bool          `json:"is_anonymous"`
		IsKnownAttacker bool          `json:"is_known_attacker"`
		IsKnownAbuser   bool          `json:"is_known_abuser"`
		IsThreat        bool          `json:"is_threat"`
		IsBogon         bool          `json:"is_bogon"`
		Blocklists      []interface{} `json:"blocklists"`
	} `json:"threat"`
	Count string `json:"count"`
}
