// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import "time"

type (
	BuildPayload struct {
		Builds []*Build `json:"builds"`
	}
	DeploymentPayload struct {
		Deployments []*Deployment `json:"deployments"`
	}
	DevelopmentInformationPayload struct {
		Repositories []*Repository `json:"repositories"`
	}

	// build provides the build details.
	Build struct {
		BuildNumber int       `json:"buildNumber"`
		Description string    `json:"description"`
		DisplayName string    `json:"displayName"`
		IssueKeys   []string  `json:"issueKeys"`
		Label       string    `json:"label"`
		LastUpdated time.Time `json:"lastUpdated"`
		PipelineID  string    `json:"pipelineId"`
		References  []struct {
			Commit struct {
				ID            string `json:"id"`
				RepositoryURI string `json:"repositoryUri"`
			} `json:"commit"`
			Ref struct {
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"ref"`
		} `json:"references"`
		SchemaVersion string `json:"schemaVersion"`
		State         string `json:"state"`
		TestInfo      struct {
			NumberFailed  int64 `json:"numberFailed"`
			NumberPassed  int64 `json:"numberPassed"`
			NumberSkipped int64 `json:"numberSkipped"`
			TotalNumber   int64 `json:"totalNumber"`
		} `json:"testInfo"`
		UpdateSequenceNumber int    `json:"updateSequenceNumber"`
		URL                  string `json:"url"`
	}

	// Deployment provides the Deployment details.
	Deployment struct {
		Deploymentsequencenumber int           `json:"deploymentSequenceNumber"`
		Updatesequencenumber     int           `json:"updateSequenceNumber"`
		Associations             []Association `json:"associations"`
		Displayname              string        `json:"displayName"`
		URL                      string        `json:"url"`
		Description              string        `json:"description"`
		Lastupdated              time.Time     `json:"lastUpdated"`
		State                    string        `json:"state"`
		Pipeline                 JiraPipeline  `json:"pipeline"`
		Environment              Environment   `json:"environment"`
	}

	// Association provides the association details.
	Association struct {
		Associationtype string   `json:"associationType"`
		Values          []string `json:"values"`
	}

	// Environment provides the environment details.
	Environment struct {
		ID          string `json:"id"`
		Displayname string `json:"displayName"`
		Type        string `json:"type"`
	}

	// JiraPipeline provides the jira pipeline details.
	JiraPipeline struct {
		ID          string `json:"id"`
		Displayname string `json:"displayName"`
		URL         string `json:"url"`
	}

	// Repository provides commit, branches, and pull request details
	Repository struct {
		Name             string `json:"name"`
		Url              string `json:"url"`
		Id               string `json:"id"`
		UpdateSequenceId int    `json:"updateSequenceId"`

		PullRequests []*PullRequest `json:"pullRequests"`
	}

	PullRequest struct {
		Id               string   `json:"id"`
		IssueKeys        []string `json:"issueKeys"`
		UpdateSequenceId int      `json:"updateSequenceId"`
		// OPEN, MERGED, DECLINED, UNKNOWN
		Status            string `json:"status"`
		Title             string `json:"title"`
		Author            Author `json:"author"`
		CommentCount      int    `json:"commentCount"`
		SourceBranch      string `json:"sourceBranch"`
		DestinationBranch string `json:"destinationBranch"`
		LastUpdate        string `json:"lastUpdate"` // ISO8601 / RFC3339 fromat
		Url               string `json:"url"`
		DisplayId         string `json:"displayId"`
	}

	Author struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Url      string `json:"url"`
		Avatar   string `json:"avatar"`
	}

	// Tenant provides the jira instance tenant details.
	Tenant struct {
		ID string `json:"cloudId"`
	}

	// struct for adaptive card
	Card struct {
		Pipeline    string `json:"pipeline"`
		Instance    string `json:"instance"`
		Project     string `json:"project"`
		State       string `json:"state"`
		Version     string `json:"version"`
		Environment string `json:"environment"`
		URL         string `json:"url"`
	}
)
