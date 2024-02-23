package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/validator"
)

func init() {
	_ = validator.RegisterRule("PageVerify",
		validator.Rules{
			"Page":     {validator.NotEmpty()},
			"PageSize": {validator.NotEmpty()},
		},
	)
	_ = validator.RegisterRule("IdVerify",
		validator.Rules{
			"Id": {validator.NotEmpty()},
		},
	)
	_ = validator.RegisterRule("AuthorityIdVerify",
		validator.Rules{
			"AuthorityId": {validator.NotEmpty()},
		},
	)
}
