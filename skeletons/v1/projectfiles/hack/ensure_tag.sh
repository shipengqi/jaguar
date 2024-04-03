#!/bin/bash

version="${VERSION}"
if [[ "${version}" == "" ]];then
  version=v`gsemver bump` # such as 0.0.0+24.be1f3ad
fi

if [[ -z "`git tag -l ${version}`" ]];then
  git tag -a -m "release version ${version}" ${version}
fi
