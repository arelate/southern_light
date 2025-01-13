package vangogh_integration

import "encoding/xml"

type ValidationChunk struct {
	XMLName xml.Name `xml:"chunk"`
	ID      int      `xml:"id,attr"`
	From    int      `xml:"from,attr"`
	To      int      `xml:"to,attr"`
	Method  string   `xml:"method,attr"`
	Value   string   `xml:",innerxml"`
}

type ValidationFile struct {
	XMLName             xml.Name          `xml:"file"`
	Name                string            `xml:"name,attr"`
	Available           int               `xml:"available,attr"`
	NotAvailableMessage string            `xml:"notavailablemsg,attr"`
	MD5                 string            `xml:"md5,attr"`
	Chunks              int               `xml:"chunks,attr"`
	Timestamp           string            `xml:"timestamp,attr"`
	TotalSize           int               `xml:"total_size,attr"`
	ValidationChunks    []ValidationChunk `xml:"chunk"`
}
