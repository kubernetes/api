set -o errexit
set -o nounset
set -o pipefail

# TODO: use the real one
kubernetes_remote="https://github.com/caesarxuchao/kubernetes"

git remote add upstream-kube "${kubernetes_remote}" || true
git fetch upstream-kube
git checkout master
git branch -D kube-sync || true
git checkout upstream-kube/master -b kube-sync

#git filter-branch --index-filter 'git rm --cached -qr --ignore-unmatch -- . && git reset -q $GIT_COMMIT -- pkg/api pkg/apis' --prune-empty HEAD
git filter-branch -f --index-filter 'git rm --cached -qr --ignore-unmatch -- . && git reset -q $GIT_COMMIT -- pkg/api pkg/apis' --msg-filter 'awk 1 && echo && echo "Kubernetes-commit: ${GIT_COMMIT}"' --prune-empty HEAD

mkdir -p core/v1
mv pkg/api/v1/* core/v1/

dirs=$(find pkg/apis/ -maxdepth 2 -mindepth 2 -type d -name v*)

for dir in $dirs; do
    no_pkg="${dir#pkg/apis/}"
    mkdir -p $no_pkg && mv $dir/* $no_pkg
done

git add copy-history.sh
git commit -m "the script that preserves the history"
# TODO: create the "move" commit
git add -A
git commit -am "move all apis to new directories"
# TODO: remove all contents and create another commit
rm -r ./*
git commit -am "remove all files"


# see https://stackoverflow.com/questions/9803294/prune-empty-merge-commits-from-history-in-git-repository
# and http://git.661346.n2.nabble.com/Removing-useless-merge-commit-with-quot-filter-branch-quot-td7356544.html
git checkout kube-sync || true
git filter-branch -f --prune-empty --parent-filter \
    'sed "s/-p //g" | xargs -r git show-branch --independent | sed "s/\</-p /g"'

# cherry-pick the commit that adds the LICENSE and README
git cherry-pick f9e9a31d684b2f1593609fcbdd0b7c307a9cce5d
