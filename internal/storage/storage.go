package storage

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"greeschenko/workoverlord2/internal/encriptor"
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
	return &Storage{
		encriptor: encriptor.NewEncriptor(),
	}
}

func (s *Storage) SetSecret(secret string) {
	s.secretkey = sha256.Sum256([]byte(secret))
}

func (s *Storage) Descrypt(data []byte) ([]byte, error) {
	data, err := s.encriptor.Descrypt(data, s.secretkey)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Storage) Load() ([]byte, error) {
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
			return nil, err
		}
		data, err := s.Descrypt(tmpdata)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, nil
}

func (s *Storage) Save() {
	log.Println("storage saved data")
}
