#! /usr/bin/env bash

# Substitua todas as referências ao módulo antigo
CURR_MODULE='github.com/go-web-templates/api'
NEW_MODULE='github.com/mateusnasciment/APIGOLANG'

echo "Substituindo o nome do módulo"
find . -type f -name "*.go" -print0 | \
  xargs -0 sed -i '' -e "s#${CURR_MODULE}#${NEW_MODULE}#g" \
  2>/dev/null

echo "Excluindo os antigos arquivos de módulo Go"
rm -rf go.mod go.sum

echo "Criando um novo módulo Go"
go mod init ${NEW_MODULE}
go mod tidy
