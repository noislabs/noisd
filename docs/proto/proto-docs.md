<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [nois/allocation/v1/params.proto](#nois/allocation/v1/params.proto)
    - [DistributionProportions](#nois.allocation.v1.DistributionProportions)
    - [Params](#nois.allocation.v1.Params)
    - [WeightedAddress](#nois.allocation.v1.WeightedAddress)
  
- [nois/allocation/v1/rewards.proto](#nois/allocation/v1/rewards.proto)
    - [ValidatorReward](#nois.allocation.v1.ValidatorReward)
  
- [nois/allocation/v1/genesis.proto](#nois/allocation/v1/genesis.proto)
    - [GenesisState](#nois.allocation.v1.GenesisState)
  
- [nois/allocation/v1/query.proto](#nois/allocation/v1/query.proto)
    - [QueryClaimableRewardsRequest](#nois.allocation.v1.QueryClaimableRewardsRequest)
    - [QueryClaimableRewardsResponse](#nois.allocation.v1.QueryClaimableRewardsResponse)
    - [QueryParamsRequest](#nois.allocation.v1.QueryParamsRequest)
    - [QueryParamsResponse](#nois.allocation.v1.QueryParamsResponse)
  
    - [Query](#nois.allocation.v1.Query)
  
- [nois/allocation/v1/tx.proto](#nois/allocation/v1/tx.proto)
    - [MsgClaimRewards](#nois.allocation.v1.MsgClaimRewards)
    - [MsgClaimRewardsResponse](#nois.allocation.v1.MsgClaimRewardsResponse)
  
    - [Msg](#nois.allocation.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="nois/allocation/v1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nois/allocation/v1/params.proto



<a name="nois.allocation.v1.DistributionProportions"></a>

### DistributionProportions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `developer_rewards` | [string](#string) |  |  |
| `validator_rewards` | [string](#string) |  |  |
| `randomness_rewards` | [string](#string) |  |  |






<a name="nois.allocation.v1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `distribution_proportions` | [DistributionProportions](#nois.allocation.v1.DistributionProportions) |  | distribution_proportions defines the proportion of the minted denom |
| `weighted_developer_rewards_receivers` | [WeightedAddress](#nois.allocation.v1.WeightedAddress) | repeated | address to receive developer rewards |
| `randomness_rewards_receiver` | [string](#string) |  | address to receive randomness rewards |






<a name="nois.allocation.v1.WeightedAddress"></a>

### WeightedAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `weight` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nois/allocation/v1/rewards.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nois/allocation/v1/rewards.proto



<a name="nois.allocation.v1.ValidatorReward"></a>

### ValidatorReward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | validator address |
| `rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | accumulated validator rewards |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nois/allocation/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nois/allocation/v1/genesis.proto



<a name="nois.allocation.v1.GenesisState"></a>

### GenesisState
GenesisState defines the allocation module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nois.allocation.v1.Params) |  |  |
| `validator_rewards` | [ValidatorReward](#nois.allocation.v1.ValidatorReward) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nois/allocation/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nois/allocation/v1/query.proto



<a name="nois.allocation.v1.QueryClaimableRewardsRequest"></a>

### QueryClaimableRewardsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="nois.allocation.v1.QueryClaimableRewardsResponse"></a>

### QueryClaimableRewardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nois.allocation.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="nois.allocation.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nois.allocation.v1.Params) |  | params defines the parameters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nois.allocation.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nois.allocation.v1.QueryParamsRequest) | [QueryParamsResponse](#nois.allocation.v1.QueryParamsResponse) |  | GET|/nois/alocation/v1/params|
| `ClaimableRewards` | [QueryClaimableRewardsRequest](#nois.allocation.v1.QueryClaimableRewardsRequest) | [QueryClaimableRewardsResponse](#nois.allocation.v1.QueryClaimableRewardsResponse) |  | GET|/nois/alocation/v1/claimable_rewards|

 <!-- end services -->



<a name="nois/allocation/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nois/allocation/v1/tx.proto



<a name="nois.allocation.v1.MsgClaimRewards"></a>

### MsgClaimRewards



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |






<a name="nois.allocation.v1.MsgClaimRewardsResponse"></a>

### MsgClaimRewardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `claimed_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | claimed rewards amount |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nois.allocation.v1.Msg"></a>

### Msg
Msg defines the allocation Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ClaimRewards` | [MsgClaimRewards](#nois.allocation.v1.MsgClaimRewards) | [MsgClaimRewardsResponse](#nois.allocation.v1.MsgClaimRewardsResponse) |  | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

