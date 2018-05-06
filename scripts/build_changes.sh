
# LAST_COMMIT=$(git log --name-status HEAD^..HEAD --pretty=oneline | awk 'FNR==1 {print $1}') #last commit at all
#  PREV_COMMIT=$(git rev-list --min-parents=2 --max-count=2 HEAD | awk 'FNR==2') #Last merge
