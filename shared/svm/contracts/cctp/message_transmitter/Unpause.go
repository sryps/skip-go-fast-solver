// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package message_transmitter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Unpause is the `unpause` instruction.
type Unpause struct {
	Params *UnpauseParams

	// [0] = [SIGNER] pauser
	//
	// [1] = [WRITE] messageTransmitter
	//
	// [2] = [] eventAuthority
	//
	// [3] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUnpauseInstructionBuilder creates a new `Unpause` instruction builder.
func NewUnpauseInstructionBuilder() *Unpause {
	nd := &Unpause{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *Unpause) SetParams(params UnpauseParams) *Unpause {
	inst.Params = &params
	return inst
}

// SetPauserAccount sets the "pauser" account.
func (inst *Unpause) SetPauserAccount(pauser ag_solanago.PublicKey) *Unpause {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(pauser).SIGNER()
	return inst
}

// GetPauserAccount gets the "pauser" account.
func (inst *Unpause) GetPauserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMessageTransmitterAccount sets the "messageTransmitter" account.
func (inst *Unpause) SetMessageTransmitterAccount(messageTransmitter ag_solanago.PublicKey) *Unpause {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(messageTransmitter).WRITE()
	return inst
}

// GetMessageTransmitterAccount gets the "messageTransmitter" account.
func (inst *Unpause) GetMessageTransmitterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *Unpause) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *Unpause {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *Unpause) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetProgramAccount sets the "program" account.
func (inst *Unpause) SetProgramAccount(program ag_solanago.PublicKey) *Unpause {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *Unpause) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst Unpause) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Unpause,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Unpause) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Unpause) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Pauser is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MessageTransmitter is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *Unpause) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Unpause")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("            pauser", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("messageTransmitter", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    eventAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           program", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj Unpause) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Unpause) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewUnpauseInstruction declares a new Unpause instruction with the provided parameters and accounts.
func NewUnpauseInstruction(
	// Parameters:
	params UnpauseParams,
	// Accounts:
	pauser ag_solanago.PublicKey,
	messageTransmitter ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *Unpause {
	return NewUnpauseInstructionBuilder().
		SetParams(params).
		SetPauserAccount(pauser).
		SetMessageTransmitterAccount(messageTransmitter).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
