package definitions

type ProtoDataTransfer struct {
	Type     [2]byte
	Encoding [1]byte
	Checksum [64]byte
	Length   [5]byte
	Payload  []byte
}
