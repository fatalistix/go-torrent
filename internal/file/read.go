package file

import (
	"fmt"
	"github.com/jackpal/bencode-go"
	"os"
)

func Read(path string) (*TorrentFile, error) {
	const op = "file.Read"

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to read file %s: %w", op, path, err)
	}

	defer func() {
		_ = f.Close()
	}()

	raw, err := bencode.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("%s: file %s is not a bencode file: %w", op, path, err)
	}

	tf, err := parse(raw)
	if err != nil {
		return nil, fmt.Errorf("%s: file %s is not a torrent file: %w", op, path, err)
	}

	return tf, nil
}
