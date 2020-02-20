package libgen

// id = —the LibGen ID
// title = —the title of the text
// volumeinfo = —the volume number, if the text is part of a multi-volume series
// series = —the series that the text is part of
// periodical
// author = —the author of the text
// year = —the publication date of the text
// edition = —the edition of the text
// publisher = —the publisher of the text
// city = —the location of the publisher
// pages = —the number of pages in the text
// language = —the language of the text
// topic = —A number corresponding to the topic of the text; for example, 130 is “Mathematics/Logic”
// library
// issue
// identifier = —the text’s short and long International Standard Book Numbers (not necessarily in that order)
// issn = —the text’s International Standard Serial Number
// asin = —the text’s Amazon Standard Identification Number
// udc = —the text’s Universal Decimal Classification number
// lbc
// ddc = —the text’s Dewey Decimal Classification number
// lcc = —the text’s Library of Congress Classification number
// doi = —the file’s Digital Object Identifier
// googlebookid = —the text’s Google Books ID
// openlibraryid = —the text’s Open Library ID
// commentary
// dpi
// color
// cleaned
// orientation
// paginated = —the text is paginated (1) or not (0)
// scanned = —the text is scanned from a physical copy (1) or not (0)
// bookmarked = —the text has bookmarks (1) or not (0)
// searchable = —the text is searchable (1) or not (0)
// filesize = —the size of the file in bytes
// extension = —the extension of the file (.pdf, .epub, .mobi, etc.)
// md5 = —the MD5 hash of the file
// crc32 = —the file’s CRC32 checksum
// edonkey = —the file’s eDonkey hash
// aich = —the text’s eMule file hash
// sha1 = —the file’s SHA-1 hash
// tth = —the file’s Tiger tree hash
// generic
// filename = —the name of the file in the LibGen database, in the form directory/md5. The directory name is the text’s LibGen ID rounded to the nearest thousand, and the MD5 hash is in lowercase. (The directory that each file is located in is also included in the file name.)
// visible
// locator = —As far as I can tell, this is the file path of the original file on the machine of whoever uploaded it.
// local
// timeadded = —the date/time when the text was added to the database, formatted as YYYY-MM-DD HH:MM:SS
// timelastmodified = —the date/time when the text’s entry in the database was edited, formatted as YYYY-MM-DD HH:MM:SS
// coverurl = —the path to the cover image for the text: the filename followed by a lowercase letter (there’s a function to determine the letter for each cover, but I don’t know enough PHP to understand it).
