package response

import "encoding/json"

func (response *Response) SendResponse() {
	jsonData, _ := json.Marshal(response.Data)

	response.ResponseWriter.Header().Set("Content-Type", "application/json")
	response.ResponseWriter.WriteHeader(response.StatusCode)
	response.ResponseWriter.Write(jsonData)
}
