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

// IViraCampaignsCampaign is an auto generated low-level Go binding around an user-defined struct.
type IViraCampaignsCampaign struct {
	Id          string
	Start       *big.Int
	End         *big.Int
	RelInv      *big.Int
	Slice       *big.Int
	PctPerSlice *big.Int
	Locked      bool
	Revoked     bool
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"ApprovalValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"campaignId\",\"type\":\"string\"}],\"name\":\"CampaignUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReleasedAllCampaignsProfit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReleasedInvestment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReleasedTokenProfit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"metadataDescriptor\",\"type\":\"address\"}],\"name\":\"SetMetadataDescriptor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_oldSlot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_newSlot\",\"type\":\"uint256\"}],\"name\":\"SlotChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"TransferValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"campaignId_\",\"type\":\"string\"}],\"name\":\"campaignData\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctPerSlice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"}],\"internalType\":\"structIViraCampaigns.Campaign\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot_\",\"type\":\"uint256\"}],\"name\":\"campaignData\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctPerSlice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"}],\"internalType\":\"structIViraCampaigns.Campaign\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataDescriptor\",\"outputs\":[{\"internalType\":\"contractIERC3525MetadataDescriptorUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"campaignId_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"mintValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder_\",\"type\":\"address\"}],\"name\":\"releasableAddressProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"releasableTokenProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseAllCampaignsProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"releaseCampaignProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"releaseInvestment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved_\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"slotOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot_\",\"type\":\"uint256\"}],\"name\":\"slotURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCampaigns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctPerSlice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"}],\"internalType\":\"structIViraCampaigns.Campaign\",\"name\":\"campaign_\",\"type\":\"tuple\"}],\"name\":\"updateCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Main *MainCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Main *MainSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Main.Contract.DEFAULTADMINROLE(&_Main.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Main *MainCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Main.Contract.DEFAULTADMINROLE(&_Main.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Main *MainCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Main *MainSession) MINTERROLE() ([32]byte, error) {
	return _Main.Contract.MINTERROLE(&_Main.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Main *MainCallerSession) MINTERROLE() ([32]byte, error) {
	return _Main.Contract.MINTERROLE(&_Main.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Main *MainCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Main *MainSession) UPGRADERROLE() ([32]byte, error) {
	return _Main.Contract.UPGRADERROLE(&_Main.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Main *MainCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Main.Contract.UPGRADERROLE(&_Main.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Main *MainCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Main *MainSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Main.Contract.UPGRADEINTERFACEVERSION(&_Main.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Main *MainCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Main.Contract.UPGRADEINTERFACEVERSION(&_Main.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xe345e0bc.
//
// Solidity: function allowance(uint256 tokenId_, address operator_) view returns(uint256)
func (_Main *MainCaller) Allowance(opts *bind.CallOpts, tokenId_ *big.Int, operator_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "allowance", tokenId_, operator_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xe345e0bc.
//
// Solidity: function allowance(uint256 tokenId_, address operator_) view returns(uint256)
func (_Main *MainSession) Allowance(tokenId_ *big.Int, operator_ common.Address) (*big.Int, error) {
	return _Main.Contract.Allowance(&_Main.CallOpts, tokenId_, operator_)
}

// Allowance is a free data retrieval call binding the contract method 0xe345e0bc.
//
// Solidity: function allowance(uint256 tokenId_, address operator_) view returns(uint256)
func (_Main *MainCallerSession) Allowance(tokenId_ *big.Int, operator_ common.Address) (*big.Int, error) {
	return _Main.Contract.Allowance(&_Main.CallOpts, tokenId_, operator_)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner_) view returns(uint256 balance)
func (_Main *MainCaller) BalanceOf(opts *bind.CallOpts, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "balanceOf", owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner_) view returns(uint256 balance)
func (_Main *MainSession) BalanceOf(owner_ common.Address) (*big.Int, error) {
	return _Main.Contract.BalanceOf(&_Main.CallOpts, owner_)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner_) view returns(uint256 balance)
func (_Main *MainCallerSession) BalanceOf(owner_ common.Address) (*big.Int, error) {
	return _Main.Contract.BalanceOf(&_Main.CallOpts, owner_)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 tokenId_) view returns(uint256)
func (_Main *MainCaller) BalanceOf0(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "balanceOf0", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf0 is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 tokenId_) view returns(uint256)
func (_Main *MainSession) BalanceOf0(tokenId_ *big.Int) (*big.Int, error) {
	return _Main.Contract.BalanceOf0(&_Main.CallOpts, tokenId_)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 tokenId_) view returns(uint256)
func (_Main *MainCallerSession) BalanceOf0(tokenId_ *big.Int) (*big.Int, error) {
	return _Main.Contract.BalanceOf0(&_Main.CallOpts, tokenId_)
}

// CampaignData is a free data retrieval call binding the contract method 0x5798e048.
//
// Solidity: function campaignData(string campaignId_) view returns((string,uint256,uint256,uint256,uint256,uint256,bool,bool))
func (_Main *MainCaller) CampaignData(opts *bind.CallOpts, campaignId_ string) (IViraCampaignsCampaign, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "campaignData", campaignId_)

	if err != nil {
		return *new(IViraCampaignsCampaign), err
	}

	out0 := *abi.ConvertType(out[0], new(IViraCampaignsCampaign)).(*IViraCampaignsCampaign)

	return out0, err

}

// CampaignData is a free data retrieval call binding the contract method 0x5798e048.
//
// Solidity: function campaignData(string campaignId_) view returns((string,uint256,uint256,uint256,uint256,uint256,bool,bool))
func (_Main *MainSession) CampaignData(campaignId_ string) (IViraCampaignsCampaign, error) {
	return _Main.Contract.CampaignData(&_Main.CallOpts, campaignId_)
}

// CampaignData is a free data retrieval call binding the contract method 0x5798e048.
//
// Solidity: function campaignData(string campaignId_) view returns((string,uint256,uint256,uint256,uint256,uint256,bool,bool))
func (_Main *MainCallerSession) CampaignData(campaignId_ string) (IViraCampaignsCampaign, error) {
	return _Main.Contract.CampaignData(&_Main.CallOpts, campaignId_)
}

// CampaignData0 is a free data retrieval call binding the contract method 0xd5abb06b.
//
// Solidity: function campaignData(uint256 slot_) view returns((string,uint256,uint256,uint256,uint256,uint256,bool,bool))
func (_Main *MainCaller) CampaignData0(opts *bind.CallOpts, slot_ *big.Int) (IViraCampaignsCampaign, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "campaignData0", slot_)

	if err != nil {
		return *new(IViraCampaignsCampaign), err
	}

	out0 := *abi.ConvertType(out[0], new(IViraCampaignsCampaign)).(*IViraCampaignsCampaign)

	return out0, err

}

// CampaignData0 is a free data retrieval call binding the contract method 0xd5abb06b.
//
// Solidity: function campaignData(uint256 slot_) view returns((string,uint256,uint256,uint256,uint256,uint256,bool,bool))
func (_Main *MainSession) CampaignData0(slot_ *big.Int) (IViraCampaignsCampaign, error) {
	return _Main.Contract.CampaignData0(&_Main.CallOpts, slot_)
}

// CampaignData0 is a free data retrieval call binding the contract method 0xd5abb06b.
//
// Solidity: function campaignData(uint256 slot_) view returns((string,uint256,uint256,uint256,uint256,uint256,bool,bool))
func (_Main *MainCallerSession) CampaignData0(slot_ *big.Int) (IViraCampaignsCampaign, error) {
	return _Main.Contract.CampaignData0(&_Main.CallOpts, slot_)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Main *MainCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Main *MainSession) ContractURI() (string, error) {
	return _Main.Contract.ContractURI(&_Main.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Main *MainCallerSession) ContractURI() (string, error) {
	return _Main.Contract.ContractURI(&_Main.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId_) view returns(address)
func (_Main *MainCaller) GetApproved(opts *bind.CallOpts, tokenId_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getApproved", tokenId_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId_) view returns(address)
func (_Main *MainSession) GetApproved(tokenId_ *big.Int) (common.Address, error) {
	return _Main.Contract.GetApproved(&_Main.CallOpts, tokenId_)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId_) view returns(address)
func (_Main *MainCallerSession) GetApproved(tokenId_ *big.Int) (common.Address, error) {
	return _Main.Contract.GetApproved(&_Main.CallOpts, tokenId_)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Main *MainCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Main *MainSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Main.Contract.GetRoleAdmin(&_Main.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Main *MainCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Main.Contract.GetRoleAdmin(&_Main.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Main *MainCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Main *MainSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Main.Contract.HasRole(&_Main.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Main *MainCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Main.Contract.HasRole(&_Main.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner_, address operator_) view returns(bool)
func (_Main *MainCaller) IsApprovedForAll(opts *bind.CallOpts, owner_ common.Address, operator_ common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "isApprovedForAll", owner_, operator_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner_, address operator_) view returns(bool)
func (_Main *MainSession) IsApprovedForAll(owner_ common.Address, operator_ common.Address) (bool, error) {
	return _Main.Contract.IsApprovedForAll(&_Main.CallOpts, owner_, operator_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner_, address operator_) view returns(bool)
func (_Main *MainCallerSession) IsApprovedForAll(owner_ common.Address, operator_ common.Address) (bool, error) {
	return _Main.Contract.IsApprovedForAll(&_Main.CallOpts, owner_, operator_)
}

// MetadataDescriptor is a free data retrieval call binding the contract method 0x840f7113.
//
// Solidity: function metadataDescriptor() view returns(address)
func (_Main *MainCaller) MetadataDescriptor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "metadataDescriptor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataDescriptor is a free data retrieval call binding the contract method 0x840f7113.
//
// Solidity: function metadataDescriptor() view returns(address)
func (_Main *MainSession) MetadataDescriptor() (common.Address, error) {
	return _Main.Contract.MetadataDescriptor(&_Main.CallOpts)
}

// MetadataDescriptor is a free data retrieval call binding the contract method 0x840f7113.
//
// Solidity: function metadataDescriptor() view returns(address)
func (_Main *MainCallerSession) MetadataDescriptor() (common.Address, error) {
	return _Main.Contract.MetadataDescriptor(&_Main.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainSession) Name() (string, error) {
	return _Main.Contract.Name(&_Main.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainCallerSession) Name() (string, error) {
	return _Main.Contract.Name(&_Main.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId_) view returns(address owner_)
func (_Main *MainCaller) OwnerOf(opts *bind.CallOpts, tokenId_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "ownerOf", tokenId_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId_) view returns(address owner_)
func (_Main *MainSession) OwnerOf(tokenId_ *big.Int) (common.Address, error) {
	return _Main.Contract.OwnerOf(&_Main.CallOpts, tokenId_)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId_) view returns(address owner_)
func (_Main *MainCallerSession) OwnerOf(tokenId_ *big.Int) (common.Address, error) {
	return _Main.Contract.OwnerOf(&_Main.CallOpts, tokenId_)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Main *MainCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Main *MainSession) ProxiableUUID() ([32]byte, error) {
	return _Main.Contract.ProxiableUUID(&_Main.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Main *MainCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Main.Contract.ProxiableUUID(&_Main.CallOpts)
}

// ReleasableAddressProfit is a free data retrieval call binding the contract method 0x94374147.
//
// Solidity: function releasableAddressProfit(address holder_) view returns(uint256)
func (_Main *MainCaller) ReleasableAddressProfit(opts *bind.CallOpts, holder_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "releasableAddressProfit", holder_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleasableAddressProfit is a free data retrieval call binding the contract method 0x94374147.
//
// Solidity: function releasableAddressProfit(address holder_) view returns(uint256)
func (_Main *MainSession) ReleasableAddressProfit(holder_ common.Address) (*big.Int, error) {
	return _Main.Contract.ReleasableAddressProfit(&_Main.CallOpts, holder_)
}

// ReleasableAddressProfit is a free data retrieval call binding the contract method 0x94374147.
//
// Solidity: function releasableAddressProfit(address holder_) view returns(uint256)
func (_Main *MainCallerSession) ReleasableAddressProfit(holder_ common.Address) (*big.Int, error) {
	return _Main.Contract.ReleasableAddressProfit(&_Main.CallOpts, holder_)
}

// ReleasableTokenProfit is a free data retrieval call binding the contract method 0x51d661e9.
//
// Solidity: function releasableTokenProfit(uint256 tokenId_) view returns(uint256)
func (_Main *MainCaller) ReleasableTokenProfit(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "releasableTokenProfit", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleasableTokenProfit is a free data retrieval call binding the contract method 0x51d661e9.
//
// Solidity: function releasableTokenProfit(uint256 tokenId_) view returns(uint256)
func (_Main *MainSession) ReleasableTokenProfit(tokenId_ *big.Int) (*big.Int, error) {
	return _Main.Contract.ReleasableTokenProfit(&_Main.CallOpts, tokenId_)
}

// ReleasableTokenProfit is a free data retrieval call binding the contract method 0x51d661e9.
//
// Solidity: function releasableTokenProfit(uint256 tokenId_) view returns(uint256)
func (_Main *MainCallerSession) ReleasableTokenProfit(tokenId_ *big.Int) (*big.Int, error) {
	return _Main.Contract.ReleasableTokenProfit(&_Main.CallOpts, tokenId_)
}

// SlotOf is a free data retrieval call binding the contract method 0x263f3e7e.
//
// Solidity: function slotOf(uint256 tokenId_) view returns(uint256)
func (_Main *MainCaller) SlotOf(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "slotOf", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotOf is a free data retrieval call binding the contract method 0x263f3e7e.
//
// Solidity: function slotOf(uint256 tokenId_) view returns(uint256)
func (_Main *MainSession) SlotOf(tokenId_ *big.Int) (*big.Int, error) {
	return _Main.Contract.SlotOf(&_Main.CallOpts, tokenId_)
}

// SlotOf is a free data retrieval call binding the contract method 0x263f3e7e.
//
// Solidity: function slotOf(uint256 tokenId_) view returns(uint256)
func (_Main *MainCallerSession) SlotOf(tokenId_ *big.Int) (*big.Int, error) {
	return _Main.Contract.SlotOf(&_Main.CallOpts, tokenId_)
}

// SlotURI is a free data retrieval call binding the contract method 0x09c3dd87.
//
// Solidity: function slotURI(uint256 slot_) view returns(string)
func (_Main *MainCaller) SlotURI(opts *bind.CallOpts, slot_ *big.Int) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "slotURI", slot_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SlotURI is a free data retrieval call binding the contract method 0x09c3dd87.
//
// Solidity: function slotURI(uint256 slot_) view returns(string)
func (_Main *MainSession) SlotURI(slot_ *big.Int) (string, error) {
	return _Main.Contract.SlotURI(&_Main.CallOpts, slot_)
}

// SlotURI is a free data retrieval call binding the contract method 0x09c3dd87.
//
// Solidity: function slotURI(uint256 slot_) view returns(string)
func (_Main *MainCallerSession) SlotURI(slot_ *big.Int) (string, error) {
	return _Main.Contract.SlotURI(&_Main.CallOpts, slot_)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Main *MainCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Main *MainSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Main.Contract.SupportsInterface(&_Main.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Main *MainCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Main.Contract.SupportsInterface(&_Main.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainSession) Symbol() (string, error) {
	return _Main.Contract.Symbol(&_Main.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainCallerSession) Symbol() (string, error) {
	return _Main.Contract.Symbol(&_Main.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index_) view returns(uint256)
func (_Main *MainCaller) TokenByIndex(opts *bind.CallOpts, index_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "tokenByIndex", index_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index_) view returns(uint256)
func (_Main *MainSession) TokenByIndex(index_ *big.Int) (*big.Int, error) {
	return _Main.Contract.TokenByIndex(&_Main.CallOpts, index_)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index_) view returns(uint256)
func (_Main *MainCallerSession) TokenByIndex(index_ *big.Int) (*big.Int, error) {
	return _Main.Contract.TokenByIndex(&_Main.CallOpts, index_)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner_, uint256 index_) view returns(uint256)
func (_Main *MainCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner_ common.Address, index_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner_, index_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner_, uint256 index_) view returns(uint256)
func (_Main *MainSession) TokenOfOwnerByIndex(owner_ common.Address, index_ *big.Int) (*big.Int, error) {
	return _Main.Contract.TokenOfOwnerByIndex(&_Main.CallOpts, owner_, index_)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner_, uint256 index_) view returns(uint256)
func (_Main *MainCallerSession) TokenOfOwnerByIndex(owner_ common.Address, index_ *big.Int) (*big.Int, error) {
	return _Main.Contract.TokenOfOwnerByIndex(&_Main.CallOpts, owner_, index_)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Main *MainCaller) TokenURI(opts *bind.CallOpts, tokenId_ *big.Int) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "tokenURI", tokenId_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Main *MainSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _Main.Contract.TokenURI(&_Main.CallOpts, tokenId_)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Main *MainCallerSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _Main.Contract.TokenURI(&_Main.CallOpts, tokenId_)
}

// TotalCampaigns is a free data retrieval call binding the contract method 0x02932f56.
//
// Solidity: function totalCampaigns() view returns(uint256)
func (_Main *MainCaller) TotalCampaigns(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "totalCampaigns")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCampaigns is a free data retrieval call binding the contract method 0x02932f56.
//
// Solidity: function totalCampaigns() view returns(uint256)
func (_Main *MainSession) TotalCampaigns() (*big.Int, error) {
	return _Main.Contract.TotalCampaigns(&_Main.CallOpts)
}

// TotalCampaigns is a free data retrieval call binding the contract method 0x02932f56.
//
// Solidity: function totalCampaigns() view returns(uint256)
func (_Main *MainCallerSession) TotalCampaigns() (*big.Int, error) {
	return _Main.Contract.TotalCampaigns(&_Main.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainSession) TotalSupply() (*big.Int, error) {
	return _Main.Contract.TotalSupply(&_Main.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainCallerSession) TotalSupply() (*big.Int, error) {
	return _Main.Contract.TotalSupply(&_Main.CallOpts)
}

// ValueDecimals is a free data retrieval call binding the contract method 0x3e7e8669.
//
// Solidity: function valueDecimals() view returns(uint8)
func (_Main *MainCaller) ValueDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "valueDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValueDecimals is a free data retrieval call binding the contract method 0x3e7e8669.
//
// Solidity: function valueDecimals() view returns(uint8)
func (_Main *MainSession) ValueDecimals() (uint8, error) {
	return _Main.Contract.ValueDecimals(&_Main.CallOpts)
}

// ValueDecimals is a free data retrieval call binding the contract method 0x3e7e8669.
//
// Solidity: function valueDecimals() view returns(uint8)
func (_Main *MainCallerSession) ValueDecimals() (uint8, error) {
	return _Main.Contract.ValueDecimals(&_Main.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to_, uint256 tokenId_) payable returns()
func (_Main *MainTransactor) Approve(opts *bind.TransactOpts, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approve", to_, tokenId_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to_, uint256 tokenId_) payable returns()
func (_Main *MainSession) Approve(to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Approve(&_Main.TransactOpts, to_, tokenId_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to_, uint256 tokenId_) payable returns()
func (_Main *MainTransactorSession) Approve(to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Approve(&_Main.TransactOpts, to_, tokenId_)
}

// Approve0 is a paid mutator transaction binding the contract method 0x8cb0a511.
//
// Solidity: function approve(uint256 tokenId_, address to_, uint256 value_) payable returns()
func (_Main *MainTransactor) Approve0(opts *bind.TransactOpts, tokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approve0", tokenId_, to_, value_)
}

// Approve0 is a paid mutator transaction binding the contract method 0x8cb0a511.
//
// Solidity: function approve(uint256 tokenId_, address to_, uint256 value_) payable returns()
func (_Main *MainSession) Approve0(tokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Approve0(&_Main.TransactOpts, tokenId_, to_, value_)
}

// Approve0 is a paid mutator transaction binding the contract method 0x8cb0a511.
//
// Solidity: function approve(uint256 tokenId_, address to_, uint256 value_) payable returns()
func (_Main *MainTransactorSession) Approve0(tokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Approve0(&_Main.TransactOpts, tokenId_, to_, value_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Main *MainTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Main *MainSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Main.Contract.GrantRole(&_Main.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Main *MainTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Main.Contract.GrantRole(&_Main.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address minter_, address token_) returns()
func (_Main *MainTransactor) Initialize(opts *bind.TransactOpts, minter_ common.Address, token_ common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "initialize", minter_, token_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address minter_, address token_) returns()
func (_Main *MainSession) Initialize(minter_ common.Address, token_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.Initialize(&_Main.TransactOpts, minter_, token_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address minter_, address token_) returns()
func (_Main *MainTransactorSession) Initialize(minter_ common.Address, token_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.Initialize(&_Main.TransactOpts, minter_, token_)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address to_, string campaignId_, uint256 value_) returns()
func (_Main *MainTransactor) Mint(opts *bind.TransactOpts, to_ common.Address, campaignId_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "mint", to_, campaignId_, value_)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address to_, string campaignId_, uint256 value_) returns()
func (_Main *MainSession) Mint(to_ common.Address, campaignId_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Mint(&_Main.TransactOpts, to_, campaignId_, value_)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address to_, string campaignId_, uint256 value_) returns()
func (_Main *MainTransactorSession) Mint(to_ common.Address, campaignId_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Mint(&_Main.TransactOpts, to_, campaignId_, value_)
}

// MintValue is a paid mutator transaction binding the contract method 0xf0e88e7f.
//
// Solidity: function mintValue(uint256 tokenId_, uint256 value_) returns()
func (_Main *MainTransactor) MintValue(opts *bind.TransactOpts, tokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "mintValue", tokenId_, value_)
}

// MintValue is a paid mutator transaction binding the contract method 0xf0e88e7f.
//
// Solidity: function mintValue(uint256 tokenId_, uint256 value_) returns()
func (_Main *MainSession) MintValue(tokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.MintValue(&_Main.TransactOpts, tokenId_, value_)
}

// MintValue is a paid mutator transaction binding the contract method 0xf0e88e7f.
//
// Solidity: function mintValue(uint256 tokenId_, uint256 value_) returns()
func (_Main *MainTransactorSession) MintValue(tokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.MintValue(&_Main.TransactOpts, tokenId_, value_)
}

// ReleaseAllCampaignsProfit is a paid mutator transaction binding the contract method 0x9670263c.
//
// Solidity: function releaseAllCampaignsProfit() returns()
func (_Main *MainTransactor) ReleaseAllCampaignsProfit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "releaseAllCampaignsProfit")
}

// ReleaseAllCampaignsProfit is a paid mutator transaction binding the contract method 0x9670263c.
//
// Solidity: function releaseAllCampaignsProfit() returns()
func (_Main *MainSession) ReleaseAllCampaignsProfit() (*types.Transaction, error) {
	return _Main.Contract.ReleaseAllCampaignsProfit(&_Main.TransactOpts)
}

// ReleaseAllCampaignsProfit is a paid mutator transaction binding the contract method 0x9670263c.
//
// Solidity: function releaseAllCampaignsProfit() returns()
func (_Main *MainTransactorSession) ReleaseAllCampaignsProfit() (*types.Transaction, error) {
	return _Main.Contract.ReleaseAllCampaignsProfit(&_Main.TransactOpts)
}

// ReleaseCampaignProfit is a paid mutator transaction binding the contract method 0x4af81954.
//
// Solidity: function releaseCampaignProfit(uint256 tokenId_, uint256 amount_) returns()
func (_Main *MainTransactor) ReleaseCampaignProfit(opts *bind.TransactOpts, tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "releaseCampaignProfit", tokenId_, amount_)
}

// ReleaseCampaignProfit is a paid mutator transaction binding the contract method 0x4af81954.
//
// Solidity: function releaseCampaignProfit(uint256 tokenId_, uint256 amount_) returns()
func (_Main *MainSession) ReleaseCampaignProfit(tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ReleaseCampaignProfit(&_Main.TransactOpts, tokenId_, amount_)
}

// ReleaseCampaignProfit is a paid mutator transaction binding the contract method 0x4af81954.
//
// Solidity: function releaseCampaignProfit(uint256 tokenId_, uint256 amount_) returns()
func (_Main *MainTransactorSession) ReleaseCampaignProfit(tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ReleaseCampaignProfit(&_Main.TransactOpts, tokenId_, amount_)
}

// ReleaseInvestment is a paid mutator transaction binding the contract method 0x886a53ab.
//
// Solidity: function releaseInvestment(uint256 tokenId_) returns()
func (_Main *MainTransactor) ReleaseInvestment(opts *bind.TransactOpts, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "releaseInvestment", tokenId_)
}

// ReleaseInvestment is a paid mutator transaction binding the contract method 0x886a53ab.
//
// Solidity: function releaseInvestment(uint256 tokenId_) returns()
func (_Main *MainSession) ReleaseInvestment(tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ReleaseInvestment(&_Main.TransactOpts, tokenId_)
}

// ReleaseInvestment is a paid mutator transaction binding the contract method 0x886a53ab.
//
// Solidity: function releaseInvestment(uint256 tokenId_) returns()
func (_Main *MainTransactorSession) ReleaseInvestment(tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ReleaseInvestment(&_Main.TransactOpts, tokenId_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Main *MainTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Main *MainSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Main.Contract.RenounceRole(&_Main.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Main *MainTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Main.Contract.RenounceRole(&_Main.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Main *MainTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Main *MainSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Main.Contract.RevokeRole(&_Main.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Main *MainTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Main.Contract.RevokeRole(&_Main.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Main *MainTransactor) SafeTransferFrom(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "safeTransferFrom", from_, to_, tokenId_)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Main *MainSession) SafeTransferFrom(from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SafeTransferFrom(&_Main.TransactOpts, from_, to_, tokenId_)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Main *MainTransactorSession) SafeTransferFrom(from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SafeTransferFrom(&_Main.TransactOpts, from_, to_, tokenId_)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_, bytes data_) payable returns()
func (_Main *MainTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, tokenId_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "safeTransferFrom0", from_, to_, tokenId_, data_)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_, bytes data_) payable returns()
func (_Main *MainSession) SafeTransferFrom0(from_ common.Address, to_ common.Address, tokenId_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Main.Contract.SafeTransferFrom0(&_Main.TransactOpts, from_, to_, tokenId_, data_)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_, bytes data_) payable returns()
func (_Main *MainTransactorSession) SafeTransferFrom0(from_ common.Address, to_ common.Address, tokenId_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Main.Contract.SafeTransferFrom0(&_Main.TransactOpts, from_, to_, tokenId_, data_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved_) returns()
func (_Main *MainTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator_ common.Address, approved_ bool) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setApprovalForAll", operator_, approved_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved_) returns()
func (_Main *MainSession) SetApprovalForAll(operator_ common.Address, approved_ bool) (*types.Transaction, error) {
	return _Main.Contract.SetApprovalForAll(&_Main.TransactOpts, operator_, approved_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved_) returns()
func (_Main *MainTransactorSession) SetApprovalForAll(operator_ common.Address, approved_ bool) (*types.Transaction, error) {
	return _Main.Contract.SetApprovalForAll(&_Main.TransactOpts, operator_, approved_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0f485c02.
//
// Solidity: function transferFrom(uint256 fromTokenId_, address to_, uint256 value_) payable returns(uint256 newTokenId)
func (_Main *MainTransactor) TransferFrom(opts *bind.TransactOpts, fromTokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferFrom", fromTokenId_, to_, value_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0f485c02.
//
// Solidity: function transferFrom(uint256 fromTokenId_, address to_, uint256 value_) payable returns(uint256 newTokenId)
func (_Main *MainSession) TransferFrom(fromTokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom(&_Main.TransactOpts, fromTokenId_, to_, value_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0f485c02.
//
// Solidity: function transferFrom(uint256 fromTokenId_, address to_, uint256 value_) payable returns(uint256 newTokenId)
func (_Main *MainTransactorSession) TransferFrom(fromTokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom(&_Main.TransactOpts, fromTokenId_, to_, value_)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Main *MainTransactor) TransferFrom0(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferFrom0", from_, to_, tokenId_)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Main *MainSession) TransferFrom0(from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom0(&_Main.TransactOpts, from_, to_, tokenId_)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Main *MainTransactorSession) TransferFrom0(from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom0(&_Main.TransactOpts, from_, to_, tokenId_)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0x310ed7f0.
//
// Solidity: function transferFrom(uint256 fromTokenId_, uint256 toTokenId_, uint256 value_) payable returns()
func (_Main *MainTransactor) TransferFrom1(opts *bind.TransactOpts, fromTokenId_ *big.Int, toTokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferFrom1", fromTokenId_, toTokenId_, value_)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0x310ed7f0.
//
// Solidity: function transferFrom(uint256 fromTokenId_, uint256 toTokenId_, uint256 value_) payable returns()
func (_Main *MainSession) TransferFrom1(fromTokenId_ *big.Int, toTokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom1(&_Main.TransactOpts, fromTokenId_, toTokenId_, value_)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0x310ed7f0.
//
// Solidity: function transferFrom(uint256 fromTokenId_, uint256 toTokenId_, uint256 value_) payable returns()
func (_Main *MainTransactorSession) TransferFrom1(fromTokenId_ *big.Int, toTokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom1(&_Main.TransactOpts, fromTokenId_, toTokenId_, value_)
}

// UpdateCampaign is a paid mutator transaction binding the contract method 0x895040d2.
//
// Solidity: function updateCampaign((string,uint256,uint256,uint256,uint256,uint256,bool,bool) campaign_) returns()
func (_Main *MainTransactor) UpdateCampaign(opts *bind.TransactOpts, campaign_ IViraCampaignsCampaign) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "updateCampaign", campaign_)
}

// UpdateCampaign is a paid mutator transaction binding the contract method 0x895040d2.
//
// Solidity: function updateCampaign((string,uint256,uint256,uint256,uint256,uint256,bool,bool) campaign_) returns()
func (_Main *MainSession) UpdateCampaign(campaign_ IViraCampaignsCampaign) (*types.Transaction, error) {
	return _Main.Contract.UpdateCampaign(&_Main.TransactOpts, campaign_)
}

// UpdateCampaign is a paid mutator transaction binding the contract method 0x895040d2.
//
// Solidity: function updateCampaign((string,uint256,uint256,uint256,uint256,uint256,bool,bool) campaign_) returns()
func (_Main *MainTransactorSession) UpdateCampaign(campaign_ IViraCampaignsCampaign) (*types.Transaction, error) {
	return _Main.Contract.UpdateCampaign(&_Main.TransactOpts, campaign_)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Main *MainTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Main *MainSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Main.Contract.UpgradeToAndCall(&_Main.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Main *MainTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Main.Contract.UpgradeToAndCall(&_Main.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_Main *MainTransactor) Withdraw(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "withdraw", amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_Main *MainSession) Withdraw(amount_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Withdraw(&_Main.TransactOpts, amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_Main *MainTransactorSession) Withdraw(amount_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Withdraw(&_Main.TransactOpts, amount_)
}

// MainApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Main contract.
type MainApprovalIterator struct {
	Event *MainApproval // Event containing the contract specifics and raw log

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
func (it *MainApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainApproval)
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
		it.Event = new(MainApproval)
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
func (it *MainApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainApproval represents a Approval event raised by the Main contract.
type MainApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Main *MainFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*MainApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MainApprovalIterator{contract: _Main.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Main *MainFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MainApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainApproval)
				if err := _Main.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Main *MainFilterer) ParseApproval(log types.Log) (*MainApproval, error) {
	event := new(MainApproval)
	if err := _Main.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Main contract.
type MainApprovalForAllIterator struct {
	Event *MainApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MainApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainApprovalForAll)
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
		it.Event = new(MainApprovalForAll)
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
func (it *MainApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainApprovalForAll represents a ApprovalForAll event raised by the Main contract.
type MainApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Main *MainFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*MainApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &MainApprovalForAllIterator{contract: _Main.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Main *MainFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MainApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainApprovalForAll)
				if err := _Main.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Main *MainFilterer) ParseApprovalForAll(log types.Log) (*MainApprovalForAll, error) {
	event := new(MainApprovalForAll)
	if err := _Main.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainApprovalValueIterator is returned from FilterApprovalValue and is used to iterate over the raw logs and unpacked data for ApprovalValue events raised by the Main contract.
type MainApprovalValueIterator struct {
	Event *MainApprovalValue // Event containing the contract specifics and raw log

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
func (it *MainApprovalValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainApprovalValue)
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
		it.Event = new(MainApprovalValue)
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
func (it *MainApprovalValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainApprovalValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainApprovalValue represents a ApprovalValue event raised by the Main contract.
type MainApprovalValue struct {
	TokenId  *big.Int
	Operator common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalValue is a free log retrieval operation binding the contract event 0x621b050de0ad08b51d19b48b3e6df75348c4de6bdd93e81b252ca62e28265b1b.
//
// Solidity: event ApprovalValue(uint256 indexed _tokenId, address indexed _operator, uint256 _value)
func (_Main *MainFilterer) FilterApprovalValue(opts *bind.FilterOpts, _tokenId []*big.Int, _operator []common.Address) (*MainApprovalValueIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ApprovalValue", _tokenIdRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &MainApprovalValueIterator{contract: _Main.contract, event: "ApprovalValue", logs: logs, sub: sub}, nil
}

// WatchApprovalValue is a free log subscription operation binding the contract event 0x621b050de0ad08b51d19b48b3e6df75348c4de6bdd93e81b252ca62e28265b1b.
//
// Solidity: event ApprovalValue(uint256 indexed _tokenId, address indexed _operator, uint256 _value)
func (_Main *MainFilterer) WatchApprovalValue(opts *bind.WatchOpts, sink chan<- *MainApprovalValue, _tokenId []*big.Int, _operator []common.Address) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ApprovalValue", _tokenIdRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainApprovalValue)
				if err := _Main.contract.UnpackLog(event, "ApprovalValue", log); err != nil {
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

// ParseApprovalValue is a log parse operation binding the contract event 0x621b050de0ad08b51d19b48b3e6df75348c4de6bdd93e81b252ca62e28265b1b.
//
// Solidity: event ApprovalValue(uint256 indexed _tokenId, address indexed _operator, uint256 _value)
func (_Main *MainFilterer) ParseApprovalValue(log types.Log) (*MainApprovalValue, error) {
	event := new(MainApprovalValue)
	if err := _Main.contract.UnpackLog(event, "ApprovalValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainCampaignUpdatedIterator is returned from FilterCampaignUpdated and is used to iterate over the raw logs and unpacked data for CampaignUpdated events raised by the Main contract.
type MainCampaignUpdatedIterator struct {
	Event *MainCampaignUpdated // Event containing the contract specifics and raw log

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
func (it *MainCampaignUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainCampaignUpdated)
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
		it.Event = new(MainCampaignUpdated)
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
func (it *MainCampaignUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainCampaignUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainCampaignUpdated represents a CampaignUpdated event raised by the Main contract.
type MainCampaignUpdated struct {
	CampaignId common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCampaignUpdated is a free log retrieval operation binding the contract event 0x97ebe72425499cfa387c633444ff604615975ceb676d866e2fb7fecd677554c4.
//
// Solidity: event CampaignUpdated(string indexed campaignId)
func (_Main *MainFilterer) FilterCampaignUpdated(opts *bind.FilterOpts, campaignId []string) (*MainCampaignUpdatedIterator, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "CampaignUpdated", campaignIdRule)
	if err != nil {
		return nil, err
	}
	return &MainCampaignUpdatedIterator{contract: _Main.contract, event: "CampaignUpdated", logs: logs, sub: sub}, nil
}

// WatchCampaignUpdated is a free log subscription operation binding the contract event 0x97ebe72425499cfa387c633444ff604615975ceb676d866e2fb7fecd677554c4.
//
// Solidity: event CampaignUpdated(string indexed campaignId)
func (_Main *MainFilterer) WatchCampaignUpdated(opts *bind.WatchOpts, sink chan<- *MainCampaignUpdated, campaignId []string) (event.Subscription, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "CampaignUpdated", campaignIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainCampaignUpdated)
				if err := _Main.contract.UnpackLog(event, "CampaignUpdated", log); err != nil {
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

// ParseCampaignUpdated is a log parse operation binding the contract event 0x97ebe72425499cfa387c633444ff604615975ceb676d866e2fb7fecd677554c4.
//
// Solidity: event CampaignUpdated(string indexed campaignId)
func (_Main *MainFilterer) ParseCampaignUpdated(log types.Log) (*MainCampaignUpdated, error) {
	event := new(MainCampaignUpdated)
	if err := _Main.contract.UnpackLog(event, "CampaignUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Main contract.
type MainInitializedIterator struct {
	Event *MainInitialized // Event containing the contract specifics and raw log

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
func (it *MainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainInitialized)
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
		it.Event = new(MainInitialized)
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
func (it *MainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainInitialized represents a Initialized event raised by the Main contract.
type MainInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Main *MainFilterer) FilterInitialized(opts *bind.FilterOpts) (*MainInitializedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MainInitializedIterator{contract: _Main.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Main *MainFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MainInitialized) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainInitialized)
				if err := _Main.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Main *MainFilterer) ParseInitialized(log types.Log) (*MainInitialized, error) {
	event := new(MainInitialized)
	if err := _Main.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainReleasedAllCampaignsProfitIterator is returned from FilterReleasedAllCampaignsProfit and is used to iterate over the raw logs and unpacked data for ReleasedAllCampaignsProfit events raised by the Main contract.
type MainReleasedAllCampaignsProfitIterator struct {
	Event *MainReleasedAllCampaignsProfit // Event containing the contract specifics and raw log

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
func (it *MainReleasedAllCampaignsProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainReleasedAllCampaignsProfit)
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
		it.Event = new(MainReleasedAllCampaignsProfit)
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
func (it *MainReleasedAllCampaignsProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainReleasedAllCampaignsProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainReleasedAllCampaignsProfit represents a ReleasedAllCampaignsProfit event raised by the Main contract.
type MainReleasedAllCampaignsProfit struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReleasedAllCampaignsProfit is a free log retrieval operation binding the contract event 0xcd6a39600e90ad2871fdfe0e74b19bfdb288a47916da78a3b333825f3c53a2b2.
//
// Solidity: event ReleasedAllCampaignsProfit(address indexed to, uint256 value)
func (_Main *MainFilterer) FilterReleasedAllCampaignsProfit(opts *bind.FilterOpts, to []common.Address) (*MainReleasedAllCampaignsProfitIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ReleasedAllCampaignsProfit", toRule)
	if err != nil {
		return nil, err
	}
	return &MainReleasedAllCampaignsProfitIterator{contract: _Main.contract, event: "ReleasedAllCampaignsProfit", logs: logs, sub: sub}, nil
}

// WatchReleasedAllCampaignsProfit is a free log subscription operation binding the contract event 0xcd6a39600e90ad2871fdfe0e74b19bfdb288a47916da78a3b333825f3c53a2b2.
//
// Solidity: event ReleasedAllCampaignsProfit(address indexed to, uint256 value)
func (_Main *MainFilterer) WatchReleasedAllCampaignsProfit(opts *bind.WatchOpts, sink chan<- *MainReleasedAllCampaignsProfit, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ReleasedAllCampaignsProfit", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainReleasedAllCampaignsProfit)
				if err := _Main.contract.UnpackLog(event, "ReleasedAllCampaignsProfit", log); err != nil {
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

// ParseReleasedAllCampaignsProfit is a log parse operation binding the contract event 0xcd6a39600e90ad2871fdfe0e74b19bfdb288a47916da78a3b333825f3c53a2b2.
//
// Solidity: event ReleasedAllCampaignsProfit(address indexed to, uint256 value)
func (_Main *MainFilterer) ParseReleasedAllCampaignsProfit(log types.Log) (*MainReleasedAllCampaignsProfit, error) {
	event := new(MainReleasedAllCampaignsProfit)
	if err := _Main.contract.UnpackLog(event, "ReleasedAllCampaignsProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainReleasedInvestmentIterator is returned from FilterReleasedInvestment and is used to iterate over the raw logs and unpacked data for ReleasedInvestment events raised by the Main contract.
type MainReleasedInvestmentIterator struct {
	Event *MainReleasedInvestment // Event containing the contract specifics and raw log

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
func (it *MainReleasedInvestmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainReleasedInvestment)
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
		it.Event = new(MainReleasedInvestment)
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
func (it *MainReleasedInvestmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainReleasedInvestmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainReleasedInvestment represents a ReleasedInvestment event raised by the Main contract.
type MainReleasedInvestment struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReleasedInvestment is a free log retrieval operation binding the contract event 0x0ab5ede1200a85ab88ecbbf26df740178a0562bd86203fa564fcbc201ed3ef67.
//
// Solidity: event ReleasedInvestment(address indexed to, uint256 value)
func (_Main *MainFilterer) FilterReleasedInvestment(opts *bind.FilterOpts, to []common.Address) (*MainReleasedInvestmentIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ReleasedInvestment", toRule)
	if err != nil {
		return nil, err
	}
	return &MainReleasedInvestmentIterator{contract: _Main.contract, event: "ReleasedInvestment", logs: logs, sub: sub}, nil
}

// WatchReleasedInvestment is a free log subscription operation binding the contract event 0x0ab5ede1200a85ab88ecbbf26df740178a0562bd86203fa564fcbc201ed3ef67.
//
// Solidity: event ReleasedInvestment(address indexed to, uint256 value)
func (_Main *MainFilterer) WatchReleasedInvestment(opts *bind.WatchOpts, sink chan<- *MainReleasedInvestment, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ReleasedInvestment", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainReleasedInvestment)
				if err := _Main.contract.UnpackLog(event, "ReleasedInvestment", log); err != nil {
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

// ParseReleasedInvestment is a log parse operation binding the contract event 0x0ab5ede1200a85ab88ecbbf26df740178a0562bd86203fa564fcbc201ed3ef67.
//
// Solidity: event ReleasedInvestment(address indexed to, uint256 value)
func (_Main *MainFilterer) ParseReleasedInvestment(log types.Log) (*MainReleasedInvestment, error) {
	event := new(MainReleasedInvestment)
	if err := _Main.contract.UnpackLog(event, "ReleasedInvestment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainReleasedTokenProfitIterator is returned from FilterReleasedTokenProfit and is used to iterate over the raw logs and unpacked data for ReleasedTokenProfit events raised by the Main contract.
type MainReleasedTokenProfitIterator struct {
	Event *MainReleasedTokenProfit // Event containing the contract specifics and raw log

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
func (it *MainReleasedTokenProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainReleasedTokenProfit)
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
		it.Event = new(MainReleasedTokenProfit)
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
func (it *MainReleasedTokenProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainReleasedTokenProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainReleasedTokenProfit represents a ReleasedTokenProfit event raised by the Main contract.
type MainReleasedTokenProfit struct {
	To      common.Address
	TokenId *big.Int
	Slot    *big.Int
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReleasedTokenProfit is a free log retrieval operation binding the contract event 0x4b54a2d813bd017c0741faa683e2acf7012a515b4fc258ffcf2597acf7117134.
//
// Solidity: event ReleasedTokenProfit(address indexed to, uint256 indexed tokenId, uint256 slot, uint256 value)
func (_Main *MainFilterer) FilterReleasedTokenProfit(opts *bind.FilterOpts, to []common.Address, tokenId []*big.Int) (*MainReleasedTokenProfitIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ReleasedTokenProfit", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MainReleasedTokenProfitIterator{contract: _Main.contract, event: "ReleasedTokenProfit", logs: logs, sub: sub}, nil
}

// WatchReleasedTokenProfit is a free log subscription operation binding the contract event 0x4b54a2d813bd017c0741faa683e2acf7012a515b4fc258ffcf2597acf7117134.
//
// Solidity: event ReleasedTokenProfit(address indexed to, uint256 indexed tokenId, uint256 slot, uint256 value)
func (_Main *MainFilterer) WatchReleasedTokenProfit(opts *bind.WatchOpts, sink chan<- *MainReleasedTokenProfit, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ReleasedTokenProfit", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainReleasedTokenProfit)
				if err := _Main.contract.UnpackLog(event, "ReleasedTokenProfit", log); err != nil {
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

// ParseReleasedTokenProfit is a log parse operation binding the contract event 0x4b54a2d813bd017c0741faa683e2acf7012a515b4fc258ffcf2597acf7117134.
//
// Solidity: event ReleasedTokenProfit(address indexed to, uint256 indexed tokenId, uint256 slot, uint256 value)
func (_Main *MainFilterer) ParseReleasedTokenProfit(log types.Log) (*MainReleasedTokenProfit, error) {
	event := new(MainReleasedTokenProfit)
	if err := _Main.contract.UnpackLog(event, "ReleasedTokenProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Main contract.
type MainRoleAdminChangedIterator struct {
	Event *MainRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MainRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRoleAdminChanged)
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
		it.Event = new(MainRoleAdminChanged)
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
func (it *MainRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRoleAdminChanged represents a RoleAdminChanged event raised by the Main contract.
type MainRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Main *MainFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MainRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MainRoleAdminChangedIterator{contract: _Main.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Main *MainFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MainRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRoleAdminChanged)
				if err := _Main.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Main *MainFilterer) ParseRoleAdminChanged(log types.Log) (*MainRoleAdminChanged, error) {
	event := new(MainRoleAdminChanged)
	if err := _Main.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Main contract.
type MainRoleGrantedIterator struct {
	Event *MainRoleGranted // Event containing the contract specifics and raw log

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
func (it *MainRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRoleGranted)
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
		it.Event = new(MainRoleGranted)
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
func (it *MainRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRoleGranted represents a RoleGranted event raised by the Main contract.
type MainRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Main *MainFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MainRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MainRoleGrantedIterator{contract: _Main.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Main *MainFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MainRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRoleGranted)
				if err := _Main.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Main *MainFilterer) ParseRoleGranted(log types.Log) (*MainRoleGranted, error) {
	event := new(MainRoleGranted)
	if err := _Main.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Main contract.
type MainRoleRevokedIterator struct {
	Event *MainRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MainRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRoleRevoked)
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
		it.Event = new(MainRoleRevoked)
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
func (it *MainRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRoleRevoked represents a RoleRevoked event raised by the Main contract.
type MainRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Main *MainFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MainRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MainRoleRevokedIterator{contract: _Main.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Main *MainFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MainRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRoleRevoked)
				if err := _Main.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Main *MainFilterer) ParseRoleRevoked(log types.Log) (*MainRoleRevoked, error) {
	event := new(MainRoleRevoked)
	if err := _Main.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainSetMetadataDescriptorIterator is returned from FilterSetMetadataDescriptor and is used to iterate over the raw logs and unpacked data for SetMetadataDescriptor events raised by the Main contract.
type MainSetMetadataDescriptorIterator struct {
	Event *MainSetMetadataDescriptor // Event containing the contract specifics and raw log

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
func (it *MainSetMetadataDescriptorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSetMetadataDescriptor)
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
		it.Event = new(MainSetMetadataDescriptor)
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
func (it *MainSetMetadataDescriptorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSetMetadataDescriptorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSetMetadataDescriptor represents a SetMetadataDescriptor event raised by the Main contract.
type MainSetMetadataDescriptor struct {
	MetadataDescriptor common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetMetadataDescriptor is a free log retrieval operation binding the contract event 0x5252f52e45fc8ee6a7b43cef3645d23e9a470a34182b8b3a12627556635bfc9c.
//
// Solidity: event SetMetadataDescriptor(address indexed metadataDescriptor)
func (_Main *MainFilterer) FilterSetMetadataDescriptor(opts *bind.FilterOpts, metadataDescriptor []common.Address) (*MainSetMetadataDescriptorIterator, error) {

	var metadataDescriptorRule []interface{}
	for _, metadataDescriptorItem := range metadataDescriptor {
		metadataDescriptorRule = append(metadataDescriptorRule, metadataDescriptorItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "SetMetadataDescriptor", metadataDescriptorRule)
	if err != nil {
		return nil, err
	}
	return &MainSetMetadataDescriptorIterator{contract: _Main.contract, event: "SetMetadataDescriptor", logs: logs, sub: sub}, nil
}

// WatchSetMetadataDescriptor is a free log subscription operation binding the contract event 0x5252f52e45fc8ee6a7b43cef3645d23e9a470a34182b8b3a12627556635bfc9c.
//
// Solidity: event SetMetadataDescriptor(address indexed metadataDescriptor)
func (_Main *MainFilterer) WatchSetMetadataDescriptor(opts *bind.WatchOpts, sink chan<- *MainSetMetadataDescriptor, metadataDescriptor []common.Address) (event.Subscription, error) {

	var metadataDescriptorRule []interface{}
	for _, metadataDescriptorItem := range metadataDescriptor {
		metadataDescriptorRule = append(metadataDescriptorRule, metadataDescriptorItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "SetMetadataDescriptor", metadataDescriptorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSetMetadataDescriptor)
				if err := _Main.contract.UnpackLog(event, "SetMetadataDescriptor", log); err != nil {
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

// ParseSetMetadataDescriptor is a log parse operation binding the contract event 0x5252f52e45fc8ee6a7b43cef3645d23e9a470a34182b8b3a12627556635bfc9c.
//
// Solidity: event SetMetadataDescriptor(address indexed metadataDescriptor)
func (_Main *MainFilterer) ParseSetMetadataDescriptor(log types.Log) (*MainSetMetadataDescriptor, error) {
	event := new(MainSetMetadataDescriptor)
	if err := _Main.contract.UnpackLog(event, "SetMetadataDescriptor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainSlotChangedIterator is returned from FilterSlotChanged and is used to iterate over the raw logs and unpacked data for SlotChanged events raised by the Main contract.
type MainSlotChangedIterator struct {
	Event *MainSlotChanged // Event containing the contract specifics and raw log

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
func (it *MainSlotChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSlotChanged)
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
		it.Event = new(MainSlotChanged)
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
func (it *MainSlotChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSlotChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSlotChanged represents a SlotChanged event raised by the Main contract.
type MainSlotChanged struct {
	TokenId *big.Int
	OldSlot *big.Int
	NewSlot *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlotChanged is a free log retrieval operation binding the contract event 0xe4f48c240d3b994948aa54f3e2f5fca59263dfe1d52b6e4cf39a5d249b5ccb65.
//
// Solidity: event SlotChanged(uint256 indexed _tokenId, uint256 indexed _oldSlot, uint256 indexed _newSlot)
func (_Main *MainFilterer) FilterSlotChanged(opts *bind.FilterOpts, _tokenId []*big.Int, _oldSlot []*big.Int, _newSlot []*big.Int) (*MainSlotChangedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _oldSlotRule []interface{}
	for _, _oldSlotItem := range _oldSlot {
		_oldSlotRule = append(_oldSlotRule, _oldSlotItem)
	}
	var _newSlotRule []interface{}
	for _, _newSlotItem := range _newSlot {
		_newSlotRule = append(_newSlotRule, _newSlotItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "SlotChanged", _tokenIdRule, _oldSlotRule, _newSlotRule)
	if err != nil {
		return nil, err
	}
	return &MainSlotChangedIterator{contract: _Main.contract, event: "SlotChanged", logs: logs, sub: sub}, nil
}

// WatchSlotChanged is a free log subscription operation binding the contract event 0xe4f48c240d3b994948aa54f3e2f5fca59263dfe1d52b6e4cf39a5d249b5ccb65.
//
// Solidity: event SlotChanged(uint256 indexed _tokenId, uint256 indexed _oldSlot, uint256 indexed _newSlot)
func (_Main *MainFilterer) WatchSlotChanged(opts *bind.WatchOpts, sink chan<- *MainSlotChanged, _tokenId []*big.Int, _oldSlot []*big.Int, _newSlot []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _oldSlotRule []interface{}
	for _, _oldSlotItem := range _oldSlot {
		_oldSlotRule = append(_oldSlotRule, _oldSlotItem)
	}
	var _newSlotRule []interface{}
	for _, _newSlotItem := range _newSlot {
		_newSlotRule = append(_newSlotRule, _newSlotItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "SlotChanged", _tokenIdRule, _oldSlotRule, _newSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSlotChanged)
				if err := _Main.contract.UnpackLog(event, "SlotChanged", log); err != nil {
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

// ParseSlotChanged is a log parse operation binding the contract event 0xe4f48c240d3b994948aa54f3e2f5fca59263dfe1d52b6e4cf39a5d249b5ccb65.
//
// Solidity: event SlotChanged(uint256 indexed _tokenId, uint256 indexed _oldSlot, uint256 indexed _newSlot)
func (_Main *MainFilterer) ParseSlotChanged(log types.Log) (*MainSlotChanged, error) {
	event := new(MainSlotChanged)
	if err := _Main.contract.UnpackLog(event, "SlotChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the Main contract.
type MainTokenMintedIterator struct {
	Event *MainTokenMinted // Event containing the contract specifics and raw log

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
func (it *MainTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTokenMinted)
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
		it.Event = new(MainTokenMinted)
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
func (it *MainTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTokenMinted represents a TokenMinted event raised by the Main contract.
type MainTokenMinted struct {
	To      common.Address
	TokenId *big.Int
	Slot    *big.Int
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenMinted is a free log retrieval operation binding the contract event 0x10546b1a6f5245ff0ffa18c256b9e46859c585cbb473b453fcd4c2dc39ae08db.
//
// Solidity: event TokenMinted(address indexed to, uint256 indexed tokenId, uint256 slot, uint256 value)
func (_Main *MainFilterer) FilterTokenMinted(opts *bind.FilterOpts, to []common.Address, tokenId []*big.Int) (*MainTokenMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "TokenMinted", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MainTokenMintedIterator{contract: _Main.contract, event: "TokenMinted", logs: logs, sub: sub}, nil
}

// WatchTokenMinted is a free log subscription operation binding the contract event 0x10546b1a6f5245ff0ffa18c256b9e46859c585cbb473b453fcd4c2dc39ae08db.
//
// Solidity: event TokenMinted(address indexed to, uint256 indexed tokenId, uint256 slot, uint256 value)
func (_Main *MainFilterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *MainTokenMinted, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "TokenMinted", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTokenMinted)
				if err := _Main.contract.UnpackLog(event, "TokenMinted", log); err != nil {
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

// ParseTokenMinted is a log parse operation binding the contract event 0x10546b1a6f5245ff0ffa18c256b9e46859c585cbb473b453fcd4c2dc39ae08db.
//
// Solidity: event TokenMinted(address indexed to, uint256 indexed tokenId, uint256 slot, uint256 value)
func (_Main *MainFilterer) ParseTokenMinted(log types.Log) (*MainTokenMinted, error) {
	event := new(MainTokenMinted)
	if err := _Main.contract.UnpackLog(event, "TokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Main contract.
type MainTransferIterator struct {
	Event *MainTransfer // Event containing the contract specifics and raw log

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
func (it *MainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransfer)
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
		it.Event = new(MainTransfer)
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
func (it *MainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransfer represents a Transfer event raised by the Main contract.
type MainTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Main *MainFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*MainTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MainTransferIterator{contract: _Main.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Main *MainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MainTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransfer)
				if err := _Main.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Main *MainFilterer) ParseTransfer(log types.Log) (*MainTransfer, error) {
	event := new(MainTransfer)
	if err := _Main.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransferValueIterator is returned from FilterTransferValue and is used to iterate over the raw logs and unpacked data for TransferValue events raised by the Main contract.
type MainTransferValueIterator struct {
	Event *MainTransferValue // Event containing the contract specifics and raw log

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
func (it *MainTransferValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransferValue)
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
		it.Event = new(MainTransferValue)
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
func (it *MainTransferValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransferValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransferValue represents a TransferValue event raised by the Main contract.
type MainTransferValue struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferValue is a free log retrieval operation binding the contract event 0x0b2aac84f3ec956911fd78eae5311062972ff949f38412e8da39069d9f068cc6.
//
// Solidity: event TransferValue(uint256 indexed _fromTokenId, uint256 indexed _toTokenId, uint256 _value)
func (_Main *MainFilterer) FilterTransferValue(opts *bind.FilterOpts, _fromTokenId []*big.Int, _toTokenId []*big.Int) (*MainTransferValueIterator, error) {

	var _fromTokenIdRule []interface{}
	for _, _fromTokenIdItem := range _fromTokenId {
		_fromTokenIdRule = append(_fromTokenIdRule, _fromTokenIdItem)
	}
	var _toTokenIdRule []interface{}
	for _, _toTokenIdItem := range _toTokenId {
		_toTokenIdRule = append(_toTokenIdRule, _toTokenIdItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "TransferValue", _fromTokenIdRule, _toTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MainTransferValueIterator{contract: _Main.contract, event: "TransferValue", logs: logs, sub: sub}, nil
}

// WatchTransferValue is a free log subscription operation binding the contract event 0x0b2aac84f3ec956911fd78eae5311062972ff949f38412e8da39069d9f068cc6.
//
// Solidity: event TransferValue(uint256 indexed _fromTokenId, uint256 indexed _toTokenId, uint256 _value)
func (_Main *MainFilterer) WatchTransferValue(opts *bind.WatchOpts, sink chan<- *MainTransferValue, _fromTokenId []*big.Int, _toTokenId []*big.Int) (event.Subscription, error) {

	var _fromTokenIdRule []interface{}
	for _, _fromTokenIdItem := range _fromTokenId {
		_fromTokenIdRule = append(_fromTokenIdRule, _fromTokenIdItem)
	}
	var _toTokenIdRule []interface{}
	for _, _toTokenIdItem := range _toTokenId {
		_toTokenIdRule = append(_toTokenIdRule, _toTokenIdItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "TransferValue", _fromTokenIdRule, _toTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransferValue)
				if err := _Main.contract.UnpackLog(event, "TransferValue", log); err != nil {
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

// ParseTransferValue is a log parse operation binding the contract event 0x0b2aac84f3ec956911fd78eae5311062972ff949f38412e8da39069d9f068cc6.
//
// Solidity: event TransferValue(uint256 indexed _fromTokenId, uint256 indexed _toTokenId, uint256 _value)
func (_Main *MainFilterer) ParseTransferValue(log types.Log) (*MainTransferValue, error) {
	event := new(MainTransferValue)
	if err := _Main.contract.UnpackLog(event, "TransferValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Main contract.
type MainUpgradedIterator struct {
	Event *MainUpgraded // Event containing the contract specifics and raw log

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
func (it *MainUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainUpgraded)
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
		it.Event = new(MainUpgraded)
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
func (it *MainUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainUpgraded represents a Upgraded event raised by the Main contract.
type MainUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Main *MainFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MainUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MainUpgradedIterator{contract: _Main.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Main *MainFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MainUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainUpgraded)
				if err := _Main.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Main *MainFilterer) ParseUpgraded(log types.Log) (*MainUpgraded, error) {
	event := new(MainUpgraded)
	if err := _Main.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Main contract.
type MainWithdrawnIterator struct {
	Event *MainWithdrawn // Event containing the contract specifics and raw log

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
func (it *MainWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainWithdrawn)
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
		it.Event = new(MainWithdrawn)
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
func (it *MainWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainWithdrawn represents a Withdrawn event raised by the Main contract.
type MainWithdrawn struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x430648de173157e069201c943adb2d4e340e7cf5b27b1b09c9cb852f03d63b56.
//
// Solidity: event Withdrawn(uint256 value)
func (_Main *MainFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*MainWithdrawnIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &MainWithdrawnIterator{contract: _Main.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x430648de173157e069201c943adb2d4e340e7cf5b27b1b09c9cb852f03d63b56.
//
// Solidity: event Withdrawn(uint256 value)
func (_Main *MainFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *MainWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainWithdrawn)
				if err := _Main.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x430648de173157e069201c943adb2d4e340e7cf5b27b1b09c9cb852f03d63b56.
//
// Solidity: event Withdrawn(uint256 value)
func (_Main *MainFilterer) ParseWithdrawn(log types.Log) (*MainWithdrawn, error) {
	event := new(MainWithdrawn)
	if err := _Main.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
