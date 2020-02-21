package libgen

type BookInfo struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	MD5           string `json:"md5"`
	Year          string `json:"year"`
	FileExtension string `json:"extension"`
	FileSize      string `json:"filesize"`
	DownloadLink  string
}
