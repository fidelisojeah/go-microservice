eploy-service(){
  echo "eploying service----"
  kubectl create -f ./deployment.yml
  kubectl create -f ./service.yml
}
