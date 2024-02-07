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

// ViraMetaData contains all meta data concerning the Vira contract.
var ViraMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"ApprovalValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"campaignId\",\"type\":\"string\"}],\"name\":\"CampaignUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReleasedAllCampaignsProfit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReleasedInvestment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReleasedTokenProfit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"metadataDescriptor\",\"type\":\"address\"}],\"name\":\"SetMetadataDescriptor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_oldSlot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_newSlot\",\"type\":\"uint256\"}],\"name\":\"SlotChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"TransferValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder_\",\"type\":\"address\"}],\"name\":\"addressReleasableProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"campaignId_\",\"type\":\"string\"}],\"name\":\"campaignData\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctPerSlice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"}],\"internalType\":\"structIViraCampaigns.Campaign\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataDescriptor\",\"outputs\":[{\"internalType\":\"contractIERC3525MetadataDescriptorUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"campaignId_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"mintValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"releasableTokenProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseAllCampaignsProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"releaseCampaignProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"releaseInvestment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved_\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"slotOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot_\",\"type\":\"uint256\"}],\"name\":\"slotURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCampaigns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctPerSlice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"}],\"internalType\":\"structIViraCampaigns.Campaign\",\"name\":\"campaign_\",\"type\":\"tuple\"}],\"name\":\"updateCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ViraABI is the input ABI used to generate the binding from.
// Deprecated: Use ViraMetaData.ABI instead.
var ViraABI = ViraMetaData.ABI

// Vira is an auto generated Go binding around an Ethereum contract.
type Vira struct {
	ViraCaller     // Read-only binding to the contract
	ViraTransactor // Write-only binding to the contract
	ViraFilterer   // Log filterer for contract events
}

