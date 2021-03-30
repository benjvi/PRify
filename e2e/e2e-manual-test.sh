#!/bin/sh
set -euo pipefail

cd examples/minimal
mkdir aaa; mkdir bbb; mkdir ccc; echo test > {aaa,bbb,ccc}/test
prify run; git checkout main
echo "PRs should have been created from minimal prify.yml, check all is as expected"
cd -

cd examples/mvp
mkdir aaa; mkdir bbb; mkdir ccc; echo test > {aaa,bbb,ccc}/test
prify run; git checkout main
echo "PRs should have been created from mvp prify.yml, check all is as expected"
cd -
