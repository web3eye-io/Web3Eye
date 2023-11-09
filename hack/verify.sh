#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# cd to the repo root
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "${REPO_ROOT}"

# Some useful colors.
if [[ -z "${color_start-}" ]]; then
  declare -r color_start="\033["
  declare -r color_red="${color_start}0;31m"
  declare -r color_green="${color_start}0;32m"
  declare -r color_norm="${color_start}0m"
fi

# Excluded check patterns are always skipped.
EXCLUDED_PATTERNS=(
  "verify.sh"                # this script calls the make rule and would cause a loop
  "verify-*-dockerized.sh"   # Don't run any scripts that intended to be run dockerized
  "verify-shellcheck.sh"
  "verify-build.sh"
)

EXCLUDED_CHECKS=$(ls ${EXCLUDED_PATTERNS[@]/#/${REPO_ROOT}\/hack\/} 2>/dev/null || true)

function is-excluded {
  for e in ${EXCLUDED_CHECKS[@]}; do
    if [[ $1 -ef "$e" ]]; then
      return
    fi
  done
  return 1
}

function run-cmd {
  if ${SILENT}; then
    "$@" &> /dev/null
  else
    "$@"
  fi
}

# Collect failed tests in this array; initialize it to nil
FAILED_TESTS=()

function print-failed-tests {
  echo -e "========================"
  echo -e "${color_red}FAILED TESTS${color_norm}"
  echo -e "========================"
  for t in "${FAILED_TESTS[@]}"; do
      echo -e "${color_red}${t}${color_norm}"
  done
}

function run-checks {
  local -r pattern=$1
  local -r runner=$2

  local t
  local check_name
  local start

  for t in $(ls ${pattern})
  do
    check_name="$(basename "${t}")"
    if is-excluded "${t}" ; then
      echo "Skipping ${check_name}"
      continue
    fi

    echo -e "Verifying ${check_name}"

    start=$(date +%s)
    run-cmd "${runner}" "${t}" && tr=$? || tr=$?
    local elapsed=$(($(date +%s) - start))

    if [[ ${tr} -eq 0 ]]; then
      echo -e "${color_green}SUCCESS${color_norm}  ${check_name}\t${elapsed}s"
    else
      echo -e "${color_red}FAILED${color_norm}   ${check_name}\t${elapsed}s"
      ret=1

      FAILED_TESTS+=("${t}")
    fi
  done
}

SILENT=false

while getopts ":s" opt; do
  case ${opt} in
    s)
      SILENT=true
      ;;
    \?)
      echo "Invalid flag: -${OPTARG}" >&2
      exit 1
      ;;
  esac
done

if ${SILENT} ; then
  echo "Running in silent mode, run without -s if you want to see script logs."
fi

ret=0
run-checks "${REPO_ROOT}/hack/*-verify-*.sh" bash

if [[ ${ret} -eq 1 ]]; then
  print-failed-tests
fi

exit ${ret}

# ex: ts=2 sw=2 et filetype=sh
