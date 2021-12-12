package ingress

import (
	"github.com/venturemark/permission/pkg/label"
	"github.com/venturemark/permission/pkg/label/action"
	"github.com/venturemark/permission/pkg/label/resource"
	"github.com/venturemark/permission/pkg/label/role"
	"github.com/venturemark/permission/pkg/label/visibility"
)

var (
	policy = [][]label.Label{
		{role.Any /*******/, action.Create /***/, visibility.Any /*******/, resource.User /*******/},
		{role.Any /*******/, action.Create /***/, visibility.Any /*******/, resource.Venture /****/},
		{role.Member /****/, action.Create /***/, visibility.Any /*******/, resource.Message /****/},
		{role.Member /****/, action.Create /***/, visibility.Any /*******/, resource.Update /*****/},
		{role.Member /****/, action.Create /***/, visibility.Any /*******/, resource.Timeline /***/},
		{role.Owner /*****/, action.Create /***/, visibility.Any /*******/, resource.Invite /*****/},
		{role.Owner /*****/, action.Create /***/, visibility.Any /*******/, resource.Message /****/},
		{role.Owner /*****/, action.Create /***/, visibility.Any /*******/, resource.Update /*****/},
		{role.Owner /*****/, action.Create /***/, visibility.Any /*******/, resource.Timeline /***/},

		{role.Subject /***/, action.Filter /***/, visibility.Any /*******/, resource.Any /********/},

		{role.Owner /*****/, action.Delete /***/, visibility.Any /*******/, resource.Any /********/},

		{role.Any /*******/, action.Search /***/, visibility.Public /****/, resource.Any /********/},
		{role.Member /****/, action.Search /***/, visibility.Any /*******/, resource.Any /********/},
		{role.Owner /*****/, action.Search /***/, visibility.Any /*******/, resource.Any /********/},
		{role.Reader /****/, action.Search /***/, visibility.Any /*******/, resource.Any /********/},

		{role.Owner /*****/, action.Update /***/, visibility.Any /*******/, resource.Any /********/},
	}
)
