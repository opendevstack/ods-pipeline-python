# ods-pipeline-python

[![Tests](https://github.com/opendevstack/ods-pipeline-python/actions/workflows/main.yaml/badge.svg)](https://github.com/opendevstack/ods-pipeline-python/actions/workflows/main.yaml)

Tekton task for use with [ODS Pipeline](https://github.com/opendevstack/ods-pipeline) to build applications with Python.

## Usage

```yaml
tasks:
- name: build
  taskRef:
    resolver: git
    params:
    - { name: url, value: https://github.com/opendevstack/ods-pipeline-python.git }
    - { name: revision, value: v0.1.0 }
    - { name: pathInRepo, value: tasks/build.yaml }
    workspaces:
    - { name: source, workspace: shared-workspace }
```

See the [documentation](https://github.com/opendevstack/ods-pipeline-python/blob/main/docs/build.adoc) for details and available parameters.

## About this repository

`docs` and `tasks` are generated directories from recipes located in `build`. See the `Makefile` target for how everything fits together.
