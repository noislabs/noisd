package main

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/bech32"
)

// Cmd creates a main CLI command
func Bech32Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bech32",
		Short: "Tool for helping encoding/decoding bech32 addresses",
		RunE:  client.ValidateCmd,
	}

	cmd.AddCommand(Bech32EncodeCmd())
	cmd.AddCommand(Bech32DecodeCmd())

	return cmd
}

func Bech32EncodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encode [prefix] [bech32 string]",
		Short: "Encode any bech32 or hex string to the [prefix] address",
		Long: `Encode any bech32 or hex string to the [prefix] address
Example:
	noisd bech32 encode nois cosmos18afy9mwgrmtpsl4vhgfvq0r3knm4tw2yftqlpg
	noisd bech32 encode nois 3F5242EDC81ED6187EACBA12C03C71B4F755B944
	noisd bech32 encode noisvaloper nois18afy9mwgrmtpsl4vhgfvq0r3knm4tw2ycra4ds
	noisd bech32 encode nois noisvaloper18afy9mwgrmtpsl4vhgfvq0r3knm4tw2yer055y
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

var flagHexFormat = "hex"

func Bech32DecodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decode [bech32 string]",
		Short: "Decode any bech32 or hex string raw bytes",
		Long: `Decode any bech32 or hex string raw bytes
Example:
	noisd bech32 decode cosmos18afy9mwgrmtpsl4vhgfvq0r3knm4tw2yftqlpg
	noisd bech32 decode 3F5242EDC81ED6187EACBA12C03C71B4F755B944
	noisd bech32 decode nois18afy9mwgrmtpsl4vhgfvq0r3knm4tw2ycra4ds
	noisd bech32 decode noisvaloper18afy9mwgrmtpsl4vhgfvq0r3knm4tw2yer055y
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
			format := "%d\n"
			hex, _ := cmd.Flags().GetBool(flagHexFormat)
			if err != nil {
				return err
			}

			if hex {
				format = "%X\n"
			}
			fmt.Printf(format, bz)
			return nil
		},
	}
	cmd.Flags().Bool(flagHexFormat, false, "Output raw bytes in hex format")
	return cmd
}
