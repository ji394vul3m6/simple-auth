package data

type Return struct {
	ReturnCode    int         `json:"ret_code"`
	ReturnMessage string      `json:"ret_msg"`
	ReturnObj     interface{} `json:"result"`
}

type ListReturn struct {
	ReturnCode    int           `json:"ret_code"`
	ReturnMessage string        `json:"ret_msg"`
	ReturnObj     []interface{} `json:"results"`
}
