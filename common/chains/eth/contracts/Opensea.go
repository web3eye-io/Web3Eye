// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AdditionalRecipient is an auto generated low-level Go binding around an user-defined struct.
type AdditionalRecipient struct {
	Amount    *big.Int
	Recipient common.Address
}

// AdvancedOrder is an auto generated low-level Go binding around an user-defined struct.
type AdvancedOrder struct {
	Parameters  OrderParameters
	Numerator   *big.Int
	Denominator *big.Int
	Signature   []byte
	ExtraData   []byte
}

// BasicOrderParameters is an auto generated low-level Go binding around an user-defined struct.
type BasicOrderParameters struct {
	ConsiderationToken                common.Address
	ConsiderationIdentifier           *big.Int
	ConsiderationAmount               *big.Int
	Offerer                           common.Address
	Zone                              common.Address
	OfferToken                        common.Address
	OfferIdentifier                   *big.Int
	OfferAmount                       *big.Int
	BasicOrderType                    uint8
	StartTime                         *big.Int
	EndTime                           *big.Int
	ZoneHash                          [32]byte
	Salt                              *big.Int
	OffererConduitKey                 [32]byte
	FulfillerConduitKey               [32]byte
	TotalOriginalAdditionalRecipients *big.Int
	AdditionalRecipients              []AdditionalRecipient
	Signature                         []byte
}

// ConsiderationItem is an auto generated low-level Go binding around an user-defined struct.
type ConsiderationItem struct {
	ItemType             uint8
	Token                common.Address
	IdentifierOrCriteria *big.Int
	StartAmount          *big.Int
	EndAmount            *big.Int
	Recipient            common.Address
}

// CriteriaResolver is an auto generated low-level Go binding around an user-defined struct.
type CriteriaResolver struct {
	OrderIndex    *big.Int
	Side          uint8
	Index         *big.Int
	Identifier    *big.Int
	CriteriaProof [][32]byte
}

// Execution is an auto generated low-level Go binding around an user-defined struct.
type Execution struct {
	Item       ReceivedItem
	Offerer    common.Address
	ConduitKey [32]byte
}

// Fulfillment is an auto generated low-level Go binding around an user-defined struct.
type Fulfillment struct {
	OfferComponents         []FulfillmentComponent
	ConsiderationComponents []FulfillmentComponent
}

// FulfillmentComponent is an auto generated low-level Go binding around an user-defined struct.
type FulfillmentComponent struct {
	OrderIndex *big.Int
	ItemIndex  *big.Int
}

// OfferItem is an auto generated low-level Go binding around an user-defined struct.
type OfferItem struct {
	ItemType             uint8
	Token                common.Address
	IdentifierOrCriteria *big.Int
	StartAmount          *big.Int
	EndAmount            *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Parameters OrderParameters
	Signature  []byte
}

// OrderComponents is an auto generated low-level Go binding around an user-defined struct.
type OrderComponents struct {
	Offerer       common.Address
	Zone          common.Address
	Offer         []OfferItem
	Consideration []ConsiderationItem
	OrderType     uint8
	StartTime     *big.Int
	EndTime       *big.Int
	ZoneHash      [32]byte
	Salt          *big.Int
	ConduitKey    [32]byte
	Counter       *big.Int
}

// OrderParameters is an auto generated low-level Go binding around an user-defined struct.
type OrderParameters struct {
	Offerer                         common.Address
	Zone                            common.Address
	Offer                           []OfferItem
	Consideration                   []ConsiderationItem
	OrderType                       uint8
	StartTime                       *big.Int
	EndTime                         *big.Int
	ZoneHash                        [32]byte
	Salt                            *big.Int
	ConduitKey                      [32]byte
	TotalOriginalConsiderationItems *big.Int
}

// ReceivedItem is an auto generated low-level Go binding around an user-defined struct.
type ReceivedItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
	Recipient  common.Address
}

// SpentItem is an auto generated low-level Go binding around an user-defined struct.
type SpentItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
}

