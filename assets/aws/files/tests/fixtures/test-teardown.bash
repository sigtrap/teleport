#!/bin/bash
unset TELEPORT_TEST_FIPS_MODE
unset GENERATE_EXIT_CODE
rm -f ${TELEPORT_CONFIG_PATH?}
rm -rf ${TELEPORT_CONFD_DIR?}
