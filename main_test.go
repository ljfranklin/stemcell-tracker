package main_test

import (
	. "github.com/cloudfoundry-incubator/stemcell-tracker/vendor/_nuts/github.com/onsi/ginkgo"
	. "github.com/cloudfoundry-incubator/stemcell-tracker/vendor/_nuts/github.com/onsi/gomega"
)

var _ = Describe("StemcellTracker", func() {

	It("passes fake test", func() {
		Expect(true).To(BeTrue())
	})
})