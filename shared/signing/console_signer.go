package signing

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"os"
	"strings"
)

type ConsoleSigner struct {
	Signer
}

func NewConsoleSigner(signer Signer) *ConsoleSigner {
	return &ConsoleSigner{
		Signer: signer,
	}
}

func (s *ConsoleSigner) Sign(chainID string, tx Transaction) (Transaction, error) {
	txJSON, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	fmt.Println()
	fmt.Println(">>> Sending Transaction <<<")
	fmt.Println("  Chain ID:", chainID)
	fmt.Println("  Tx Data:")
	fmt.Println(string(txJSON))
	fmt.Println()

	if !promptYesNo("Do you want to send this transaction?") {
		return nil, errors.New("user abandoned transaction")
	}

	return s.Signer.Sign(context.Background(), chainID, tx)
}

func promptYesNo(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s (y/n): ", question)

		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response. Please try again.")
			continue
		}

		response = strings.TrimSpace(response)
		response = strings.ToLower(response)

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		} else {
			fmt.Println("Invalid response. Please answer 'y' or 'n'.")
		}
	}
}
