APP_NAME = application-broker
APP_PATH = components/$(APP_NAME)
BUILDPACK = eu.gcr.io/kyma-project/test-infra/buildpack-golang-toolbox:v20190930-d28d219
SCRIPTS_DIR = $(realpath $(shell pwd)/../..)/scripts

include $(SCRIPTS_DIR)/go-dep.mk

MAIN_PATH:=cmd/broker/main.go
VERIFY_IGNORE := /vendor\|/automock\|/testdata\|/pkg

