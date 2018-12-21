package daemon

const (
	// CurrentMachineConfigAnnotationKey is used to fetch current MachineConfig for a machine
	CurrentMachineConfigAnnotationKey = "machineconfiguration.openshift.io/currentConfig"
	// DesiredMachineConfigAnnotationKey is used to specify the desired MachineConfig for a machine
	DesiredMachineConfigAnnotationKey = "machineconfiguration.openshift.io/desiredConfig"
	// MachineConfigDaemonStateAnnotationKey is used to fetch the state of the daemon on the machine.
	MachineConfigDaemonStateAnnotationKey = "machineconfiguration.openshift.io/state"
	// MachineConfigDaemonStateWorking is set by daemon when it is applying an update.
	MachineConfigDaemonStateWorking = "Working"
	// MachineConfigDaemonStateDone is set by daemon when it is done applying an update.
	MachineConfigDaemonStateDone = "Done"
	// MachineConfigDaemonStateDegraded is set by daemon when update cannot be applied.
	MachineConfigDaemonStateDegraded = "Degraded"

	// This label was introduced in addition to state/Degraded because we
	// want operators and admins to be able to easily query this state.
	MachineConfigDaemonLabelDegraded = "machineconfiguration.openshift.io/degraded"

	// MachineConfigDaemonOSRHCOS denotes RHCOS
	MachineConfigDaemonOSRHCOS = "RHCOS"
	// MachineConfigDaemonOSRHEL denotes RHEL
	MachineConfigDaemonOSRHEL = "RHEL"
	// MachineConfigDaemonOSCENTOS denotes CENTOS
	MachineConfigDaemonOSCENTOS = "CENTOS"

	// MachineConfigMCFileType denotes when an MC config has been provided
	MachineConfigMCFileType = "MACHINECONFIG"
	// MachineConfigIgnitionFileType denotes when an Ignition config has provided
	MachineConfigIgnitionFileType = "IGNITION"

	// MachineConfigOnceFromRemoteConfig denotes that the config was pulled from a remote source
	MachineConfigOnceFromRemoteConfig = "REMOTE"
	// MachineConfigOnceFromLocalConfig denotes that the config was found locally
	MachineConfigOnceFromLocalConfig = "LOCAL"
)
