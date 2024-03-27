package filemanager

type Handler struct {
	FileLimit int64
}

func (h Handler) Register() {
	h.upload(nil, nil)
}