// ViraCaller is an auto generated read-only Go binding around an Ethereum contract.
type ViraCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViraTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ViraTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViraFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ViraFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViraSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ViraSession struct {
	Contract     *Vira             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ViraCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ViraCallerSession struct {
	Contract *ViraCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ViraTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ViraTransactorSession struct {
	Contract     *ViraTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ViraRaw is an auto generated low-level Go binding around an Ethereum contract.
type ViraRaw struct {
	Contract *Vira // Generic contract binding to access the raw methods on
}

// ViraCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ViraCallerRaw struct {
	Contract *ViraCaller // Generic read-only contract binding to access the raw methods on
}

// ViraTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ViraTransactorRaw struct {
	Contract *ViraTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVira creates a new instance of Vira, bound to a specific deployed contract.
func NewVira(address common.Address, backend bind.ContractBackend) (*Vira, error) {
	contract, err := bindVira(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vira{ViraCaller: ViraCaller{contract: contract}, ViraTransactor: ViraTransactor{contract: contract}, ViraFilterer: ViraFilterer{contract: contract}}, nil
}

// NewViraCaller creates a new read-only instance of Vira, bound to a specific deployed contract.
func NewViraCaller(address common.Address, caller bind.ContractCaller) (*ViraCaller, error) {
	contract, err := bindVira(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ViraCaller{contract: contract}, nil
}

// NewViraTransactor creates a new write-only instance of Vira, bound to a specific deployed contract.
func NewViraTransactor(address common.Address, transactor bind.ContractTransactor) (*ViraTransactor, error) {
	contract, err := bindVira(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ViraTransactor{contract: contract}, nil
}

// NewViraFilterer creates a new log filterer instance of Vira, bound to a specific deployed contract.
func NewViraFilterer(address common.Address, filterer bind.ContractFilterer) (*ViraFilterer, error) {
	contract, err := bindVira(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ViraFilterer{contract: contract}, nil
}

// bindVira binds a generic wrapper to an already deployed contract.
func bindVira(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ViraMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vira *ViraRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vira.Contract.ViraCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vira *ViraRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vira.Contract.ViraTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vira *ViraRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vira.Contract.ViraTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vira *ViraCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vira.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vira *ViraTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vira.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vira *ViraTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vira.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vira *ViraCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vira *ViraSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vira.Contract.DEFAULTADMINROLE(&_Vira.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vira *ViraCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vira.Contract.DEFAULTADMINROLE(&_Vira.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Vira *ViraCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Vira *ViraSession) MINTERROLE() ([32]byte, error) {
	return _Vira.Contract.MINTERROLE(&_Vira.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Vira *ViraCallerSession) MINTERROLE() ([32]byte, error) {
	return _Vira.Contract.MINTERROLE(&_Vira.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Vira *ViraCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Vira *ViraSession) UPGRADERROLE() ([32]byte, error) {
	return _Vira.Contract.UPGRADERROLE(&_Vira.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Vira *ViraCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Vira.Contract.UPGRADERROLE(&_Vira.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Vira *ViraCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Vira *ViraSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Vira.Contract.UPGRADEINTERFACEVERSION(&_Vira.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Vira *ViraCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Vira.Contract.UPGRADEINTERFACEVERSION(&_Vira.CallOpts)
}

// AddressReleasableProfit is a free data retrieval call binding the contract method 0xd3a6afc1.
//
// Solidity: function addressReleasableProfit(address holder_) view returns(uint256)
func (_Vira *ViraCaller) AddressReleasableProfit(opts *bind.CallOpts, holder_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "addressReleasableProfit", holder_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressReleasableProfit is a free data retrieval call binding the contract method 0xd3a6afc1.
//
// Solidity: function addressReleasableProfit(address holder_) view returns(uint256)
func (_Vira *ViraSession) AddressReleasableProfit(holder_ common.Address) (*big.Int, error) {
	return _Vira.Contract.AddressReleasableProfit(&_Vira.CallOpts, holder_)
}

// AddressReleasableProfit is a free data retrieval call binding the contract method 0xd3a6afc1.
//
// Solidity: function addressReleasableProfit(address holder_) view returns(uint256)
func (_Vira *ViraCallerSession) AddressReleasableProfit(holder_ common.Address) (*big.Int, error) {
	return _Vira.Contract.AddressReleasableProfit(&_Vira.CallOpts, holder_)
}

// Allowance is a free data retrieval call binding the contract method 0xe345e0bc.
//
// Solidity: function allowance(uint256 tokenId_, address operator_) view returns(uint256)
func (_Vira *ViraCaller) Allowance(opts *bind.CallOpts, tokenId_ *big.Int, operator_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "allowance", tokenId_, operator_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xe345e0bc.
//
// Solidity: function allowance(uint256 tokenId_, address operator_) view returns(uint256)
func (_Vira *ViraSession) Allowance(tokenId_ *big.Int, operator_ common.Address) (*big.Int, error) {
	return _Vira.Contract.Allowance(&_Vira.CallOpts, tokenId_, operator_)
}

// Allowance is a free data retrieval call binding the contract method 0xe345e0bc.
//
// Solidity: function allowance(uint256 tokenId_, address operator_) view returns(uint256)
func (_Vira *ViraCallerSession) Allowance(tokenId_ *big.Int, operator_ common.Address) (*big.Int, error) {
	return _Vira.Contract.Allowance(&_Vira.CallOpts, tokenId_, operator_)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner_) view returns(uint256 balance)
func (_Vira *ViraCaller) BalanceOf(opts *bind.CallOpts, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "balanceOf", owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner_) view returns(uint256 balance)
func (_Vira *ViraSession) BalanceOf(owner_ common.Address) (*big.Int, error) {
	return _Vira.Contract.BalanceOf(&_Vira.CallOpts, owner_)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner_) view returns(uint256 balance)
func (_Vira *ViraCallerSession) BalanceOf(owner_ common.Address) (*big.Int, error) {
	return _Vira.Contract.BalanceOf(&_Vira.CallOpts, owner_)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 tokenId_) view returns(uint256)
func (_Vira *ViraCaller) BalanceOf0(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "balanceOf0", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf0 is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 tokenId_) view returns(uint256)
func (_Vira *ViraSession) BalanceOf0(tokenId_ *big.Int) (*big.Int, error) {
	return _Vira.Contract.BalanceOf0(&_Vira.CallOpts, tokenId_)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 tokenId_) view returns(uint256)
func (_Vira *ViraCallerSession) BalanceOf0(tokenId_ *big.Int) (*big.Int, error) {
	return _Vira.Contract.BalanceOf0(&_Vira.CallOpts, tokenId_)
}

// CampaignData is a free data retrieval call binding the contract method 0x5798e048.
//
// Solidity: function campaignData(string campaignId_) view returns((string,uint256,uint256,uint256,uint256,uint256,bool,bool))
func (_Vira *ViraCaller) CampaignData(opts *bind.CallOpts, campaignId_ string) (IViraCampaignsCampaign, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "campaignData", campaignId_)

	if err != nil {
		return *new(IViraCampaignsCampaign), err
	}

	out0 := *abi.ConvertType(out[0], new(IViraCampaignsCampaign)).(*IViraCampaignsCampaign)

	return out0, err

}

// CampaignData is a free data retrieval call binding the contract method 0x5798e048.
//
// Solidity: function campaignData(string campaignId_) view returns((string,uint256,uint256,uint256,uint256,uint256,bool,bool))
func (_Vira *ViraSession) CampaignData(campaignId_ string) (IViraCampaignsCampaign, error) {
	return _Vira.Contract.CampaignData(&_Vira.CallOpts, campaignId_)
}

// CampaignData is a free data retrieval call binding the contract method 0x5798e048.
//
// Solidity: function campaignData(string campaignId_) view returns((string,uint256,uint256,uint256,uint256,uint256,bool,bool))
func (_Vira *ViraCallerSession) CampaignData(campaignId_ string) (IViraCampaignsCampaign, error) {
	return _Vira.Contract.CampaignData(&_Vira.CallOpts, campaignId_)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Vira *ViraCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Vira *ViraSession) ContractURI() (string, error) {
	return _Vira.Contract.ContractURI(&_Vira.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Vira *ViraCallerSession) ContractURI() (string, error) {
	return _Vira.Contract.ContractURI(&_Vira.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId_) view returns(address)
func (_Vira *ViraCaller) GetApproved(opts *bind.CallOpts, tokenId_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "getApproved", tokenId_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId_) view returns(address)
func (_Vira *ViraSession) GetApproved(tokenId_ *big.Int) (common.Address, error) {
	return _Vira.Contract.GetApproved(&_Vira.CallOpts, tokenId_)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId_) view returns(address)
func (_Vira *ViraCallerSession) GetApproved(tokenId_ *big.Int) (common.Address, error) {
	return _Vira.Contract.GetApproved(&_Vira.CallOpts, tokenId_)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vira *ViraCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vira *ViraSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vira.Contract.GetRoleAdmin(&_Vira.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vira *ViraCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vira.Contract.GetRoleAdmin(&_Vira.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vira *ViraCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vira *ViraSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vira.Contract.HasRole(&_Vira.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vira *ViraCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vira.Contract.HasRole(&_Vira.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner_, address operator_) view returns(bool)
func (_Vira *ViraCaller) IsApprovedForAll(opts *bind.CallOpts, owner_ common.Address, operator_ common.Address) (bool, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "isApprovedForAll", owner_, operator_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner_, address operator_) view returns(bool)
func (_Vira *ViraSession) IsApprovedForAll(owner_ common.Address, operator_ common.Address) (bool, error) {
	return _Vira.Contract.IsApprovedForAll(&_Vira.CallOpts, owner_, operator_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner_, address operator_) view returns(bool)
func (_Vira *ViraCallerSession) IsApprovedForAll(owner_ common.Address, operator_ common.Address) (bool, error) {
	return _Vira.Contract.IsApprovedForAll(&_Vira.CallOpts, owner_, operator_)
}

// MetadataDescriptor is a free data retrieval call binding the contract method 0x840f7113.
//
// Solidity: function metadataDescriptor() view returns(address)
func (_Vira *ViraCaller) MetadataDescriptor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "metadataDescriptor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataDescriptor is a free data retrieval call binding the contract method 0x840f7113.
//
// Solidity: function metadataDescriptor() view returns(address)
func (_Vira *ViraSession) MetadataDescriptor() (common.Address, error) {
	return _Vira.Contract.MetadataDescriptor(&_Vira.CallOpts)
}

// MetadataDescriptor is a free data retrieval call binding the contract method 0x840f7113.
//
// Solidity: function metadataDescriptor() view returns(address)
func (_Vira *ViraCallerSession) MetadataDescriptor() (common.Address, error) {
	return _Vira.Contract.MetadataDescriptor(&_Vira.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vira *ViraCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vira *ViraSession) Name() (string, error) {
	return _Vira.Contract.Name(&_Vira.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vira *ViraCallerSession) Name() (string, error) {
	return _Vira.Contract.Name(&_Vira.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId_) view returns(address owner_)
func (_Vira *ViraCaller) OwnerOf(opts *bind.CallOpts, tokenId_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "ownerOf", tokenId_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId_) view returns(address owner_)
func (_Vira *ViraSession) OwnerOf(tokenId_ *big.Int) (common.Address, error) {
	return _Vira.Contract.OwnerOf(&_Vira.CallOpts, tokenId_)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId_) view returns(address owner_)
func (_Vira *ViraCallerSession) OwnerOf(tokenId_ *big.Int) (common.Address, error) {
	return _Vira.Contract.OwnerOf(&_Vira.CallOpts, tokenId_)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vira *ViraCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vira *ViraSession) ProxiableUUID() ([32]byte, error) {
	return _Vira.Contract.ProxiableUUID(&_Vira.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vira *ViraCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Vira.Contract.ProxiableUUID(&_Vira.CallOpts)
}

// ReleasableTokenProfit is a free data retrieval call binding the contract method 0x51d661e9.
//
// Solidity: function releasableTokenProfit(uint256 tokenId_) view returns(uint256)
func (_Vira *ViraCaller) ReleasableTokenProfit(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "releasableTokenProfit", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleasableTokenProfit is a free data retrieval call binding the contract method 0x51d661e9.
//
// Solidity: function releasableTokenProfit(uint256 tokenId_) view returns(uint256)
func (_Vira *ViraSession) ReleasableTokenProfit(tokenId_ *big.Int) (*big.Int, error) {
	return _Vira.Contract.ReleasableTokenProfit(&_Vira.CallOpts, tokenId_)
}

// ReleasableTokenProfit is a free data retrieval call binding the contract method 0x51d661e9.
//
// Solidity: function releasableTokenProfit(uint256 tokenId_) view returns(uint256)
func (_Vira *ViraCallerSession) ReleasableTokenProfit(tokenId_ *big.Int) (*big.Int, error) {
	return _Vira.Contract.ReleasableTokenProfit(&_Vira.CallOpts, tokenId_)
}

// SlotOf is a free data retrieval call binding the contract method 0x263f3e7e.
//
// Solidity: function slotOf(uint256 tokenId_) view returns(uint256)
func (_Vira *ViraCaller) SlotOf(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "slotOf", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotOf is a free data retrieval call binding the contract method 0x263f3e7e.
//
// Solidity: function slotOf(uint256 tokenId_) view returns(uint256)
func (_Vira *ViraSession) SlotOf(tokenId_ *big.Int) (*big.Int, error) {
	return _Vira.Contract.SlotOf(&_Vira.CallOpts, tokenId_)
}

// SlotOf is a free data retrieval call binding the contract method 0x263f3e7e.
//
// Solidity: function slotOf(uint256 tokenId_) view returns(uint256)
func (_Vira *ViraCallerSession) SlotOf(tokenId_ *big.Int) (*big.Int, error) {
	return _Vira.Contract.SlotOf(&_Vira.CallOpts, tokenId_)
}

// SlotURI is a free data retrieval call binding the contract method 0x09c3dd87.
//
// Solidity: function slotURI(uint256 slot_) view returns(string)
func (_Vira *ViraCaller) SlotURI(opts *bind.CallOpts, slot_ *big.Int) (string, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "slotURI", slot_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SlotURI is a free data retrieval call binding the contract method 0x09c3dd87.
//
// Solidity: function slotURI(uint256 slot_) view returns(string)
func (_Vira *ViraSession) SlotURI(slot_ *big.Int) (string, error) {
	return _Vira.Contract.SlotURI(&_Vira.CallOpts, slot_)
}

// SlotURI is a free data retrieval call binding the contract method 0x09c3dd87.
//
// Solidity: function slotURI(uint256 slot_) view returns(string)
func (_Vira *ViraCallerSession) SlotURI(slot_ *big.Int) (string, error) {
	return _Vira.Contract.SlotURI(&_Vira.CallOpts, slot_)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vira *ViraCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vira *ViraSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vira.Contract.SupportsInterface(&_Vira.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vira *ViraCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vira.Contract.SupportsInterface(&_Vira.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vira *ViraCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vira *ViraSession) Symbol() (string, error) {
	return _Vira.Contract.Symbol(&_Vira.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vira *ViraCallerSession) Symbol() (string, error) {
	return _Vira.Contract.Symbol(&_Vira.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index_) view returns(uint256)
func (_Vira *ViraCaller) TokenByIndex(opts *bind.CallOpts, index_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "tokenByIndex", index_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index_) view returns(uint256)
func (_Vira *ViraSession) TokenByIndex(index_ *big.Int) (*big.Int, error) {
	return _Vira.Contract.TokenByIndex(&_Vira.CallOpts, index_)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index_) view returns(uint256)
func (_Vira *ViraCallerSession) TokenByIndex(index_ *big.Int) (*big.Int, error) {
	return _Vira.Contract.TokenByIndex(&_Vira.CallOpts, index_)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner_, uint256 index_) view returns(uint256)
func (_Vira *ViraCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner_ common.Address, index_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner_, index_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner_, uint256 index_) view returns(uint256)
func (_Vira *ViraSession) TokenOfOwnerByIndex(owner_ common.Address, index_ *big.Int) (*big.Int, error) {
	return _Vira.Contract.TokenOfOwnerByIndex(&_Vira.CallOpts, owner_, index_)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner_, uint256 index_) view returns(uint256)
func (_Vira *ViraCallerSession) TokenOfOwnerByIndex(owner_ common.Address, index_ *big.Int) (*big.Int, error) {
	return _Vira.Contract.TokenOfOwnerByIndex(&_Vira.CallOpts, owner_, index_)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Vira *ViraCaller) TokenURI(opts *bind.CallOpts, tokenId_ *big.Int) (string, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "tokenURI", tokenId_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Vira *ViraSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _Vira.Contract.TokenURI(&_Vira.CallOpts, tokenId_)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Vira *ViraCallerSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _Vira.Contract.TokenURI(&_Vira.CallOpts, tokenId_)
}

// TotalCampaigns is a free data retrieval call binding the contract method 0x02932f56.
//
// Solidity: function totalCampaigns() view returns(uint256)
func (_Vira *ViraCaller) TotalCampaigns(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "totalCampaigns")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCampaigns is a free data retrieval call binding the contract method 0x02932f56.
//
// Solidity: function totalCampaigns() view returns(uint256)
func (_Vira *ViraSession) TotalCampaigns() (*big.Int, error) {
	return _Vira.Contract.TotalCampaigns(&_Vira.CallOpts)
}

// TotalCampaigns is a free data retrieval call binding the contract method 0x02932f56.
//
// Solidity: function totalCampaigns() view returns(uint256)
func (_Vira *ViraCallerSession) TotalCampaigns() (*big.Int, error) {
	return _Vira.Contract.TotalCampaigns(&_Vira.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vira *ViraCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vira *ViraSession) TotalSupply() (*big.Int, error) {
	return _Vira.Contract.TotalSupply(&_Vira.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vira *ViraCallerSession) TotalSupply() (*big.Int, error) {
	return _Vira.Contract.TotalSupply(&_Vira.CallOpts)
}

// ValueDecimals is a free data retrieval call binding the contract method 0x3e7e8669.
//
// Solidity: function valueDecimals() view returns(uint8)
func (_Vira *ViraCaller) ValueDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vira.contract.Call(opts, &out, "valueDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValueDecimals is a free data retrieval call binding the contract method 0x3e7e8669.
//
// Solidity: function valueDecimals() view returns(uint8)
func (_Vira *ViraSession) ValueDecimals() (uint8, error) {
	return _Vira.Contract.ValueDecimals(&_Vira.CallOpts)
}

// ValueDecimals is a free data retrieval call binding the contract method 0x3e7e8669.
//
// Solidity: function valueDecimals() view returns(uint8)
func (_Vira *ViraCallerSession) ValueDecimals() (uint8, error) {
	return _Vira.Contract.ValueDecimals(&_Vira.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to_, uint256 tokenId_) payable returns()
func (_Vira *ViraTransactor) Approve(opts *bind.TransactOpts, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "approve", to_, tokenId_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to_, uint256 tokenId_) payable returns()
func (_Vira *ViraSession) Approve(to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.Approve(&_Vira.TransactOpts, to_, tokenId_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to_, uint256 tokenId_) payable returns()
func (_Vira *ViraTransactorSession) Approve(to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.Approve(&_Vira.TransactOpts, to_, tokenId_)
}

// Approve0 is a paid mutator transaction binding the contract method 0x8cb0a511.
//
// Solidity: function approve(uint256 tokenId_, address to_, uint256 value_) payable returns()
func (_Vira *ViraTransactor) Approve0(opts *bind.TransactOpts, tokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "approve0", tokenId_, to_, value_)
}

// Approve0 is a paid mutator transaction binding the contract method 0x8cb0a511.
//
// Solidity: function approve(uint256 tokenId_, address to_, uint256 value_) payable returns()
func (_Vira *ViraSession) Approve0(tokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.Approve0(&_Vira.TransactOpts, tokenId_, to_, value_)
}

// Approve0 is a paid mutator transaction binding the contract method 0x8cb0a511.
//
// Solidity: function approve(uint256 tokenId_, address to_, uint256 value_) payable returns()
func (_Vira *ViraTransactorSession) Approve0(tokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.Approve0(&_Vira.TransactOpts, tokenId_, to_, value_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vira *ViraTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vira *ViraSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vira.Contract.GrantRole(&_Vira.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vira *ViraTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vira.Contract.GrantRole(&_Vira.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address minter_, address token_) returns()
func (_Vira *ViraTransactor) Initialize(opts *bind.TransactOpts, minter_ common.Address, token_ common.Address) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "initialize", minter_, token_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address minter_, address token_) returns()
func (_Vira *ViraSession) Initialize(minter_ common.Address, token_ common.Address) (*types.Transaction, error) {
	return _Vira.Contract.Initialize(&_Vira.TransactOpts, minter_, token_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address minter_, address token_) returns()
func (_Vira *ViraTransactorSession) Initialize(minter_ common.Address, token_ common.Address) (*types.Transaction, error) {
	return _Vira.Contract.Initialize(&_Vira.TransactOpts, minter_, token_)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address to_, string campaignId_, uint256 value_) returns()
func (_Vira *ViraTransactor) Mint(opts *bind.TransactOpts, to_ common.Address, campaignId_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "mint", to_, campaignId_, value_)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address to_, string campaignId_, uint256 value_) returns()
func (_Vira *ViraSession) Mint(to_ common.Address, campaignId_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.Mint(&_Vira.TransactOpts, to_, campaignId_, value_)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address to_, string campaignId_, uint256 value_) returns()
func (_Vira *ViraTransactorSession) Mint(to_ common.Address, campaignId_ string, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.Mint(&_Vira.TransactOpts, to_, campaignId_, value_)
}

// MintValue is a paid mutator transaction binding the contract method 0xf0e88e7f.
//
// Solidity: function mintValue(uint256 tokenId_, uint256 value_) returns()
func (_Vira *ViraTransactor) MintValue(opts *bind.TransactOpts, tokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "mintValue", tokenId_, value_)
}

// MintValue is a paid mutator transaction binding the contract method 0xf0e88e7f.
//
// Solidity: function mintValue(uint256 tokenId_, uint256 value_) returns()
func (_Vira *ViraSession) MintValue(tokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.MintValue(&_Vira.TransactOpts, tokenId_, value_)
}

// MintValue is a paid mutator transaction binding the contract method 0xf0e88e7f.
//
// Solidity: function mintValue(uint256 tokenId_, uint256 value_) returns()
func (_Vira *ViraTransactorSession) MintValue(tokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.MintValue(&_Vira.TransactOpts, tokenId_, value_)
}

// ReleaseAllCampaignsProfit is a paid mutator transaction binding the contract method 0x9670263c.
//
// Solidity: function releaseAllCampaignsProfit() returns()
func (_Vira *ViraTransactor) ReleaseAllCampaignsProfit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "releaseAllCampaignsProfit")
}

// ReleaseAllCampaignsProfit is a paid mutator transaction binding the contract method 0x9670263c.
//
// Solidity: function releaseAllCampaignsProfit() returns()
func (_Vira *ViraSession) ReleaseAllCampaignsProfit() (*types.Transaction, error) {
	return _Vira.Contract.ReleaseAllCampaignsProfit(&_Vira.TransactOpts)
}

// ReleaseAllCampaignsProfit is a paid mutator transaction binding the contract method 0x9670263c.
//
// Solidity: function releaseAllCampaignsProfit() returns()
func (_Vira *ViraTransactorSession) ReleaseAllCampaignsProfit() (*types.Transaction, error) {
	return _Vira.Contract.ReleaseAllCampaignsProfit(&_Vira.TransactOpts)
}

// ReleaseCampaignProfit is a paid mutator transaction binding the contract method 0x4af81954.
//
// Solidity: function releaseCampaignProfit(uint256 tokenId_, uint256 amount_) returns()
func (_Vira *ViraTransactor) ReleaseCampaignProfit(opts *bind.TransactOpts, tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "releaseCampaignProfit", tokenId_, amount_)
}

// ReleaseCampaignProfit is a paid mutator transaction binding the contract method 0x4af81954.
//
// Solidity: function releaseCampaignProfit(uint256 tokenId_, uint256 amount_) returns()
func (_Vira *ViraSession) ReleaseCampaignProfit(tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.ReleaseCampaignProfit(&_Vira.TransactOpts, tokenId_, amount_)
}

// ReleaseCampaignProfit is a paid mutator transaction binding the contract method 0x4af81954.
//
// Solidity: function releaseCampaignProfit(uint256 tokenId_, uint256 amount_) returns()
func (_Vira *ViraTransactorSession) ReleaseCampaignProfit(tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.ReleaseCampaignProfit(&_Vira.TransactOpts, tokenId_, amount_)
}

// ReleaseInvestment is a paid mutator transaction binding the contract method 0x886a53ab.
//
// Solidity: function releaseInvestment(uint256 tokenId_) returns()
func (_Vira *ViraTransactor) ReleaseInvestment(opts *bind.TransactOpts, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "releaseInvestment", tokenId_)
}

// ReleaseInvestment is a paid mutator transaction binding the contract method 0x886a53ab.
//
// Solidity: function releaseInvestment(uint256 tokenId_) returns()
func (_Vira *ViraSession) ReleaseInvestment(tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.ReleaseInvestment(&_Vira.TransactOpts, tokenId_)
}

// ReleaseInvestment is a paid mutator transaction binding the contract method 0x886a53ab.
//
// Solidity: function releaseInvestment(uint256 tokenId_) returns()
func (_Vira *ViraTransactorSession) ReleaseInvestment(tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.ReleaseInvestment(&_Vira.TransactOpts, tokenId_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Vira *ViraTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Vira *ViraSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Vira.Contract.RenounceRole(&_Vira.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Vira *ViraTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Vira.Contract.RenounceRole(&_Vira.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vira *ViraTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vira *ViraSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vira.Contract.RevokeRole(&_Vira.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vira *ViraTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vira.Contract.RevokeRole(&_Vira.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Vira *ViraTransactor) SafeTransferFrom(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "safeTransferFrom", from_, to_, tokenId_)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Vira *ViraSession) SafeTransferFrom(from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.SafeTransferFrom(&_Vira.TransactOpts, from_, to_, tokenId_)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Vira *ViraTransactorSession) SafeTransferFrom(from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.SafeTransferFrom(&_Vira.TransactOpts, from_, to_, tokenId_)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_, bytes data_) payable returns()
func (_Vira *ViraTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, tokenId_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "safeTransferFrom0", from_, to_, tokenId_, data_)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_, bytes data_) payable returns()
func (_Vira *ViraSession) SafeTransferFrom0(from_ common.Address, to_ common.Address, tokenId_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Vira.Contract.SafeTransferFrom0(&_Vira.TransactOpts, from_, to_, tokenId_, data_)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from_, address to_, uint256 tokenId_, bytes data_) payable returns()
func (_Vira *ViraTransactorSession) SafeTransferFrom0(from_ common.Address, to_ common.Address, tokenId_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Vira.Contract.SafeTransferFrom0(&_Vira.TransactOpts, from_, to_, tokenId_, data_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved_) returns()
func (_Vira *ViraTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator_ common.Address, approved_ bool) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "setApprovalForAll", operator_, approved_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved_) returns()
func (_Vira *ViraSession) SetApprovalForAll(operator_ common.Address, approved_ bool) (*types.Transaction, error) {
	return _Vira.Contract.SetApprovalForAll(&_Vira.TransactOpts, operator_, approved_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved_) returns()
func (_Vira *ViraTransactorSession) SetApprovalForAll(operator_ common.Address, approved_ bool) (*types.Transaction, error) {
	return _Vira.Contract.SetApprovalForAll(&_Vira.TransactOpts, operator_, approved_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0f485c02.
//
// Solidity: function transferFrom(uint256 fromTokenId_, address to_, uint256 value_) payable returns(uint256 newTokenId)
func (_Vira *ViraTransactor) TransferFrom(opts *bind.TransactOpts, fromTokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "transferFrom", fromTokenId_, to_, value_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0f485c02.
//
// Solidity: function transferFrom(uint256 fromTokenId_, address to_, uint256 value_) payable returns(uint256 newTokenId)
func (_Vira *ViraSession) TransferFrom(fromTokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.TransferFrom(&_Vira.TransactOpts, fromTokenId_, to_, value_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0f485c02.
//
// Solidity: function transferFrom(uint256 fromTokenId_, address to_, uint256 value_) payable returns(uint256 newTokenId)
func (_Vira *ViraTransactorSession) TransferFrom(fromTokenId_ *big.Int, to_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.TransferFrom(&_Vira.TransactOpts, fromTokenId_, to_, value_)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Vira *ViraTransactor) TransferFrom0(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "transferFrom0", from_, to_, tokenId_)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Vira *ViraSession) TransferFrom0(from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.TransferFrom0(&_Vira.TransactOpts, from_, to_, tokenId_)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 tokenId_) payable returns()
func (_Vira *ViraTransactorSession) TransferFrom0(from_ common.Address, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.TransferFrom0(&_Vira.TransactOpts, from_, to_, tokenId_)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0x310ed7f0.
//
// Solidity: function transferFrom(uint256 fromTokenId_, uint256 toTokenId_, uint256 value_) payable returns()
func (_Vira *ViraTransactor) TransferFrom1(opts *bind.TransactOpts, fromTokenId_ *big.Int, toTokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "transferFrom1", fromTokenId_, toTokenId_, value_)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0x310ed7f0.
//
// Solidity: function transferFrom(uint256 fromTokenId_, uint256 toTokenId_, uint256 value_) payable returns()
func (_Vira *ViraSession) TransferFrom1(fromTokenId_ *big.Int, toTokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.TransferFrom1(&_Vira.TransactOpts, fromTokenId_, toTokenId_, value_)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0x310ed7f0.
//
// Solidity: function transferFrom(uint256 fromTokenId_, uint256 toTokenId_, uint256 value_) payable returns()
func (_Vira *ViraTransactorSession) TransferFrom1(fromTokenId_ *big.Int, toTokenId_ *big.Int, value_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.TransferFrom1(&_Vira.TransactOpts, fromTokenId_, toTokenId_, value_)
}

// UpdateCampaign is a paid mutator transaction binding the contract method 0x895040d2.
//
// Solidity: function updateCampaign((string,uint256,uint256,uint256,uint256,uint256,bool,bool) campaign_) returns()
func (_Vira *ViraTransactor) UpdateCampaign(opts *bind.TransactOpts, campaign_ IViraCampaignsCampaign) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "updateCampaign", campaign_)
}

// UpdateCampaign is a paid mutator transaction binding the contract method 0x895040d2.
//
// Solidity: function updateCampaign((string,uint256,uint256,uint256,uint256,uint256,bool,bool) campaign_) returns()
func (_Vira *ViraSession) UpdateCampaign(campaign_ IViraCampaignsCampaign) (*types.Transaction, error) {
	return _Vira.Contract.UpdateCampaign(&_Vira.TransactOpts, campaign_)
}

// UpdateCampaign is a paid mutator transaction binding the contract method 0x895040d2.
//
// Solidity: function updateCampaign((string,uint256,uint256,uint256,uint256,uint256,bool,bool) campaign_) returns()
func (_Vira *ViraTransactorSession) UpdateCampaign(campaign_ IViraCampaignsCampaign) (*types.Transaction, error) {
	return _Vira.Contract.UpdateCampaign(&_Vira.TransactOpts, campaign_)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vira *ViraTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vira *ViraSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vira.Contract.UpgradeToAndCall(&_Vira.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vira *ViraTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vira.Contract.UpgradeToAndCall(&_Vira.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_Vira *ViraTransactor) Withdraw(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _Vira.contract.Transact(opts, "withdraw", amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_Vira *ViraSession) Withdraw(amount_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.Withdraw(&_Vira.TransactOpts, amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_Vira *ViraTransactorSession) Withdraw(amount_ *big.Int) (*types.Transaction, error) {
	return _Vira.Contract.Withdraw(&_Vira.TransactOpts, amount_)
}

// ViraApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vira contract.
type ViraApprovalIterator struct {
	Event *ViraApproval // Event containing the contract specifics and raw log

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
func (it *ViraApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraApproval)
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
		it.Event = new(ViraApproval)
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
func (it *ViraApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraApproval represents a Approval event raised by the Vira contract.
type ViraApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Vira *ViraFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ViraApprovalIterator, error) {

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

	logs, sub, err := _Vira.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ViraApprovalIterator{contract: _Vira.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Vira *ViraFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ViraApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Vira.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraApproval)
				if err := _Vira.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Vira *ViraFilterer) ParseApproval(log types.Log) (*ViraApproval, error) {
	event := new(ViraApproval)
	if err := _Vira.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Vira contract.
type ViraApprovalForAllIterator struct {
	Event *ViraApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ViraApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraApprovalForAll)
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
		it.Event = new(ViraApprovalForAll)
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
func (it *ViraApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraApprovalForAll represents a ApprovalForAll event raised by the Vira contract.
type ViraApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Vira *ViraFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ViraApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Vira.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ViraApprovalForAllIterator{contract: _Vira.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Vira *ViraFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ViraApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Vira.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraApprovalForAll)
				if err := _Vira.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Vira *ViraFilterer) ParseApprovalForAll(log types.Log) (*ViraApprovalForAll, error) {
	event := new(ViraApprovalForAll)
	if err := _Vira.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraApprovalValueIterator is returned from FilterApprovalValue and is used to iterate over the raw logs and unpacked data for ApprovalValue events raised by the Vira contract.
type ViraApprovalValueIterator struct {
	Event *ViraApprovalValue // Event containing the contract specifics and raw log

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
func (it *ViraApprovalValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraApprovalValue)
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
		it.Event = new(ViraApprovalValue)
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
func (it *ViraApprovalValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraApprovalValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraApprovalValue represents a ApprovalValue event raised by the Vira contract.
type ViraApprovalValue struct {
	TokenId  *big.Int
	Operator common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalValue is a free log retrieval operation binding the contract event 0x621b050de0ad08b51d19b48b3e6df75348c4de6bdd93e81b252ca62e28265b1b.
//
// Solidity: event ApprovalValue(uint256 indexed _tokenId, address indexed _operator, uint256 _value)
func (_Vira *ViraFilterer) FilterApprovalValue(opts *bind.FilterOpts, _tokenId []*big.Int, _operator []common.Address) (*ViraApprovalValueIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Vira.contract.FilterLogs(opts, "ApprovalValue", _tokenIdRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ViraApprovalValueIterator{contract: _Vira.contract, event: "ApprovalValue", logs: logs, sub: sub}, nil
}

// WatchApprovalValue is a free log subscription operation binding the contract event 0x621b050de0ad08b51d19b48b3e6df75348c4de6bdd93e81b252ca62e28265b1b.
//
// Solidity: event ApprovalValue(uint256 indexed _tokenId, address indexed _operator, uint256 _value)
func (_Vira *ViraFilterer) WatchApprovalValue(opts *bind.WatchOpts, sink chan<- *ViraApprovalValue, _tokenId []*big.Int, _operator []common.Address) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Vira.contract.WatchLogs(opts, "ApprovalValue", _tokenIdRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraApprovalValue)
				if err := _Vira.contract.UnpackLog(event, "ApprovalValue", log); err != nil {
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
func (_Vira *ViraFilterer) ParseApprovalValue(log types.Log) (*ViraApprovalValue, error) {
	event := new(ViraApprovalValue)
	if err := _Vira.contract.UnpackLog(event, "ApprovalValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraCampaignUpdatedIterator is returned from FilterCampaignUpdated and is used to iterate over the raw logs and unpacked data for CampaignUpdated events raised by the Vira contract.
type ViraCampaignUpdatedIterator struct {
	Event *ViraCampaignUpdated // Event containing the contract specifics and raw log

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
func (it *ViraCampaignUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraCampaignUpdated)
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
		it.Event = new(ViraCampaignUpdated)
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
func (it *ViraCampaignUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraCampaignUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraCampaignUpdated represents a CampaignUpdated event raised by the Vira contract.
type ViraCampaignUpdated struct {
	CampaignId common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCampaignUpdated is a free log retrieval operation binding the contract event 0x97ebe72425499cfa387c633444ff604615975ceb676d866e2fb7fecd677554c4.
//
// Solidity: event CampaignUpdated(string indexed campaignId)
func (_Vira *ViraFilterer) FilterCampaignUpdated(opts *bind.FilterOpts, campaignId []string) (*ViraCampaignUpdatedIterator, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}

	logs, sub, err := _Vira.contract.FilterLogs(opts, "CampaignUpdated", campaignIdRule)
	if err != nil {
		return nil, err
	}
	return &ViraCampaignUpdatedIterator{contract: _Vira.contract, event: "CampaignUpdated", logs: logs, sub: sub}, nil
}

// WatchCampaignUpdated is a free log subscription operation binding the contract event 0x97ebe72425499cfa387c633444ff604615975ceb676d866e2fb7fecd677554c4.
//
// Solidity: event CampaignUpdated(string indexed campaignId)
func (_Vira *ViraFilterer) WatchCampaignUpdated(opts *bind.WatchOpts, sink chan<- *ViraCampaignUpdated, campaignId []string) (event.Subscription, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}

	logs, sub, err := _Vira.contract.WatchLogs(opts, "CampaignUpdated", campaignIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraCampaignUpdated)
				if err := _Vira.contract.UnpackLog(event, "CampaignUpdated", log); err != nil {
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
func (_Vira *ViraFilterer) ParseCampaignUpdated(log types.Log) (*ViraCampaignUpdated, error) {
	event := new(ViraCampaignUpdated)
	if err := _Vira.contract.UnpackLog(event, "CampaignUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Vira contract.
type ViraInitializedIterator struct {
	Event *ViraInitialized // Event containing the contract specifics and raw log

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
func (it *ViraInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraInitialized)
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
		it.Event = new(ViraInitialized)
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
func (it *ViraInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraInitialized represents a Initialized event raised by the Vira contract.
type ViraInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Vira *ViraFilterer) FilterInitialized(opts *bind.FilterOpts) (*ViraInitializedIterator, error) {

	logs, sub, err := _Vira.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ViraInitializedIterator{contract: _Vira.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Vira *ViraFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ViraInitialized) (event.Subscription, error) {

	logs, sub, err := _Vira.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraInitialized)
				if err := _Vira.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Vira *ViraFilterer) ParseInitialized(log types.Log) (*ViraInitialized, error) {
	event := new(ViraInitialized)
	if err := _Vira.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraReleasedAllCampaignsProfitIterator is returned from FilterReleasedAllCampaignsProfit and is used to iterate over the raw logs and unpacked data for ReleasedAllCampaignsProfit events raised by the Vira contract.
type ViraReleasedAllCampaignsProfitIterator struct {
	Event *ViraReleasedAllCampaignsProfit // Event containing the contract specifics and raw log

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
func (it *ViraReleasedAllCampaignsProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraReleasedAllCampaignsProfit)
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
		it.Event = new(ViraReleasedAllCampaignsProfit)
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
func (it *ViraReleasedAllCampaignsProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraReleasedAllCampaignsProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraReleasedAllCampaignsProfit represents a ReleasedAllCampaignsProfit event raised by the Vira contract.
type ViraReleasedAllCampaignsProfit struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReleasedAllCampaignsProfit is a free log retrieval operation binding the contract event 0xcd6a39600e90ad2871fdfe0e74b19bfdb288a47916da78a3b333825f3c53a2b2.
//
// Solidity: event ReleasedAllCampaignsProfit(address indexed to, uint256 value)
func (_Vira *ViraFilterer) FilterReleasedAllCampaignsProfit(opts *bind.FilterOpts, to []common.Address) (*ViraReleasedAllCampaignsProfitIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vira.contract.FilterLogs(opts, "ReleasedAllCampaignsProfit", toRule)
	if err != nil {
		return nil, err
	}
	return &ViraReleasedAllCampaignsProfitIterator{contract: _Vira.contract, event: "ReleasedAllCampaignsProfit", logs: logs, sub: sub}, nil
}

// WatchReleasedAllCampaignsProfit is a free log subscription operation binding the contract event 0xcd6a39600e90ad2871fdfe0e74b19bfdb288a47916da78a3b333825f3c53a2b2.
//
// Solidity: event ReleasedAllCampaignsProfit(address indexed to, uint256 value)
func (_Vira *ViraFilterer) WatchReleasedAllCampaignsProfit(opts *bind.WatchOpts, sink chan<- *ViraReleasedAllCampaignsProfit, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vira.contract.WatchLogs(opts, "ReleasedAllCampaignsProfit", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraReleasedAllCampaignsProfit)
				if err := _Vira.contract.UnpackLog(event, "ReleasedAllCampaignsProfit", log); err != nil {
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
func (_Vira *ViraFilterer) ParseReleasedAllCampaignsProfit(log types.Log) (*ViraReleasedAllCampaignsProfit, error) {
	event := new(ViraReleasedAllCampaignsProfit)
	if err := _Vira.contract.UnpackLog(event, "ReleasedAllCampaignsProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraReleasedInvestmentIterator is returned from FilterReleasedInvestment and is used to iterate over the raw logs and unpacked data for ReleasedInvestment events raised by the Vira contract.
type ViraReleasedInvestmentIterator struct {
	Event *ViraReleasedInvestment // Event containing the contract specifics and raw log

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
func (it *ViraReleasedInvestmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraReleasedInvestment)
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
		it.Event = new(ViraReleasedInvestment)
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
func (it *ViraReleasedInvestmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraReleasedInvestmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraReleasedInvestment represents a ReleasedInvestment event raised by the Vira contract.
type ViraReleasedInvestment struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReleasedInvestment is a free log retrieval operation binding the contract event 0x0ab5ede1200a85ab88ecbbf26df740178a0562bd86203fa564fcbc201ed3ef67.
//
// Solidity: event ReleasedInvestment(address indexed to, uint256 value)
func (_Vira *ViraFilterer) FilterReleasedInvestment(opts *bind.FilterOpts, to []common.Address) (*ViraReleasedInvestmentIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vira.contract.FilterLogs(opts, "ReleasedInvestment", toRule)
	if err != nil {
		return nil, err
	}
	return &ViraReleasedInvestmentIterator{contract: _Vira.contract, event: "ReleasedInvestment", logs: logs, sub: sub}, nil
}

// WatchReleasedInvestment is a free log subscription operation binding the contract event 0x0ab5ede1200a85ab88ecbbf26df740178a0562bd86203fa564fcbc201ed3ef67.
//
// Solidity: event ReleasedInvestment(address indexed to, uint256 value)
func (_Vira *ViraFilterer) WatchReleasedInvestment(opts *bind.WatchOpts, sink chan<- *ViraReleasedInvestment, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vira.contract.WatchLogs(opts, "ReleasedInvestment", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraReleasedInvestment)
				if err := _Vira.contract.UnpackLog(event, "ReleasedInvestment", log); err != nil {
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
func (_Vira *ViraFilterer) ParseReleasedInvestment(log types.Log) (*ViraReleasedInvestment, error) {
	event := new(ViraReleasedInvestment)
	if err := _Vira.contract.UnpackLog(event, "ReleasedInvestment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraReleasedTokenProfitIterator is returned from FilterReleasedTokenProfit and is used to iterate over the raw logs and unpacked data for ReleasedTokenProfit events raised by the Vira contract.
type ViraReleasedTokenProfitIterator struct {
	Event *ViraReleasedTokenProfit // Event containing the contract specifics and raw log

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
func (it *ViraReleasedTokenProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraReleasedTokenProfit)
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
		it.Event = new(ViraReleasedTokenProfit)
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
func (it *ViraReleasedTokenProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraReleasedTokenProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraReleasedTokenProfit represents a ReleasedTokenProfit event raised by the Vira contract.
type ViraReleasedTokenProfit struct {
	To      common.Address
	TokenId *big.Int
	Slot    *big.Int
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReleasedTokenProfit is a free log retrieval operation binding the contract event 0x4b54a2d813bd017c0741faa683e2acf7012a515b4fc258ffcf2597acf7117134.
//
// Solidity: event ReleasedTokenProfit(address indexed to, uint256 indexed tokenId, uint256 slot, uint256 value)
func (_Vira *ViraFilterer) FilterReleasedTokenProfit(opts *bind.FilterOpts, to []common.Address, tokenId []*big.Int) (*ViraReleasedTokenProfitIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Vira.contract.FilterLogs(opts, "ReleasedTokenProfit", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ViraReleasedTokenProfitIterator{contract: _Vira.contract, event: "ReleasedTokenProfit", logs: logs, sub: sub}, nil
}

// WatchReleasedTokenProfit is a free log subscription operation binding the contract event 0x4b54a2d813bd017c0741faa683e2acf7012a515b4fc258ffcf2597acf7117134.
//
// Solidity: event ReleasedTokenProfit(address indexed to, uint256 indexed tokenId, uint256 slot, uint256 value)
func (_Vira *ViraFilterer) WatchReleasedTokenProfit(opts *bind.WatchOpts, sink chan<- *ViraReleasedTokenProfit, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Vira.contract.WatchLogs(opts, "ReleasedTokenProfit", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraReleasedTokenProfit)
				if err := _Vira.contract.UnpackLog(event, "ReleasedTokenProfit", log); err != nil {
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
func (_Vira *ViraFilterer) ParseReleasedTokenProfit(log types.Log) (*ViraReleasedTokenProfit, error) {
	event := new(ViraReleasedTokenProfit)
	if err := _Vira.contract.UnpackLog(event, "ReleasedTokenProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Vira contract.
type ViraRoleAdminChangedIterator struct {
	Event *ViraRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ViraRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraRoleAdminChanged)
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
		it.Event = new(ViraRoleAdminChanged)
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
func (it *ViraRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraRoleAdminChanged represents a RoleAdminChanged event raised by the Vira contract.
type ViraRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vira *ViraFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ViraRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Vira.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ViraRoleAdminChangedIterator{contract: _Vira.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vira *ViraFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ViraRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Vira.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraRoleAdminChanged)
				if err := _Vira.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Vira *ViraFilterer) ParseRoleAdminChanged(log types.Log) (*ViraRoleAdminChanged, error) {
	event := new(ViraRoleAdminChanged)
	if err := _Vira.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Vira contract.
type ViraRoleGrantedIterator struct {
	Event *ViraRoleGranted // Event containing the contract specifics and raw log

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
func (it *ViraRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraRoleGranted)
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
		it.Event = new(ViraRoleGranted)
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
func (it *ViraRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraRoleGranted represents a RoleGranted event raised by the Vira contract.
type ViraRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vira *ViraFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ViraRoleGrantedIterator, error) {

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

	logs, sub, err := _Vira.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ViraRoleGrantedIterator{contract: _Vira.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vira *ViraFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ViraRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Vira.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraRoleGranted)
				if err := _Vira.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Vira *ViraFilterer) ParseRoleGranted(log types.Log) (*ViraRoleGranted, error) {
	event := new(ViraRoleGranted)
	if err := _Vira.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Vira contract.
type ViraRoleRevokedIterator struct {
	Event *ViraRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ViraRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraRoleRevoked)
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
		it.Event = new(ViraRoleRevoked)
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
func (it *ViraRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraRoleRevoked represents a RoleRevoked event raised by the Vira contract.
type ViraRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vira *ViraFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ViraRoleRevokedIterator, error) {

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

	logs, sub, err := _Vira.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ViraRoleRevokedIterator{contract: _Vira.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vira *ViraFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ViraRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Vira.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraRoleRevoked)
				if err := _Vira.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Vira *ViraFilterer) ParseRoleRevoked(log types.Log) (*ViraRoleRevoked, error) {
	event := new(ViraRoleRevoked)
	if err := _Vira.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraSetMetadataDescriptorIterator is returned from FilterSetMetadataDescriptor and is used to iterate over the raw logs and unpacked data for SetMetadataDescriptor events raised by the Vira contract.
type ViraSetMetadataDescriptorIterator struct {
	Event *ViraSetMetadataDescriptor // Event containing the contract specifics and raw log

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
func (it *ViraSetMetadataDescriptorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraSetMetadataDescriptor)
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
		it.Event = new(ViraSetMetadataDescriptor)
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
func (it *ViraSetMetadataDescriptorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraSetMetadataDescriptorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraSetMetadataDescriptor represents a SetMetadataDescriptor event raised by the Vira contract.
type ViraSetMetadataDescriptor struct {
	MetadataDescriptor common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetMetadataDescriptor is a free log retrieval operation binding the contract event 0x5252f52e45fc8ee6a7b43cef3645d23e9a470a34182b8b3a12627556635bfc9c.
//
// Solidity: event SetMetadataDescriptor(address indexed metadataDescriptor)
func (_Vira *ViraFilterer) FilterSetMetadataDescriptor(opts *bind.FilterOpts, metadataDescriptor []common.Address) (*ViraSetMetadataDescriptorIterator, error) {

	var metadataDescriptorRule []interface{}
	for _, metadataDescriptorItem := range metadataDescriptor {
		metadataDescriptorRule = append(metadataDescriptorRule, metadataDescriptorItem)
	}

	logs, sub, err := _Vira.contract.FilterLogs(opts, "SetMetadataDescriptor", metadataDescriptorRule)
	if err != nil {
		return nil, err
	}
	return &ViraSetMetadataDescriptorIterator{contract: _Vira.contract, event: "SetMetadataDescriptor", logs: logs, sub: sub}, nil
}

// WatchSetMetadataDescriptor is a free log subscription operation binding the contract event 0x5252f52e45fc8ee6a7b43cef3645d23e9a470a34182b8b3a12627556635bfc9c.
//
// Solidity: event SetMetadataDescriptor(address indexed metadataDescriptor)
func (_Vira *ViraFilterer) WatchSetMetadataDescriptor(opts *bind.WatchOpts, sink chan<- *ViraSetMetadataDescriptor, metadataDescriptor []common.Address) (event.Subscription, error) {

	var metadataDescriptorRule []interface{}
	for _, metadataDescriptorItem := range metadataDescriptor {
		metadataDescriptorRule = append(metadataDescriptorRule, metadataDescriptorItem)
	}

	logs, sub, err := _Vira.contract.WatchLogs(opts, "SetMetadataDescriptor", metadataDescriptorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraSetMetadataDescriptor)
				if err := _Vira.contract.UnpackLog(event, "SetMetadataDescriptor", log); err != nil {
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
func (_Vira *ViraFilterer) ParseSetMetadataDescriptor(log types.Log) (*ViraSetMetadataDescriptor, error) {
	event := new(ViraSetMetadataDescriptor)
	if err := _Vira.contract.UnpackLog(event, "SetMetadataDescriptor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraSlotChangedIterator is returned from FilterSlotChanged and is used to iterate over the raw logs and unpacked data for SlotChanged events raised by the Vira contract.
type ViraSlotChangedIterator struct {
	Event *ViraSlotChanged // Event containing the contract specifics and raw log

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
func (it *ViraSlotChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraSlotChanged)
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
		it.Event = new(ViraSlotChanged)
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
func (it *ViraSlotChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraSlotChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraSlotChanged represents a SlotChanged event raised by the Vira contract.
type ViraSlotChanged struct {
	TokenId *big.Int
	OldSlot *big.Int
	NewSlot *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlotChanged is a free log retrieval operation binding the contract event 0xe4f48c240d3b994948aa54f3e2f5fca59263dfe1d52b6e4cf39a5d249b5ccb65.
//
// Solidity: event SlotChanged(uint256 indexed _tokenId, uint256 indexed _oldSlot, uint256 indexed _newSlot)
func (_Vira *ViraFilterer) FilterSlotChanged(opts *bind.FilterOpts, _tokenId []*big.Int, _oldSlot []*big.Int, _newSlot []*big.Int) (*ViraSlotChangedIterator, error) {

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

	logs, sub, err := _Vira.contract.FilterLogs(opts, "SlotChanged", _tokenIdRule, _oldSlotRule, _newSlotRule)
	if err != nil {
		return nil, err
	}
	return &ViraSlotChangedIterator{contract: _Vira.contract, event: "SlotChanged", logs: logs, sub: sub}, nil
}

// WatchSlotChanged is a free log subscription operation binding the contract event 0xe4f48c240d3b994948aa54f3e2f5fca59263dfe1d52b6e4cf39a5d249b5ccb65.
//
// Solidity: event SlotChanged(uint256 indexed _tokenId, uint256 indexed _oldSlot, uint256 indexed _newSlot)
func (_Vira *ViraFilterer) WatchSlotChanged(opts *bind.WatchOpts, sink chan<- *ViraSlotChanged, _tokenId []*big.Int, _oldSlot []*big.Int, _newSlot []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Vira.contract.WatchLogs(opts, "SlotChanged", _tokenIdRule, _oldSlotRule, _newSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraSlotChanged)
				if err := _Vira.contract.UnpackLog(event, "SlotChanged", log); err != nil {
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
func (_Vira *ViraFilterer) ParseSlotChanged(log types.Log) (*ViraSlotChanged, error) {
	event := new(ViraSlotChanged)
	if err := _Vira.contract.UnpackLog(event, "SlotChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraTokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the Vira contract.
type ViraTokenMintedIterator struct {
	Event *ViraTokenMinted // Event containing the contract specifics and raw log

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
func (it *ViraTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraTokenMinted)
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
		it.Event = new(ViraTokenMinted)
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
func (it *ViraTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraTokenMinted represents a TokenMinted event raised by the Vira contract.
type ViraTokenMinted struct {
	To      common.Address
	TokenId *big.Int
	Slot    *big.Int
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenMinted is a free log retrieval operation binding the contract event 0x10546b1a6f5245ff0ffa18c256b9e46859c585cbb473b453fcd4c2dc39ae08db.
//
// Solidity: event TokenMinted(address indexed to, uint256 indexed tokenId, uint256 slot, uint256 value)
func (_Vira *ViraFilterer) FilterTokenMinted(opts *bind.FilterOpts, to []common.Address, tokenId []*big.Int) (*ViraTokenMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Vira.contract.FilterLogs(opts, "TokenMinted", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ViraTokenMintedIterator{contract: _Vira.contract, event: "TokenMinted", logs: logs, sub: sub}, nil
}

// WatchTokenMinted is a free log subscription operation binding the contract event 0x10546b1a6f5245ff0ffa18c256b9e46859c585cbb473b453fcd4c2dc39ae08db.
//
// Solidity: event TokenMinted(address indexed to, uint256 indexed tokenId, uint256 slot, uint256 value)
func (_Vira *ViraFilterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *ViraTokenMinted, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Vira.contract.WatchLogs(opts, "TokenMinted", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraTokenMinted)
				if err := _Vira.contract.UnpackLog(event, "TokenMinted", log); err != nil {
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
func (_Vira *ViraFilterer) ParseTokenMinted(log types.Log) (*ViraTokenMinted, error) {
	event := new(ViraTokenMinted)
	if err := _Vira.contract.UnpackLog(event, "TokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vira contract.
type ViraTransferIterator struct {
	Event *ViraTransfer // Event containing the contract specifics and raw log

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
func (it *ViraTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraTransfer)
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
		it.Event = new(ViraTransfer)
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
func (it *ViraTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraTransfer represents a Transfer event raised by the Vira contract.
type ViraTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Vira *ViraFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ViraTransferIterator, error) {

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

	logs, sub, err := _Vira.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ViraTransferIterator{contract: _Vira.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Vira *ViraFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ViraTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Vira.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraTransfer)
				if err := _Vira.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Vira *ViraFilterer) ParseTransfer(log types.Log) (*ViraTransfer, error) {
	event := new(ViraTransfer)
	if err := _Vira.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraTransferValueIterator is returned from FilterTransferValue and is used to iterate over the raw logs and unpacked data for TransferValue events raised by the Vira contract.
type ViraTransferValueIterator struct {
	Event *ViraTransferValue // Event containing the contract specifics and raw log

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
func (it *ViraTransferValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraTransferValue)
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
		it.Event = new(ViraTransferValue)
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
func (it *ViraTransferValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraTransferValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraTransferValue represents a TransferValue event raised by the Vira contract.
type ViraTransferValue struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferValue is a free log retrieval operation binding the contract event 0x0b2aac84f3ec956911fd78eae5311062972ff949f38412e8da39069d9f068cc6.
//
// Solidity: event TransferValue(uint256 indexed _fromTokenId, uint256 indexed _toTokenId, uint256 _value)
func (_Vira *ViraFilterer) FilterTransferValue(opts *bind.FilterOpts, _fromTokenId []*big.Int, _toTokenId []*big.Int) (*ViraTransferValueIterator, error) {

	var _fromTokenIdRule []interface{}
	for _, _fromTokenIdItem := range _fromTokenId {
		_fromTokenIdRule = append(_fromTokenIdRule, _fromTokenIdItem)
	}
	var _toTokenIdRule []interface{}
	for _, _toTokenIdItem := range _toTokenId {
		_toTokenIdRule = append(_toTokenIdRule, _toTokenIdItem)
	}

	logs, sub, err := _Vira.contract.FilterLogs(opts, "TransferValue", _fromTokenIdRule, _toTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ViraTransferValueIterator{contract: _Vira.contract, event: "TransferValue", logs: logs, sub: sub}, nil
}

// WatchTransferValue is a free log subscription operation binding the contract event 0x0b2aac84f3ec956911fd78eae5311062972ff949f38412e8da39069d9f068cc6.
//
// Solidity: event TransferValue(uint256 indexed _fromTokenId, uint256 indexed _toTokenId, uint256 _value)
func (_Vira *ViraFilterer) WatchTransferValue(opts *bind.WatchOpts, sink chan<- *ViraTransferValue, _fromTokenId []*big.Int, _toTokenId []*big.Int) (event.Subscription, error) {

	var _fromTokenIdRule []interface{}
	for _, _fromTokenIdItem := range _fromTokenId {
		_fromTokenIdRule = append(_fromTokenIdRule, _fromTokenIdItem)
	}
	var _toTokenIdRule []interface{}
	for _, _toTokenIdItem := range _toTokenId {
		_toTokenIdRule = append(_toTokenIdRule, _toTokenIdItem)
	}

	logs, sub, err := _Vira.contract.WatchLogs(opts, "TransferValue", _fromTokenIdRule, _toTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraTransferValue)
				if err := _Vira.contract.UnpackLog(event, "TransferValue", log); err != nil {
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
func (_Vira *ViraFilterer) ParseTransferValue(log types.Log) (*ViraTransferValue, error) {
	event := new(ViraTransferValue)
	if err := _Vira.contract.UnpackLog(event, "TransferValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Vira contract.
type ViraUpgradedIterator struct {
	Event *ViraUpgraded // Event containing the contract specifics and raw log

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
func (it *ViraUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraUpgraded)
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
		it.Event = new(ViraUpgraded)
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
func (it *ViraUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraUpgraded represents a Upgraded event raised by the Vira contract.
type ViraUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vira *ViraFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ViraUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vira.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ViraUpgradedIterator{contract: _Vira.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vira *ViraFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ViraUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vira.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraUpgraded)
				if err := _Vira.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Vira *ViraFilterer) ParseUpgraded(log types.Log) (*ViraUpgraded, error) {
	event := new(ViraUpgraded)
	if err := _Vira.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViraWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Vira contract.
type ViraWithdrawnIterator struct {
	Event *ViraWithdrawn // Event containing the contract specifics and raw log

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
func (it *ViraWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ViraWithdrawn)
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
		it.Event = new(ViraWithdrawn)
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
func (it *ViraWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ViraWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ViraWithdrawn represents a Withdrawn event raised by the Vira contract.
type ViraWithdrawn struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x430648de173157e069201c943adb2d4e340e7cf5b27b1b09c9cb852f03d63b56.
//
// Solidity: event Withdrawn(uint256 value)
func (_Vira *ViraFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*ViraWithdrawnIterator, error) {

	logs, sub, err := _Vira.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &ViraWithdrawnIterator{contract: _Vira.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x430648de173157e069201c943adb2d4e340e7cf5b27b1b09c9cb852f03d63b56.
//
// Solidity: event Withdrawn(uint256 value)
func (_Vira *ViraFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ViraWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Vira.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ViraWithdrawn)
				if err := _Vira.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_Vira *ViraFilterer) ParseWithdrawn(log types.Log) (*ViraWithdrawn, error) {
	event := new(ViraWithdrawn)
	if err := _Vira.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
