// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resourcegroup

const (
	CLUSTER_RESOURCE_GROUP_TYPE_FILE_SERVER            int32 = 100
	CLUSTER_RESOURCE_GROUP_TYPE_PRINT_SERVER           int32 = 101
	CLUSTER_RESOURCE_GROUP_TYPE_DHCP_SERVER            int32 = 102
	CLUSTER_RESOURCE_GROUP_TYPE_DTC                    int32 = 103
	CLUSTER_RESOURCE_GROUP_TYPE_MESSAGE_QUEUING        int32 = 104
	CLUSTER_RESOURCE_GROUP_TYPE_WINS_SERVER            int32 = 105
	CLUSTER_RESOURCE_GROUP_TYPE_DFS_NAMESPACE_SERVER   int32 = 106
	CLUSTER_RESOURCE_GROUP_TYPE_GENERIC_APPLICATION    int32 = 107
	CLUSTER_RESOURCE_GROUP_TYPE_GENERIC_SERVICE        int32 = 108
	CLUSTER_RESOURCE_GROUP_TYPE_GENERIC_SCRIPT         int32 = 109
	CLUSTER_RESOURCE_GROUP_TYPE_ISNS_CLUSTER_RESOURCE  int32 = 110
	CLUSTER_RESOURCE_GROUP_TYPE_VIRTUAL_MACHINE        int32 = 111
	CLUSTER_RESOURCE_GROUP_TYPE_TS_SESSION_BROKER      int32 = 112
	CLUSTER_RESOURCE_GROUP_TYPE_ISCSI_TARGET_SERVER    int32 = 113
	CLUSTER_RESOURCE_GROUP_TYPE_SCALE_OUT_FILE_SERVER  int32 = 114
	CLUSTER_RESOURCE_GROUP_TYPE_HYPER_V_REPLICA_BROKER int32 = 115
	CLUSTER_RESOURCE_GROUP_TYPE_UNKNOWN                int32 = 9999
)

const (
	CLUSTER_RESOURCE_GROUP_STATE_UNKNOWN        int32 = -1
	CLUSTER_RESOURCE_GROUP_STATE_ONLINE         int32 = 0
	CLUSTER_RESOURCE_GROUP_STATE_OFFLINE        int32 = 1
	CLUSTER_RESOURCE_GROUP_STATE_FAILED         int32 = 2
	CLUSTER_RESOURCE_GROUP_STATE_PARTIAL_ONLINE int32 = 3
	CLUSTER_RESOURCE_GROUP_STATE_PENDING        int32 = 4
)