// OpenseaMetaData contains all meta data concerning the Opensea contract.
var OpenseaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadContractSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"BadSignatureV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotCancelOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsiderationCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsiderationLengthNotEqualToTotalOriginal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfallAmount\",\"type\":\"uint256\"}],\"name\":\"ConsiderationNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CriteriaNotEnabledForItem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InexactFraction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientNativeTokensSupplied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invalid1155BatchTransferEncoding\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBasicOrderParameterEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidCallToConduit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidConduit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidContractOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFulfillmentComponentData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNativeOfferItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidRestrictedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"InvalidTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fulfillmentIndex\",\"type\":\"uint256\"}],\"name\":\"MismatchedFulfillmentOfferAndConsiderationComponents\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"MissingFulfillmentComponentOnAggregation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingOriginalConsiderationItems\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeTokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReentrantCalls\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpecifiedOrdersAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferAndConsiderationRequiredOnFulfillment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"OrderCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderIsCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderPartiallyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillsNotEnabledForOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"}],\"name\":\"UnresolvedConsiderationCriteria\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"}],\"name\":\"UnresolvedOfferCriteria\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnusedItemParameters\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"CounterIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSpentItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structReceivedItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOrderParameters\",\"name\":\"orderParameters\",\"type\":\"tuple\"}],\"name\":\"OrderValidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"fulfillAdvancedOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableAdvancedOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"considerationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"considerationIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBasicOrderType\",\"name\":\"basicOrderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offererConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalAdditionalRecipients\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structAdditionalRecipient[]\",\"name\":\"additionalRecipients\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"fulfillBasicOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"considerationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"considerationIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBasicOrderType\",\"name\":\"basicOrderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offererConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalAdditionalRecipients\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structAdditionalRecipient[]\",\"name\":\"additionalRecipients\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"fulfillBasicOrder_efficient_6GL6yc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"}],\"name\":\"fulfillOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractOfferer\",\"type\":\"address\"}],\"name\":\"getContractOffererNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"getCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"information\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"matchAdvancedOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"matchOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OpenseaABI is the input ABI used to generate the binding from.
// Deprecated: Use OpenseaMetaData.ABI instead.
var OpenseaABI = OpenseaMetaData.ABI

// Opensea is an auto generated Go binding around an Ethereum contract.
type Opensea struct {
	OpenseaCaller     // Read-only binding to the contract
	OpenseaTransactor // Write-only binding to the contract
	OpenseaFilterer   // Log filterer for contract events
}

// OpenseaCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpenseaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpenseaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpenseaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpenseaSession struct {
	Contract     *Opensea          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpenseaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpenseaCallerSession struct {
	Contract *OpenseaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OpenseaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpenseaTransactorSession struct {
	Contract     *OpenseaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OpenseaRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpenseaRaw struct {
	Contract *Opensea // Generic contract binding to access the raw methods on
}

// OpenseaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpenseaCallerRaw struct {
	Contract *OpenseaCaller // Generic read-only contract binding to access the raw methods on
}

// OpenseaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpenseaTransactorRaw struct {
	Contract *OpenseaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpensea creates a new instance of Opensea, bound to a specific deployed contract.
func NewOpensea(address common.Address, backend bind.ContractBackend) (*Opensea, error) {
	contract, err := bindOpensea(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Opensea{OpenseaCaller: OpenseaCaller{contract: contract}, OpenseaTransactor: OpenseaTransactor{contract: contract}, OpenseaFilterer: OpenseaFilterer{contract: contract}}, nil
}

// NewOpenseaCaller creates a new read-only instance of Opensea, bound to a specific deployed contract.
func NewOpenseaCaller(address common.Address, caller bind.ContractCaller) (*OpenseaCaller, error) {
	contract, err := bindOpensea(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpenseaCaller{contract: contract}, nil
}

// NewOpenseaTransactor creates a new write-only instance of Opensea, bound to a specific deployed contract.
func NewOpenseaTransactor(address common.Address, transactor bind.ContractTransactor) (*OpenseaTransactor, error) {
	contract, err := bindOpensea(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpenseaTransactor{contract: contract}, nil
}

// NewOpenseaFilterer creates a new log filterer instance of Opensea, bound to a specific deployed contract.
func NewOpenseaFilterer(address common.Address, filterer bind.ContractFilterer) (*OpenseaFilterer, error) {
	contract, err := bindOpensea(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpenseaFilterer{contract: contract}, nil
}

// bindOpensea binds a generic wrapper to an already deployed contract.
func bindOpensea(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OpenseaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opensea *OpenseaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opensea.Contract.OpenseaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opensea *OpenseaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opensea.Contract.OpenseaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opensea *OpenseaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opensea.Contract.OpenseaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opensea *OpenseaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opensea.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opensea *OpenseaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opensea.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opensea *OpenseaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opensea.Contract.contract.Transact(opts, method, params...)
}

// GetContractOffererNonce is a free data retrieval call binding the contract method 0xa900866b.
//
// Solidity: function getContractOffererNonce(address contractOfferer) view returns(uint256 nonce)
func (_Opensea *OpenseaCaller) GetContractOffererNonce(opts *bind.CallOpts, contractOfferer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getContractOffererNonce", contractOfferer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractOffererNonce is a free data retrieval call binding the contract method 0xa900866b.
//
// Solidity: function getContractOffererNonce(address contractOfferer) view returns(uint256 nonce)
func (_Opensea *OpenseaSession) GetContractOffererNonce(contractOfferer common.Address) (*big.Int, error) {
	return _Opensea.Contract.GetContractOffererNonce(&_Opensea.CallOpts, contractOfferer)
}

// GetContractOffererNonce is a free data retrieval call binding the contract method 0xa900866b.
//
// Solidity: function getContractOffererNonce(address contractOfferer) view returns(uint256 nonce)
func (_Opensea *OpenseaCallerSession) GetContractOffererNonce(contractOfferer common.Address) (*big.Int, error) {
	return _Opensea.Contract.GetContractOffererNonce(&_Opensea.CallOpts, contractOfferer)
}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Opensea *OpenseaCaller) GetCounter(opts *bind.CallOpts, offerer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getCounter", offerer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Opensea *OpenseaSession) GetCounter(offerer common.Address) (*big.Int, error) {
	return _Opensea.Contract.GetCounter(&_Opensea.CallOpts, offerer)
}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Opensea *OpenseaCallerSession) GetCounter(offerer common.Address) (*big.Int, error) {
	return _Opensea.Contract.GetCounter(&_Opensea.CallOpts, offerer)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) ) view returns(bytes32 orderHash)
func (_Opensea *OpenseaCaller) GetOrderHash(opts *bind.CallOpts, arg0 OrderComponents) ([32]byte, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getOrderHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) ) view returns(bytes32 orderHash)
func (_Opensea *OpenseaSession) GetOrderHash(arg0 OrderComponents) ([32]byte, error) {
	return _Opensea.Contract.GetOrderHash(&_Opensea.CallOpts, arg0)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) ) view returns(bytes32 orderHash)
func (_Opensea *OpenseaCallerSession) GetOrderHash(arg0 OrderComponents) ([32]byte, error) {
	return _Opensea.Contract.GetOrderHash(&_Opensea.CallOpts, arg0)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_Opensea *OpenseaCaller) GetOrderStatus(opts *bind.CallOpts, orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getOrderStatus", orderHash)

	outstruct := new(struct {
		IsValidated bool
		IsCancelled bool
		TotalFilled *big.Int
		TotalSize   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValidated = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsCancelled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.TotalFilled = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalSize = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_Opensea *OpenseaSession) GetOrderStatus(orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	return _Opensea.Contract.GetOrderStatus(&_Opensea.CallOpts, orderHash)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_Opensea *OpenseaCallerSession) GetOrderStatus(orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	return _Opensea.Contract.GetOrderStatus(&_Opensea.CallOpts, orderHash)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_Opensea *OpenseaCaller) Information(opts *bind.CallOpts) (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "information")

	outstruct := new(struct {
		Version           string
		DomainSeparator   [32]byte
		ConduitController common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Version = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DomainSeparator = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ConduitController = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_Opensea *OpenseaSession) Information() (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	return _Opensea.Contract.Information(&_Opensea.CallOpts)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_Opensea *OpenseaCallerSession) Information() (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	return _Opensea.Contract.Information(&_Opensea.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Opensea *OpenseaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Opensea *OpenseaSession) Name() (string, error) {
	return _Opensea.Contract.Name(&_Opensea.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Opensea *OpenseaCallerSession) Name() (string, error) {
	return _Opensea.Contract.Name(&_Opensea.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Opensea *OpenseaTransactor) Cancel(opts *bind.TransactOpts, orders []OrderComponents) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "cancel", orders)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Opensea *OpenseaSession) Cancel(orders []OrderComponents) (*types.Transaction, error) {
	return _Opensea.Contract.Cancel(&_Opensea.TransactOpts, orders)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Opensea *OpenseaTransactorSession) Cancel(orders []OrderComponents) (*types.Transaction, error) {
	return _Opensea.Contract.Cancel(&_Opensea.TransactOpts, orders)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) , (uint256,uint8,uint256,uint256,bytes32[])[] , bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Opensea *OpenseaTransactor) FulfillAdvancedOrder(opts *bind.TransactOpts, arg0 AdvancedOrder, arg1 []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "fulfillAdvancedOrder", arg0, arg1, fulfillerConduitKey, recipient)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) , (uint256,uint8,uint256,uint256,bytes32[])[] , bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Opensea *OpenseaSession) FulfillAdvancedOrder(arg0 AdvancedOrder, arg1 []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillAdvancedOrder(&_Opensea.TransactOpts, arg0, arg1, fulfillerConduitKey, recipient)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) , (uint256,uint8,uint256,uint256,bytes32[])[] , bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Opensea *OpenseaTransactorSession) FulfillAdvancedOrder(arg0 AdvancedOrder, arg1 []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillAdvancedOrder(&_Opensea.TransactOpts, arg0, arg1, fulfillerConduitKey, recipient)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaTransactor) FulfillAvailableAdvancedOrders(opts *bind.TransactOpts, arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 [][]FulfillmentComponent, arg3 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "fulfillAvailableAdvancedOrders", arg0, arg1, arg2, arg3, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaSession) FulfillAvailableAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 [][]FulfillmentComponent, arg3 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillAvailableAdvancedOrders(&_Opensea.TransactOpts, arg0, arg1, arg2, arg3, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaTransactorSession) FulfillAvailableAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 [][]FulfillmentComponent, arg3 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillAvailableAdvancedOrders(&_Opensea.TransactOpts, arg0, arg1, arg2, arg3, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaTransactor) FulfillAvailableOrders(opts *bind.TransactOpts, arg0 []Order, arg1 [][]FulfillmentComponent, arg2 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "fulfillAvailableOrders", arg0, arg1, arg2, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaSession) FulfillAvailableOrders(arg0 []Order, arg1 [][]FulfillmentComponent, arg2 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillAvailableOrders(&_Opensea.TransactOpts, arg0, arg1, arg2, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaTransactorSession) FulfillAvailableOrders(arg0 []Order, arg1 [][]FulfillmentComponent, arg2 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillAvailableOrders(&_Opensea.TransactOpts, arg0, arg1, arg2, fulfillerConduitKey, maximumFulfilled)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Opensea *OpenseaTransactor) FulfillBasicOrder(opts *bind.TransactOpts, parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "fulfillBasicOrder", parameters)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Opensea *OpenseaSession) FulfillBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillBasicOrder(&_Opensea.TransactOpts, parameters)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Opensea *OpenseaTransactorSession) FulfillBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillBasicOrder(&_Opensea.TransactOpts, parameters)
}

