// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"errors"
	"os"
	"testing"
)

func TestDefaultTokenCredential_ExcludeEnvCredential(t *testing.T) {
	msiEndpoint := os.Getenv("MSI_ENDPOINT")
	if len(msiEndpoint) == 0 {
		t.Skip()
	}
	cred, err := NewDefaultTokenCredential(&DefaultTokenCredentialOptions{ExcludeEnvironmentCredential: true})
	if err != nil {
		t.Fatalf("Did not expect to receive an error in creating the credential")
	}

	if len(cred.sources) != 1 {
		t.Fatalf("Length of ChainedTokenCredential sources for DefaultTokenCredential. Expected: 1, Received: %d", len(cred.sources))
	}

}

func TestDefaultTokenCredential_ExcludeMSICredential(t *testing.T) {
	err := initEnvironmentVarsForTest()
	if err != nil {
		t.Fatalf("Unexpected error when initializing environment variables: %v", err)
	}
	cred, err := NewDefaultTokenCredential(&DefaultTokenCredentialOptions{ExcludeMSICredential: true})
	if err != nil {
		t.Fatalf("Did not expect to receive an error in creating the credential")
	}
	if len(cred.sources) != 1 {
		t.Fatalf("Length of ChainedTokenCredential sources for DefaultTokenCredential. Expected: 1, Received: %d", len(cred.sources))
	}

}

func TestDefaultTokenCredential_ExcludeAllCredentials(t *testing.T) {
	err := resetEnvironmentVarsForTest()
	if err != nil {
		t.Fatalf("Unexpected error when initializing environment variables: %v", err)
	}
	var credUnavailable *CredentialUnavailableError
	_, err = NewDefaultTokenCredential(&DefaultTokenCredentialOptions{ExcludeEnvironmentCredential: true, ExcludeMSICredential: true})
	if err == nil {
		t.Fatalf("Expected an error but received nil")
	}
	if !errors.As(err, &credUnavailable) {
		t.Fatalf("Expected: CredentialUnavailableError, Received: %T", err)
	}

}

func TestDefaultTokenCredential_NilOptions(t *testing.T) {
	err := initEnvironmentVarsForTest()
	if err != nil {
		t.Fatalf("Unexpected error when initializing environment variables: %v", err)
	}
	cred, err := NewDefaultTokenCredential(&DefaultTokenCredentialOptions{ExcludeMSICredential: true})
	if err != nil {
		t.Fatalf("Did not expect to receive an error in creating the credential")
	}
	if len(cred.sources) != 1 {
		t.Fatalf("Length of ChainedTokenCredential sources for DefaultTokenCredential. Expected: 1, Received: %d", len(cred.sources))
	}
}

// func TestDefaultTokenCredential_IncludeMSICred(t *testing.T) {
// 	err := resetEnvironmentVarsForTest()
// 	if err != nil {
// 		t.Fatalf("Unable to set environment variables")
// 	}
// 	srv, close := mock.NewServer()
// 	defer close()
// 	srv.AppendResponse(mock.WithBody([]byte(appServiceTokenSuccessResp)))
// 	testURL := srv.URL()
// 	if len(os.Getenv("MSI_ENDPOINT")) == 0 {
// 		_ = os.Setenv("MSI_ENDPOINT", testURL.String())
// 		_ = os.Setenv("MSI_SECRET", "secret")
// 	}
// 	_, err = NewDefaultTokenCredential(&DefaultTokenCredentialOptions{ExcludeEnvironmentCredential: true})
// 	if err != nil {
// 		var credUnavailable *CredentialUnavailableError
// 		if !errors.As(err, &credUnavailable) {
// 			t.Fatalf("Expected a CredentialUnavailableError, received: %v", err)
// 		}
// 	}
// }
