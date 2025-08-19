package file

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"github.com/jackpal/bencode-go"
	"io"
	"os"
)

type infoContainer struct {
	info []byte
}

func Read(path string) (*TorrentFile, [20]byte, error) {
	const op = "file.Read"

	f, err := os.Open(path)
	if err != nil {
		return nil, [20]byte{}, fmt.Errorf("%s: unable to read file %s: %w", op, path, err)
	}

	defer func() {
		_ = f.Close()
	}()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, [20]byte{}, fmt.Errorf("%s: unable to read file %s: %w", op, path, err)
	}

	r := bytes.NewReader(b)

	raw, err := bencode.Decode(r)
	if err != nil {
		return nil, [20]byte{}, fmt.Errorf("%s: file %s is not a bencode file: %w", op, path, err)
	}

	tf, err := parse(raw)
	if err != nil {
		return nil, [20]byte{}, fmt.Errorf("%s: file %s is not a torrent file: %w", op, path, err)
	}

	r = bytes.NewReader(b)

	i := infoContainer{}
	if err = bencode.Unmarshal(r, &i); err != nil {
		return nil, [20]byte{}, fmt.Errorf("%s: file %s is not a torrent file: %w", op, path, err)
	}

	infoHash := sha1.Sum(i.info)

	return tf, infoHash, nil
}
