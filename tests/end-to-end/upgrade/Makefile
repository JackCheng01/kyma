APP_NAME = e2e-upgrade-test
APP_PATH = tests/end-to-end/upgrade
BUILDPACK = eu.gcr.io/kyma-project/test-infra/buildpack-golang-toolbox:v20190930-d28d219
SCRIPTS_DIR = $(realpath $(shell pwd)/../../..)/scripts

include $(SCRIPTS_DIR)/go-dep.mk

VERIFY_IGNORE := /vendor\|/automock\|/testdata\|/pkg
