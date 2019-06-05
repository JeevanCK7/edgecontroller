// Copyright 2019 Smart-Edge.com, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	cce "github.com/smartedgemec/controller-ce"
	"github.com/smartedgemec/controller-ce/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("/apps", func() {
	Describe("POST /apps", func() {
		DescribeTable("201 Created",
			func(req string) {
				By("Sending a POST /apps request")
				resp, err := apiCli.Post(
					"http://127.0.0.1:8080/apps",
					"application/json",
					strings.NewReader(req))
				Expect(err).ToNot(HaveOccurred())
				defer resp.Body.Close()

				By("Verifying a 201 response")
				Expect(resp.StatusCode).To(Equal(http.StatusCreated))

				By("Reading the response body")
				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())

				var respBody struct {
					ID string
				}

				By("Unmarshalling the response")
				Expect(json.Unmarshal(body, &respBody)).To(Succeed())

				By("Verifying a UUID was returned")
				Expect(uuid.IsValid(respBody.ID)).To(BeTrue())
			},
			Entry(
				"POST /apps",
				`
				{
					"name": "container app",
					"version": "latest",
					"type": "container",
					"vendor": "smart edge",
					"description": "my container app",
					"cores": 4,
					"memory": 1024,
					"ports": [{"port": 80, "protocol": "tcp"}],
					"source": "http://www.test.com/my_container_app.tar.gz"
				}`),
			Entry(
				"POST /apps without description",
				`
				{
					"name": "container app",
					"version": "latest",
					"type": "container",
					"vendor": "smart edge",		
					"cores": 4,
					"memory": 1024,
					"ports": [{"port": 80, "protocol": "tcp"}],
					"source": "http://www.test.com/my_container_app.tar.gz"
				}`),
		)

		DescribeTable("400 Bad Request",
			func(req, expectedResp string) {
				By("Sending a POST /apps request")
				resp, err := apiCli.Post(
					"http://127.0.0.1:8080/apps",
					"application/json",
					strings.NewReader(req))
				Expect(err).ToNot(HaveOccurred())
				defer resp.Body.Close()

				By("Verifying a 400 Bad Request response")
				Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))

				By("Reading the response body")
				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the response body")
				Expect(string(body)).To(Equal(expectedResp))
			},
			Entry(
				"POST /apps with id",
				`
				{
					"id": "123"
				}`,
				"Validation failed: id cannot be specified in POST request"),
			Entry(
				"POST /apps without type",
				`
				{
					"name": "container app",
					"version": "latest",
					"vendor": "smart edge",
					"description": "my container app",
					"cores": 4,
					"memory": 1024,
					"ports": [{"port": 80, "protocol": "tcp"}],
					"source": "http://www.test.com/my_container_app.tar.gz"
				}`,
				`Validation failed: type must be either "container" or "vm"`),
			Entry(
				"POST /apps without name",
				`
				{
					"type": "container",
					"vendor": "smart edge",
					"version": "latest",
					"description": "my container app",
					"cores": 4,
					"memory": 1024,
					"ports": [{"port": 80, "protocol": "tcp"}],
					"source": "http://www.test.com/my_container_app.tar.gz"
				}`,
				"Validation failed: name cannot be empty"),
			Entry(
				"POST /apps without version",
				`
					{
						"type": "container",
						"name": "container app",
						"vendor": "smart edge",
						"description": "my container app",
						"cores": 4,
						"memory": 1024,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"source": "http://www.test.com/my_container_app.tar.gz"
					}`,
				"Validation failed: version cannot be empty"),
			Entry(
				"POST /apps without vendor",
				`
				{
					"type": "container",
					"name": "container app",
					"version": "latest",
					"description": "my container app",
					"cores": 4,
					"memory": 1024,
					"ports": [{"port": 80, "protocol": "tcp"}],
					"source": "http://www.test.com/my_container_app.tar.gz"
				}`,
				"Validation failed: vendor cannot be empty"),
			Entry("POST /apps with cores not in [1..8]",
				`
				{
					"type": "container",
					"name": "container app",
					"version": "latest",
					"vendor": "smart edge",
					"description": "my container app",
					"cores": 9,
					"memory": 1024,
					"ports": [{"port": 80, "protocol": "tcp"}],
					"source": "http://www.test.com/my_container_app.tar.gz"
				}`,
				"Validation failed: cores must be in [1..8]"),
			Entry("POST /apps with memory not in [1..16384]",
				`
				{
					"type": "container",
					"name": "container app",
					"version": "latest",
					"vendor": "smart edge",
					"description": "my container app",
					"cores": 8,
					"memory": 16385,
					"ports": [{"port": 80, "protocol": "tcp"}],
					"source": "http://www.test.com/my_container_app.tar.gz"
				}`,
				"Validation failed: memory must be in [1..16384]"),
			Entry("POST /apps with ports not in [1..65535]",
				`
				{
					"type": "container",
					"name": "container app",
					"version": "latest",
					"vendor": "smart edge",
					"description": "my container app",
					"cores": 8,
					"memory": 1024,
					"ports": [{"port": 99999, "protocol": "tcp"}],
					"source": "http://www.test.com/my_container_app.tar.gz"
				}`,
				"Validation failed: port must be in [1..65535]"),
			Entry("POST /apps with protocol not tcp, udp, sctp or icmp",
				`
				{
					"type": "container",
					"name": "container app",
					"version": "latest",
					"vendor": "smart edge",
					"description": "my container app",
					"cores": 8,
					"memory": 1024,
					"ports": [{"port": 80, "protocol": "thisisnotaprotocol"}],
					"source": "http://www.test.com/my_container_app.tar.gz"
				}`,
				"Validation failed: protocol must be tcp, udp, sctp or icmp"),
			Entry(
				"POST /apps without source",
				`
						{
							"type": "container",
							"name": "container app",
							"version": "latest",
							"vendor": "smart edge",
							"description": "my container app",
							"cores": 4,
							"ports": [{"port": 80, "protocol": "tcp"}],
							"memory": 1024
						}`,
				"Validation failed: source cannot be empty"),
			Entry(
				"POST /apps without source",
				`
							{
								"type": "container",
								"name": "container app",
								"version": "latest",
								"vendor": "smart edge",
								"description": "my container app",
								"cores": 4,
								"memory": 1024,
								"ports": [{"port": 80, "protocol": "tcp"}],
								"source": "invalid.url"
							}`,
				"Validation failed: source cannot be parsed as a URI"),
		)
	})

	Describe("GET /apps", func() {
		var (
			containerAppID string
			vmAppID        string
		)

		BeforeEach(func() {
			containerAppID = postApps("container")
			vmAppID = postApps("vm")
		})

		DescribeTable("200 OK",
			func() {
				By("Sending a GET /apps request")
				resp, err := apiCli.Get("http://127.0.0.1:8080/apps")
				Expect(err).ToNot(HaveOccurred())
				defer resp.Body.Close()

				By("Verifying a 200 OK response")
				Expect(resp.StatusCode).To(Equal(http.StatusOK))

				By("Reading the response body")
				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())

				var apps []*cce.App

				By("Unmarshalling the response")
				Expect(json.Unmarshal(body, &apps)).To(Succeed())

				By("Verifying the 2 created apps were returned")
				Expect(apps).To(ContainElement(
					&cce.App{
						ID:          containerAppID,
						Type:        "container",
						Name:        "container app",
						Version:     "latest",
						Vendor:      "smart edge",
						Description: "my container app",
						Cores:       4,
						Memory:      1024,
						Ports:       []cce.PortProto{{Port: 80, Protocol: "tcp"}},
						Source:      "http://www.test.com/my_container_app.tar.gz",
					}))
				Expect(apps).To(ContainElement(
					&cce.App{
						ID:          vmAppID,
						Type:        "vm",
						Name:        "vm app",
						Version:     "latest",
						Vendor:      "smart edge",
						Description: "my vm app",
						Cores:       4,
						Memory:      1024,
						Ports:       []cce.PortProto{{Port: 80, Protocol: "tcp"}},
						Source:      "http://www.test.com/my_vm_app.tar.gz",
					}))
			},
			Entry("GET /apps"),
		)
	})

	Describe("GET /apps/{id}", func() {
		var (
			containerAppID string
		)

		BeforeEach(func() {
			containerAppID = postApps("container")
		})

		DescribeTable("200 OK",
			func() {
				app := getApp(containerAppID)

				By("Verifying the created app was returned")
				Expect(app).To(Equal(
					&cce.App{
						ID:          containerAppID,
						Type:        "container",
						Name:        "container app",
						Version:     "latest",
						Vendor:      "smart edge",
						Description: "my container app",
						Cores:       4,
						Memory:      1024,
						Ports:       []cce.PortProto{{Port: 80, Protocol: "tcp"}},
						Source:      "http://www.test.com/my_container_app.tar.gz",
					},
				))
			},
			Entry("GET /apps/{id}"),
		)

		DescribeTable("404 Not Found",
			func() {
				By("Sending a GET /apps/{id} request")
				resp, err := apiCli.Get(
					fmt.Sprintf("http://127.0.0.1:8080/apps/%s",
						uuid.New()))
				Expect(err).ToNot(HaveOccurred())
				defer resp.Body.Close()

				By("Verifying a 404 Not Found response")
				Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
			},
			Entry("GET /apps/{id} with nonexistent ID"),
		)
	})

	Describe("PATCH /apps", func() {
		var (
			containerAppID string
		)

		BeforeEach(func() {
			containerAppID = postApps("container")
		})

		DescribeTable("204 No Content",
			func(reqStr string, expectedApp *cce.App) {
				By("Sending a PATCH /apps request")
				resp, err := apiCli.Patch(
					"http://127.0.0.1:8080/apps",
					"application/json",
					strings.NewReader(fmt.Sprintf(reqStr, containerAppID)))
				Expect(err).ToNot(HaveOccurred())
				defer resp.Body.Close()

				By("Verifying a 204 No Content response")
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusNoContent))

				By("Getting the updated app")
				updatedApp := getApp(containerAppID)

				By("Verifying the app was updated")
				expectedApp.SetID(containerAppID)
				Expect(updatedApp).To(Equal(expectedApp))
			},
			Entry(
				"PATCH /apps",
				`
				[
					{
						"id": "%s",
						"type": "container",
						"name": "container app2",
						"version": "latest",
						"vendor": "smart edge",
						"description": "my container app",	
						"cores": 4,
						"memory": 1024,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"source": "http://www.test.com/my_container_app.tar.gz"
					}
				]`,
				&cce.App{
					Type:        "container",
					Name:        "container app2",
					Version:     "latest",
					Vendor:      "smart edge",
					Description: "my container app",
					Cores:       4,
					Memory:      1024,
					Ports:       []cce.PortProto{{Port: 80, Protocol: "tcp"}},
					Source:      "http://www.test.com/my_container_app.tar.gz",
				}),
			Entry("PATCH /apps with no description",
				`
				[
					{
						"id": "%s",
						"type": "container",
						"name": "container app2",
						"version": "latest",
						"vendor": "smart edge",
						"cores": 4,
						"memory": 1024,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"source": "http://www.test.com/my_container_app.tar.gz"
					}
				]`,
				&cce.App{
					Name:        "container app2",
					Type:        "container",
					Vendor:      "smart edge",
					Description: "",
					Version:     "latest",
					Cores:       4,
					Memory:      1024,
					Ports:       []cce.PortProto{{Port: 80, Protocol: "tcp"}},
					Source:      "http://www.test.com/my_container_app.tar.gz",
				}),
		)

		DescribeTable("400 Bad Request",
			func(reqStr string, expectedResp string) {
				By("Sending a PATCH /apps request")
				if strings.Contains(reqStr, "%s") {
					reqStr = fmt.Sprintf(reqStr, containerAppID)
				}
				resp, err := apiCli.Patch(
					"http://127.0.0.1:8080/apps",
					"application/json",
					strings.NewReader(reqStr))
				Expect(err).ToNot(HaveOccurred())
				defer resp.Body.Close()

				By("Verifying a 400 Bad Request")
				Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))

				By("Reading the response body")
				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the response body")
				Expect(string(body)).To(Equal(expectedResp))
			},
			Entry(
				"PATCH /apps without id",
				`
				[
					{
						"type": "container",
						"name": "container app2",
						"version": "latest",
						"vendor": "smart edge",
						"description": "my container app",
						"cores": 4,
						"memory": 1024,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"source": "http://www.test.com/my_container_app.tar.gz"
					}
				]`,
				"Validation failed: id not a valid uuid"),
			Entry(
				"PATCH /apps without type",
				`
				[
					{
						"id": "%s",
						"name": "container app2",
						"version": "latest",
						"vendor": "smart edge",
						"description": "my container app",	
						"cores": 4,
						"memory": 1024,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"source": "http://www.test.com/my_container_app.tar.gz"
					}
				]`,
				`Validation failed: type must be either "container" or "vm"`),
			Entry(
				"PATCH /apps without name",
				`
				[
					{
						"id": "%s",
						"type": "container",
						"version": "latest",
						"vendor": "smart edge",
						"description": "my container app",
						"cores": 4,
						"memory": 1024,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"source": "http://www.test.com/my_container_app.tar.gz"
					}
				]`,
				"Validation failed: name cannot be empty"),
			Entry("PATCH /apps without version",
				`
				[
					{
						"id": "%s",
						"type": "container",
						"name": "container app2",
						"vendor": "smart edge",
						"description": "my container app",
						"cores": 4,
						"memory": 1024,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"source": "http://www.test.com/my_container_app.tar.gz"
					}
				]`,
				"Validation failed: version cannot be empty"),

			Entry("PATCH /apps without vendor",
				`
				[
					{
						"id": "%s",
						"type": "container",
						"name": "container app2",
						"version": "latest",
						"description": "my container app",
						"cores": 4,
						"memory": 1024,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"source": "http://www.test.com/my_container_app.tar.gz"
					}
				]`,
				"Validation failed: vendor cannot be empty"),
			Entry("PATCH /apps with cores not in [1..8]",
				`
				[
					{
						"id": "%s",
						"type": "container",
						"name": "container app2",
						"version": "latest",
						"vendor": "smart edge",
						"description": "my container app",
						"cores": 9,
						"memory": 1024,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"source": "http://www.test.com/my_container_app.tar.gz"
					}
				]`,
				"Validation failed: cores must be in [1..8]"),
			Entry("PATCH /apps with memory not in [1..16384]",
				`
				[
					{
						"id": "%s",
						"type": "container",
						"name": "container app2",
						"version": "latest",
						"vendor": "smart edge",
						"description": "my container app",
						"cores": 4,
						"memory": 16385,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"source": "http://www.test.com/my_container_app.tar.gz"
					}
				]`,
				"Validation failed: memory must be in [1..16384]"),
			Entry("PATCH /apps without source",
				`
				[
					{
						"id": "%s",
						"type": "container",
						"name": "container app2",
						"version": "latest",
						"vendor": "smart edge",
						"description": "my container app",
						"cores": 4,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"memory": 1024
					}
				]`,
				"Validation failed: source cannot be empty"),
			Entry("PATCH /apps without source",
				`
				[
					{
						"id": "%s",
						"type": "container",
						"name": "container app2",
						"version": "latest",
						"vendor": "smart edge",
						"description": "my container app",
						"cores": 4,
						"memory": 1024,
						"ports": [{"port": 80, "protocol": "tcp"}],
						"source": "invalid.url"
					}
				]`,
				"Validation failed: source cannot be parsed as a URI"),
		)
	})

	Describe("DELETE /apps/{id}", func() {
		var (
			containerAppID string
		)

		BeforeEach(func() {
			containerAppID = postApps("container")
		})

		DescribeTable("200 OK",
			func() {
				By("Sending a DELETE /apps/{id} request")
				resp, err := apiCli.Delete(
					fmt.Sprintf("http://127.0.0.1:8080/apps/%s",
						containerAppID))
				Expect(err).ToNot(HaveOccurred())
				defer resp.Body.Close()

				By("Verifying a 200 OK response")
				Expect(resp.StatusCode).To(Equal(http.StatusOK))

				By("Verifying the app was deleted")

				By("Sending a GET /apps/{id} request")
				resp, err = apiCli.Get(
					fmt.Sprintf("http://127.0.0.1:8080/apps/%s",
						containerAppID))
				Expect(err).ToNot(HaveOccurred())
				defer resp.Body.Close()

				By("Verifying a 404 Not Found response")
				Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
			},
			Entry("DELETE /apps/{id}"),
		)

		DescribeTable("404 Not Found",
			func(id string) {
				By("Sending a DELETE /apps/{id} request")
				resp, err := apiCli.Delete(
					fmt.Sprintf("http://127.0.0.1:8080/apps/%s", id))
				Expect(err).ToNot(HaveOccurred())
				defer resp.Body.Close()

				By("Verifying a 404 Not Found response")
				Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
			},
			Entry(
				"DELETE /apps/{id} with nonexistent ID",
				uuid.New()),
		)

		DescribeTable("422 Unprocessable Entity",
			func(resource, expectedResp string) {
				switch resource {
				case "dns_configs_app_aliases":
					postDNSConfigsAppAliases(
						postDNSConfigs(),
						containerAppID)
				case "nodes_apps":
					clearGRPCTargetsTable()
					nodeCfg := createAndRegisterNode()
					containerAppID = postApps("container")
					postNodesApps(
						nodeCfg.nodeID,
						containerAppID)
				}

				By("Sending a DELETE /apps/{id} request")
				resp, err := apiCli.Delete(
					fmt.Sprintf("http://127.0.0.1:8080/apps/%s",
						containerAppID))
				Expect(err).ToNot(HaveOccurred())
				defer resp.Body.Close()

				By("Verifying a 422 response")
				Expect(resp.StatusCode).To(Equal(
					http.StatusUnprocessableEntity))

				By("Reading the response body")
				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the response body")
				Expect(string(body)).To(Equal(
					fmt.Sprintf(expectedResp, containerAppID)))
			},
			Entry(
				"DELETE /apps/{id} with dns_configs_app_aliases record",
				"dns_configs_app_aliases",
				"cannot delete app_id %s: record in use in dns_configs_app_aliases",
			),
			Entry(
				"DELETE /apps/{id} with nodes_apps record",
				"nodes_apps",
				"cannot delete app_id %s: record in use in nodes_apps",
			),
		)
	})
})
