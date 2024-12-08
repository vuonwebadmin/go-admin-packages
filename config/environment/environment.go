package environment

import (
	"github.com/vuonwebadmin/go-admin-packages/constants"
)

type Environment string

var (
	Development Environment = Environment(constants.EnvDevelopment)
	Production  Environment = Environment(constants.EnvProduction)
	Test        Environment = Environment(constants.EnvTest)
)