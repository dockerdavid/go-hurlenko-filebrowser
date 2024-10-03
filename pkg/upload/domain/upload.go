package upload

type File struct {
	Token    string `json:"token"`
	BaseURL  string `json:"baseURL"`
	Folder   string `json:"folder"`
	ShareURL string `json:"shareURL"`
	File     []byte `json:"file"`
	Ext      string `json:"ext"`
}
