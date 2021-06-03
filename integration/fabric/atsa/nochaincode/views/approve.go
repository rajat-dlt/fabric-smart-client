/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package views

import (
	"github.com/hyperledger-labs/fabric-smart-client/platform/fabric/services/state"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/assert"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
)

type ApproveView struct{}

func (a *ApproveView) Call(context view.Context) (interface{}, error) {
	tx, err := state.ReceiveTransaction(context)
	assert.NoError(err, "failed receiving transaction")

	assert.Equal(1, tx.Commands().Count(), "expected one command, got [%d]", tx.Commands().Count())
	assert.Equal(1, tx.Namespaces().Count(), "expected one namespace, got [%d]", tx.Namespaces().Count())
	assert.Equal("asset_transfer", tx.Namespaces().At(0), "expected 'asset_transfer', got [%s]", tx.Namespaces().At(0))

	switch cmd := tx.Commands().At(0); cmd.Name {
	case "issue":
		assert.Equal(0, tx.Inputs().Count(), "expected zero inputs in issue")
		assert.Equal(1, tx.Outputs().Count(), "expected one output in issue")

		assert.Equal(2, cmd.Ids.Count(), "expected two identities in issue")
		assert.False(cmd.Ids[0].Equal(cmd.Ids.Others(cmd.Ids[0])[0]), "expected two different identities in issue")
		assert.NoError(tx.HasBeenEndorsedBy(cmd.Ids...), "expected two valid signatures in issue")

		asset := &Asset{}
		assert.NoError(tx.Outputs().At(0).State(asset))
		assert.True(cmd.Ids.Contain(asset.Owner), "expected asset to contain one of the two command signers")
		// TODO: check asset
	case "agreeToSell":
		assert.Equal(0, tx.Inputs().Count(), "expected zero inputs in agreeToSell")
		assert.Equal(1, tx.Outputs().Count(), "expected one output in agreeToSell")

		assert.Equal(1, cmd.Ids.Count(), "expected one identity in agreeToSell")
		assert.NoError(tx.HasBeenEndorsedBy(cmd.Ids[0]), "expected a valid signature in agreeToSell")
		assert.True(tx.Outputs().At(0).ID().HasPrefix(TypeAssetForSale), "expected agreeToSell prefix")

		agreeToSell := &AgreementToSell{}
		assert.NoError(tx.Outputs().At(0).State(agreeToSell))
		assert.True(cmd.Ids.Contain(agreeToSell.Owner), "expected agree to sell to contain the command signers")
	case "agreeToBuy":
		assert.Equal(0, tx.Inputs().Count(), "expected zero inputs in agreeToBuy")
		assert.Equal(1, tx.Outputs().Count(), "expected one output in agreeToBuy")

		assert.Equal(1, cmd.Ids.Count(), "expected one identity in agreeToBuy")
		assert.NoError(tx.HasBeenEndorsedBy(cmd.Ids[0]), "expected a valid signature in agreeToBuy")
		assert.True(tx.Outputs().At(0).ID().HasPrefix(TypeAssetBid), "expected agreeToBuy prefix")

		agreeToBuy := &AgreementToBuy{}
		assert.NoError(tx.Outputs().At(0).State(agreeToBuy))
		assert.True(cmd.Ids.Contain(agreeToBuy.Owner), "expected agree to buy to contain the command signers")
	case "transfer":
		assert.Equal(3, tx.Inputs().Count(), "expected three input in transfer")
		assert.Equal(3, tx.Outputs().Count(), "expected three output in transfer")
		assert.Equal(2, tx.Outputs().Deleted().Count(), "expected two delete in transfer")

		assert.Equal(1, tx.Inputs().IDs().Filter(state.IDHasPrefixFilter(TypeAssetForSale)).Count(), "expected one agreeToSell input")
		assert.Equal(1, tx.Inputs().IDs().Filter(state.IDHasPrefixFilter(TypeAssetBid)).Count(), "expected one agreeToBuy input")

		assert.True(tx.Outputs().IDs().Match(tx.Inputs().IDs()))

		assetIn := &Asset{}
		assert.NoError(tx.Inputs().Filter(state.InputHasIDPrefixFilter(TypeAsset)).At(0).State(assetIn))
		assetOut := &Asset{}
		assert.NoError(tx.Outputs().Written().At(0).State(assetOut))

		assert.Equal(2, cmd.Ids.Count(), "expected two identities in transfer")
		assert.NoError(tx.HasBeenEndorsedBy(cmd.Ids...), "expected two valid signatures in transfer")
		assert.True(cmd.Ids.Match(state.Identities{assetIn.Owner, assetOut.Owner}), "expected asset owners")

		//assert.EqualMod(assetIn, assetOut, []string{"Owner"}, "assets do not match")
	default:
		assert.Fail("expected a valid command, got [%s]", cmd)
	}

	// Accept and send back
	_, err = context.RunView(state.NewEndorseView(tx))
	assert.NoError(err)

	// Wait for finality
	_, err = context.RunView(state.NewFinalityView(tx))
	assert.NoError(err)
	return nil, nil
}