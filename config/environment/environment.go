package environment

import (
	"github.com/vuonwebadmin/go-admin-packages/constants"
)

type Environment string

var (
	Development Environment = constants.EnvDevelopment
	Production  Environment = constants.EnvProduction
	Test        Environment = constants.EnvTest
)