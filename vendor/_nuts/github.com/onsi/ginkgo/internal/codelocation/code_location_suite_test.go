package codelocation_test

import (
	. "github.com/cloudfoundry-incubator/stemcell-tracker/vendor/_nuts/github.com/onsi/ginkgo"
	. "github.com/cloudfoundry-incubator/stemcell-tracker/vendor/_nuts/github.com/onsi/gomega"

	"testing"
)

func TestCodelocation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CodeLocation Suite")
}
