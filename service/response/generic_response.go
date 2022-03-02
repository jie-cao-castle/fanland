package response

type GenericResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
