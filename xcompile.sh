#!/bin/bash
work="onamaeddns/exec"
export GOPATH="`pwd`"
ls -1 src/${work} | while read row ; do
  echo "compile ${row}"
  GOOS=linux GOARCH=amd64 go install ${work}/${row}
  GOOS=windows GOARCH=amd64 go install ${work}/${row}
  GOOS=darwin GOARCH=amd64 go install ${work}/${row}
done
echo "done"
exit 0
