package utils

import (
	v1 "k8s.io/api/core/v1"
)

func IsInclude(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func RemoveItem(slice []string, target string) (result []string) {
	for _, v := range slice {
		if v == target {
			continue
		}
		result = append(result, v)
	}
	return
}

func IsContainerWaiting(status v1.ContainerStatus) bool {
	if status.State.Waiting != nil {
		return true
	}
	return false
}

func IsContainerRunning(status v1.ContainerStatus) bool {
	if status.Ready && status.State.Running != nil {
		return true
	}
	return false
}

func IsContainerSucceeded(status v1.ContainerStatus) bool {
	if !status.Ready && status.State.Terminated != nil {
		if status.State.Terminated.ExitCode == 0 && status.State.Terminated.Reason != "Error" {
			return true
		}
	}
	return false
}

func IsContainerFailed(status v1.ContainerStatus) bool {
	if status.State.Terminated != nil {
		if status.State.Terminated.Reason == "Error" {
			return true
		}
	}
	return false
}

func GetContainerStatus(pod *v1.Pod, containerName string) *v1.ContainerStatus {
	if pod == nil {
		return nil
	}

	for _, condition := range pod.Status.ContainerStatuses {
		if condition.Name == containerName {
			return &condition
		}
	}
	return nil
}

func ExtractBooleanValue(defaultValue, value *bool) bool {
	if value != nil {
		return *value
	}
	if defaultValue != nil {
		return *defaultValue
	}
	return false
}
