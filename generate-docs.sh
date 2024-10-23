#!/bin/bash

go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest

${GOPATH}/bin/gomarkdoc . ./v1/models \
	--repository.default-branch master \
	--repository.url https://github.com/boonlogic/amber-go-sdk > docs/user-docs-v1.md 

${GOPATH}/bin/gomarkdoc . ./v2/models \
	--repository.default-branch master \
	--repository.url https://github.com/boonlogic/amber-go-sdk > docs/user-docs-v2.md 

pandoc docs/user-docs-v1.md --toc --metadata title="amber-go-sdk v1 User Guide" -c https://unpkg.com/sakura.css/css/sakura.css --self-contained -o docs/user-docs-v1.html
pandoc docs/user-docs-v2.md --toc --metadata title="amber-go-sdk v2 User Guide" -c https://unpkg.com/sakura.css/css/sakura.css --self-contained -o docs/user-docs-v2.html
pandoc docs/user-docs.md --toc --metadata title="amber-go-sdk User Guide" -c https://unpkg.com/sakura.css/css/sakura.css --self-contained -o docs/user-docs.html