#!/usr/bin/env bash
# Copyright 2024 The KubeStellar Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This is an end to end test of multi cluster deployement.  
# For readable instructions, please visit https://github.com/kubestellar/kubestellar/tree/ks-0.20/docs/content/v0.20

set -x # so users can see what is going on
set -e # exit on error

./cleanup.sh
./setup-kubestellar.sh
./workload-deployment.sh
