package handler

type Config struct {
	ApiKey    string `required:"true"`
	ApiSecret string `required:"true"`
	RestUrl   string `default:"https://api.bitvavo.com/v2"`
	WsUrl     string `default:"wss://ws.bitvavo.com/v2/"`
}
