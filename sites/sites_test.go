package sites

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"

	"github.com/appwrite/sdk-for-go/v3/client"
	"github.com/appwrite/sdk-for-go/v3/file")

func TestSites(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	_, _ = tmpFile.WriteString("test content")
	_ = tmpFile.Close()
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test List", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "sites": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Site",
            "enabled": true,
            "live": true,
            "logging": true,
            "framework": "react",
            "deploymentRetention": 7,
            "deploymentId": "5e5ea5c16897e",
            "deploymentCreatedAt": "2020-10-15T06:38:00.000+00:00",
            "deploymentScreenshotLight": "5e5ea5c16897e",
            "deploymentScreenshotDark": "5e5ea5c16897e",
            "latestDeploymentId": "5e5ea5c16897e",
            "latestDeploymentCreatedAt": "2020-10-15T06:38:00.000+00:00",
            "latestDeploymentStatus": "ready",
            "vars": [
                {
                    "$id": "5e5ea5c16897e",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "key": "API_KEY",
                    "value": "myPa$$word1",
                    "secret": true,
                    "resourceType": "function",
                    "resourceId": "myAwesomeFunction"
                }
            ],
            "timeout": 300,
            "installCommand": "npm install",
            "buildCommand": "npm run build",
            "startCommand": "node custom-server.mjs",
            "outputDirectory": "build",
            "installationId": "6m40at4ejk5h2u9s1hboo",
            "providerRepositoryId": "appwrite",
            "providerBranch": "main",
            "providerRootDirectory": "sites/helloWorld",
            "providerSilentMode": true,
            "buildSpecification": "s-1vcpu-512mb",
            "runtimeSpecification": "s-1vcpu-512mb",
            "buildRuntime": "node-22",
            "adapter": "static",
            "fallbackFile": "index.html"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.List()
		if err != nil {
			t.Errorf("Method List failed: %v", err)
		}
	})

	t.Run("Test Create", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Site",
    "enabled": true,
    "live": true,
    "logging": true,
    "framework": "react",
    "deploymentRetention": 7,
    "deploymentId": "5e5ea5c16897e",
    "deploymentCreatedAt": "2020-10-15T06:38:00.000+00:00",
    "deploymentScreenshotLight": "5e5ea5c16897e",
    "deploymentScreenshotDark": "5e5ea5c16897e",
    "latestDeploymentId": "5e5ea5c16897e",
    "latestDeploymentCreatedAt": "2020-10-15T06:38:00.000+00:00",
    "latestDeploymentStatus": "ready",
    "vars": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "API_KEY",
            "value": "myPa$$word1",
            "secret": true,
            "resourceType": "function",
            "resourceId": "myAwesomeFunction"
        }
    ],
    "timeout": 300,
    "installCommand": "npm install",
    "buildCommand": "npm run build",
    "startCommand": "node custom-server.mjs",
    "outputDirectory": "build",
    "installationId": "6m40at4ejk5h2u9s1hboo",
    "providerRepositoryId": "appwrite",
    "providerBranch": "main",
    "providerRootDirectory": "sites/helloWorld",
    "providerSilentMode": true,
    "buildSpecification": "s-1vcpu-512mb",
    "runtimeSpecification": "s-1vcpu-512mb",
    "buildRuntime": "node-22",
    "adapter": "static",
    "fallbackFile": "index.html"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Create("<SITE_ID>", "<NAME>", "analog", "node-14.5")
		if err != nil {
			t.Errorf("Method Create failed: %v", err)
		}
	})

	t.Run("Test ListFrameworks", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "frameworks": [
        {
            "key": "sveltekit",
            "name": "SvelteKit",
            "buildRuntime": "node-22",
            "runtimes": [],
            "adapters": [
                {
                    "key": "static",
                    "installCommand": "npm install",
                    "buildCommand": "npm run build",
                    "outputDirectory": "./dist",
                    "fallbackFile": "index.html"
                }
            ]
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListFrameworks()
		if err != nil {
			t.Errorf("Method ListFrameworks failed: %v", err)
		}
	})

	t.Run("Test ListSpecifications", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "specifications": [
        {
            "memory": 512,
            "cpus": 1,
            "enabled": true,
            "slug": "s-1vcpu-512mb"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListSpecifications()
		if err != nil {
			t.Errorf("Method ListSpecifications failed: %v", err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Site",
    "enabled": true,
    "live": true,
    "logging": true,
    "framework": "react",
    "deploymentRetention": 7,
    "deploymentId": "5e5ea5c16897e",
    "deploymentCreatedAt": "2020-10-15T06:38:00.000+00:00",
    "deploymentScreenshotLight": "5e5ea5c16897e",
    "deploymentScreenshotDark": "5e5ea5c16897e",
    "latestDeploymentId": "5e5ea5c16897e",
    "latestDeploymentCreatedAt": "2020-10-15T06:38:00.000+00:00",
    "latestDeploymentStatus": "ready",
    "vars": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "API_KEY",
            "value": "myPa$$word1",
            "secret": true,
            "resourceType": "function",
            "resourceId": "myAwesomeFunction"
        }
    ],
    "timeout": 300,
    "installCommand": "npm install",
    "buildCommand": "npm run build",
    "startCommand": "node custom-server.mjs",
    "outputDirectory": "build",
    "installationId": "6m40at4ejk5h2u9s1hboo",
    "providerRepositoryId": "appwrite",
    "providerBranch": "main",
    "providerRootDirectory": "sites/helloWorld",
    "providerSilentMode": true,
    "buildSpecification": "s-1vcpu-512mb",
    "runtimeSpecification": "s-1vcpu-512mb",
    "buildRuntime": "node-22",
    "adapter": "static",
    "fallbackFile": "index.html"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Get("<SITE_ID>")
		if err != nil {
			t.Errorf("Method Get failed: %v", err)
		}
	})

	t.Run("Test Update", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Site",
    "enabled": true,
    "live": true,
    "logging": true,
    "framework": "react",
    "deploymentRetention": 7,
    "deploymentId": "5e5ea5c16897e",
    "deploymentCreatedAt": "2020-10-15T06:38:00.000+00:00",
    "deploymentScreenshotLight": "5e5ea5c16897e",
    "deploymentScreenshotDark": "5e5ea5c16897e",
    "latestDeploymentId": "5e5ea5c16897e",
    "latestDeploymentCreatedAt": "2020-10-15T06:38:00.000+00:00",
    "latestDeploymentStatus": "ready",
    "vars": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "API_KEY",
            "value": "myPa$$word1",
            "secret": true,
            "resourceType": "function",
            "resourceId": "myAwesomeFunction"
        }
    ],
    "timeout": 300,
    "installCommand": "npm install",
    "buildCommand": "npm run build",
    "startCommand": "node custom-server.mjs",
    "outputDirectory": "build",
    "installationId": "6m40at4ejk5h2u9s1hboo",
    "providerRepositoryId": "appwrite",
    "providerBranch": "main",
    "providerRootDirectory": "sites/helloWorld",
    "providerSilentMode": true,
    "buildSpecification": "s-1vcpu-512mb",
    "runtimeSpecification": "s-1vcpu-512mb",
    "buildRuntime": "node-22",
    "adapter": "static",
    "fallbackFile": "index.html"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Update("<SITE_ID>", "<NAME>", "analog")
		if err != nil {
			t.Errorf("Method Update failed: %v", err)
		}
	})

	t.Run("Test Delete", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Delete("<SITE_ID>")
		if err != nil {
			t.Errorf("Method Delete failed: %v", err)
		}
	})

	t.Run("Test UpdateSiteDeployment", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Site",
    "enabled": true,
    "live": true,
    "logging": true,
    "framework": "react",
    "deploymentRetention": 7,
    "deploymentId": "5e5ea5c16897e",
    "deploymentCreatedAt": "2020-10-15T06:38:00.000+00:00",
    "deploymentScreenshotLight": "5e5ea5c16897e",
    "deploymentScreenshotDark": "5e5ea5c16897e",
    "latestDeploymentId": "5e5ea5c16897e",
    "latestDeploymentCreatedAt": "2020-10-15T06:38:00.000+00:00",
    "latestDeploymentStatus": "ready",
    "vars": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "API_KEY",
            "value": "myPa$$word1",
            "secret": true,
            "resourceType": "function",
            "resourceId": "myAwesomeFunction"
        }
    ],
    "timeout": 300,
    "installCommand": "npm install",
    "buildCommand": "npm run build",
    "startCommand": "node custom-server.mjs",
    "outputDirectory": "build",
    "installationId": "6m40at4ejk5h2u9s1hboo",
    "providerRepositoryId": "appwrite",
    "providerBranch": "main",
    "providerRootDirectory": "sites/helloWorld",
    "providerSilentMode": true,
    "buildSpecification": "s-1vcpu-512mb",
    "runtimeSpecification": "s-1vcpu-512mb",
    "buildRuntime": "node-22",
    "adapter": "static",
    "fallbackFile": "index.html"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateSiteDeployment("<SITE_ID>", "<DEPLOYMENT_ID>")
		if err != nil {
			t.Errorf("Method UpdateSiteDeployment failed: %v", err)
		}
	})

	t.Run("Test ListDeployments", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "deployments": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "vcs",
            "resourceId": "5e5ea6g16897e",
            "resourceType": "functions",
            "entrypoint": "index.js",
            "sourceSize": 128,
            "buildSize": 128,
            "totalSize": 128,
            "buildId": "5e5ea5c16897e",
            "activate": true,
            "screenshotLight": "5e5ea5c16897e",
            "screenshotDark": "5e5ea5c16897e",
            "status": "ready",
            "buildLogs": "Compiling source files...",
            "buildDuration": 128,
            "providerRepositoryName": "database",
            "providerRepositoryOwner": "utopia",
            "providerRepositoryUrl": "https://github.com/vermakhushboo/g4-node-function",
            "providerCommitHash": "7c3f25d",
            "providerCommitAuthorUrl": "https://github.com/vermakhushboo",
            "providerCommitAuthor": "Khushboo Verma",
            "providerCommitMessage": "Update index.js",
            "providerCommitUrl": "https://github.com/vermakhushboo/g4-node-function/commit/60c0416257a9cbcdd96b2d370c38d8f8d150ccfb",
            "providerBranch": "0.7.x",
            "providerBranchUrl": "https://github.com/vermakhushboo/appwrite/tree/0.7.x"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListDeployments("<SITE_ID>")
		if err != nil {
			t.Errorf("Method ListDeployments failed: %v", err)
		}
	})

	t.Run("Test CreateDeployment", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "vcs",
    "resourceId": "5e5ea6g16897e",
    "resourceType": "functions",
    "entrypoint": "index.js",
    "sourceSize": 128,
    "buildSize": 128,
    "totalSize": 128,
    "buildId": "5e5ea5c16897e",
    "activate": true,
    "screenshotLight": "5e5ea5c16897e",
    "screenshotDark": "5e5ea5c16897e",
    "status": "ready",
    "buildLogs": "Compiling source files...",
    "buildDuration": 128,
    "providerRepositoryName": "database",
    "providerRepositoryOwner": "utopia",
    "providerRepositoryUrl": "https://github.com/vermakhushboo/g4-node-function",
    "providerCommitHash": "7c3f25d",
    "providerCommitAuthorUrl": "https://github.com/vermakhushboo",
    "providerCommitAuthor": "Khushboo Verma",
    "providerCommitMessage": "Update index.js",
    "providerCommitUrl": "https://github.com/vermakhushboo/g4-node-function/commit/60c0416257a9cbcdd96b2d370c38d8f8d150ccfb",
    "providerBranch": "0.7.x",
    "providerBranchUrl": "https://github.com/vermakhushboo/appwrite/tree/0.7.x"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" && "POST" != "GET" {
				// Handle file upload resume check
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"chunksUploaded": 0}`))
				return
			}
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateDeployment("<SITE_ID>", file.NewInputFile(tmpFile.Name(), "test.txt"))
		if err != nil {
			t.Errorf("Method CreateDeployment failed: %v", err)
		}
	})

	t.Run("Test CreateDuplicateDeployment", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "vcs",
    "resourceId": "5e5ea6g16897e",
    "resourceType": "functions",
    "entrypoint": "index.js",
    "sourceSize": 128,
    "buildSize": 128,
    "totalSize": 128,
    "buildId": "5e5ea5c16897e",
    "activate": true,
    "screenshotLight": "5e5ea5c16897e",
    "screenshotDark": "5e5ea5c16897e",
    "status": "ready",
    "buildLogs": "Compiling source files...",
    "buildDuration": 128,
    "providerRepositoryName": "database",
    "providerRepositoryOwner": "utopia",
    "providerRepositoryUrl": "https://github.com/vermakhushboo/g4-node-function",
    "providerCommitHash": "7c3f25d",
    "providerCommitAuthorUrl": "https://github.com/vermakhushboo",
    "providerCommitAuthor": "Khushboo Verma",
    "providerCommitMessage": "Update index.js",
    "providerCommitUrl": "https://github.com/vermakhushboo/g4-node-function/commit/60c0416257a9cbcdd96b2d370c38d8f8d150ccfb",
    "providerBranch": "0.7.x",
    "providerBranchUrl": "https://github.com/vermakhushboo/appwrite/tree/0.7.x"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateDuplicateDeployment("<SITE_ID>", "<DEPLOYMENT_ID>")
		if err != nil {
			t.Errorf("Method CreateDuplicateDeployment failed: %v", err)
		}
	})

	t.Run("Test CreateTemplateDeployment", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "vcs",
    "resourceId": "5e5ea6g16897e",
    "resourceType": "functions",
    "entrypoint": "index.js",
    "sourceSize": 128,
    "buildSize": 128,
    "totalSize": 128,
    "buildId": "5e5ea5c16897e",
    "activate": true,
    "screenshotLight": "5e5ea5c16897e",
    "screenshotDark": "5e5ea5c16897e",
    "status": "ready",
    "buildLogs": "Compiling source files...",
    "buildDuration": 128,
    "providerRepositoryName": "database",
    "providerRepositoryOwner": "utopia",
    "providerRepositoryUrl": "https://github.com/vermakhushboo/g4-node-function",
    "providerCommitHash": "7c3f25d",
    "providerCommitAuthorUrl": "https://github.com/vermakhushboo",
    "providerCommitAuthor": "Khushboo Verma",
    "providerCommitMessage": "Update index.js",
    "providerCommitUrl": "https://github.com/vermakhushboo/g4-node-function/commit/60c0416257a9cbcdd96b2d370c38d8f8d150ccfb",
    "providerBranch": "0.7.x",
    "providerBranchUrl": "https://github.com/vermakhushboo/appwrite/tree/0.7.x"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateTemplateDeployment("<SITE_ID>", "<REPOSITORY>", "<OWNER>", "<ROOT_DIRECTORY>", "branch", "<REFERENCE>")
		if err != nil {
			t.Errorf("Method CreateTemplateDeployment failed: %v", err)
		}
	})

	t.Run("Test CreateVcsDeployment", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "vcs",
    "resourceId": "5e5ea6g16897e",
    "resourceType": "functions",
    "entrypoint": "index.js",
    "sourceSize": 128,
    "buildSize": 128,
    "totalSize": 128,
    "buildId": "5e5ea5c16897e",
    "activate": true,
    "screenshotLight": "5e5ea5c16897e",
    "screenshotDark": "5e5ea5c16897e",
    "status": "ready",
    "buildLogs": "Compiling source files...",
    "buildDuration": 128,
    "providerRepositoryName": "database",
    "providerRepositoryOwner": "utopia",
    "providerRepositoryUrl": "https://github.com/vermakhushboo/g4-node-function",
    "providerCommitHash": "7c3f25d",
    "providerCommitAuthorUrl": "https://github.com/vermakhushboo",
    "providerCommitAuthor": "Khushboo Verma",
    "providerCommitMessage": "Update index.js",
    "providerCommitUrl": "https://github.com/vermakhushboo/g4-node-function/commit/60c0416257a9cbcdd96b2d370c38d8f8d150ccfb",
    "providerBranch": "0.7.x",
    "providerBranchUrl": "https://github.com/vermakhushboo/appwrite/tree/0.7.x"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateVcsDeployment("<SITE_ID>", "branch", "<REFERENCE>")
		if err != nil {
			t.Errorf("Method CreateVcsDeployment failed: %v", err)
		}
	})

	t.Run("Test GetDeployment", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "vcs",
    "resourceId": "5e5ea6g16897e",
    "resourceType": "functions",
    "entrypoint": "index.js",
    "sourceSize": 128,
    "buildSize": 128,
    "totalSize": 128,
    "buildId": "5e5ea5c16897e",
    "activate": true,
    "screenshotLight": "5e5ea5c16897e",
    "screenshotDark": "5e5ea5c16897e",
    "status": "ready",
    "buildLogs": "Compiling source files...",
    "buildDuration": 128,
    "providerRepositoryName": "database",
    "providerRepositoryOwner": "utopia",
    "providerRepositoryUrl": "https://github.com/vermakhushboo/g4-node-function",
    "providerCommitHash": "7c3f25d",
    "providerCommitAuthorUrl": "https://github.com/vermakhushboo",
    "providerCommitAuthor": "Khushboo Verma",
    "providerCommitMessage": "Update index.js",
    "providerCommitUrl": "https://github.com/vermakhushboo/g4-node-function/commit/60c0416257a9cbcdd96b2d370c38d8f8d150ccfb",
    "providerBranch": "0.7.x",
    "providerBranchUrl": "https://github.com/vermakhushboo/appwrite/tree/0.7.x"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetDeployment("<SITE_ID>", "<DEPLOYMENT_ID>")
		if err != nil {
			t.Errorf("Method GetDeployment failed: %v", err)
		}
	})

	t.Run("Test DeleteDeployment", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.DeleteDeployment("<SITE_ID>", "<DEPLOYMENT_ID>")
		if err != nil {
			t.Errorf("Method DeleteDeployment failed: %v", err)
		}
	})

	t.Run("Test GetDeploymentDownload", func(t *testing.T) {
		mockData := []byte("image_data")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(mockData)
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetDeploymentDownload("<SITE_ID>", "<DEPLOYMENT_ID>")
		if err != nil {
			t.Errorf("Method GetDeploymentDownload failed: %v", err)
		}
	})

	t.Run("Test UpdateDeploymentStatus", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "vcs",
    "resourceId": "5e5ea6g16897e",
    "resourceType": "functions",
    "entrypoint": "index.js",
    "sourceSize": 128,
    "buildSize": 128,
    "totalSize": 128,
    "buildId": "5e5ea5c16897e",
    "activate": true,
    "screenshotLight": "5e5ea5c16897e",
    "screenshotDark": "5e5ea5c16897e",
    "status": "ready",
    "buildLogs": "Compiling source files...",
    "buildDuration": 128,
    "providerRepositoryName": "database",
    "providerRepositoryOwner": "utopia",
    "providerRepositoryUrl": "https://github.com/vermakhushboo/g4-node-function",
    "providerCommitHash": "7c3f25d",
    "providerCommitAuthorUrl": "https://github.com/vermakhushboo",
    "providerCommitAuthor": "Khushboo Verma",
    "providerCommitMessage": "Update index.js",
    "providerCommitUrl": "https://github.com/vermakhushboo/g4-node-function/commit/60c0416257a9cbcdd96b2d370c38d8f8d150ccfb",
    "providerBranch": "0.7.x",
    "providerBranchUrl": "https://github.com/vermakhushboo/appwrite/tree/0.7.x"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateDeploymentStatus("<SITE_ID>", "<DEPLOYMENT_ID>")
		if err != nil {
			t.Errorf("Method UpdateDeploymentStatus failed: %v", err)
		}
	})

	t.Run("Test ListLogs", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "executions": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": [],
            "functionId": "5e5ea6g16897e",
            "deploymentId": "5e5ea5c16897e",
            "trigger": "http",
            "status": "processing",
            "requestMethod": "GET",
            "requestPath": "/articles?id=5",
            "requestHeaders": [
                {
                    "name": "Content-Type",
                    "value": "application/json"
                }
            ],
            "responseStatusCode": 200,
            "responseBody": "string",
            "responseHeaders": [
                {
                    "name": "Content-Type",
                    "value": "application/json"
                }
            ],
            "logs": "string",
            "errors": "string",
            "duration": 0.4
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListLogs("<SITE_ID>")
		if err != nil {
			t.Errorf("Method ListLogs failed: %v", err)
		}
	})

	t.Run("Test GetLog", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "functionId": "5e5ea6g16897e",
    "deploymentId": "5e5ea5c16897e",
    "trigger": "http",
    "status": "processing",
    "requestMethod": "GET",
    "requestPath": "/articles?id=5",
    "requestHeaders": [
        {
            "name": "Content-Type",
            "value": "application/json"
        }
    ],
    "responseStatusCode": 200,
    "responseBody": "string",
    "responseHeaders": [
        {
            "name": "Content-Type",
            "value": "application/json"
        }
    ],
    "logs": "string",
    "errors": "string",
    "duration": 0.4
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetLog("<SITE_ID>", "<LOG_ID>")
		if err != nil {
			t.Errorf("Method GetLog failed: %v", err)
		}
	})

	t.Run("Test DeleteLog", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.DeleteLog("<SITE_ID>", "<LOG_ID>")
		if err != nil {
			t.Errorf("Method DeleteLog failed: %v", err)
		}
	})

	t.Run("Test ListVariables", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "variables": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "API_KEY",
            "value": "myPa$$word1",
            "secret": true,
            "resourceType": "function",
            "resourceId": "myAwesomeFunction"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListVariables("<SITE_ID>")
		if err != nil {
			t.Errorf("Method ListVariables failed: %v", err)
		}
	})

	t.Run("Test CreateVariable", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "key": "API_KEY",
    "value": "myPa$$word1",
    "secret": true,
    "resourceType": "function",
    "resourceId": "myAwesomeFunction"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateVariable("<SITE_ID>", "<KEY>", "<VALUE>")
		if err != nil {
			t.Errorf("Method CreateVariable failed: %v", err)
		}
	})

	t.Run("Test GetVariable", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "key": "API_KEY",
    "value": "myPa$$word1",
    "secret": true,
    "resourceType": "function",
    "resourceId": "myAwesomeFunction"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetVariable("<SITE_ID>", "<VARIABLE_ID>")
		if err != nil {
			t.Errorf("Method GetVariable failed: %v", err)
		}
	})

	t.Run("Test UpdateVariable", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "key": "API_KEY",
    "value": "myPa$$word1",
    "secret": true,
    "resourceType": "function",
    "resourceId": "myAwesomeFunction"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateVariable("<SITE_ID>", "<VARIABLE_ID>", "<KEY>")
		if err != nil {
			t.Errorf("Method UpdateVariable failed: %v", err)
		}
	})

	t.Run("Test DeleteVariable", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.DeleteVariable("<SITE_ID>", "<VARIABLE_ID>")
		if err != nil {
			t.Errorf("Method DeleteVariable failed: %v", err)
		}
	})
}
