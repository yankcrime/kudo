package install

import (
	"fmt"

	"github.com/thoas/go-funk"

	"github.com/kudobuilder/kudo/pkg/kudoctl/clog"
	"github.com/kudobuilder/kudo/pkg/kudoctl/packages"
	"github.com/kudobuilder/kudo/pkg/kudoctl/util/kudo"
)

func installOperatorAndOperatorVersion(client *kudo.Client, resources packages.Resources) error {
	if !client.OperatorExistsInCluster(resources.Operator.Name, resources.Operator.Namespace) {
		if _, err := client.InstallOperatorObjToCluster(resources.Operator, resources.Operator.Namespace); err != nil {
			return fmt.Errorf(
				"failed to install operator %s/%s: %v",
				resources.Operator.Namespace,
				resources.Operator.Name,
				err)
		}
		clog.Printf(
			"operator %s/%s created",
			resources.Operator.Namespace,
			resources.Operator.Name)
	}

	versionsInstalled, err := client.OperatorVersionsInstalled(resources.Operator.Name, resources.Operator.Namespace)
	if err != nil {
		return fmt.Errorf(
			"failed to retrieve existing operator versions of operator %s/%s: %v",
			resources.Operator.Namespace,
			resources.Operator.Name,
			err)
	}

	if !funk.ContainsString(versionsInstalled, resources.OperatorVersion.Spec.Version) {
		if _, err := client.InstallOperatorVersionObjToCluster(
			resources.OperatorVersion,
			resources.OperatorVersion.Namespace); err != nil {
			return fmt.Errorf(
				"failed to install operatorversion %s/%s: %v",
				resources.OperatorVersion.Namespace,
				resources.OperatorVersion.Name,
				err)
		}
		clog.Printf(
			"operatorversion %s/%s created",
			resources.OperatorVersion.Namespace,
			resources.OperatorVersion.Name)
	} else {
		clog.Printf(
			"operatorversion %s/%s already installed",
			resources.OperatorVersion.Namespace,
			resources.OperatorVersion.Name)
	}

	return nil
}
