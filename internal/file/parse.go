package file

import (
	"fmt"
	"reflect"
	"time"
)

func parse(raw any) (*TorrentFile, error) {
	data, ok := raw.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("not a torrent file: expected dictionary, but got %s", raw)
	}

	info, err := parseInfo(data[infoField])
	if err != nil {
		return nil, err
	}

	announce, err := parseAnnounce(data[announceField])
	if err != nil {
		return nil, err
	}

	announceList, err := parseAnnounceList(data[announceListField])
	if err != nil {
		return nil, err
	}

	creationDate, err := parseCreationDate(data[creationDateField])
	if err != nil {
		return nil, err
	}

	comment, err := parseComment(data[commentField])
	if err != nil {
		return nil, err
	}

	createdBy, err := parseCreatedBy(data[createdByField])
	if err != nil {
		return nil, err
	}

	encoding, err := parseEncoding(data[encodingField])
	if err != nil {
		return nil, err
	}

	return &TorrentFile{
		Info:         info,
		Announce:     announce,
		AnnounceList: announceList,
		CreationDate: creationDate,
		Comment:      comment,
		CreatedBy:    createdBy,
		Encoding:     encoding,
	}, nil
}

func parseInfo(raw any) (Info, error) {
	if raw == nil {
		return SingleFileInfo{}, fmt.Errorf("%s is not present", infoField)
	}

	info, ok := raw.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%s is present as %s, not a dictionary", infoField, reflect.TypeOf(raw))
	}

	pieceLength, err := parseAbstractInfoPieceLength(info[infoPieceLengthField])
	if err != nil {
		return nil, fmt.Errorf("failed to parse piece length: %w", err)
	}

	pieces, err := parseAbstractInfoPieces(info[infoPiecesField])
	if err != nil {
		return nil, fmt.Errorf("failed to parse pieces: %w", err)
	}

	private, err := parseAbstractInfoPrivate(info[infoPrivateField])
	if err != nil {
		return nil, fmt.Errorf("failed to parse private: %w", err)
	}

	_, ok = info[singleFileInfoLengthField]
	if ok {
		name, err := parseSingleFileInfoName(info[singleFileInfoLengthField])
		if err != nil {
			return nil, fmt.Errorf("failed to parse single file name: %w", err)
		}

		length, err := parseSingleFileInfoLength(info[singleFileInfoLengthField])
		if err != nil {
			return nil, fmt.Errorf("failed to parse single file length: %w", err)
		}

		md5Sum, err := parseSingleFileInfoMD5Sum(info[singleFileInfoMD5SumField])
		if err != nil {
			return nil, fmt.Errorf("failed to parse single file MD5 sum: %w", err)
		}

		s := SingleFileInfo{
			abstractInfo: abstractInfo{
				PieceLength: pieceLength,
				Pieces:      pieces,
				Private:     private,
			},
			Name:   name,
			Length: length,
			MD5Sum: md5Sum,
		}
		return s, nil
	} else {
		name, err := parseMultipleFileInfoName(info[multipleFileInfoNameField])
		if err != nil {
			return nil, fmt.Errorf("failed to parse multiple file name: %w", err)
		}

		files, err := parseMultipleFileInfoFiles(info[multipleFileInfoFilesField])
		if err != nil {
			return nil, fmt.Errorf("failed to parse multiple file files: %w", err)
		}

		m := MultipleFileInfo{
			abstractInfo: abstractInfo{
				PieceLength: pieceLength,
				Pieces:      pieces,
				Private:     private,
			},
			Name:  name,
			Files: files,
		}
		return m, nil
	}
}

func parseAbstractInfoPieceLength(raw any) (int64, error) {
	if raw == nil {
		return 0, fmt.Errorf("%s.%s is not present", infoField, infoPieceLengthField)
	}

	pieceLength, ok := raw.(int64)
	if !ok {
		return 0, fmt.Errorf("%s.%s is present as %s, not a int", infoField, infoPieceLengthField, reflect.TypeOf(raw))
	}

	return pieceLength, nil
}

func parseAbstractInfoPieces(raw any) (string, error) {
	if raw == nil {
		return "", fmt.Errorf("%s.%s is not present", infoField, infoPiecesField)
	}

	pieces, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf("%s.%s is present as %s, not a string", infoField, infoPiecesField, reflect.TypeOf(raw))
	}

	return pieces, nil
}

func parseAbstractInfoPrivate(raw any) (*int64, error) {
	if raw == nil {
		return nil, nil
	}

	private, ok := raw.(int64)
	if !ok {
		return nil, fmt.Errorf("%s.%s is present as %s, not an int", infoField, infoPrivateField, reflect.TypeOf(raw))
	}

	return &private, nil
}

func parseSingleFileInfoName(raw any) (string, error) {
	if raw == nil {
		return "", fmt.Errorf("%s.%s is not present", infoField, singleFileInfoNameField)
	}

	name, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf("%s.%s is present as %s, not a string", infoField, singleFileInfoNameField, reflect.TypeOf(raw))
	}

	return name, nil
}

func parseSingleFileInfoLength(raw any) (int64, error) {
	if raw == nil {
		return 0, fmt.Errorf("%s.%s is not present", infoField, singleFileInfoLengthField)
	}

	length, ok := raw.(int64)
	if !ok {
		return 0, fmt.Errorf("%s.%s is present as %s, not an int", infoField, singleFileInfoLengthField, reflect.TypeOf(raw))
	}

	return length, nil
}

func parseSingleFileInfoMD5Sum(raw any) (*string, error) {
	if raw == nil {
		return nil, nil
	}

	md5Sum, ok := raw.(string)
	if !ok {
		return nil, fmt.Errorf("%s.%s is present as %s, not a string", infoField, singleFileInfoMD5SumField, reflect.TypeOf(raw))
	}

	return &md5Sum, nil
}

