package file

import "time"

const (
	infoField         = "info"
	announceField     = "announce"
	announceListField = "announce-list"
	creationDateField = "creation date"
	commentField      = "comment"
	createdByField    = "created by"
	encodingField     = "encoding"

	infoPieceLengthField = "piece length"
	infoPiecesField      = "pieces"
	infoPrivateField     = "private"

	singleFileInfoNameField   = "name"
	singleFileInfoLengthField = "length"
	singleFileInfoMD5SumField = "md5sum"

	multipleFileInfoNameField  = "name"
	multipleFileInfoFilesField = "files"

	multipleFileInfoFileLengthField = "length"
	multipleFileInfoFileMD5SumField = "md5sum"
	multipleFileInfoFilePathField   = "path"
)

type TorrentFile struct {
	Info         Info
	Announce     string
	AnnounceList *[][]string
	CreationDate *time.Time
	Comment      *string
	CreatedBy    *string
	Encoding     *string
}

type Info interface {
	PieceLength() int64
	Pieces() string
	Private() *int64
}

type abstractInfo struct {
	Info
	PieceLength int64
	Pieces      string
	Private     *int64
}

type SingleFileInfo struct {
	abstractInfo
	Name   string
	Length int64
	MD5Sum *string
}

type MultipleFileInfo struct {
	abstractInfo
	Name  string
	Files []File
}

type File struct {
	Length int64
	MD5Sum *string
	Path   []string
}

func (i SingleFileInfo) PieceLength() int64 {
	return i.abstractInfo.PieceLength
}

func (i SingleFileInfo) Pieces() string {
	return i.abstractInfo.Pieces
}

func (i SingleFileInfo) Private() *int64 {
	return i.abstractInfo.Private
}

func (i MultipleFileInfo) PieceLength() int64 {
	return i.abstractInfo.PieceLength
}

func (i MultipleFileInfo) Pieces() string {
	return i.abstractInfo.Pieces
}

func (i MultipleFileInfo) Private() *int64 {
	return i.abstractInfo.Private
}