// FulfillBasicOrderEfficient6GL6yc is a paid mutator transaction binding the contract method 0x00000000.
//
// Solidity: function fulfillBasicOrder_efficient_6GL6yc((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Opensea *OpenseaTransactor) FulfillBasicOrderEfficient6GL6yc(opts *bind.TransactOpts, parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "fulfillBasicOrder_efficient_6GL6yc", parameters)
}

// FulfillBasicOrderEfficient6GL6yc is a paid mutator transaction binding the contract method 0x00000000.
//
// Solidity: function fulfillBasicOrder_efficient_6GL6yc((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Opensea *OpenseaSession) FulfillBasicOrderEfficient6GL6yc(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillBasicOrderEfficient6GL6yc(&_Opensea.TransactOpts, parameters)
}

// FulfillBasicOrderEfficient6GL6yc is a paid mutator transaction binding the contract method 0x00000000.
//
// Solidity: function fulfillBasicOrder_efficient_6GL6yc((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Opensea *OpenseaTransactorSession) FulfillBasicOrderEfficient6GL6yc(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillBasicOrderEfficient6GL6yc(&_Opensea.TransactOpts, parameters)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) , bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Opensea *OpenseaTransactor) FulfillOrder(opts *bind.TransactOpts, arg0 Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "fulfillOrder", arg0, fulfillerConduitKey)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) , bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Opensea *OpenseaSession) FulfillOrder(arg0 Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillOrder(&_Opensea.TransactOpts, arg0, fulfillerConduitKey)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) , bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Opensea *OpenseaTransactorSession) FulfillOrder(arg0 Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Opensea.Contract.FulfillOrder(&_Opensea.TransactOpts, arg0, fulfillerConduitKey)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Opensea *OpenseaTransactor) IncrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "incrementCounter")
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Opensea *OpenseaSession) IncrementCounter() (*types.Transaction, error) {
	return _Opensea.Contract.IncrementCounter(&_Opensea.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Opensea *OpenseaTransactorSession) IncrementCounter() (*types.Transaction, error) {
	return _Opensea.Contract.IncrementCounter(&_Opensea.TransactOpts)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0xf2d12b12.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , ((uint256,uint256)[],(uint256,uint256)[])[] , address recipient) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaTransactor) MatchAdvancedOrders(opts *bind.TransactOpts, arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 []Fulfillment, recipient common.Address) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "matchAdvancedOrders", arg0, arg1, arg2, recipient)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0xf2d12b12.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , ((uint256,uint256)[],(uint256,uint256)[])[] , address recipient) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaSession) MatchAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 []Fulfillment, recipient common.Address) (*types.Transaction, error) {
	return _Opensea.Contract.MatchAdvancedOrders(&_Opensea.TransactOpts, arg0, arg1, arg2, recipient)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0xf2d12b12.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , ((uint256,uint256)[],(uint256,uint256)[])[] , address recipient) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaTransactorSession) MatchAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 []Fulfillment, recipient common.Address) (*types.Transaction, error) {
	return _Opensea.Contract.MatchAdvancedOrders(&_Opensea.TransactOpts, arg0, arg1, arg2, recipient)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , ((uint256,uint256)[],(uint256,uint256)[])[] ) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaTransactor) MatchOrders(opts *bind.TransactOpts, arg0 []Order, arg1 []Fulfillment) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "matchOrders", arg0, arg1)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , ((uint256,uint256)[],(uint256,uint256)[])[] ) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaSession) MatchOrders(arg0 []Order, arg1 []Fulfillment) (*types.Transaction, error) {
	return _Opensea.Contract.MatchOrders(&_Opensea.TransactOpts, arg0, arg1)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , ((uint256,uint256)[],(uint256,uint256)[])[] ) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Opensea *OpenseaTransactorSession) MatchOrders(arg0 []Order, arg1 []Fulfillment) (*types.Transaction, error) {
	return _Opensea.Contract.MatchOrders(&_Opensea.TransactOpts, arg0, arg1)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] ) returns(bool)
