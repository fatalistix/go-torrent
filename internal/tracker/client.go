package tracker

import (
	"github.com/fatalistix/go-torrent/internal/file"
	"net/http"
	"net/url"
)

const (
	infoHashParam   = "info_hash"
	peerIdParam     = "peer_id"
	portParam       = "port"
	uploadedParam   = "uploaded"
	downloadedParam = "downloaded"
	leftParam       = "left"
	compactParam    = "compact"
	noPeerIdParam   = "no_peer_id"
	eventParam      = "event"
	ipParam         = "ip"
	numWantParam    = "numwant"
	keyParam        = "key"
	trackeridParam  = "trackerid"
)

func send(file file.TorrentFile) error {
	fullUrl, err := url.Parse(file.Announce)
	if err != nil {
		return err
	}

	queryParams := url.Values{}
	//queryParams.Set()

	fullUrl.RawQuery = queryParams.Encode()

	req, err := http.NewRequest(http.MethodGet, fullUrl.String(), nil)
	if err != nil {
		return err
	}

	_ = req

	return nil
}
