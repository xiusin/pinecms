package common

import "io"

type Uploader interface {
	Upload(storageName string, LocalFile io.Reader) (string, error)
}
