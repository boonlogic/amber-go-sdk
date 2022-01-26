#!/bin/bash

${GOPATH}/bin/gomarkdoc . ./models \
	--repository.default-branch master \
	--repository.url https://github.com/boonlogic/amber-go-sdk > docs/user-docs.md 

pandoc docs/user-docs.md --toc --metadata title="amber-go-sdk User Guide" -c https://unpkg.com/sakura.css/css/sakura.css --self-contained -o docs/user-docs.html
