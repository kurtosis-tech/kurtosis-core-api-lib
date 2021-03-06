package binding_constructors

import (
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/kurtosis_core_rpc_api_bindings"
)

// The generated bindings don't come with constructors (leaving it up to the user to initialize all the fields), so we
// add them so that our code is safer

// ==============================================================================================
//                           Shared Objects (Used By Multiple Endpoints)
// ==============================================================================================
func NewPort(number uint32, protocol kurtosis_core_rpc_api_bindings.Port_Protocol) *kurtosis_core_rpc_api_bindings.Port {
	return &kurtosis_core_rpc_api_bindings.Port{
		Number:   number,
		Protocol: protocol,
	}
}

// ==============================================================================================
//                                     Load Module
// ==============================================================================================
func NewLoadModuleArgs(moduleId string, containerImage string, serializedParams string) *kurtosis_core_rpc_api_bindings.LoadModuleArgs {
	return &kurtosis_core_rpc_api_bindings.LoadModuleArgs{
		ModuleId:         moduleId,
		ContainerImage:   containerImage,
		SerializedParams: serializedParams,
	}
}

func NewLoadModuleResponse(
	guid string,
	privateIpAddr string,
	privatePort *kurtosis_core_rpc_api_bindings.Port,
	publicIpAddr string,
	publicPort *kurtosis_core_rpc_api_bindings.Port,
) *kurtosis_core_rpc_api_bindings.LoadModuleResponse {
	return &kurtosis_core_rpc_api_bindings.LoadModuleResponse{
		Guid:          guid,
		PrivateIpAddr: privateIpAddr,
		PrivatePort:   privatePort,
		PublicIpAddr:  publicIpAddr,
		PublicPort:    publicPort,
	}
}

// ==============================================================================================
//                                     Unload Module
// ==============================================================================================
func NewUnloadModuleArgs(moduleId string) *kurtosis_core_rpc_api_bindings.UnloadModuleArgs {
	return &kurtosis_core_rpc_api_bindings.UnloadModuleArgs{
		ModuleId: moduleId,
	}
}

func NewUnloadModuleResponse(moduleGuid string) *kurtosis_core_rpc_api_bindings.UnloadModuleResponse {
	return &kurtosis_core_rpc_api_bindings.UnloadModuleResponse{
		ModuleGuid: moduleGuid,
	}
}

// ==============================================================================================
//                                     Execute Module
// ==============================================================================================
func NewExecuteModuleArgs(moduleId string, serializedParams string) *kurtosis_core_rpc_api_bindings.ExecuteModuleArgs {
	return &kurtosis_core_rpc_api_bindings.ExecuteModuleArgs{
		ModuleId:         moduleId,
		SerializedParams: serializedParams,
	}
}

func NewExecuteModuleResponse(serializedResult string) *kurtosis_core_rpc_api_bindings.ExecuteModuleResponse {
	return &kurtosis_core_rpc_api_bindings.ExecuteModuleResponse{
		SerializedResult: serializedResult,
	}
}

// ==============================================================================================
//                                     Get Module Info
// ==============================================================================================
func NewGetModulesArgs(moduleIds map[string]bool) *kurtosis_core_rpc_api_bindings.GetModulesArgs {
	return &kurtosis_core_rpc_api_bindings.GetModulesArgs{
		Ids: moduleIds,
	}
}

func NewGetModulesResponse(
	moduleInfoMap map[string]*kurtosis_core_rpc_api_bindings.ModuleInfo,
) *kurtosis_core_rpc_api_bindings.GetModulesResponse {
	return &kurtosis_core_rpc_api_bindings.GetModulesResponse{
		ModuleInfo: moduleInfoMap,
	}
}

func NewModuleInfo(
	guid string,
	privateIpAddr string,
	privateGrpcPort *kurtosis_core_rpc_api_bindings.Port,
	maybePublicIpAddr string,
	maybePublicGrpcPort *kurtosis_core_rpc_api_bindings.Port,
) *kurtosis_core_rpc_api_bindings.ModuleInfo {
	return &kurtosis_core_rpc_api_bindings.ModuleInfo{
		Guid:                guid,
		PrivateIpAddr:       privateIpAddr,
		PrivateGrpcPort:     privateGrpcPort,
		MaybePublicIpAddr:   maybePublicIpAddr,
		MaybePublicGrpcPort: maybePublicGrpcPort,
	}
}

// ==============================================================================================
//                                     Register Service
// ==============================================================================================
func NewRegisterServiceArgs(serviceId string, partitionId string) *kurtosis_core_rpc_api_bindings.RegisterServiceArgs {
	return &kurtosis_core_rpc_api_bindings.RegisterServiceArgs{
		ServiceId:   serviceId,
		PartitionId: partitionId,
	}
}

func NewRegisterServiceResponse(privateIpAddr string) *kurtosis_core_rpc_api_bindings.RegisterServiceResponse {
	return &kurtosis_core_rpc_api_bindings.RegisterServiceResponse{
		PrivateIpAddr: privateIpAddr,
	}
}

