.DEFAULT_GOAL := help

# ----------------------------------------------------------------------------------------------------------------------
# Variables
# ----------------------------------------------------------------------------------------------------------------------

SHELLBITS_DIR?=$(shell dirname $(shell which shellbits))

# ----------------------------------------------------------------------------------------------------------------------
# Utilities functions
# ----------------------------------------------------------------------------------------------------------------------

include ${SHELLBITS_DIR}/utils.mk

# ----------------------------------------------------------------------------------------------------------------------
# Utilities targets
# ----------------------------------------------------------------------------------------------------------------------

include ${SHELLBITS_DIR}/help.mk

# ----------------------------------------------------------------------------------------------------------------------
# QA targets
# ----------------------------------------------------------------------------------------------------------------------

include ${SHELLBITS_DIR}/file.mk
include ${SHELLBITS_DIR}/golang.mk
include ${SHELLBITS_DIR}/goreleaser.mk
include ${SHELLBITS_DIR}/markdown.mk
include ${SHELLBITS_DIR}/mise.mk
include ${SHELLBITS_DIR}/mkfile.mk

# ----------------------------------------------------------------------------------------------------------------------
# Development targets
# ----------------------------------------------------------------------------------------------------------------------

# N/A
