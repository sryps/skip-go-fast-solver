// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package message_transmitter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SendMessageWithCaller is the `sendMessageWithCaller` instruction.
type SendMessageWithCaller struct {
	Params *SendMessageWithCallerParams

	// [0] = [WRITE, SIGNER] eventRentPayer
	//
	// [1] = [SIGNER] senderAuthorityPda
	//
	// [2] = [WRITE] messageTransmitter
	//
	// [3] = [WRITE, SIGNER] messageSentEventData
	//
	// [4] = [] senderProgram
	//
	// [5] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSendMessageWithCallerInstructionBuilder creates a new `SendMessageWithCaller` instruction builder.
func NewSendMessageWithCallerInstructionBuilder() *SendMessageWithCaller {
	nd := &SendMessageWithCaller{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *SendMessageWithCaller) SetParams(params SendMessageWithCallerParams) *SendMessageWithCaller {
	inst.Params = &params
	return inst
}

// SetEventRentPayerAccount sets the "eventRentPayer" account.
func (inst *SendMessageWithCaller) SetEventRentPayerAccount(eventRentPayer ag_solanago.PublicKey) *SendMessageWithCaller {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(eventRentPayer).WRITE().SIGNER()
	return inst
}

// GetEventRentPayerAccount gets the "eventRentPayer" account.
func (inst *SendMessageWithCaller) GetEventRentPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSenderAuthorityPdaAccount sets the "senderAuthorityPda" account.
func (inst *SendMessageWithCaller) SetSenderAuthorityPdaAccount(senderAuthorityPda ag_solanago.PublicKey) *SendMessageWithCaller {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(senderAuthorityPda).SIGNER()
	return inst
}

// GetSenderAuthorityPdaAccount gets the "senderAuthorityPda" account.
func (inst *SendMessageWithCaller) GetSenderAuthorityPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMessageTransmitterAccount sets the "messageTransmitter" account.
func (inst *SendMessageWithCaller) SetMessageTransmitterAccount(messageTransmitter ag_solanago.PublicKey) *SendMessageWithCaller {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(messageTransmitter).WRITE()
	return inst
}

// GetMessageTransmitterAccount gets the "messageTransmitter" account.
func (inst *SendMessageWithCaller) GetMessageTransmitterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMessageSentEventDataAccount sets the "messageSentEventData" account.
func (inst *SendMessageWithCaller) SetMessageSentEventDataAccount(messageSentEventData ag_solanago.PublicKey) *SendMessageWithCaller {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(messageSentEventData).WRITE().SIGNER()
	return inst
}

// GetMessageSentEventDataAccount gets the "messageSentEventData" account.
func (inst *SendMessageWithCaller) GetMessageSentEventDataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSenderProgramAccount sets the "senderProgram" account.
func (inst *SendMessageWithCaller) SetSenderProgramAccount(senderProgram ag_solanago.PublicKey) *SendMessageWithCaller {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(senderProgram)
	return inst
}

// GetSenderProgramAccount gets the "senderProgram" account.
func (inst *SendMessageWithCaller) GetSenderProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *SendMessageWithCaller) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SendMessageWithCaller {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *SendMessageWithCaller) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst SendMessageWithCaller) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SendMessageWithCaller,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SendMessageWithCaller) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SendMessageWithCaller) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.EventRentPayer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SenderAuthorityPda is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MessageTransmitter is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MessageSentEventData is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SenderProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *SendMessageWithCaller) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SendMessageWithCaller")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      eventRentPayer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  senderAuthorityPda", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  messageTransmitter", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("messageSentEventData", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       senderProgram", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj SendMessageWithCaller) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SendMessageWithCaller) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewSendMessageWithCallerInstruction declares a new SendMessageWithCaller instruction with the provided parameters and accounts.
func NewSendMessageWithCallerInstruction(
	// Parameters:
	params SendMessageWithCallerParams,
	// Accounts:
	eventRentPayer ag_solanago.PublicKey,
	senderAuthorityPda ag_solanago.PublicKey,
	messageTransmitter ag_solanago.PublicKey,
	messageSentEventData ag_solanago.PublicKey,
	senderProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *SendMessageWithCaller {
	return NewSendMessageWithCallerInstructionBuilder().
		SetParams(params).
		SetEventRentPayerAccount(eventRentPayer).
		SetSenderAuthorityPdaAccount(senderAuthorityPda).
		SetMessageTransmitterAccount(messageTransmitter).
		SetMessageSentEventDataAccount(messageSentEventData).
		SetSenderProgramAccount(senderProgram).
		SetSystemProgramAccount(systemProgram)
}
