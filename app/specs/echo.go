package specs

type EchoReq struct {
	Message string `json:"message"`
}

type Resp struct {
	Message string `json:"message"`
}
