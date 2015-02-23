package fs2

import (
	"errors"
	"fmt"
	"io"
	"os"
	"encoding/binary"
)

const (
	fileHeaderLength = uint32(4096)
)

type Store struct {
	Path string
	file os.File
	alignment uint64
	index map[string]uint64
}

type HeaderStart struct {
	Marker [4]byte
	Length uint32
}

type FileHeader struct {
	HeaderStart
	Alignment byte
	PathLength uint16
}

func bread(w io.Reader, data interface{}) error {
	return binary.Read(w, binary.LittleEndian, data)
}

func bwrite(w io.Writer, data interface{}) error {
	return binary.Write(w, binary.LittleEndian, data)
}

func NewAlignment(path string, alignment_size byte) (*Store, error) {
	file, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil { return nil, err }

	info, err := file.Stat()
	if err != nil { return nil, err }
	size := info.Size()

	var magic = [4]byte{'f', 's', '2', ' '}
	if size == 0 {

		err = bwrite(
			file,
			FileHeader{HeaderStart{magic, fileHeaderLength}, alignment_size, 0},
		)
		if err != nil { return nil, err }

		_, err = file.Seek(int64(fileHeaderLength-4), 0)
		if err != nil { return nil, err }

		err = bwrite(file, fileHeaderLength)
		if err != nil { return nil, err }

		err = file.Sync()
		if err != nil { return nil, err }
	} else {
		if size < 4096 {
			return nil, errors.New("Invalid file content")
		}
		var header FileHeader
		bread(file, &header)
		if header.Marker != magic {
			return nil, fmt.Errorf("Bad magic %s", header.Marker)
		}
		if header.Length != 4096{
			return nil, errors.New("Bad header length")
		}
		// TODO: prev path
		alignment_size = header.Alignment
	}

	store := Store{
		path, *file, uint64(1 << alignment_size), make(map[string]uint64)}
	return &store, nil
}

func New(path string) (*Store, error) {
	return NewAlignment(path, 28)
}

func (store *Store) Close() error {
	return store.file.Close()
}

//func Store(oid string, tid string, data string)

