package models

// アプリケーションのVersion
type Version struct {
	Version string `json:"version"`
}

var Ver Version

func init() {
	Ver = Version{
		Version: "0.0.1",
	}
}
