package conf

var (
	Env2url = map[string]string{
		//"development:boe:intranet": "http://10.94.93.158:6789",
		"development:boe:intranet": "http://apaas-faasinfra-dev.byted.org",
		"staging:boe:intranet":     "http://apaas-faasinfra-staging-boe.bytedance.net",
		"staging::intranet":        "https://apaas-faasinfra-staging.bytedance.net",
		"staging::extranet":        "https://apaas-faasinfra-staging.bytedance.com",
		"gray::intranet":           "https://apaas-faasinfra-gray.bytedance.net",
		"gray::extranet":           "https://apaas-faasinfra-gray.kundou.cn",
		"online::intranet":         "https://apaas-faasinfra.bytedance.net",
		"online::extranet":         "https://apaas-faasinfra.kundou.cn",
	}
)
