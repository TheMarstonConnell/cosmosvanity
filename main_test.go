package main

import "testing"

import "github.com/stretchr/testify/require"

func TestGenerateWallet(t *testing.T) {
	w := generateWallet()
	require.Equal(t, w.Address[:7], "cosmos1", "Incorrect bech32 prefix")
	require.Equal(t, len(w.Address), 45, "Incorrect privkey length")
	require.Equal(t, len(w.Pubkey), 33, "Incorrect pubkey length")
	require.Equal(t, len(w.Privkey), 32, "Incorrect privkey length")
}

func TestFindMatchingWallet(t *testing.T) {
	m := matcher{StartsWith: "a"}
	w := findMatchingWallet(m)
	require.Equal(t, w.Address[:8], "cosmos1a", "Incorrect address prefix")
}

func TestFindMatchingWalletMultiProcess(t *testing.T) {
	m := matcher{EndsWith: "a"}
	w := findMatchingWalletMultiProcess(m)
	require.Equal(t, w.Address[len(w.Address)-1:], "a", "Incorrect address suffix")
}