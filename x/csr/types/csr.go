package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ethermint "github.com/evmos/ethermint/types"
)

// Creates a new instance of the CSR object
func NewCSR(contracts []string, id uint64) CSR {
	return CSR{
		Contracts: contracts,
		Id:        id,
	}
}

// Validate performs stateless validation of a CSR object
func (csr CSR) Validate() error {
	seenSmartContracts := make(map[string]bool)
	for _, smartContract := range csr.Contracts {
		if err := ethermint.ValidateNonZeroAddress(smartContract); err != nil {
			return sdkerrors.Wrapf(ErrInvalidSmartContractAddress, "CSR::Validate one or more of the entered smart contract address are invalid.")
		}

		if seenSmartContracts[smartContract] {
			return sdkerrors.Wrapf(ErrDuplicateSmartContracts, "CSR::Validate there are duplicate smart contracts in this CSR.")
		}
	}

	// Ensure that there is at least one smart contract in the CSR Pool
	numSmartContracts := len(csr.Contracts)
	if numSmartContracts < 1 {
		return sdkerrors.Wrapf(ErrSmartContractSupply, "CSR::Validate # of smart contracts must be greater than 0 got: %d", numSmartContracts)
	}
	return nil
}
