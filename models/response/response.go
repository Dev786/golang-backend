package response

type Response struct {
	Message string `json:"message"`
	Data    []byte `json:"data"`
	Success bool   `json:"success"`
}

func (r Response) FailedResponse(message string) []byte {
	r.Data = []byte{}
	r.Message = message
	r.Success = false
	// data, err := dataConversion.ParseStructToJson(r)
	// if err != nil {
	// 	return []byte{}
	// }
	return r.Data
}

function(r Response) ParseToJSON() []byte{

}

func (r Response) SuccessResponse(message string) []byte {
	r.Data = []byte{}
	r.Message = message
	r.Success = true
	// data, err := dataConversion.ParseStructToJson(r)
	// if err != nil {
	// return []byte{}
	// }
	return r.Data
}
