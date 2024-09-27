## Fun Facts

Did you know that to use private github repositories as Go modules, all you have to do is
```sh
git config --global url."git@github.com:meb-ru/".insteadOf "https://github.com/meb-ru/"
echo export GOPRIVATE=github.com/meb-ru/,\$GOPRIVATE >> . ~/.profile
. ~/.profile
go clean -cache -modcache
```