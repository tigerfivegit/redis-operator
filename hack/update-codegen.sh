#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

#//项目示例使用 vendor进行的包管理方式,需要下载包到本地,所有直接使用sdk方式
#SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
#CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
# 因为 generate-groups.sh 是通过sdk获得，需要需改相对路径 [将代码上传后 go mod tidy下载sdk到本地]
#位置在$GOPATH/pkg/mod/k8s.io/code-generator\@v0.26.3/
#同时需要修改项目pkg的位置因为在远程调试，使用go.mod项目名称 redis-controller
#同时需要修改group:version、output-base、go-header-file等内容
#根据位置定义CODEGEN_PKG变量
CODEGEN_PKG=$GOPATH/pkg/mod/k8s.io/code-generator\@v0.26.3/
"${CODEGEN_PKG}/generate-groups.sh" "deepcopy,client,informer,lister" \
  redis-controller/pkg/generated \
  redis-controller/pkg/apis \
  stable:v1beta1 \
  --output-base ./ \
  --go-header-file ./boilerplate.go.txt

# To use your own boilerplate text append:
#   --go-header-file "${SCRIPT_ROOT}"/hack/custom-boilerplate.go.txt