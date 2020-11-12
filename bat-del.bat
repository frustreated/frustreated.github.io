set http_proxy=http://127.0.0.1:1080
set https_proxy=http://127.0.0.1
git checkout --orphan latest_branch
git add -A
git commit -am "commit message"
git branch -D master
git branch -m master
git push -f origin master