func (_Opensea *OpenseaTransactor) Validate(opts *bind.TransactOpts, arg0 []Order) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "validate", arg0)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] ) returns(bool)
func (_Opensea *OpenseaSession) Validate(arg0 []Order) (*types.Transaction, error) {
	return _Opensea.Contract.Validate(&_Opensea.TransactOpts, arg0)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] ) returns(bool)
func (_Opensea *OpenseaTransactorSession) Validate(arg0 []Order) (*types.Transaction, error) {
	return _Opensea.Contract.Validate(&_Opensea.TransactOpts, arg0)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Opensea *OpenseaTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opensea.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Opensea *OpenseaSession) Receive() (*types.Transaction, error) {
	return _Opensea.Contract.Receive(&_Opensea.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Opensea *OpenseaTransactorSession) Receive() (*types.Transaction, error) {
	return _Opensea.Contract.Receive(&_Opensea.TransactOpts)
}

// OpenseaCounterIncrementedIterator is returned from FilterCounterIncremented and is used to iterate over the raw logs and unpacked data for CounterIncremented events raised by the Opensea contract.
type OpenseaCounterIncrementedIterator struct {
	Event *OpenseaCounterIncremented // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OpenseaCounterIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaCounterIncremented)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OpenseaCounterIncremented)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OpenseaCounterIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaCounterIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaCounterIncremented represents a CounterIncremented event raised by the Opensea contract.
type OpenseaCounterIncremented struct {
	NewCounter *big.Int
	Offerer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCounterIncremented is a free log retrieval operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Opensea *OpenseaFilterer) FilterCounterIncremented(opts *bind.FilterOpts, offerer []common.Address) (*OpenseaCounterIncrementedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaCounterIncrementedIterator{contract: _Opensea.contract, event: "CounterIncremented", logs: logs, sub: sub}, nil
}

// WatchCounterIncremented is a free log subscription operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Opensea *OpenseaFilterer) WatchCounterIncremented(opts *bind.WatchOpts, sink chan<- *OpenseaCounterIncremented, offerer []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaCounterIncremented)
				if err := _Opensea.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCounterIncremented is a log parse operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Opensea *OpenseaFilterer) ParseCounterIncremented(log types.Log) (*OpenseaCounterIncremented, error) {
	event := new(OpenseaCounterIncremented)
	if err := _Opensea.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Opensea contract.
type OpenseaOrderCancelledIterator struct {
	Event *OpenseaOrderCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OpenseaOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaOrderCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OpenseaOrderCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OpenseaOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaOrderCancelled represents a OrderCancelled event raised by the Opensea contract.
type OpenseaOrderCancelled struct {
	OrderHash [32]byte
	Offerer   common.Address
	Zone      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Opensea *OpenseaFilterer) FilterOrderCancelled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*OpenseaOrderCancelledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaOrderCancelledIterator{contract: _Opensea.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Opensea *OpenseaFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *OpenseaOrderCancelled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaOrderCancelled)
				if err := _Opensea.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderCancelled is a log parse operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Opensea *OpenseaFilterer) ParseOrderCancelled(log types.Log) (*OpenseaOrderCancelled, error) {
	event := new(OpenseaOrderCancelled)
	if err := _Opensea.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the Opensea contract.
type OpenseaOrderFulfilledIterator struct {
	Event *OpenseaOrderFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OpenseaOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaOrderFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OpenseaOrderFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OpenseaOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaOrderFulfilled represents a OrderFulfilled event raised by the Opensea contract.
type OpenseaOrderFulfilled struct {
	OrderHash     [32]byte
	Offerer       common.Address
	Zone          common.Address
	Recipient     common.Address
	Offer         []SpentItem
	Consideration []ReceivedItem
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_Opensea *OpenseaFilterer) FilterOrderFulfilled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*OpenseaOrderFulfilledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaOrderFulfilledIterator{contract: _Opensea.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_Opensea *OpenseaFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *OpenseaOrderFulfilled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaOrderFulfilled)
				if err := _Opensea.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderFulfilled is a log parse operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_Opensea *OpenseaFilterer) ParseOrderFulfilled(log types.Log) (*OpenseaOrderFulfilled, error) {
	event := new(OpenseaOrderFulfilled)
	if err := _Opensea.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaOrderValidatedIterator is returned from FilterOrderValidated and is used to iterate over the raw logs and unpacked data for OrderValidated events raised by the Opensea contract.
type OpenseaOrderValidatedIterator struct {
	Event *OpenseaOrderValidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OpenseaOrderValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaOrderValidated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OpenseaOrderValidated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OpenseaOrderValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaOrderValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaOrderValidated represents a OrderValidated event raised by the Opensea contract.
type OpenseaOrderValidated struct {
	OrderHash       [32]byte
	OrderParameters OrderParameters
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderValidated is a free log retrieval operation binding the contract event 0xf280791efe782edcf06ce15c8f4dff17601db3b88eb3805a0db7d77faf757f04.
//
// Solidity: event OrderValidated(bytes32 orderHash, (address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) orderParameters)
func (_Opensea *OpenseaFilterer) FilterOrderValidated(opts *bind.FilterOpts) (*OpenseaOrderValidatedIterator, error) {

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "OrderValidated")
	if err != nil {
		return nil, err
	}
	return &OpenseaOrderValidatedIterator{contract: _Opensea.contract, event: "OrderValidated", logs: logs, sub: sub}, nil
}

// WatchOrderValidated is a free log subscription operation binding the contract event 0xf280791efe782edcf06ce15c8f4dff17601db3b88eb3805a0db7d77faf757f04.
//
// Solidity: event OrderValidated(bytes32 orderHash, (address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) orderParameters)
func (_Opensea *OpenseaFilterer) WatchOrderValidated(opts *bind.WatchOpts, sink chan<- *OpenseaOrderValidated) (event.Subscription, error) {

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "OrderValidated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaOrderValidated)
				if err := _Opensea.contract.UnpackLog(event, "OrderValidated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderValidated is a log parse operation binding the contract event 0xf280791efe782edcf06ce15c8f4dff17601db3b88eb3805a0db7d77faf757f04.
//
// Solidity: event OrderValidated(bytes32 orderHash, (address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) orderParameters)
func (_Opensea *OpenseaFilterer) ParseOrderValidated(log types.Log) (*OpenseaOrderValidated, error) {
	event := new(OpenseaOrderValidated)
	if err := _Opensea.contract.UnpackLog(event, "OrderValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the Opensea contract.
type OpenseaOrdersMatchedIterator struct {
	Event *OpenseaOrdersMatched // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OpenseaOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaOrdersMatched)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OpenseaOrdersMatched)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OpenseaOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaOrdersMatched represents a OrdersMatched event raised by the Opensea contract.
type OpenseaOrdersMatched struct {
	OrderHashes [][32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x4b9f2d36e1b4c93de62cc077b00b1a91d84b6c31b4a14e012718dcca230689e7.
//
// Solidity: event OrdersMatched(bytes32[] orderHashes)
func (_Opensea *OpenseaFilterer) FilterOrdersMatched(opts *bind.FilterOpts) (*OpenseaOrdersMatchedIterator, error) {

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "OrdersMatched")
	if err != nil {
		return nil, err
	}
	return &OpenseaOrdersMatchedIterator{contract: _Opensea.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x4b9f2d36e1b4c93de62cc077b00b1a91d84b6c31b4a14e012718dcca230689e7.
//
// Solidity: event OrdersMatched(bytes32[] orderHashes)
func (_Opensea *OpenseaFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *OpenseaOrdersMatched) (event.Subscription, error) {

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "OrdersMatched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaOrdersMatched)
				if err := _Opensea.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrdersMatched is a log parse operation binding the contract event 0x4b9f2d36e1b4c93de62cc077b00b1a91d84b6c31b4a14e012718dcca230689e7.
//
// Solidity: event OrdersMatched(bytes32[] orderHashes)
func (_Opensea *OpenseaFilterer) ParseOrdersMatched(log types.Log) (*OpenseaOrdersMatched, error) {
	event := new(OpenseaOrdersMatched)
	if err := _Opensea.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
