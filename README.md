# Go SDK For Akchain

* [Go SDK For Akchain](#go-sdk-for-akchain)
	* [1. Overview](#1-overview)
	* [2. How to use?](#2-how-to-use)
		* [2.1 Blockchain API](#21-blockchain-api)
			* [2.1.1 Get current block height](#211-get-current-block-height)
			* [2.1.2 Get current block hash](#212-get-current-block-hash)
			* [2.1.3 Get block by height](#213-get-block-hash-by-height)
			* [2.1.4 Check blockchain healthy](#214-check-blockchain-healthy)
			* [2.1.5 Get network info](#215-get-network-info)
			* [2.1.6 Get node info](#216-get-node-info)
			* [2.1.7 Get version of Akchain](#217-get-version-of-akchain)
			* [2.1.8 Get node status](#218-get-node-status)
			* [2.1.9 Get genesis info](#219-get-genesis-info)
			* [2.1.10 Get config](#2110-get-config)
			* [2.1.11 List node peers](#2111-list-node-peers)
			* [2.1.12 Get hardware usage](#2112-get-hardware-usage)
			* [2.1.13 Get validators by height](#2113-get-validators-by-height)
			* [2.1.14 Get node type by height](#2114-get-node-type-by-height)
			* [2.1.15 Get block header by height](#2115-get-block-header-by-height)
			* [2.1.16 Get block detail by height or hash](#2116-get-block-detail-by-height-or-hash)
			* [2.1.17 Dump consensus status](#2117-dump-consensus-status)
            * [2.1.18 Get consensus params by height](#2118-get-consensus-params-by-height)
            * [2.1.19 Get consensus state](#2119-get-consensus-state)
		* [2.2 Account API](#22-account-api)
			* [2.2.1 Create new account](#221-create-new-account)
			* [2.2.2 Create new account from private key](#222-create-new-account-from-private-key)
			* [2.2.3 Generate mnemonic](#223-generate-mnemonic)
			* [2.2.4 Get keys from mnemonic](#224-get-keys-from-mnemonic)
			* [2.2.5 Load account from serialization string](#225-load-account-from-serialization-string)
			* [2.2.6 Serializing account](#226-serializing-account)
			* [2.2.7 List account balance](#227-list-account-balance)
			* [2.2.8 Get account info](#228-get-account-info)
			* [2.2.9 List all accounts](#229-list-all-accounts)
			* [2.2.10 Sign message](#2210-sign-message)
			* [2.2.11 Verify message](#2211-verify-message)
			* [2.2.12 Validate address](#2212-validate-address)	
			* [2.2.13 Recovery account](#2213-recovery-account)			
			* [2.2.14 List public keys](#2214-list-public-keys)	
		* [2.3 Asset API](#23-asset-api)
			* [2.3.1 Create asset](#231-create-asset)
			* [2.3.2 Get asset info](#232-get-asset-info)
			* [2.3.3 Update asset alias](#233-update-asset-alias)
			* [2.3.4 Issue asset(sync or async)](#234-issue-asset-sync-or-async)
			* [2.3.5 Retire asset(sync or async)](#235-retire-asset-sync-or-async)
			* [2.3.6 Transfer asset(sync or async)](#236-transfer-asset-sync-or-async)
		* [2.4 Transaction API](#24-transaction-api)
			* [2.4.1 Get number of unconfirmed transactions](#241-get-number-of-unconfirmed-transactions)
			* [2.4.2 List unconfirmed transactions](#242-list-unconfirmed-transactions)
			* [2.4.3 Search transaction](#243-search-transaction)
		* [2.5 Contract API](#25-contract-api)
			* [2.5.1 Upload wasm contract file](#251-upload-wasm-contract-file)
			* [2.5.2 Deploy contract](#252-deploy-contract)
			* [2.5.3 Upgrade contract](#253-upgrade-contract)
			* [2.5.4 List contracts by account ID](#254-list-contracts-by-account-id)
			* [2.5.5 Invoke contract method](#255-invoke-contract-method)
			* [2.5.6 Query contract](#256-query-contract)
* [License](#license)

## 1. Overview
This is a comprehensive Go library for the Akchain blockchain. Currently, it supports account management, digital asset management and communication with the Akchain. In the future it will also support more rich functions and applications.

## 2. How to use?

First, create an `AkchainSdk` instance with the `NewAkchainSdk` method.

```
akcSdk = NewAkchainSdk()
```

Next, create a rpc or rest client.

```
akcSdk.NewRestClient().SetAddress("http://localhost:30000")
```

Then, call the rpc server through the sdk instance.


### 2.1 Blockchain API

#### 2.1.1 Get current block height

```
akcSdk.GetBlockHeight() (uint64, error)
```

#### 2.1.2 Get current block hash

```
akcSdk.GetBlockHash() (string, error)
```

#### 2.1.3 Get block hash by height

```
akcSdk.GetBlockHashByHeight(height uint64) (string, error)
```

#### 2.1.4 Check blockchain healthy

```
akcSdk.CheckHealth() (error)
```

#### 2.1.5 Get network info

```
akcSdk.GetNetInfo() (*types.NetInfo, error)
```

#### 2.1.6 Get node info

```
akcSdk.GetNodeInfo() (*types.DefaultNodeInfo, error)
```

#### 2.1.7 Get version of Akchain

```
akcSdk.GetVersion() (*types.VersionInfo, error)
```

#### 2.1.8 Get node status

```
akcSdk.GetStatus() (*types.StatusInfo, error)
```

#### 2.1.9 Get genesis info

```
akcSdk.GetGenesis() (*types.ResultGenesis, error)
```

#### 2.1.10 Get config

```
akcSdk.GetConfig() (*types.BaseConfig, error)
```

#### 2.1.11 List node peers

```
akcSdk.ListPeers() (*types.ListPeersInfo, error)
```

#### 2.1.12 Get hardware usage

```
akcSdk.GetHardwareUsage() (*types.HardwareUsage, error)
```

#### 2.1.13 Get validators by height

```
akcSdk.GetValidatorsByHeight(height uint64) (*types.HeightValidators, error)
```

#### 2.1.14 Get node type by height

```
akcSdk.GetNodeTypeByHeight(height uint64) (*types.NodeType, errorr)
```

#### 2.1.15 Get block header by height

```
akcSdk.GetBlockHeader(minHeight uint64, maxHeight uint64) (*types.BlockChainInfo, error)
```

#### 2.1.16 Get block detail by height or hash

```
akcSdk.GetBlockDetail(height uint64, hash string) (*types.BlockResp, error)
```

#### 2.1.17 Dump consensus status

```
akcSdk.DumpConsensusState()(*types.DumpConsensusState, error) 
```

#### 2.1.18 Get consensus params by height

```
akcSdk.GetConsensusParamsByHeight(height uint64) (*types.ConsensusParams, error)
```

#### 2.1.19 Get consensus state

```
akcSdk.GetConsensusState() (*types.ConsensusStateInfo, error)
```

### 2.2 Account API

#### 2.2.1 Create new account

```
akcSdk.NewAccount() (*account.Account, error)
```

#### 2.2.2 Create new account from private key

```
akcSdk.NewAccountFromPrivateKey(xprv chainkd.XPrv) (*account.Account, error)
```

#### 2.2.3 Generate mnemonic

```
akcSdk.GenerateMnemonicCodesStr() (string, error)
```

#### 2.2.4 Get keys from mnemonic

```
akcSdk.GetKeysFromMnemonic(mnemonicCodesStr string) (string, string, error)
```

The default settings for an account uses ECDSA with SHA256withECDSA as signature scheme.

#### 2.2.5 Load account from serialization string

```
LoadAcount(accountStr string) (*Account, error)
```

#### 2.2.6 Serializing account

```
account.ToString()
```

#### 2.2.7 List account balance

```
akcSdk.ListAccountBalances(accountID string) (*[]types.AccountBalanceInfo, error)
```

#### 2.2.8 Get account info

```
akcSdk.GetAccountInfo(accountID string) (*[]types.AccountInfo, error)
```

#### 2.2.9 List all accounts

```
akcSdk.ListAccounts()(*[]types.CreateAccountInfo, error)
```

#### 2.2.10 Sign message

```
account.Sign(data []byte) []byte
```

#### 2.2.11 Verify message

```
account.Verify(msg []byte, sig []byte) bool
```

#### 2.2.12 Validate address

```
akcSdk.ValidateAddress(address string) (*types.ValidateAddressInfo, error)
```

#### 2.2.13 Recovery account

```
akcSdk.RecoveryAccount(xpubs []string) error
```

#### 2.2.14 List public keys

```
akcSdk.ListPubKeys(accountID string) (*types.ListPublicKeysInfo, error)
```

### 2.3 Asset API

#### 2.3.1 Create asset

```
akcSdk.CreateAsset(xpub string, alias string) (*types.CreateAssetInfo, error)
```

#### 2.3.2 Get asset info

```
akcSdk.GetAsset(id string) (*types.AssetInfo, error)
```

#### 2.3.3 Update asset alias

```
akcSdk.akcSdk.UpdateAssetAlias(id string, alias string) error
```

#### 2.3.4 Issue asset (sync or async)

```
akcSdk.IssueAsset(amount uint64, assetID string, destAddress string, singer *account.Account) (string, error)
```

```
akcSdk.IssueAssetAsync(amount uint64, assetID string, destAddress string, singer *account.Account) (string, error)
```

#### 2.3.5 Retire asset (sync or async)

```
akcSdk.RetireAsset(amount uint64, assetID string, singer *account.Account) (string, error)
```

```
akcSdk.RetireAssetAsync(amount uint64, assetID string, singer *account.Account) (string, error)
```

#### 2.3.6 Transfer asset (sync or async)

```
akcSdk.TransferAsset(amount uint64, assetID string, destAddress string, singer *account.Account) (string, error)
```

```
akcSdk.TransferAssetAsync(amount uint64, assetID string, destAddress string, singer *account.Account) (string, error)
```

### 2.4 Transaction API


#### 2.4.1 Get number of unconfirmed transactions

```
akcSdk.GetNumUnconfirmedTx() (int, error)
```

#### 2.4.2 List unconfirmed transactions

```
akcSdk.ListUnconfirmedTx() (*types.UnconfirmedTx, error)
```

#### 2.4.3 Search transaction

```
akcSdk.SearchTx(tx string) (*types.Tx, error)
```

### 2.5 Contract API


#### 2.5.1 Upload wasm contract file

```
akcSdk.UploadContract(filePath string) (string, error)
```

#### 2.5.2 Deploy contract

```
akcSdk.DeployContract(contractPath string, contractName string, args string, signer *account.Account) (string, error)
```

#### 2.5.3 Upgrade contract

```
akcSdk.UpgradeContract(contractPath string, contractName string, signer *account.Account) (string, error)
```

#### 2.5.4 List contracts by account ID

```
akcSdk.ListContracts(accountID string) (*[]types.ContractInfo, error)
```

#### 2.5.5 Invoke contract method

```
akcSdk.InvokeMethod(contractName string, method string, args string, signer *account.Account) (string, error)
```

#### 2.5.6 Query contract 

```
akcSdk.QueryContract(contractName string, method string, args string) ([]*types.ContractResponse, error)
```

# License

The Akchain library (i.e. all the code outside of the cmd directory) is licensed under the GNU Lesser General Public License v3.0, also included in our repository in the License file.
