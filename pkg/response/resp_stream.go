package response

type StreamVideoResponse struct {
	Base
	VideoName string `json:"videoName"`
	Source    string `json:"source"`
}
