#!/usr/bin/env bash

# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -eEuo pipefail

ROOT="$(cd "$(dirname "$0")/.." &>/dev/null; pwd -P)"
SOURCE_DIRS="cmd internal tools"


echo "๐ณ Set up environment variables"
eval $(${ROOT}/scripts/dev init)


echo "๐ Verify Protobufs are up to date"
${ROOT}/scripts/dev protoc
# Don't verify generated pb files here as they are tidied later.


echo "๐งฝ Verify goimports formattting"
set +e
which goimports >/dev/null 2>&1
if [ $? -ne 0 ]; then
   echo "โ No 'goimports' found. Please use"
   echo "โ   go get golang.org/x/tools/cmd/goimports"
   echo "โ to enable import cleanup. Import cleanup skipped."
else
   echo "๐งฝ Format with goimports"
   goimports -w $(echo $SOURCE_DIRS)
   # Check if there were uncommited changes.
   # Ignore comment line changes as sometimes proto gen just updates versions
   # of the generator
   git diff -G'(^\s+[^/])' *.go | tee /dev/stderr | (! read)
   if [ $? -ne 0 ]; then
      echo "โ Found uncommited changes after goimports."
      echo "โ Commit these changes before merging."
      exit 1
   fi
fi
set -e


echo "๐งน Verify gofmt format"
set +e
diff -u <(echo -n) <(gofmt -d -s .)
git diff -G'(^\s+[^/])' *.go | tee /dev/stderr | (! read)
if [ $? -ne 0 ]; then
   echo "โ Found uncommited changes after gofmt."
   echo "โ Commit these changes before merging."
   exit 1
fi
set -e


echo "๐ Go mod verify"
set +e
go mod verify
if [ $? -ne 0 ]; then
   echo "โ go mod verify failed."
   exit 1
fi
set -e

# Fail if a dependency was added without the necessary go.mod/go.sum change
# being part of the commit.
echo "๐ Go mod tidy"
set +e
go mod tidy;
git diff go.mod | tee /dev/stderr | (! read)
if [ $? -ne 0 ]; then
   echo "โ Found uncommited go.mod changes after go mod tidy."
   exit 1
fi
git diff go.sum | tee /dev/stderr | (! read)
if [ $? -ne 0 ]; then
   echo "โ Found uncommited go.sum changes after go mod tidy."
   exit 1
fi
set -e

echo "๐จ Running 'go vet'..."
go vet ./...


echo "๐ง Compile"
go build ./...


echo "๐งช Test"
go test ./... \
  -coverprofile=coverage.out \
  -count=1 \
  -parallel=20 \
  -timeout=5m


echo "๐งโ๐ฌ Test Coverage"
go tool cover -func coverage.out | grep total | awk '{print $NF}'
