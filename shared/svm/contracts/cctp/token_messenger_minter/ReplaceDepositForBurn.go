// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_messenger_minter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ReplaceDepositForBurn is the `replaceDepositForBurn` instruction.
type ReplaceDepositForBurn struct {
	Params *ReplaceDepositForBurnParams

	// [0] = [SIGNER] owner
	//
	// [1] = [WRITE, SIGNER] eventRentPayer
	//
	// [2] = [] senderAuthorityPda
	//
	// [3] = [WRITE] messageTransmitter
	//
	// [4] = [] tokenMessenger
	//
	// [5] = [WRITE, SIGNER] messageSentEventData
	//
	// [6] = [] messageTransmitterProgram
	//
	// [7] = [] tokenMessengerMinterProgram
	//
	// [8] = [] systemProgram
	//
	// [9] = [] eventAuthority
	//
	// [10] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewReplaceDepositForBurnInstructionBuilder creates a new `ReplaceDepositForBurn` instruction builder.
func NewReplaceDepositForBurnInstructionBuilder() *ReplaceDepositForBurn {
	nd := &ReplaceDepositForBurn{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *ReplaceDepositForBurn) SetParams(params ReplaceDepositForBurnParams) *ReplaceDepositForBurn {
	inst.Params = &params
	return inst
}

// SetOwnerAccount sets the "owner" account.
func (inst *ReplaceDepositForBurn) SetOwnerAccount(owner ag_solanago.PublicKey) *ReplaceDepositForBurn {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *ReplaceDepositForBurn) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetEventRentPayerAccount sets the "eventRentPayer" account.
func (inst *ReplaceDepositForBurn) SetEventRentPayerAccount(eventRentPayer ag_solanago.PublicKey) *ReplaceDepositForBurn {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(eventRentPayer).WRITE().SIGNER()
	return inst
}

// GetEventRentPayerAccount gets the "eventRentPayer" account.
func (inst *ReplaceDepositForBurn) GetEventRentPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSenderAuthorityPdaAccount sets the "senderAuthorityPda" account.
func (inst *ReplaceDepositForBurn) SetSenderAuthorityPdaAccount(senderAuthorityPda ag_solanago.PublicKey) *ReplaceDepositForBurn {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(senderAuthorityPda)
	return inst
}

// GetSenderAuthorityPdaAccount gets the "senderAuthorityPda" account.
func (inst *ReplaceDepositForBurn) GetSenderAuthorityPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMessageTransmitterAccount sets the "messageTransmitter" account.
func (inst *ReplaceDepositForBurn) SetMessageTransmitterAccount(messageTransmitter ag_solanago.PublicKey) *ReplaceDepositForBurn {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(messageTransmitter).WRITE()
	return inst
}

// GetMessageTransmitterAccount gets the "messageTransmitter" account.
func (inst *ReplaceDepositForBurn) GetMessageTransmitterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenMessengerAccount sets the "tokenMessenger" account.
func (inst *ReplaceDepositForBurn) SetTokenMessengerAccount(tokenMessenger ag_solanago.PublicKey) *ReplaceDepositForBurn {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenMessenger)
	return inst
}

// GetTokenMessengerAccount gets the "tokenMessenger" account.
func (inst *ReplaceDepositForBurn) GetTokenMessengerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMessageSentEventDataAccount sets the "messageSentEventData" account.
func (inst *ReplaceDepositForBurn) SetMessageSentEventDataAccount(messageSentEventData ag_solanago.PublicKey) *ReplaceDepositForBurn {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(messageSentEventData).WRITE().SIGNER()
	return inst
}

// GetMessageSentEventDataAccount gets the "messageSentEventData" account.
func (inst *ReplaceDepositForBurn) GetMessageSentEventDataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMessageTransmitterProgramAccount sets the "messageTransmitterProgram" account.
func (inst *ReplaceDepositForBurn) SetMessageTransmitterProgramAccount(messageTransmitterProgram ag_solanago.PublicKey) *ReplaceDepositForBurn {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(messageTransmitterProgram)
	return inst
}

// GetMessageTransmitterProgramAccount gets the "messageTransmitterProgram" account.
func (inst *ReplaceDepositForBurn) GetMessageTransmitterProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenMessengerMinterProgramAccount sets the "tokenMessengerMinterProgram" account.
func (inst *ReplaceDepositForBurn) SetTokenMessengerMinterProgramAccount(tokenMessengerMinterProgram ag_solanago.PublicKey) *ReplaceDepositForBurn {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenMessengerMinterProgram)
	return inst
}

// GetTokenMessengerMinterProgramAccount gets the "tokenMessengerMinterProgram" account.
func (inst *ReplaceDepositForBurn) GetTokenMessengerMinterProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *ReplaceDepositForBurn) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ReplaceDepositForBurn {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *ReplaceDepositForBurn) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *ReplaceDepositForBurn) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *ReplaceDepositForBurn {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *ReplaceDepositForBurn) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetProgramAccount sets the "program" account.
func (inst *ReplaceDepositForBurn) SetProgramAccount(program ag_solanago.PublicKey) *ReplaceDepositForBurn {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *ReplaceDepositForBurn) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst ReplaceDepositForBurn) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ReplaceDepositForBurn,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ReplaceDepositForBurn) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ReplaceDepositForBurn) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.EventRentPayer is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SenderAuthorityPda is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MessageTransmitter is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenMessenger is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.MessageSentEventData is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.MessageTransmitterProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenMessengerMinterProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
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

func (inst *ReplaceDepositForBurn) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ReplaceDepositForBurn")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                      owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             eventRentPayer", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         senderAuthorityPda", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         messageTransmitter", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             tokenMessenger", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       messageSentEventData", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("  messageTransmitterProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("tokenMessengerMinterProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("              systemProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("             eventAuthority", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                    program", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj ReplaceDepositForBurn) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ReplaceDepositForBurn) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewReplaceDepositForBurnInstruction declares a new ReplaceDepositForBurn instruction with the provided parameters and accounts.
func NewReplaceDepositForBurnInstruction(
	// Parameters:
	params ReplaceDepositForBurnParams,
	// Accounts:
	owner ag_solanago.PublicKey,
	eventRentPayer ag_solanago.PublicKey,
	senderAuthorityPda ag_solanago.PublicKey,
	messageTransmitter ag_solanago.PublicKey,
	tokenMessenger ag_solanago.PublicKey,
	messageSentEventData ag_solanago.PublicKey,
	messageTransmitterProgram ag_solanago.PublicKey,
	tokenMessengerMinterProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *ReplaceDepositForBurn {
	return NewReplaceDepositForBurnInstructionBuilder().
		SetParams(params).
		SetOwnerAccount(owner).
		SetEventRentPayerAccount(eventRentPayer).
		SetSenderAuthorityPdaAccount(senderAuthorityPda).
		SetMessageTransmitterAccount(messageTransmitter).
		SetTokenMessengerAccount(tokenMessenger).
		SetMessageSentEventDataAccount(messageSentEventData).
		SetMessageTransmitterProgramAccount(messageTransmitterProgram).
		SetTokenMessengerMinterProgramAccount(tokenMessengerMinterProgram).
		SetSystemProgramAccount(systemProgram).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
