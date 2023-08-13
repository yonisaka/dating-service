package consts

const (
	ExtJpeg = "image/jpeg"
	ExtPng  = "image/png"
	ExtGif  = "image/gif"
)

var (
	ExtImageBase64Allowed = []string{
		ExtJpeg,
		ExtPng,
		ExtGif,
	}
)
