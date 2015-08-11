package main_test

import (
	. "github.com/cloudfoundry-incubator/stemcell-tracker/vendor/_nuts/github.com/onsi/ginkgo"
	. "github.com/cloudfoundry-incubator/stemcell-tracker/vendor/_nuts/github.com/onsi/gomega"
	"net/http"
	"io/ioutil"
	"os/exec"
	"github.com/cloudfoundry-incubator/stemcell-tracker/vendor/_nuts/github.com/onsi/gomega/gexec"
	"time"
	"fmt"
	"strings"
)

var _ = Describe("StemcellTracker", func() {

	var session *gexec.Session

	BeforeEach(func() {
		command := exec.Command(pathToBinary)
		var err error
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ShouldNot(HaveOccurred())
		time.Sleep(1 * time.Second)
	})

	AfterEach(func() {
		session.Terminate().Wait()
	})

	It("PUTs then GETs the latest stemcell for the given products", func() {

		products := map[string]string{
			"cf-mysql": "3026",
			"cf-riak-cs": "3030",
		}

		for product, expectedStemcell := range products {
			url := fmt.Sprintf("http://localhost:8181/stemcell?product_name=%s", product)
			req, err := http.NewRequest("PUT", url, strings.NewReader(expectedStemcell))
			Expect(err).ToNot(HaveOccurred())

			resp, err := http.DefaultClient.Do(req)
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusCreated))

			resp, err = http.Get(url)
			Expect(err).ToNot(HaveOccurred())

			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			contents, err := ioutil.ReadAll(resp.Body)
			Expect(err).ToNot(HaveOccurred())

			Expect(string(contents)).To(Equal(expectedStemcell))
		}
	})

	It("PUTs then GETs the latest stemcell for the given products and versions", func() {
		products := map[string]map[string]string{
			"cf-mysql": map[string]string {
				"21": "3019",
				"22": "3030",
			},
			"cf-riak-cs": map[string]string {
				"10": "3008",
				"11": "3009",
			},
		}

		for product, versionMap := range products {

			//PUT all versions
			for version, expectedStemcell := range versionMap {
				url := fmt.Sprintf("http://localhost:8181/stemcell?product_name=%s&product_version=%s", product, version)
				req, err := http.NewRequest("PUT", url, strings.NewReader(expectedStemcell))
				Expect(err).ToNot(HaveOccurred())

				resp, err := http.DefaultClient.Do(req)
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusCreated))
			}

			//GET all versions
			for version, expectedStemcell := range versionMap {
				url := fmt.Sprintf("http://localhost:8181/stemcell?product_name=%s&product_version=%s", product, version)
				resp, err := http.Get(url)
				Expect(err).ToNot(HaveOccurred())

				Expect(resp.StatusCode).To(Equal(http.StatusOK))

				contents, err := ioutil.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())

				Expect(string(contents)).To(Equal(expectedStemcell))
			}
		}
	})
})