// ==============================================================================================
//                                        Start Service
// ==============================================================================================
func NewStartServiceArgs(
	serviceId string,
	image string,
	privatePorts map[string]*kurtosis_core_rpc_api_bindings.Port,
	publicPorts map[string]*kurtosis_core_rpc_api_bindings.Port, //TODO this is a huge hack to temporarily enable static ports for NEAR until we have a more productized solution
	entrypointArgs []string,
	cmdArgs []string,
	envVars map[string]string,
	filesArtifactMountDirpaths map[string]string,
	cpuAllocationMillicpus uint64,
	memoryAllocationMegabytes uint64,
) *kurtosis_core_rpc_api_bindings.StartServiceArgs {
	return &kurtosis_core_rpc_api_bindings.StartServiceArgs{
		ServiceId:                serviceId,
		DockerImage:              image,
		PrivatePorts:             privatePorts,
		EntrypointArgs:           entrypointArgs,
		CmdArgs:                  cmdArgs,
		DockerEnvVars:            envVars,
		FilesArtifactMountpoints: filesArtifactMountDirpaths,
		PublicPorts:              publicPorts, //TODO this is a huge hack to temporarily enable static ports for NEAR until we have a more productized solution
		CpuAllocationMillicpus:            cpuAllocationMillicpus,
		MemoryAllocationMegabytes:         memoryAllocationMegabytes,
	}
}

func NewStartServiceResponse(privateIpAddr string, privatePorts map[string]*kurtosis_core_rpc_api_bindings.Port, publicIpAddr string, publicPorts map[string]*kurtosis_core_rpc_api_bindings.Port, serviceGuid string) *kurtosis_core_rpc_api_bindings.StartServiceResponse {
	return &kurtosis_core_rpc_api_bindings.StartServiceResponse{
		ServiceInfo: NewServiceInfo(serviceGuid, privateIpAddr, privatePorts, publicIpAddr, publicPorts),
	}
}

// ==============================================================================================
//                                       Get Service Info
// ==============================================================================================
func NewGetServicesArgs(serviceIds map[string]bool) *kurtosis_core_rpc_api_bindings.GetServicesArgs {
	return &kurtosis_core_rpc_api_bindings.GetServicesArgs{
		ServiceIds: serviceIds,
	}
}

func NewGetServicesResponse(
	serviceInfo map[string]*kurtosis_core_rpc_api_bindings.ServiceInfo,
) *kurtosis_core_rpc_api_bindings.GetServicesResponse {
	return &kurtosis_core_rpc_api_bindings.GetServicesResponse{
		ServiceInfo: serviceInfo,
	}
}

func NewServiceInfo(
	guid string,
	privateIpAddr string,
	privatePorts map[string]*kurtosis_core_rpc_api_bindings.Port,
	maybePublicIpAddr string,
	maybePublicPorts map[string]*kurtosis_core_rpc_api_bindings.Port,
) *kurtosis_core_rpc_api_bindings.ServiceInfo {
	return &kurtosis_core_rpc_api_bindings.ServiceInfo{
		ServiceGuid:       guid,
		PrivateIpAddr:     privateIpAddr,
		PrivatePorts:      privatePorts,
		MaybePublicIpAddr: maybePublicIpAddr,
		MaybePublicPorts:  maybePublicPorts,
	}
}

// ==============================================================================================
//                                        Remove Service
// ==============================================================================================
func NewRemoveServiceArgs(serviceId string, containerStopTimeoutSeconds uint64) *kurtosis_core_rpc_api_bindings.RemoveServiceArgs {
	return &kurtosis_core_rpc_api_bindings.RemoveServiceArgs{
		ServiceId:                   serviceId,
		ContainerStopTimeoutSeconds: containerStopTimeoutSeconds,
	}
}

func NewRemoveServiceResponse(serviceGuid string) *kurtosis_core_rpc_api_bindings.RemoveServiceResponse {
	return &kurtosis_core_rpc_api_bindings.RemoveServiceResponse{
		ServiceGuid: serviceGuid,
	}
}

// ==============================================================================================
//                                          Repartition
// ==============================================================================================
func NewRepartitionArgs(
	partitionServices map[string]*kurtosis_core_rpc_api_bindings.PartitionServices,
	partitionConnections map[string]*kurtosis_core_rpc_api_bindings.PartitionConnections,
	defaultConnection *kurtosis_core_rpc_api_bindings.PartitionConnectionInfo) *kurtosis_core_rpc_api_bindings.RepartitionArgs {
	return &kurtosis_core_rpc_api_bindings.RepartitionArgs{
		PartitionServices:    partitionServices,
		PartitionConnections: partitionConnections,
		DefaultConnection:    defaultConnection,
	}
}

func NewPartitionServices(serviceIdSet map[string]bool) *kurtosis_core_rpc_api_bindings.PartitionServices {
	return &kurtosis_core_rpc_api_bindings.PartitionServices{
		ServiceIdSet: serviceIdSet,
	}
}

