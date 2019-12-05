#!/bin/bash

blocklabel () {

  blue=$(tput setaf 4)
  green=$(tput setaf 2)
  normal="\033[0m"
  underline=$(tput smul)
  disableunderline=$(tput rmul)
  bold=$(tput smso)
  unbold=$(tput rmso)

  printf "${green} \n************************************************************************************************************\n"
  echo "$@"
  printf "************************************************************************************************************${normal}"
  printf "\n"

}

blocklabel "Rename Variable"
go run rename_variable/main.go

blocklabel "Inline Method"
go run inline_method/main.go

blocklabel "Extract Variable"
go run extract_variable/main.go

blocklabel "Extract Method"
go run extract_method/main.go

blocklabel "Move Method"
go run move_method/main.go

blocklabel "Parameterize Method"
go run parameterize_method/main.go
