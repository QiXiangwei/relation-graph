package library

type RelationResponse struct {
	ErrNo  int          `json:"errNo"`
	ErrStr string		`json:"errStr"`
	Data   interface{}  `json:"data"`
}

func BuildRelationResponse(errNo int, errStr string, data interface{}) RelationResponse {
	return RelationResponse{
		ErrNo:  errNo,
		ErrStr: errStr,
		Data:   data,
	}
}