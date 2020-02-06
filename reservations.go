package namegrind

import (
	"bufio"
	"github.com/cheggaaa/pb/v3"
	"github.com/mitchellh/go-homedir"
	"io"
	"net/http"
	"os"
	"path"
)

const ReservationsURL = "https://gist.githubusercontent.com/mslipper/d1a9989a8a175272f51da5e2f43e8d15/raw/72b07b1e29f08575d2feebef21c9da09c5e167cc/namehashes"

func ReservationsPath() string {
	dir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return path.Join(dir, "reservations.json")
}

func ReservationsExist() (bool, error) {
	_, err := os.Stat(ReservationsPath())
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func FetchReservations() error {
	out, err := os.Create(ReservationsPath())
	defer out.Close()
	res, err := http.Get(ReservationsURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bar := pb.Full.Start64(res.ContentLength)
	if _, err := io.Copy(out, bar.NewProxyReader(res.Body)); err != nil {
		return err
	}
	bar.Finish()
	return nil
}

func ParseReservations() (map[string]bool, error) {
	out := make(map[string]bool)
	f, err := os.Open(ReservationsPath())
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		out[line] = true
	}

	return out, nil
}
