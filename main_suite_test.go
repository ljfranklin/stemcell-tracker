package main_test

import (
	. "github.com/cloudfoundry-incubator/stemcell-tracker/vendor/_nuts/github.com/onsi/ginkgo"
	. "github.com/cloudfoundry-incubator/stemcell-tracker/vendor/_nuts/github.com/onsi/gomega"

	"testing"
	"github.com/cloudfoundry-incubator/stemcell-tracker/vendor/_nuts/github.com/onsi/gomega/gexec"
)

var pathToBinary string

var _ = BeforeSuite(func() {
	var err error
	pathToBinary, err = gexec.Build("github.com/cloudfoundry-incubator/stemcell-tracker")
	Expect(err).ShouldNot(HaveOccurred())
})


func TestStemcellTracker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StemcellTracker Suite")
}

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})