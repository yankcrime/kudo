package v1beta1

import (
	"fmt"
)

func OperatorInstanceName(operatorName string) string {
	return fmt.Sprintf("%s-instance", operatorName)
}

func OperatorVersionName(operatorName, version string) string {
	return fmt.Sprintf("%s-%s", operatorName, version)
}

func (ov *OperatorVersion) FullyQualifiedName() string {
	return fmt.Sprintf("%s-%s", ov.Name, ov.Spec.AppVersion)
}

func (ov *OperatorVersion) EqualOperatorVersion(other *OperatorVersion) bool {
	return ov.FullyQualifiedName() == other.FullyQualifiedName()
}