func NewPartitionConnections(connectionInfo map[string]*kurtosis_core_rpc_api_bindings.PartitionConnectionInfo) *kurtosis_core_rpc_api_bindings.PartitionConnections {
	return &kurtosis_core_rpc_api_bindings.PartitionConnections{
		ConnectionInfo: connectionInfo,
	}
}

func NewPartitionConnectionInfo(packetLossPercentage float32) *kurtosis_core_rpc_api_bindings.PartitionConnectionInfo {
	return &kurtosis_core_rpc_api_bindings.PartitionConnectionInfo{
		PacketLossPercentage: packetLossPercentage,
	}
}

// ==============================================================================================
//                                          Pause/Unpause Service
// ==============================================================================================

func NewPauseServiceArgs(serviceId string) *kurtosis_core_rpc_api_bindings.PauseServiceArgs {
	return &kurtosis_core_rpc_api_bindings.PauseServiceArgs{
		ServiceId: serviceId,
	}
}

func NewUnpauseServiceArgs(serviceId string) *kurtosis_core_rpc_api_bindings.UnpauseServiceArgs {
	return &kurtosis_core_rpc_api_bindings.UnpauseServiceArgs{
		ServiceId: serviceId,
	}
}

// ==============================================================================================
//                                          Exec Command
// ==============================================================================================
func NewExecCommandArgs(serviceId string, commandArgs []string) *kurtosis_core_rpc_api_bindings.ExecCommandArgs {
	return &kurtosis_core_rpc_api_bindings.ExecCommandArgs{
		ServiceId:   serviceId,
		CommandArgs: commandArgs,
	}
}

func NewExecCommandResponse(exitCode int32, logOutput string) *kurtosis_core_rpc_api_bindings.ExecCommandResponse {
	return &kurtosis_core_rpc_api_bindings.ExecCommandResponse{
		ExitCode:  exitCode,
		LogOutput: logOutput,
	}
}

// ==============================================================================================
//                           Wait For Http Get Endpoint Availability
// ==============================================================================================
func NewWaitForHttpGetEndpointAvailabilityArgs(
	serviceId string,
	port uint32,
	path string,
	initialDelayMilliseconds uint32,
	retries uint32,
	retriesDelayMilliseconds uint32,
	bodyText string) *kurtosis_core_rpc_api_bindings.WaitForHttpGetEndpointAvailabilityArgs {
	return &kurtosis_core_rpc_api_bindings.WaitForHttpGetEndpointAvailabilityArgs{
		ServiceId:                serviceId,
		Port:                     port,
		Path:                     path,
		InitialDelayMilliseconds: initialDelayMilliseconds,
		Retries:                  retries,
		RetriesDelayMilliseconds: retriesDelayMilliseconds,
		BodyText:                 bodyText,
	}
}

// ==============================================================================================
//                           Wait For Http Post Endpoint Availability
// ==============================================================================================
func NewWaitForHttpPostEndpointAvailabilityArgs(
	serviceId string,
	port uint32,
	path string,
	requestBody string,
	initialDelayMilliseconds uint32,
	retries uint32,
	retriesDelayMilliseconds uint32,
	bodyText string) *kurtosis_core_rpc_api_bindings.WaitForHttpPostEndpointAvailabilityArgs {
	return &kurtosis_core_rpc_api_bindings.WaitForHttpPostEndpointAvailabilityArgs{
		ServiceId:                serviceId,
		Port:                     port,
		Path:                     path,
		RequestBody:              requestBody,
		InitialDelayMilliseconds: initialDelayMilliseconds,
		Retries:                  retries,
		RetriesDelayMilliseconds: retriesDelayMilliseconds,
		BodyText:                 bodyText,
	}
}

// ==============================================================================================
//                                  Upload Files Artifact
// ==============================================================================================
func NewUploadFilesArtifactArgs(data []byte) *kurtosis_core_rpc_api_bindings.UploadFilesArtifactArgs {
	return &kurtosis_core_rpc_api_bindings.UploadFilesArtifactArgs{Data: data}
}

// ==============================================================================================
//                                 Store Web Files Artifact
// ==============================================================================================
func NewStoreWebFilesArtifactArgs(url string) *kurtosis_core_rpc_api_bindings.StoreWebFilesArtifactArgs {
	return &kurtosis_core_rpc_api_bindings.StoreWebFilesArtifactArgs{Url: url}
}

// ==============================================================================================
//                              Store Files Artifact From Service
// ==============================================================================================
func NewStoreFilesArtifactFromServiceArgs(serviceId string, sourcePath string) *kurtosis_core_rpc_api_bindings.StoreFilesArtifactFromServiceArgs {
	return &kurtosis_core_rpc_api_bindings.StoreFilesArtifactFromServiceArgs{ServiceId: serviceId, SourcePath: sourcePath}
}
