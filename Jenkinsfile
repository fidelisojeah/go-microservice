pipeline{
  agent{
    docker{
      image 'eu.gcr.io/fidelis-microservice/gcdeploy:latest'
      args '-u root:root'
    }
  }
  environment{
    GCLOUD_PROJECT_NAME = credentials('project_name')//'fidelis-microservice'
    GCLOUD_CLUSTER_NAME = credentials('cluster-name')//'microservice-cluster'
    CLOUDSDK_COMPUTE_ZONE = 'us-central1-a'
    GCLOUD_SERVICE_KEY = credentials('GCLOUD_SERVICE_KEY')
  }
  stages{
    stage('Build Dependencies'){
      steps{
        sh "touch ~/gcloud-service-key.json"
        echo "${env.GCLOUD_SERVICE_KEY} | base64 --decode >> ~/gcloud-service-key.json"
        sh "ls ~"
        echo '---decoding service key---'
        sh "cat ~/gcloud-service-key.json"
        echo '---decoded service key---'
        sh "gcloud auth activate-service-account --key-file ~/gcloud-service-key.json"
        sh "gcloud config set project ${env.GCLOUD_PROJECT_NAME}"
        sh "gcloud --quiet config set container/cluster ${env.GCLOUD_CLUSTER_NAME}"
        sh "gcloud config set compute/zone ${env.CLOUDSDK_COMPUTE_ZONE}"
        sh "gcloud --quiet container clusters get-credentials ${env.GCLOUD_CLUSTER_NAME}"
        echo 'credentials set'
      }
    }
    // stage('Build Changes'){

    // }
    // stage('Deploy Changes')
  }
}
