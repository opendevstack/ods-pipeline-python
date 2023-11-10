package e2e

import (
	"bytes"
	"path/filepath"
	"testing"

	ott "github.com/opendevstack/ods-pipeline/pkg/odstasktest"
	"github.com/opendevstack/ods-pipeline/pkg/pipelinectxt"
	ttr "github.com/opendevstack/ods-pipeline/pkg/tektontaskrun"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func TestNPMBuildTask(t *testing.T) {
	if err := runTask(
		ttr.WithStringParams(map[string]string{
			"cache-build": "false",
		}),
		ott.WithGitSourceWorkspace(t, "../testdata/workspaces/python-fastapi-sample-app", namespaceConfig.Name),
		ttr.AfterRun(func(config *ttr.TaskRunConfig, run *tekton.TaskRun, logs bytes.Buffer) {
			wd := config.WorkspaceConfigs["source"].Dir
			ott.AssertFilesExist(t, wd,
				filepath.Join(pipelinectxt.XUnitReportsPath, "report.xml"),
				filepath.Join(pipelinectxt.CodeCoveragesPath, "coverage.xml"),
				"src/main.py",
				"requirements.txt",
			)
		}),
	); err != nil {
		t.Fatal(err)
	}
}
