#!/bin/bash

set -e

mkdir -p ./shared/contracts/usdc
abigen --abi ./shared/abi/usdc.json --pkg usdc --out ./shared/contracts/usdc/usdc.go

mkdir -p ./shared/contracts/fast_transfer_gateway
abigen --abi ./shared/abi/fast_transfer_gateway.json --pkg fast_transfer_gateway --out ./shared/contracts/fast_transfer_gateway/fast_transfer_gateway.go

mkdir -p ./shared/contracts/hyperlane/AggregationIsm
abigen --abi ./shared/abi/hyperlane/IAggregationIsm.abi.json --pkg aggregation_ism --out ./shared/contracts/hyperlane/AggregationIsm/aggregation_ism.go

mkdir -p ./shared/contracts/hyperlane/InterchainGasPaymaster
abigen --abi ./shared/abi/hyperlane/IInterchainGasPaymaster.abi.json --pkg interchain_gas_paymaster --out ./shared/contracts/hyperlane/InterchainGasPaymaster/interchain_gas_paymaster.go

mkdir -p ./shared/contracts/hyperlane/InterchainSecurityModule
abigen --abi ./shared/abi/hyperlane/IInterchainSecurityModule.abi.json --pkg interchain_security_module --out ./shared/contracts/hyperlane/InterchainSecurityModule/interchain_security_module.go

mkdir -p ./shared/contracts/hyperlane/Mailbox
abigen --abi ./shared/abi/hyperlane/Mailbox.abi.json --pkg mailbox --out ./shared/contracts/hyperlane/Mailbox/mailbox.go

mkdir -p ./shared/contracts/hyperlane/MultisigIsm
abigen --abi ./shared/abi/hyperlane/IMultisigIsm.abi.json --pkg multisig_ism --out ./shared/contracts/hyperlane/MultisigIsm/multisig_ism.go

mkdir -p ./shared/contracts/hyperlane/ValidatorAnnounce
abigen --abi ./shared/abi/hyperlane/IValidatorAnnounce.abi.json --pkg validator_announce --out ./shared/contracts/hyperlane/ValidatorAnnounce/validator_announce.go

mkdir -p ./shared/contracts/hyperlane/MerkleTreeHook
abigen --abi ./shared/abi/hyperlane/MerkleTreeHook.abi.json --pkg merkle_tree_hook --out ./shared/contracts/hyperlane/MerkleTreeHook/merkle_tree_hook.go