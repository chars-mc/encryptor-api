package domain

type DataType int

var DataTypes = []string{
	"Text",
	"File",
}

const (
	Text = iota + 1
	File
)

func (d DataType) String() string {
	return DataTypes[d-1]
}

type Algorithm int

const (
	AES = iota + 1
	Blowsifh
)

var Algorithms = []string{
	"AES",
	"Blowsifh",
}

func (a Algorithm) String() string {
	return Algorithms[a-1]
}

type Data struct {
	ID        int
	Content   string
	DataType  DataType
	Algorithm Algorithm
}
