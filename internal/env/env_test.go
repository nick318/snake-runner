package env

import (
	"testing"

	"github.com/reconquest/snake-runner/internal/builtin"
	"github.com/reconquest/snake-runner/internal/config"
	"github.com/reconquest/snake-runner/internal/responses"
	"github.com/reconquest/snake-runner/internal/runner"
	"github.com/reconquest/snake-runner/internal/snake"
	"github.com/reconquest/snake-runner/internal/tasks"
	"github.com/stretchr/testify/assert"
)

func TestEnvBuilder(t *testing.T) {
	test := assert.New(t)

	basicPipeline := snake.Pipeline{
		ID:       123,
		Commit:   "1234567890",
		RunnerID: 80,
	}
	job := snake.PipelineJob{
		ID:    321,
		Stage: "deploy",
		Name:  "docker deploy",
	}
	runnerConfig := runner.Config{
		Name: "gotest",
	}
	task := tasks.PipelineRun{
		Pipeline: basicPipeline,
		Env:      map[string]string{"user_a": "user_a_value"},
		Project: responses.Project{
			Key:  "proj1",
			Name: "the proj1",
			ID:   11,
		},
		Repository: responses.Repository{
			Slug: "repo1",
			Name: "the repo1",
			ID:   111,
		},
	}
	task.CloneURL.SSH = "cloneurl"

	configPipeline := config.Pipeline{}
	configJob := config.Job{}

	builder := func(pipeline snake.Pipeline) *Builder {
		return NewBuilder(
			task,
			pipeline,
			job,
			configPipeline,
			configJob,
			&runnerConfig,
			"/git",
			"/ssh",
		)
	}

	expected := map[string]string{
		"user_a":                "user_a_value",
		"CI":                    "true",
		"CI_PIPELINE_ID":        "123",
		"CI_JOB_ID":             "321",
		"CI_JOB_STAGE":          "deploy",
		"CI_JOB_NAME":           "docker deploy",
		"CI_COMMIT_HASH":        "1234567890",
		"CI_COMMIT_SHORT_HASH":  "123456",
		"CI_PIPELINE_DIR":       "/git",
		"CI_PROJECT_KEY":        "proj1",
		"CI_PROJECT_NAME":       "the proj1",
		"CI_PROJECT_ID":         "11",
		"CI_REPO_SLUG":          "repo1",
		"CI_REPO_NAME":          "the repo1",
		"CI_REPO_ID":            "111",
		"CI_REPO_CLONE_URL_SSH": "cloneurl",
		"CI_RUNNER_ID":          "80",
		"CI_RUNNER_NAME":        "gotest",
		"CI_RUNNER_VERSION":     builtin.Version,
		"SSH_AUTH_SOCK":         "/ssh/ssh-agent.sock",
	}

	{
		test.EqualValues(expected, builder(basicPipeline).build())
	}

	{
		pipeline := basicPipeline
		pipeline.RefType = "BRANCH"
		pipeline.RefDisplayId = "someref"

		expected := clone(expected)
		expected["CI_BRANCH"] = "someref"

		test.EqualValues(expected, builder(pipeline).build())
	}

	{
		pipeline := basicPipeline
		pipeline.RefType = "TAG"
		pipeline.RefDisplayId = "someref"

		expected := clone(expected)
		expected["CI_TAG"] = "someref"

		test.EqualValues(expected, builder(pipeline).build())
	}

	{
		configPipeline.Variables = map[string]string{"foo": "global"}

		expected := clone(expected)
		expected["foo"] = "global"

		test.EqualValues(expected, builder(basicPipeline).build())
	}

	{
		configJob.Variables = map[string]string{"foo": "job"}

		expected := clone(expected)
		expected["foo"] = "job"

		test.EqualValues(expected, builder(basicPipeline).build())
	}

	{
		configPipeline.Variables = map[string]string{"foo": "globalfoo", "bar": "globalbar"}
		configJob.Variables = map[string]string{"foo": "foojob", "qux": "quxjob"}

		expected := clone(expected)
		expected["foo"] = "foojob"
		expected["qux"] = "quxjob"
		expected["bar"] = "globalbar"

		test.EqualValues(expected, builder(basicPipeline).build())
	}
}

func clone(original map[string]string) map[string]string {
	result := map[string]string{}
	for key, value := range original {
		result[key] = value
	}
	return result
}
