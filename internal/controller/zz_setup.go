/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	actionssecret "github.com/coopnorge/provider-github/internal/controller/actions/actionssecret"
	providerconfig "github.com/coopnorge/provider-github/internal/controller/providerconfig"
	branch "github.com/coopnorge/provider-github/internal/controller/repo/branch"
	branchprotection "github.com/coopnorge/provider-github/internal/controller/repo/branchprotection"
	defaultbranch "github.com/coopnorge/provider-github/internal/controller/repo/defaultbranch"
	repository "github.com/coopnorge/provider-github/internal/controller/repo/repository"
	repositoryfile "github.com/coopnorge/provider-github/internal/controller/repo/repositoryfile"
	pullrequest "github.com/coopnorge/provider-github/internal/controller/repository/pullrequest"
	team "github.com/coopnorge/provider-github/internal/controller/team/team"
	teamrepository "github.com/coopnorge/provider-github/internal/controller/team/teamrepository"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		actionssecret.Setup,
		providerconfig.Setup,
		branch.Setup,
		branchprotection.Setup,
		defaultbranch.Setup,
		repository.Setup,
		repositoryfile.Setup,
		pullrequest.Setup,
		team.Setup,
		teamrepository.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
