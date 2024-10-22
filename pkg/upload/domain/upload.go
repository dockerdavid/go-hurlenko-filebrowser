package upload

type File struct {
	File []byte `json:"file"`
}

type BaseData struct {
	Token    string `json:"token"`
	BaseURL  string `json:"baseURL"`
	Folder   string `json:"folder"`
	Filename string `json:"filename"`
	ShareURL string `json:"shareURL"`
}
