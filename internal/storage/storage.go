package storage

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"greeschenko/workoverlord2/internal/encriptor"
	"greeschenko/workoverlord2/internal/models"
	"log"
	"os"
	"time"
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
	data      models.MIND
}

func NewStorage() *Storage {
	return &Storage{encriptor: encriptor.NewEncriptor()}
}

func (s *Storage) SetSecret(secret string) {
	s.secretkey = sha256.Sum256([]byte(secret))
}

func (s *Storage) Descrypt(data []byte) error {
	data, err := s.encriptor.Descrypt(data, s.secretkey)
	if err != nil {
		return err
	}
	json.Unmarshal(data, &s.data)
	fmt.Println(s.data)
	return nil
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
		err = s.Descrypt(tmpdata)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) Save() {
	log.Println("storage saved data")
}

func (s *Storage) GetData() models.MIND {
	return s.data
}

func (s *Storage) AddData(newdata models.Cell) {
	newkey := time.Now().Format(time.RFC3339)
	s.data.Cells[newkey] = &newdata
	s.Save()
}
func (s *Storage) UpdateData(key string, newdata models.Cell) {
	s.data.Cells[key] = &newdata
	s.Save()
}
func (s *Storage) DeleteData(key string, newdata models.Cell) {
	delete(s.data.Cells, key)
	s.Save()
}
