package storage

// Define methods to write/read data from different providers
type StorageWriter interface {
    Write (s MemStorage) error
    RestoreData (s *MemStorage) error
    Save (t int, s MemStorage) error
    Close ()
}