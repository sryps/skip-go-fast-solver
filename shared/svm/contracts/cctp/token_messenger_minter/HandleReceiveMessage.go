// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_messenger_minter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// HandleReceiveMessage is the `handleReceiveMessage` instruction.
type HandleReceiveMessage struct {
	Params *HandleReceiveMessageParams

	// [0] = [SIGNER] authorityPda
	//
	// [1] = [] tokenMessenger
	//
	// [2] = [] remoteTokenMessenger
	//
	// [3] = [] tokenMinter
	//
	// [4] = [WRITE] localToken
	//
	// [5] = [] tokenPair
	//
	// [6] = [WRITE] recipientTokenAccount
	//
	// [7] = [WRITE] custodyTokenAccount
	//
	// [8] = [] tokenProgram
	//
	// [9] = [] eventAuthority
	//
	// [10] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewHandleReceiveMessageInstructionBuilder creates a new `HandleReceiveMessage` instruction builder.
func NewHandleReceiveMessageInstructionBuilder() *HandleReceiveMessage {
	nd := &HandleReceiveMessage{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *HandleReceiveMessage) SetParams(params HandleReceiveMessageParams) *HandleReceiveMessage {
	inst.Params = &params
	return inst
}

// SetAuthorityPdaAccount sets the "authorityPda" account.
func (inst *HandleReceiveMessage) SetAuthorityPdaAccount(authorityPda ag_solanago.PublicKey) *HandleReceiveMessage {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authorityPda).SIGNER()
	return inst
}

// GetAuthorityPdaAccount gets the "authorityPda" account.
func (inst *HandleReceiveMessage) GetAuthorityPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenMessengerAccount sets the "tokenMessenger" account.
func (inst *HandleReceiveMessage) SetTokenMessengerAccount(tokenMessenger ag_solanago.PublicKey) *HandleReceiveMessage {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenMessenger)
	return inst
}

// GetTokenMessengerAccount gets the "tokenMessenger" account.
func (inst *HandleReceiveMessage) GetTokenMessengerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetRemoteTokenMessengerAccount sets the "remoteTokenMessenger" account.
func (inst *HandleReceiveMessage) SetRemoteTokenMessengerAccount(remoteTokenMessenger ag_solanago.PublicKey) *HandleReceiveMessage {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(remoteTokenMessenger)
	return inst
}

// GetRemoteTokenMessengerAccount gets the "remoteTokenMessenger" account.
func (inst *HandleReceiveMessage) GetRemoteTokenMessengerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenMinterAccount sets the "tokenMinter" account.
func (inst *HandleReceiveMessage) SetTokenMinterAccount(tokenMinter ag_solanago.PublicKey) *HandleReceiveMessage {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenMinter)
	return inst
}

// GetTokenMinterAccount gets the "tokenMinter" account.
func (inst *HandleReceiveMessage) GetTokenMinterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetLocalTokenAccount sets the "localToken" account.
func (inst *HandleReceiveMessage) SetLocalTokenAccount(localToken ag_solanago.PublicKey) *HandleReceiveMessage {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(localToken).WRITE()
	return inst
}

// GetLocalTokenAccount gets the "localToken" account.
func (inst *HandleReceiveMessage) GetLocalTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenPairAccount sets the "tokenPair" account.
func (inst *HandleReceiveMessage) SetTokenPairAccount(tokenPair ag_solanago.PublicKey) *HandleReceiveMessage {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenPair)
	return inst
}

// GetTokenPairAccount gets the "tokenPair" account.
func (inst *HandleReceiveMessage) GetTokenPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetRecipientTokenAccountAccount sets the "recipientTokenAccount" account.
func (inst *HandleReceiveMessage) SetRecipientTokenAccountAccount(recipientTokenAccount ag_solanago.PublicKey) *HandleReceiveMessage {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(recipientTokenAccount).WRITE()
	return inst
}

// GetRecipientTokenAccountAccount gets the "recipientTokenAccount" account.
func (inst *HandleReceiveMessage) GetRecipientTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCustodyTokenAccountAccount sets the "custodyTokenAccount" account.
func (inst *HandleReceiveMessage) SetCustodyTokenAccountAccount(custodyTokenAccount ag_solanago.PublicKey) *HandleReceiveMessage {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(custodyTokenAccount).WRITE()
	return inst
}

// GetCustodyTokenAccountAccount gets the "custodyTokenAccount" account.
func (inst *HandleReceiveMessage) GetCustodyTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *HandleReceiveMessage) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *HandleReceiveMessage {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *HandleReceiveMessage) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *HandleReceiveMessage) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *HandleReceiveMessage {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *HandleReceiveMessage) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetProgramAccount sets the "program" account.
func (inst *HandleReceiveMessage) SetProgramAccount(program ag_solanago.PublicKey) *HandleReceiveMessage {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *HandleReceiveMessage) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst HandleReceiveMessage) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_HandleReceiveMessage,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst HandleReceiveMessage) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *HandleReceiveMessage) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AuthorityPda is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenMessenger is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.RemoteTokenMessenger is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenMinter is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.LocalToken is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenPair is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.RecipientTokenAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CustodyTokenAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *HandleReceiveMessage) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("HandleReceiveMessage")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        authorityPda", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      tokenMessenger", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("remoteTokenMessenger", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         tokenMinter", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          localToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           tokenPair", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("      recipientToken", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("        custodyToken", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("        tokenProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("      eventAuthority", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("             program", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj HandleReceiveMessage) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *HandleReceiveMessage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewHandleReceiveMessageInstruction declares a new HandleReceiveMessage instruction with the provided parameters and accounts.
func NewHandleReceiveMessageInstruction(
	// Parameters:
	params HandleReceiveMessageParams,
	// Accounts:
	authorityPda ag_solanago.PublicKey,
	tokenMessenger ag_solanago.PublicKey,
	remoteTokenMessenger ag_solanago.PublicKey,
	tokenMinter ag_solanago.PublicKey,
	localToken ag_solanago.PublicKey,
	tokenPair ag_solanago.PublicKey,
	recipientTokenAccount ag_solanago.PublicKey,
	custodyTokenAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *HandleReceiveMessage {
	return NewHandleReceiveMessageInstructionBuilder().
		SetParams(params).
		SetAuthorityPdaAccount(authorityPda).
		SetTokenMessengerAccount(tokenMessenger).
		SetRemoteTokenMessengerAccount(remoteTokenMessenger).
		SetTokenMinterAccount(tokenMinter).
		SetLocalTokenAccount(localToken).
		SetTokenPairAccount(tokenPair).
		SetRecipientTokenAccountAccount(recipientTokenAccount).
		SetCustodyTokenAccountAccount(custodyTokenAccount).
		SetTokenProgramAccount(tokenProgram).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
