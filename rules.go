package namegrind

import (
	"errors"
	"golang.org/x/crypto/sha3"
)

const (
	MaxNameLen             = 64
	MainnetRolloutInterval = 1008
	MainnetAuctionStart    = 2016
)

var validCharset = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 0, 0, 4,
	0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 0, 0, 0, 0, 0,
}

func ValidateName(name string) error {
	if len(name) == 0 {
		return errors.New("name must have nonzero length")
	}

	if len(name) > MaxNameLen {
		return errors.New("name over maximum length")
	}

	for i := 0; i < len(name); i++ {
		ch := name[i]

		if int(ch) > len(validCharset) {
			return errors.New("invalid character")
		}

		charType := validCharset[ch]
		switch charType {
		case 0:
			return errors.New("invalid character")
		case 1:
			continue
		case 2:
			return errors.New("name cannot contain capital letters")
		case 3:
			continue
		case 4:
			if i == 0 {
				return errors.New("name cannot start with a hyphen")
			}
			if i == len(name)-1 {
				return errors.New("name cannot end with a hyphen")
			}
			if name[i-1] == ch {
				return errors.New("name cannot contain consecutive hyphens")
			}
		}
	}

	return nil
}

func HashName(name string) ([]byte, error) {
	if err := ValidateName(name); err != nil {
		return nil, err
	}
	h := sha3.Sum256([]byte(name))
	return h[:], nil
}

func Rollout(nameHash []byte) (int, int) {
	week := modBuffer(nameHash, 52)
	height := week * MainnetRolloutInterval
	return MainnetAuctionStart + height, week
}

func modBuffer(buf []byte, num int) int {
	if (num & 0xff) != num {
		panic("invalid num")
	}
	if num == 0 {
		panic("invalid num")
	}

	p := 256 % num
	var acc int
	for i := 0; i < len(buf); i++ {
		acc = (p*acc + int(buf[i])) % num
	}

	return acc
}
