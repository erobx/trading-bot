package model

var (
	baseUrl  = "http://api.steampowered.com/"
)

type Request struct {
	config Config
}

type Config struct {
	baseUrl      string
	genInterface string
	method       string
	apiKey       string
	appId        string
}

func newConfig() Config {
	return Config{
		baseUrl: baseUrl,
	}
}

func NewRequest() Request {
	return Request{
		config: newConfig(),
	}
}

func (r Request) CombineUrl() string {
	url := r.config.baseUrl + r.config.genInterface + r.config.method + "key=" + r.config.apiKey +
		"&appid=" + r.config.appId + "&"

	return url
}

func (r Request) WithGenInterface(gi string) Request {
	r.config.genInterface = gi + "/"
	return r
}

func (r Request) WithMethod(method string) Request {
	r.config.method = method + "/v0001?"
	return r
}

func (r Request) WithApiKey(key string) Request {
	r.config.apiKey = key
	return r
}

func (r Request) WithAppId(id string) Request {
	r.config.appId = id
	return r
}