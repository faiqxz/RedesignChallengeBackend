package models

type DownloadableFile struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	FileName    string `json:"fileName"`
	Description string `json:"description"`
	FileURL     string `json:"fileURL"`
	FileSize    int64  `json:"fileSize"`
	FileType    string `json:"fileType"`
}
