//go:build integration
// +build integration

/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package posturemanagementv1_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/posturemanagementv1"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

/**
 * This file contains an integration test for the posturemanagementv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`PostureManagementV1 Integration Tests`, func() {

	const externalConfigFile = "../posture_management_v1.env"

	var (
		scopeId                  string
		collectorId              string
		credentialId             string
		scanId                   string
		profileId                string
		credentialString         string
		err                      error
		accountId                string
		apiKey                   string
		authUrl                  string
		authenticator            core.IamAuthenticator
		postureManagementService *posturemanagementv1.PostureManagementV1
		serviceURL               string
		config                   map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(posturemanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			accountId = config["POSTURE_ACCOUNT_ID"]
			if accountId == "" {
				Skip("Unable to load posture account configuration property, skipping tests")
			}

			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			authUrl = config["IAM_APIKEY_URL"]
			if authUrl == "" {
				Skip("Unable to load auth service URL configuration property, skipping tests")
			}

			apiKey = config["IAM"]
			if apiKey == "" {
				Skip("Unable to load IAM configuration property, skipping tests")
			}

			credentialString = config["CREDENTIALS"]
			if credentialString == "" {
				Skip("Unable to load credentials configuration property, skipping tests")
			}

			authenticator = core.IamAuthenticator{
				ApiKey: apiKey,
				URL:    authUrl,
			}

			credentialId = config["CREDENTIAL_ID"]
			collectorId = config["COLLECTOR_ID"]
			scopeId = config["SCOPE_ID"]
			profileId = config["PROFILE_ID"]
			scanId = config["SCAN_ID"]

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			postureManagementServiceOptions := &posturemanagementv1.PostureManagementV1Options{
				URL:           serviceURL,
				Authenticator: &authenticator,
				AccountID:     core.StringPtr(accountId),
			}

			postureManagementService, err = posturemanagementv1.NewPostureManagementV1UsingExternalConfig(postureManagementServiceOptions)

			Expect(err).To(BeNil())
			Expect(postureManagementService).ToNot(BeNil())
			Expect(postureManagementService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`ListLatestScans - List latest scans`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLatestScans(listLatestScansOptions *ListLatestScansOptions)`, func() {

			listLatestScansOptions := &posturemanagementv1.ListLatestScansOptions{
				TransactionID: core.StringPtr(uuid.NewString()),
			}

			scansList, response, err := postureManagementService.ListLatestScans(listLatestScansOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scansList).ToNot(BeNil())

		})
	})

	Describe(`CreateValidation - Initiate a validation scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateValidation(createValidationOptions *CreateValidationOptions)`, func() {

			createValidationOptions := &posturemanagementv1.CreateValidationOptions{
				ScopeID:        core.StringPtr(scopeId),
				ProfileID:      core.StringPtr(profileId),
				GroupProfileID: core.StringPtr("0"),
				TransactionID:  core.StringPtr(uuid.NewString()),
			}

			result, response, err := postureManagementService.CreateValidation(createValidationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(result).ToNot(BeNil())

		})
	})

	Describe(`ScansSummary - Retrieve the summary of a specific scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScansSummary(scansSummaryOptions *ScansSummaryOptions)`, func() {

			scansSummaryOptions := &posturemanagementv1.ScansSummaryOptions{
				ScanID:        core.StringPtr(scanId),
				ProfileID:     core.StringPtr(profileId),
				TransactionID: core.StringPtr(uuid.NewString()),
			}

			summary, response, err := postureManagementService.ScansSummary(scansSummaryOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(summary).ToNot(BeNil())

		})
	})

	Describe(`ScanSummaries - List the validation summaries for a scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScanSummaries(scanSummariesOptions *ScanSummariesOptions)`, func() {

			scanSummariesOptions := &posturemanagementv1.ScanSummariesOptions{
				ScopeID:       core.StringPtr(scopeId),
				ProfileID:     core.StringPtr(profileId),
				TransactionID: core.StringPtr(uuid.NewString()),
			}

			summariesList, response, err := postureManagementService.ScanSummaries(scanSummariesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(summariesList).ToNot(BeNil())

		})
	})

	Describe(`ListProfiles - List profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfiles(listProfilesOptions *ListProfilesOptions)`, func() {

			listProfilesOptions := &posturemanagementv1.ListProfilesOptions{
				TransactionID: core.StringPtr(uuid.NewString()),
			}

			profilesList, response, err := postureManagementService.ListProfiles(listProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profilesList).ToNot(BeNil())

		})
	})

	Describe(`CreateScope - Create a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScope(createScopeOptions *CreateScopeOptions)`, func() {

			createScopeOptions := &posturemanagementv1.CreateScopeOptions{
				ScopeName:        core.StringPtr("SDK-IT-" + uuid.NewString()),
				ScopeDescription: core.StringPtr("SDK-IT IBM Scope Example"),
				CollectorIds:     []string{collectorId},
				CredentialID:     core.StringPtr(credentialId),
				EnvironmentType:  core.StringPtr("ibm"),
				TransactionID:    core.StringPtr(uuid.NewString()),
			}

			scope, response, err := postureManagementService.CreateScope(createScopeOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(scope).ToNot(BeNil())

		})
	})

	Describe(`ListScopes - List scopes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListScopes(listScopesOptions *ListScopesOptions)`, func() {

			listScopesOptions := &posturemanagementv1.ListScopesOptions{
				TransactionID: core.StringPtr(uuid.NewString()),
			}

			scopesList, response, err := postureManagementService.ListScopes(listScopesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopesList).ToNot(BeNil())

		})
	})

	Describe(`CreateCollector - Create a collector`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCollector(createCollectorOptions *CreateCollectorOptions)`, func() {

			createCollectorOptions := &posturemanagementv1.CreateCollectorOptions{
				CollectorName:        core.StringPtr("IBMSDK-" + uuid.NewString()),
				CollectorDescription: core.StringPtr("sample collector from SDK IT"),
				IsPublic:             core.BoolPtr(true),
				ManagedBy:            core.StringPtr("customer"),
				PassPhrase:           core.StringPtr("secret"),
				TransactionID:        core.StringPtr(uuid.NewString()),
			}

			collector, response, err := postureManagementService.CreateCollector(createCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(collector).ToNot(BeNil())

		})
	})

	Describe(`CreateCredential - Create a credential`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCredential(createCredentialOptions *CreateCredentialOptions)`, func() {

			createCredentialOptions := &posturemanagementv1.CreateCredentialOptions{
				CredentialDataFile: CreateMockReader(credentialString),
				PemFile:            CreateMockReader("test"),
				TransactionID:      core.StringPtr(uuid.NewString()),
			}

			credential, response, err := postureManagementService.CreateCredential(createCredentialOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(credential).ToNot(BeNil())

		})
	})
})

//
// Utility functions are declared in the unit test file
//
