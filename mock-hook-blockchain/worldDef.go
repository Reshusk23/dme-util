package callbackblockchain

import (
	vmcommon "github.com/Reshusk23/dme-vm-common"
	"github.com/Reshusk23/dme-vm-common/data/dct"
)

// NewAddressMock allows tests to specify what new addresses to generate
type NewAddressMock struct {
	CreatorAddress []byte
	CreatorNonce   uint64
	NewAddress     []byte
}

// BlockInfo contains mock data about the corent block
type BlockInfo struct {
	BlockTimestamp uint64
	BlockNonce     uint64
	BlockRound     uint64
	BlockEpoch     uint32
}

// BlockchainHookMock provides a mock representation of the blockchain to be used in VM tests.
type BlockchainHookMock struct {
	AcctMap                      AccountMap
	PreviousBlockInfo            *BlockInfo
	CurrentBlockInfo             *BlockInfo
	Blockhashes                  [][]byte
	mockAddressGenerationEnabled bool
	NewAddressMocks              []*NewAddressMock
}

// ClearCompiledCodes implements vmcommon.BlockchainHook.
func (*BlockchainHookMock) ClearCompiledCodes() {
	panic("unimplemented")
}

// GetCode implements vmcommon.BlockchainHook.
func (*BlockchainHookMock) GetCode(vmcommon.UserAccountHandler) []byte {
	panic("unimplemented")
}

// GetCompiledCode implements vmcommon.BlockchainHook.
func (*BlockchainHookMock) GetCompiledCode(codeHash []byte) (bool, []byte) {
	panic("unimplemented")
}

// GetDCTToken implements vmcommon.BlockchainHook.
func (*BlockchainHookMock) GetDCTToken(address []byte, tokenID []byte, nonce uint64) (*dct.DCToken, error) {
	panic("unimplemented")
}

// GetSnapshot implements vmcommon.BlockchainHook.
func (*BlockchainHookMock) GetSnapshot() int {
	panic("unimplemented")
}

// IsInterfaceNil implements vmcommon.BlockchainHook.
func (*BlockchainHookMock) IsInterfaceNil() bool {
	panic("unimplemented")
}

// IsPayable implements vmcommon.BlockchainHook.
func (*BlockchainHookMock) IsPayable(address []byte) (bool, error) {
	panic("unimplemented")
}

// RevertToSnapshot implements vmcommon.BlockchainHook.
func (*BlockchainHookMock) RevertToSnapshot(snapshot int) error {
	panic("unimplemented")
}

// SaveCompiledCode implements vmcommon.BlockchainHook.
func (*BlockchainHookMock) SaveCompiledCode(codeHash []byte, code []byte) {
	panic("unimplemented")
}

// NewMock creates a new mock instance
func NewMock() *BlockchainHookMock {
	return &BlockchainHookMock{
		AcctMap:                      NewAccountMap(),
		PreviousBlockInfo:            nil,
		CurrentBlockInfo:             nil,
		Blockhashes:                  nil,
		mockAddressGenerationEnabled: false,
	}
}

// Clear resets all mock data between tests.
func (b *BlockchainHookMock) Clear() {
	b.AcctMap = NewAccountMap()
	b.Blockhashes = nil
}

// EnableMockAddressGeneration causes the mock to generate its own new addresses.
func (b *BlockchainHookMock) EnableMockAddressGeneration() {
	b.mockAddressGenerationEnabled = true
}
