// Copyright © 2023 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package coremsgs

//revive:disable
var (
	CoreSystemNSDescription = ffm("core.systemNSDescription", "FireFly system namespace (legacy - no longer used by newer versions)")

	APIParamsConfigRecordKeyUpdate          = ffm("api.params.configRecordKey.update", "The configuration key to update. This should use dot notation to reference a key documented in https://hyperledger.github.io/firefly/reference/config.html")
	APIParamsConfigRecordKeyGet             = ffm("api.params.configRecordKey.get", "The configuration key to get. This should use dot notation to reference a key documented in https://hyperledger.github.io/firefly/reference/config.html")
	APIParamsOperationIDGet                 = ffm("api.params.operationID.get", "The operation ID key to get")
	APIParamsOperationNamespacedID          = ffm("api.params.spiOperationID", "The operation ID as passed to the connector when the operation was performed, including the 'namespace:' prefix")
	APIParamsNamespace                      = ffm("api.params.namespace", "The namespace which scopes this request")
	APIParamsContractListenerNameOrID       = ffm("api.params.contractListenerNameOrID", "The contract listener name or ID")
	APIParamsContractListenerID             = ffm("api.params.contractListenerID", "The contract listener ID")
	APIParamsSubscriptionID                 = ffm("api.params.subscriptionID", "The subscription ID")
	APIParamsBatchID                        = ffm("api.params.batchId", "The batch ID")
	APIParamsBlockchainEventID              = ffm("api.params.blockchainEventID", "The blockchain event ID")
	APIParamsCollectionID                   = ffm("api.params.collectionID", "The collection ID")
	APIParamsContractAPIName                = ffm("api.params.contractAPIName", "The name of the contract API")
	APIParamsContractInterfaceName          = ffm("api.params.contractInterfaceName", "The name of the contract interface")
	APIParamsContractInterfaceVersion       = ffm("api.params.contractInterfaceVersion", "The version of the contract interface")
	APIParamsContractInterfaceID            = ffm("api.params.contractInterfaceID", "The ID of the contract interface")
	APIParamsContractInterfaceFetchChildren = ffm("api.params.contractInterfaceFetchChildren", "When set, the API will return the full FireFly Interface document including all methods, events, and parameters")
	APIParamsNSIncludeInitializing          = ffm("api.params.nsIncludeInitializing", "When set, the API will return namespaces even if they are not yet initialized, including in error cases where an initializationError is included")
	APIParamsBlobID                         = ffm("api.params.blobID", "The blob ID")
	APIParamsDataID                         = ffm("api.params.dataID", "The data item ID")
	APIParamsDatatypeName                   = ffm("api.params.datatypeName", "The name of the datatype")
	APIParamsDatatypeVersion                = ffm("api.params.datatypeVersion", "The version of the datatype")
	APIParamsDataParentPath                 = ffm("api.params.dataParentPath", "The parent path to query")
	APIParamsEventID                        = ffm("api.params.eventID", "The event ID")
	APIParamsFetchReferences                = ffm("api.params.fetchReferences", "When set, the API will return the record that this item references in its 'reference' field")
	APIParamsFetchReference                 = ffm("api.params.fetchReference", "When set, the API will return the record that this item references in its 'reference' field")
	APIParamsGroupHash                      = ffm("api.params.groupID", "The hash of the group")
	APIParamsFetchVerifiers                 = ffm("api.params.fetchVerifiers", "When set, the API will return the verifier for this identity")
	APIParamsIdentityID                     = ffm("api.params.identityID", "The identity ID, which is a UUID generated by FireFly")
	APIParamsMessageID                      = ffm("api.params.messageID", "The message ID")
	APIParamsDID                            = ffm("api.params.DID", "The identity DID")
	APIParamsNodeNameOrID                   = ffm("api.params.nodeNameOrID", "The name or ID of the node")
	APIParamsOrgNameOrID                    = ffm("api.params.orgNameOrID", "The name or ID of the org")
	APIParamsTokenAccountKey                = ffm("api.params.tokenAccountKey", "The key for the token account. The exact format may vary based on the token connector use")
	APIParamsTokenPoolNameOrID              = ffm("api.params.tokenPoolNameOrID", "The token pool name or ID")
	APIParamsTokenTransferFromOrTo          = ffm("api.params.tokenTransferFromOrTo", "The sending or receiving token account for a token transfer")
	APIParamsTokenTransferID                = ffm("api.params.tokenTransferID", "The token transfer ID")
	APIParamsTransactionID                  = ffm("api.params.transactionID", "The transaction ID")
	APIParamsVerifierHash                   = ffm("api.params.verifierID", "The hash of the verifier")
	APIParamsMethodPath                     = ffm("api.params.methodPath", "The name or uniquely generated path name of a method on a smart contract")
	APIParamsEventPath                      = ffm("api.params.eventPath", "The name or uniquely generated path name of a event on a smart contract")
	APIParamsInterfaceID                    = ffm("api.params.interfaceID", "The contract interface ID")
	APIParamsValidator                      = ffm("api.params.validator", "The validator type for this data item. Options are: \"json\", \"none\", or \"definition\"")
	APIParamsMetadata                       = ffm("api.params.metadata", "Metadata associated with this data item")
	APIParamsAutometa                       = ffm("api.params.autometa", "When set, FireFly will automatically generate JSON metadata with the upload details")
	APIParamsContractAPIID                  = ffm("api.params.contractAPIID", "The ID of the contract API")
	APIParamsFetchStatus                    = ffm("api.params.fetchStatus", "When set, the API will return additional status information if available")

	APIEndpointsAdminGetNamespaceByName = ffm("api.endpoints.adminGetNamespaceByName", "Gets a namespace by name")
	APIEndpointsAdminGetNamespaces      = ffm("api.endpoints.adminGetNamespaces", "List namespaces")
	APIEndpointsAdminGetOpByID          = ffm("api.endpoints.adminGetOpByID", "Gets an operation by ID")
	APIEndpointsAdminGetOps             = ffm("api.endpoints.adminGetOps", "Lists operations")
	APIEndpointsAdminPostReset          = ffm("api.endpoints.adminPostResetConfig", "Restarts FireFly Core HTTP servers and apply all configuration updates")
	APIEndpointsAdminPatchOpByID        = ffm("api.endpoints.adminPatchOpByID", "Updates an operation by ID")
	APIEndpointsAdminGetListenerByID    = ffm("api.endpoints.adminGetListenerByID", "Gets a contract listener by ID")
	APIEndpointsAdminGetListeners       = ffm("api.endpoints.adminGetListeners", "Lists contract listeners")

	APIEndpointsDeleteContractAPI               = ffm("api.endpoints.deleteContractAPI", "Delete a contract API")
	APIEndpointsDeleteContractInterface         = ffm("api.endpoints.deleteContractInterface", "Delete a contract interface")
	APIEndpointsDeleteContractListener          = ffm("api.endpoints.deleteContractListener", "Deletes a contract listener referenced by its name or its ID")
	APIEndpointsDeleteSubscription              = ffm("api.endpoints.deleteSubscription", "Deletes a subscription")
	APIEndpointsDeleteTokenPool                 = ffm("api.endpoints.deleteTokenPool", "Delete a token pool")
	APIEndpointsGetBatchBbyID                   = ffm("api.endpoints.getBatchByID", "Gets a message batch")
	APIEndpointsGetBatches                      = ffm("api.endpoints.getBatches", "Gets a list of message batches")
	APIEndpointsGetBlockchainEventByID          = ffm("api.endpoints.getBlockchainEventByID", "Gets a blockchain event")
	APIEndpointsListBlockchainEvents            = ffm("api.endpoints.getBlockchainEvents", "Gets a list of blockchain events")
	APIEndpointsGetChartHistogram               = ffm("api.endpoints.getChartHistogram", "Gets a JSON object containing statistics data that can be used to build a graphical representation of recent activity in a given database collection")
	APIEndpointsGetContractAPIByName            = ffm("api.endpoints.getContractAPIByName", "Gets information about a contract API, including the URLs for the OpenAPI Spec and Swagger UI for the API")
	APIEndpointsGetContractAPIs                 = ffm("api.endpoints.getContractAPIs", "Gets a list of contract APIs that have been published")
	APIEndpointsGetContractInterfaceNameVersion = ffm("api.endpoints.getContractInterfaceNameVersion", "Gets a contract interface by its name and version")
	APIEndpointsGetContractInterface            = ffm("api.endpoints.getContractInterface", "Gets a contract interface by its ID")
	APIEndpointsGetContractInterfaces           = ffm("api.endpoints.getContractInterfaces", "Gets a list of contract interfaces that have been published")
	APIEndpointsGetContractListenerByNameOrID   = ffm("api.endpoints.getContractListenerByNameOrID", "Gets a contract listener by its name or ID")
	APIEndpointsGetContractListeners            = ffm("api.endpoints.getContractListeners", "Gets a list of contract listeners")
	APIEndpointsGetDataBlob                     = ffm("api.endpoints.getDataBlob", "Downloads the original file that was previously uploaded or received")
	APIEndpointsGetDataValue                    = ffm("api.endpoints.getDataValue", "Downloads the JSON value of the data resource, without the associated metadata")
	APIEndpointsGetDataByID                     = ffm("api.endpoints.getDataByID", "Gets a data item by its ID, including metadata about this item")
	APIEndpointsDeleteData                      = ffm("api.endpoints.deleteData", "Deletes a data item by its ID, including metadata about this item")
	APIEndpointsGetDataMsgs                     = ffm("api.endpoints.getDataMsgs", "Gets a list of the messages associated with a data item")
	APIEndpointsGetData                         = ffm("api.endpoints.getData", "Gets a list of data items")
	APIEndpointsGetDataSubPaths                 = ffm("api.endpoints.getDataSubPaths", "Gets a list of path names of named blob data, underneath a given parent path ('/' path prefixes are automatically pre-prepended)")
	APIEndpointsGetDatatypeByName               = ffm("api.endpoints.getDatatypeByName", "Gets a datatype by its name and version")
	APIEndpointsGetDatatypes                    = ffm("api.endpoints.getDatatypes", "Gets a list of datatypes that have been published")
	APIEndpointsGetEventByID                    = ffm("api.endpoints.eventID", "Gets an event by its ID")
	APIEndpointsGetEvents                       = ffm("api.endpoints.getEvents", "Gets a list of events")
	APIEndpointsGetGroupByHash                  = ffm("api.endpoints.getGroupByHash", "Gets a group by its ID (hash)")
	APIEndpointsGetGroups                       = ffm("api.endpoints.getGroups", "Gets a list of groups")
	APIEndpointsGetIdentities                   = ffm("api.endpoints.getIdentities", "Gets a list of all identities that have been registered in the namespace")
	APIEndpointsGetIdentityByID                 = ffm("api.endpoints.getIdentityByID", "Gets an identity by its ID")
	APIEndpointsGetIdentityDID                  = ffm("api.endpoints.getIdentityDID", "Gets the DID for an identity based on its ID")
	APIEndpointsGetIdentityVerifiers            = ffm("api.endpoints.getIdentityVerifiers", "Gets the verifiers for an identity")
	APIEndpointsGetMsgByID                      = ffm("api.endpoints.getMsgByID", "Gets a message by its ID")
	APIEndpointsGetMsgData                      = ffm("api.endpoints.getMsgData", "Gets the list of data items that are attached to a message")
	APIEndpointsGetMsgEvents                    = ffm("api.endpoints.getMsgEvents", "Gets the list of events for a message")
	APIEndpointsGetMsgTxn                       = ffm("api.endpoints.getMsgTxn", "Gets the transaction for a message")
	APIEndpointsGetMsgs                         = ffm("api.endpoints.getMsgs", "Gets a list of messages")
	APIEndpointsGetNamespace                    = ffm("api.endpoints.getNamespace", "Gets a namespace")
	APIEndpointsGetNamespaces                   = ffm("api.endpoints.getNamespaces", "Gets a list of namespaces")
	APIEndpointsGetNetworkIdentityByDID         = ffm("api.endpoints.getNetworkIdentityByDID", "Gets an identity by its DID (deprecated - use /identities/{did} instead of /network/identities/{did})")
	APIEndpointsGetIdentityByDID                = ffm("api.endpoints.getIdentityByDID", "Gets an identity by its DID")
	APIEndpointsGetDIDDocByDID                  = ffm("api.endpoints.getDIDDocByDID", "Gets a DID document by its DID")
	APIEndpointsGetNetworkIdentities            = ffm("api.endpoints.getNetworkIdentities", "Gets the list of identities in the network (deprecated - use /identities instead of /network/identities")
	APIEndpointsGetNetworkNode                  = ffm("api.endpoints.getNetworkNode", "Gets information about a specific node in the network")
	APIEndpointsGetNetworkNodes                 = ffm("api.endpoints.getNetworkNodes", "Gets a list of nodes in the network")
	APIEndpointsGetNetworkOrg                   = ffm("api.endpoints.getNetworkOrg", "Gets information about a specific org in the network")
	APIEndpointsGetNetworkOrgs                  = ffm("api.endpoints.APIEndpointsGetNetworkOrgs", "Gets a list of orgs in the network")
	APIEndpointsGetOpByID                       = ffm("api.endpoints.getOpByID", "Gets an operation by ID")
	APIEndpointsGetOps                          = ffm("api.endpoints.getOps", "Gets a a list of operations")
	APIEndpointsGetStatusBatchManager           = ffm("api.endpoints.getStatusBatchManager", "Gets the status of the batch manager")
	APIEndpointsGetPins                         = ffm("api.endpoints.getPins", "Queries the list of pins received from the blockchain")
	APIEndpointsGetNextPins                     = ffm("api.endpoints.getNextPins", "Queries the list of next-pins that determine the next masked message sequence for each member of a privacy group, on each context/topic")
	APIEndpointsGetWebSockets                   = ffm("api.endpoints.getStatusWebSockets", "Gets a list of the current WebSocket connections to this node")
	APIEndpointsGetStatus                       = ffm("api.endpoints.getStatus", "Gets the status of this namespace")
	APIEndpointsGetSubscriptionByID             = ffm("api.endpoints.getSubscriptionByID", "Gets a subscription by its ID")
	APIEndpointsGetSubscriptions                = ffm("api.endpoints.getSubscriptions", "Gets a list of subscriptions")
	APIEndpointsGetTokenAccountPools            = ffm("api.endpoints.getTokenAccountPools", "Gets a list of token pools that contain a given token account key")
	APIEndpointsGetTokenAccounts                = ffm("api.endpoints.getTokenAccounts", "Gets a list of token accounts")
	APIEndpointsGetTokenApprovals               = ffm("api.endpoints.getTokenApprovals", "Gets a list of token approvals")
	APIEndpointsGetTokenBalances                = ffm("api.endpoints.getTokenBalances", "Gets a list of token balances")
	APIEndpointsGetTokenConnectors              = ffm("api.endpoints.getTokenConnectors", "Gets the list of token connectors currently in use")
	APIEndpointsGetTokenPoolByNameOrID          = ffm("api.endpoints.getTokenPoolByNameOrID", "Gets a token pool by its name or its ID")
	APIEndpointsGetTokenPools                   = ffm("api.endpoints.getTokenPools", "Gets a list of token pools")
	APIEndpointsGetTokenTransferByID            = ffm("api.endpoints.getTokenTransferByID", "Gets a token transfer by its ID")
	APIEndpointsGetTokenTransfers               = ffm("api.endpoints.getTokenTransfers", "Gets a list of token transfers")
	APIEndpointsGetTxnBlockchainEvents          = ffm("api.endpoints.getTxnBlockchainEvents", "Gets a list blockchain events for a specific transaction")
	APIEndpointsGetTxnByID                      = ffm("api.endpoints.getTxnByID", "Gets a transaction by its ID")
	APIEndpointsGetTxnOps                       = ffm("api.endpoints.getTxnOps", "Gets a list of operations in a specific transaction")
	APIEndpointsGetTxnStatus                    = ffm("api.endpoints.getTxnStatus", "Gets the status of a transaction")
	APIEndpointsGetTxns                         = ffm("api.endpoints.getTxns", "Gets a list of transactions")
	APIEndpointsGetVerifierByHash               = ffm("api.endpoints.getVerifierByHash", "Gets a verifier by its hash")
	APIEndpointsGetVerifiers                    = ffm("api.endpoints.getVerifiers", "Gets a list of verifiers")
	APIEndpointsPatchUpdateIdentity             = ffm("api.endpoints.patchUpdateIdentity", "Updates an identity")
	APIEndpointsPostContractDeploy              = ffm("api.endpoints.postContractDeploy", "Deploy a new smart contract")
	APIEndpointsPostContractAPIInvoke           = ffm("api.endpoints.postContractAPIInvoke", "Invokes a method on a smart contract API. Performs a blockchain transaction.")
	APIEndpointsPostContractAPIPublish          = ffm("api.endpoints.postContractAPIPublish", "Publish a contract API to all other members of the multiparty network")
	APIEndpointsPostContractAPIQuery            = ffm("api.endpoints.postContractAPIQuery", "Queries a method on a smart contract API. Performs a read-only query.")
	APIEndpointsPostContractInterfaceGenerate   = ffm("api.endpoints.postContractInterfaceGenerate", "A convenience method to convert a blockchain specific smart contract format into a FireFly Interface format. The specific blockchain plugin in use must support this functionality.")
	APIEndpointsPostContractInterfaceInvoke     = ffm("api.endpoints.postContractInterfaceInvoke", "Invokes a method on a smart contract that matches a given contract interface. Performs a blockchain transaction.")
	APIEndpointsPostContractInterfaceQuery      = ffm("api.endpoints.postContractInterfaceQuery", "Queries a method on a smart contract that matches a given contract interface. Performs a read-only query.")
	APIEndpointsPostContractInterfacePublish    = ffm("api.endpoints.postContractInterfacePublish", "Publish a contract interface to all other members of the multiparty network")
	APIEndpointsPostContractInvoke              = ffm("api.endpoints.postContractInvoke", "Invokes a method on a smart contract. Performs a blockchain transaction.")
	APIEndpointsPostContractQuery               = ffm("api.endpoints.postContractQuery", "Queries a method on a smart contract. Performs a read-only query.")
	APIEndpointsPostData                        = ffm("api.endpoints.postData", "Creates a new data item in this FireFly node")
	APIEndpointsPostDataValuePublish            = ffm("api.endpoints.postDataValuePublish", "Publishes the JSON value from the specified data resource, to shared storage")
	APIEndpointsPostDataBlobPublish             = ffm("api.endpoints.postDataBlobPublish", "Publishes the binary blob attachment stored in your local data exchange, to shared storage")
	APIEndpointsPostNewContractAPI              = ffm("api.endpoints.postNewContractAPI", "Creates and broadcasts a new custom smart contract API")
	APIEndpointsPostNewContractInterface        = ffm("api.endpoints.postNewContractInterface", "Creates and broadcasts a new custom smart contract interface")
	APIEndpointsPostNewContractListener         = ffm("api.endpoints.postNewContractListener", "Creates a new blockchain listener for events emitted by custom smart contracts")
	APIEndpointsPostNewDatatype                 = ffm("api.endpoints.postNewDatatype", "Creates and broadcasts a new datatype")
	APIEndpointsPostNewIdentity                 = ffm("api.endpoints.postNewIdentity", "Registers a new identity in the network")
	APIEndpointsPostNewMessageBroadcast         = ffm("api.endpoints.postNewMessageBroadcast", "Broadcasts a message to all members in the network")
	APIEndpointsPostNewMessagePrivate           = ffm("api.endpoints.postNewMessagePrivate", "Privately sends a message to one or more members in the network")
	APIEndpointsPostNewMessageRequestReply      = ffm("api.endpoints.postNewMessageRequestReply", "Sends a message with a blocking HTTP request, waits for a reply to that message, then sends the reply as the HTTP response.")
	APIEndpointsPostNewNamespace                = ffm("api.endpoints.postNewNamespace", "Creates and broadcasts a new namespace")
	APIEndpointsPostNodesSelf                   = ffm("api.endpoints.postNodesSelf", "Instructs this FireFly node to register itself on the network")
	APIEndpointsPostNewOrganizationSelf         = ffm("api.endpoints.postNewOrganizationSelf", "Instructs this FireFly node to register its org on the network")
	APIEndpointsPostNewOrganization             = ffm("api.endpoints.postNewOrganization", "Registers a new org in the network")
	APIEndpointsPostNewSubscription             = ffm("api.endpoints.postNewSubscription", "Creates a new subscription for an application to receive events from FireFly")
	APIEndpointsPostOpRetry                     = ffm("api.endpoints.postOpRetry", "Retries a failed operation")
	APIEndpointsPostPinsRewind                  = ffm("api.endpoints.postPinsRewind", "Force a rewind of the event aggregator to a previous position, to re-evaluate (and possibly dispatch) that pin and others after it. Only accepts a sequence or batch ID for a currently undispatched pin")
	APIEndpointsPostTokenApproval               = ffm("api.endpoints.postTokenApproval", "Creates a token approval")
	APIEndpointsPostTokenBurn                   = ffm("api.endpoints.postTokenBurn", "Burns some tokens")
	APIEndpointsPostTokenMint                   = ffm("api.endpoints.postTokenMint", "Mints some tokens")
	APIEndpointsPostTokenPool                   = ffm("api.endpoints.postTokenPool", "Creates a new token pool")
	APIEndpointsPostTokenPoolPublish            = ffm("api.endpoints.postTokenPoolPublish", "Publish a token pool to all other members of the multiparty network")
	APIEndpointsPostTokenTransfer               = ffm("api.endpoints.postTokenTransfer", "Transfers some tokens")
	APIEndpointsPutContractAPI                  = ffm("api.endpoints.putContractAPI", "Updates an existing contract API")
	APIEndpointsPutSubscription                 = ffm("api.endpoints.putSubscription", "Update an existing subscription")
	APIEndpointsGetContractAPIInterface         = ffm("api.endpoints.getContractAPIInterface", "Gets a contract interface for a contract API")
	APIEndpointsPostNetworkAction               = ffm("api.endpoints.postNetworkAction", "Notify all nodes in the network of a new governance action")
	APIEndpointsPostVerifiersResolve            = ffm("api.endpoints.postVerifiersResolve", "Resolves an input key to a signing key")

	APIFilterParamDesc         = ffm("api.filterParam", "Data filter field. Prefixes supported: > >= < <= @ ^ ! !@ !^")
	APIFilterSortDesc          = ffm("api.filterSort", "Sort field. For multi-field sort use comma separated values (or multiple query values) with '-' prefix for descending")
	APIFilterAscendingDesc     = ffm("api.filterAscending", "Ascending sort order (overrides all fields in a multi-field sort)")
	APIFilterDescendingDesc    = ffm("api.filterDescending", "Descending sort order (overrides all fields in a multi-field sort)")
	APIFilterSkipDesc          = ffm("api.filterSkip", "The number of records to skip (max: %d). Unsuitable for bulk operations")
	APIFilterLimitDesc         = ffm("api.filterLimit", "The maximum number of records to return (max: %d)")
	APIFilterCountDesc         = ffm("api.filterCount", "Return a total count as well as items (adds extra database processing)")
	APIFetchDataDesc           = ffm("api.fetchData", "Fetch the data and include it in the messages returned")
	APIConfirmQueryParam       = ffm("api.confirmQueryParam", "When true the HTTP request blocks until the message is confirmed")
	APIPublishQueryParam       = ffm("api.publishQueryParam", "When true the definition will be published to all other members of the multiparty network")
	APIHistogramStartTimeParam = ffm("api.histogramStartTime", "Start time of the data to be fetched")
	APIHistogramEndTimeParam   = ffm("api.histogramEndTime", "End time of the data to be fetched")
	APIHistogramBucketsParam   = ffm("api.histogramBuckets", "Number of buckets between start time and end time")

	APISmartContractDetails      = ffm("api.smartContractDetails", "Additional smart contract details")
	APISmartContractDetailsKey   = ffm("api.smartContractDetailsKey", "Key")
	APISmartContractDetailsValue = ffm("api.smartContractDetailsValue", "Value")
)