func parseMultipleFileInfoName(raw any) (string, error) {
	if raw == nil {
		return "", fmt.Errorf("%s.%s is not present", infoField, multipleFileInfoNameField)
	}

	name, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf("%s.%s is present as %s, not a string", infoField, multipleFileInfoNameField, reflect.TypeOf(raw))
	}

	return name, nil
}

func parseMultipleFileInfoFiles(raw any) ([]File, error) {
	if raw == nil {
		return nil, fmt.Errorf("%s.%s is not present", infoField, multipleFileInfoFilesField)
	}

	rawFiles, ok := raw.([]any)
	if !ok {
		return nil, fmt.Errorf("%s.%s is present as %s, not a list", infoField, multipleFileInfoFilesField, reflect.TypeOf(raw))
	}

	files := make([]File, 0, len(rawFiles))
	for _, rawFile := range rawFiles {
		m, ok := rawFile.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("not a map")
		}

		length, err := parseFileLength(m[multipleFileInfoFileLengthField])
		if err != nil {
			return nil, fmt.Errorf("unable to parse file length: %w", err)
		}

		md5Sum, err := parseFileMD5Sum(m[multipleFileInfoFileMD5SumField])
		if err != nil {
			return nil, fmt.Errorf("unable to parse file md5sum: %w", err)
		}

		path, err := parseFilePath(m[multipleFileInfoFilePathField])
		if err != nil {
			return nil, fmt.Errorf("unable to parse file path: %w", err)
		}

		f := File{
			Length: length,
			MD5Sum: md5Sum,
			Path:   path,
		}

		files = append(files, f)
	}

	return files, nil
}

func parseFileLength(raw any) (int64, error) {
	if raw == nil {
		return 0, fmt.Errorf("%s.%s is not present", multipleFileInfoFilesField, multipleFileInfoFileLengthField)
	}

	length, ok := raw.(int64)
	if !ok {
		return 0, fmt.Errorf(
			"%s.%s with value %s is present as %s, not a int",
			multipleFileInfoFilesField,
			multipleFileInfoFileLengthField,
			raw,
			reflect.TypeOf(raw),
		)
	}

	return length, nil
}

func parseFileMD5Sum(raw any) (*string, error) {
	if raw == nil {
		return nil, nil
	}

	md5Sum, ok := raw.(string)
	if !ok {
		return nil, fmt.Errorf(
			"%s.%s with value %s is present as %s, not a string",
			multipleFileInfoFilesField,
			multipleFileInfoFileMD5SumField,
			raw,
			reflect.TypeOf(raw),
		)
	}

	return &md5Sum, nil
}

func parseFilePath(raw any) ([]string, error) {
	if raw == nil {
		return nil, fmt.Errorf("%s.%s is not present", multipleFileInfoFilesField, multipleFileInfoFilePathField)
	}

	rawPath, ok := raw.([]any)
	if !ok {
		return nil, fmt.Errorf("not a list")
	}

	path := make([]string, len(rawPath))
	for i, r := range rawPath {
		path[i], ok = r.(string)
		if !ok {
			return nil, fmt.Errorf("not a string")
		}
	}

	return path, nil
}

func parseAnnounce(raw any) (string, error) {
	if raw == nil {
		return "", fmt.Errorf("%s is not present", announceField)
	}

	announce, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf("%s is present as %s, not a string", announceField, reflect.TypeOf(raw))
	}

	return announce, nil
}

func parseAnnounceList(raw any) (*[][]string, error) {
	if raw == nil {
		return nil, nil
	}

	rawAnnounceList, ok := raw.([]any)
	if !ok {
		return nil, fmt.Errorf("%s is present as %s, not a list", announceField, reflect.TypeOf(raw))
	}

	announceList := make([][]string, len(rawAnnounceList))

	for i, e := range rawAnnounceList {
		rawAnnounceListElement, ok := e.([]any)
		if !ok {
			return nil, fmt.Errorf("%s element is present as %s, not a list", announceField, reflect.TypeOf(e))
		}

		tmpAnnounceListElement := make([]string, len(rawAnnounceListElement))

		for j, e := range rawAnnounceListElement {
			announce, ok := e.(string)
			if !ok {
				return nil, fmt.Errorf("%s list's element is present as %s, not a string", announceField, reflect.TypeOf(e))
			}

			tmpAnnounceListElement[j] = announce
		}

		announceList[i] = tmpAnnounceListElement
	}

	return &announceList, nil
}

func parseCreationDate(raw any) (*time.Time, error) {
	if raw == nil {
		return nil, nil
	}

	creationDateInt, ok := raw.(int64)
	if !ok {
		return nil, fmt.Errorf("creation date is present as %s, not an int64", reflect.TypeOf(raw))
	}

	creationDate := time.Unix(creationDateInt, 0)
	return &creationDate, nil
}

func parseComment(raw any) (*string, error) {
	if raw == nil {
		return nil, nil
	}

	comment, ok := raw.(string)
	if !ok {
		return nil, fmt.Errorf("comment is present as %s, not a string", reflect.TypeOf(raw))
	}

	return &comment, nil
}

func parseCreatedBy(raw any) (*string, error) {
	if raw == nil {
		return nil, nil
	}

	createdBy, ok := raw.(string)
	if !ok {
		return nil, fmt.Errorf("created by is present as %s, not a string", reflect.TypeOf(raw))
	}

	return &createdBy, nil
}

func parseEncoding(raw any) (*string, error) {
	if raw == nil {
		return nil, nil
	}

	encoding, ok := raw.(string)
	if !ok {
		return nil, fmt.Errorf("encoding is present as %s, not a string", reflect.TypeOf(raw))
	}

	return &encoding, nil
}
