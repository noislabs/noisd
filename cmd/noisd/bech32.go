package main

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/types/bech32"
)

func Bech32EncodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bech32-encode [prefix] [bech32 string]",
		Short: "Encode any bech32 or hex string to the [prefix] address",
		Long: `Encode any bech32 or hex string to the [prefix] address
Example:
	noisd debug bech32-encode nois cosmos18afy9mwgrmtpsl4vhgfvq0r3knm4tw2yftqlpg
	noisd debug bech32-encode nois 3F5242EDC81ED6187EACBA12C03C71B4F755B944
	noisd debug bech32-encode noisvaloper nois18afy9mwgrmtpsl4vhgfvq0r3knm4tw2ycra4ds
	noisd debug bech32-encode nois noisvaloper18afy9mwgrmtpsl4vhgfvq0r3knm4tw2yer055y
	`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			bech32prefix := args[0]

			addrString := args[1]

			// bytes of the parsed address
			var bz []byte
			var err error

			// try decoding hex first
			bz, err = hex.DecodeString(addrString)
			if err != nil {
				// try decoding any bech32 address
				_, bz, err = bech32.DecodeAndConvert(addrString)
				if err != nil {
					return err
				}
			}

			// convert to desired bech32 prefix
			bech32Addr, err := bech32.ConvertAndEncode(bech32prefix, bz)
			if err != nil {
				return err
			}
			cmd.Println(bech32Addr)
			return nil
		},
	}

	return cmd
}

func Bech32DecodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bech32-decode [bech32 string]",
		Short: "Decode any bech32 or hex string raw bytes",
		Long: `Decode any bech32 or hex string raw bytes
Example:
	noisd debug bech32-decode cosmos18afy9mwgrmtpsl4vhgfvq0r3knm4tw2yftqlpg
	noisd debug bech32-decode 3F5242EDC81ED6187EACBA12C03C71B4F755B944
	noisd debug bech32-encode nois18afy9mwgrmtpsl4vhgfvq0r3knm4tw2ycra4ds
	noisd debug bech32-encode noisvaloper18afy9mwgrmtpsl4vhgfvq0r3knm4tw2yer055y
	`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addrString := args[0]

			// bytes of the parsed address
			var bz []byte
			var err error

			// try decoding hex first
			bz, err = hex.DecodeString(addrString)
			if err != nil {
				// try decoding any bech32 address
				_, bz, err = bech32.DecodeAndConvert(addrString)
				if err != nil {
					return err
				}
			}

			fmt.Printf("%d\n", bz)
			return nil
		},
	}

	return cmd
}
