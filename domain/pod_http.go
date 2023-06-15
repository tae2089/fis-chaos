package domain

type PodHttpSpec struct {
	Rules []*PodHttpRule `json:"rules,omitempty"`
	TLS   *PodHttpTls    `json:"tls,omitempty"`
}

type PodHttpRule struct {
	Source   string         `json:"source,omitempty"`
	Target   string         `json:"target" binding:"required"`
	Port     int            `json:"port" binding:"required"`
	Selector ChaosSelectors `json:"selector" binding:"required"`
}

type PodHttpActions struct {
	Abort   bool           `json:"abort,omitempty"`
	Delay   string         `json:"delay,omitempty"`
	Patch   PodHttpPatch   `json:"patch,omitempty"`
	Replace PodHttpReplcae `json:"replace,omitempty"`
}

type PodHttpPatch struct {
	Body    PodHttpBody `json:"body,omitempty"`
	Headers [][]string  `json:"headers,omitempty"`
	Queries [][]string  `json:"queries,omitempty"`
}

type PodHttpBody struct {
	Type  string `json:"type" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type PodHttpTls struct {
	CaName          string `json:"caName" binding:"required"`
	CertName        string `json:"certName" binding:"required"`
	KeyName         string `json:"keyName" binding:"required"`
	SecretName      string `json:"secretName" binding:"required"`
	SecretNameSpace string `json:"secretNameSpace" binding:"required"`
}

type PodHttpReplcae struct {
	Body    string            `json:"body,omitempty"`
	Code    int               `json:"code,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Method  string            `json:"method,omitempty"`
	Path    string            `json:"path,omitempty"`
	Queries map[string]string `json:"queries,omitempty"`
}
