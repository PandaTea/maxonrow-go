package nonfungible

import (
	"fmt"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var prefixAuthorised = []byte("nft/authorised")
var prefixProvider = []byte("nft/provider")
var prefixIssuer = []byte("nft/issuer")
var prefixNonFungibleOwner = []byte("nft/owner")
var prefixNonFungibleItem = []byte("nft/item")

// keys
func getAuthorisedKey() []byte {
	return prefixAuthorised
}

func getProviderKey() []byte {
	return prefixProvider
}

func getIssuerKey() []byte {
	return prefixIssuer
}

func getTokenKey(symbol string) []byte {
	return []byte(fmt.Sprintf("nft/symbol:%s", symbol))
}

func getNonFungibleItemOwnerKey(symbol string, itemID []byte) []byte {
	key := make([]byte, 0, len(prefixNonFungibleOwner)+1+len(symbol)+1+len(itemID))
	key = append(key, prefixNonFungibleOwner...)
	key = append(key, ':')
	key = append(key, []byte(symbol)...)
	key = append(key, ':')
	key = append(key, itemID...)
	return key
}

func getNonFungibleItemKey(symbol string, itemID []byte) []byte {
	key := make([]byte, 0, len(prefixNonFungibleItem)+1+len(symbol)+1+len(itemID))
	key = append(key, prefixNonFungibleItem...)
	key = append(key, ':')
	key = append(key, []byte(symbol)...)
	key = append(key, ':')
	key = append(key, itemID...)
	return key
}

func getMintItemLimitKey(symbol string, owner sdkTypes.AccAddress) []byte {
	key := make([]byte, 0, len(symbol)+1+len(owner.Bytes()))
	key = append(key, []byte(symbol)...)
	key = append(key, ':')
	key = append(key, owner.Bytes()...)
	return key
}
