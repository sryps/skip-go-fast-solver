// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_messenger_minter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// LinkTokenPair is the `linkTokenPair` instruction.
type LinkTokenPair struct {
	Params *LinkTokenPairParams

	// [0] = [WRITE, SIGNER] payer
	//
	// [1] = [SIGNER] tokenController
	//
	// [2] = [] tokenMinter
	//
	// [3] = [WRITE] tokenPair
	//
	// [4] = [] systemProgram
	//
	// [5] = [] eventAuthority
	//
	// [6] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewLinkTokenPairInstructionBuilder creates a new `LinkTokenPair` instruction builder.
func NewLinkTokenPairInstructionBuilder() *LinkTokenPair {
	nd := &LinkTokenPair{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *LinkTokenPair) SetParams(params LinkTokenPairParams) *LinkTokenPair {
	inst.Params = &params
	return inst
}

// SetPayerAccount sets the "payer" account.
func (inst *LinkTokenPair) SetPayerAccount(payer ag_solanago.PublicKey) *LinkTokenPair {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *LinkTokenPair) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenControllerAccount sets the "tokenController" account.
func (inst *LinkTokenPair) SetTokenControllerAccount(tokenController ag_solanago.PublicKey) *LinkTokenPair {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenController).SIGNER()
	return inst
}

// GetTokenControllerAccount gets the "tokenController" account.
func (inst *LinkTokenPair) GetTokenControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenMinterAccount sets the "tokenMinter" account.
func (inst *LinkTokenPair) SetTokenMinterAccount(tokenMinter ag_solanago.PublicKey) *LinkTokenPair {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenMinter)
	return inst
}

// GetTokenMinterAccount gets the "tokenMinter" account.
func (inst *LinkTokenPair) GetTokenMinterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenPairAccount sets the "tokenPair" account.
func (inst *LinkTokenPair) SetTokenPairAccount(tokenPair ag_solanago.PublicKey) *LinkTokenPair {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenPair).WRITE()
	return inst
}

// GetTokenPairAccount gets the "tokenPair" account.
func (inst *LinkTokenPair) GetTokenPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *LinkTokenPair) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *LinkTokenPair {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *LinkTokenPair) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *LinkTokenPair) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *LinkTokenPair {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *LinkTokenPair) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetProgramAccount sets the "program" account.
func (inst *LinkTokenPair) SetProgramAccount(program ag_solanago.PublicKey) *LinkTokenPair {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *LinkTokenPair) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst LinkTokenPair) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_LinkTokenPair,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst LinkTokenPair) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *LinkTokenPair) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenController is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenMinter is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenPair is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *LinkTokenPair) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("LinkTokenPair")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("tokenController", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    tokenMinter", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      tokenPair", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta(" eventAuthority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("        program", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj LinkTokenPair) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *LinkTokenPair) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewLinkTokenPairInstruction declares a new LinkTokenPair instruction with the provided parameters and accounts.
func NewLinkTokenPairInstruction(
	// Parameters:
	params LinkTokenPairParams,
	// Accounts:
	payer ag_solanago.PublicKey,
	tokenController ag_solanago.PublicKey,
	tokenMinter ag_solanago.PublicKey,
	tokenPair ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *LinkTokenPair {
	return NewLinkTokenPairInstructionBuilder().
		SetParams(params).
		SetPayerAccount(payer).
		SetTokenControllerAccount(tokenController).
		SetTokenMinterAccount(tokenMinter).
		SetTokenPairAccount(tokenPair).
		SetSystemProgramAccount(systemProgram).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
