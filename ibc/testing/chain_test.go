package ibctesting_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/x/staking/types"

	ibctesting "github.com/Canto-Network/Canto/v8/ibc/testing"
)

func TestChangeValSet(t *testing.T) {
	coord := ibctesting.NewCoordinator(t, 1, 1)
	chainA := coord.GetChain(ibctesting.GetChainID(2))
	chainB := coord.GetChain(ibctesting.GetChainIDCanto(1))

	path := ibctesting.NewTransferPath(chainA, chainB)
	coord.Setup(path)

	amount, ok := sdkmath.NewIntFromString("10000000000000000000")
	require.True(t, ok)
	amount2, ok := sdkmath.NewIntFromString("30000000000000000000")
	require.True(t, ok)

	val, err := chainA.GetSimApp().StakingKeeper.GetValidators(chainA.GetContext(), 4)
	require.NoError(t, err)

	chainA.GetSimApp().StakingKeeper.Delegate(chainA.GetContext(), chainA.SenderAccounts[1].SenderAccount.GetAddress(), //nolint:errcheck // ignore error for test
		amount, types.Unbonded, val[1], true)
	chainA.GetSimApp().StakingKeeper.Delegate(chainA.GetContext(), chainA.SenderAccounts[3].SenderAccount.GetAddress(), //nolint:errcheck // ignore error for test
		amount2, types.Unbonded, val[3], true)

	coord.CommitBlock(chainA)

	// verify that update clients works even after validator update goes into effect
	err = path.EndpointB.UpdateClient()
	require.NoError(t, err)
	err = path.EndpointB.UpdateClient()
	require.NoError(t, err)
}
