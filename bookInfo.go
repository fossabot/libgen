package libgen

type BookInfo struct {
	ID               int64  `json:"id,string"`        //—the LibGen ID
	Title            string `json:"title"`            //—the title of the text
	VolumeInfo       string `json:"volumeinfo"`       //—the volume number, if the text is part of a multi-volume series
	Series           string `json:"series"`           //—the series that the text is part of
	Author           string `json:"author"`           //—the author of the text
	Year             string `json:"year"`             //—the publication date of the text
	Edition          string `json:"edition"`          //—the edition of the text
	Publisher        string `json:"publisher"`        //—the publisher of the text
	City             string `json:"city"`             //—the location of the publisher
	Pages            string `json:"pages"`            //—the number of pages in the text
	Language         string `json:"language"`         //—the language of the text
	Topic            string `json:"topic"`            //—A number corresponding to the topic of the text; for example, 130 is “Mathematics/Logic”
	Identifier       string `json:"identifier"`       //—the text’s short and long International Standard Book Numbers (not necessarily in that order)
	Issn             string `json:"issn"`             //—the text’s International Standard Serial Number
	Asin             string `json:"asin"`             //—the text’s Amazon Standard Identification Number
	Udc              string `json:"udc"`              //—the text’s Universal Decimal Classification number
	GoogleBookID     string `json:"googlebookid"`     //—the text’s Google Books ID
	OpenLibraryID    string `json:"openlibraryid"`    //—the text’s Open Library ID
	Paginated        string `json:"paginated"`        //—the text is paginated (1) or not (0)
	Scanned          string `json:"scanned"`          //—the text is scanned from a physical copy (1) or not (0)
	Bookmarked       string `json:"bookmarked"`       //—the text has bookmarks (1) or not (0)
	Searchable       string `json:"searchable"`       //—the text is searchable (1) or not (0)
	Filesize         string `json:"filesize"`         //—the size of the file in bytes
	Extension        string `json:"extension"`        //—the extension of the file (.pdf, .epub, .mobi, etc.)
	MD5              string `json:"md5"`              //—the MD5 hash of the file
	Crc32            string `json:"crc32"`            //—the file’s CRC32 checksum
	Edonkey          string `json:"edonkey"`          //—the file’s eDonkey hash
	Aich             string `json:"aich"`             //—the text’s eMule file hash
	Sha1             string `json:"sha1"`             //—the file’s SHA-1 hash
	Tth              string `json:"tth"`              //—the file’s Tiger tree hash
	Filename         string `json:"filename"`         //—the name of the file in the LibGen database, in the form directory/md5. The directory name is the text’s LibGen ID rounded to the nearest thousand, and the MD5 hash is in lowercase. (The directory that each file is located in is also included in the file name.)
	Locator          string `json:"locator"`          //—As far as I can tell, this is the file path of the original file on the machine of whoever uploaded it.
	Timeadded        string `json:"timeadded"`        //—the date/time when the text was added to the database, formatted as YYYY-MM-DD HH:MM:SS
	TimeLastModified string `json:"timelastmodified"` //—the date/time when the text’s entry in the database was edited, formatted as YYYY-MM-DD HH:MM:SS
	Coverurl         string `json:"coverurl"`         //—the path to the cover image for the text: the filename followed by a lowercase letter (there’s a function to determine the letter for each cover, but I don’t know enough PHP to understand it).

	DownloadPageURL string
}
