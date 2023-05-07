package definitions

type ProtoDataTransfer struct {
	Type    [2]byte
	Length  [5]byte
	Payload []byte
}
