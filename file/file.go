// type File represents file access to a resource
package file

// type File represents a file
type File interface {
	Read(b []byte) (int, error)
	Write(b []byte) (int, error)
	Seek(off int64, when int) (int64, error)
	Close() error
}
