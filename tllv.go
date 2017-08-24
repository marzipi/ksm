package ksm

import (
	"crypto/rand"
	"encoding/binary"
)

type TLLVBlock struct {
	Tag         uint64
	BlockLength uint32
	ValueLength uint32
	Value       []byte
}

func NewTLLVBlock(tag uint64, value []byte) *TLLVBlock {
	valueLen := uint32(len(value))
	paddingSize := 32 - valueLen%16 // Extend to nearest 16 bytes + extra 16 bytes
	blockLen := valueLen + paddingSize

	return &TLLVBlock{
		Tag:         tag,
		BlockLength: blockLen,
		ValueLength: valueLen,
		Value:       value,
	}
}

func (t *TLLVBlock) Serialize() []byte {
	var out []byte

	tagOut := make([]byte, 8)
	blockLenOut := make([]byte, 4)
	valueLenOut := make([]byte, 4)

	valueLen := uint32(len(t.Value))
	paddingLen := 32 - valueLen%16 // Extend to nearest 16 bytes + extra 16 bytes
	blockLen := valueLen + paddingLen

	paddingOut := make([]byte, paddingLen)
	rand.Read(paddingOut)

	binary.BigEndian.PutUint64(tagOut, t.Tag)
	binary.BigEndian.PutUint32(blockLenOut, blockLen)
	binary.BigEndian.PutUint32(valueLenOut, valueLen)

	out = append(out, tagOut...)
	out = append(out, blockLenOut...)
	out = append(out, valueLenOut...)
	out = append(out, t.Value...)
	out = append(out, paddingOut...)

	return out
}

type SKR1TLLVBlock struct {
	TLLVBlock
	IV      []byte
	Payload []byte
}

type DecryptedSKR1Payload struct {
	SK             []byte //Session key
	HU             []byte
	R1             []byte
	IntegrityBytes []byte
}

type CkcR1 struct {
	R1 []byte
}

type CkcDataIv struct {
	IV []byte
}

type CkcEncryptedPayload struct {
	Payload []byte
}

const (
	Tag_SessionKey_R1                 = 0x3d1a10b8bffac2ec
	Tag_SessionKey_R1_integrity       = 0xb349d4809e910687
	Tag_AntiReplaySeed                = 0x89c90f12204106b2
	Tag_R2                            = 0x71b5595ac1521133
	Tag_ReturnRequest                 = 0x19f9d4e5ab7609cb
	Tag_AssetID                       = 0x1bf7f53f5d5d5a1f
	Tag_TransactionID                 = 0x47aa7ad3440577de
	Tag_ProtocolVersionsSupported     = 0x67b8fb79ecce1a13
	Tag_ProtocolVersionUsed           = 0x5d81bcbcc7f61703
	Tag_treamingIndicator             = 0xabb0256a31843974
	Tag_kSKDServerClientReferenceTime = 0xeb8efdf2b25ab3a0 //Media playback state

	//kSKDServerReturnTags,
	kSKDServerKeyDurationTag = 0x47acf6a418cd091a
)

const (
	Field_Tag_Length   = 8
	Field_Block_Length = 4
	Field_Value_Length = 4
)

const (
	Tag_Encrypted_CK         = 0x58b38165af0e3d5a
	Tag_R1                   = 0xea74c4645d5efee9
	Tag_Content_Key_Duration = 0x47acf6a418cd091a
	Tag_HDCP_Enforcement     = 0x2e52f1530d8ddb4a
)