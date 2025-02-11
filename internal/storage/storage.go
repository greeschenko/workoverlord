package storage

import (
	"greeschenko/workoverlord2/internal/encriptor"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	Dbpath = os.Getenv("HOME") + "/prodev/"
	Dbfile = Dbpath + "MIND2"
)

type EnkriptorInteraface interface {
	Encrypt([]byte, [32]byte) []byte
	Descrypt([]byte, [32]byte) ([]byte, error)
}

type Storage struct {
	secretkey [32]byte
	encriptor EnkriptorInteraface
}

func NewStorage() *Storage {
	return &Storage{encriptor: encriptor.NewEncriptor()}
}

func (s *Storage) SetSecret(secret string) {
	s.secretkey = sha256.Sum256([]byte(secret))
}

func (s *Storage) GetSecret() [32]byte {
	return s.secretkey
}

func (s *Storage) Load() error {
	if _, err := os.Stat(Dbpath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(Dbpath, os.ModePerm)
		fmt.Println("- db path created")
	}

	if _, err := os.Stat(Dbfile); errors.Is(err, os.ErrNotExist) {
		fmt.Println("===============================")
		fmt.Println("New MIND created")
	} else {
		tmpdata, err := os.ReadFile(Dbfile)
		if err != nil {
			return err
		}
		data, err := s.encriptor.Descrypt(tmpdata, s.secretkey)
		if err != nil {
			return err
		}
		json.Unmarshal(data, &USERMIND)
	}

	return nil
}

func (s *Storage) Save() {
	log.Println("storage saved data")
}
