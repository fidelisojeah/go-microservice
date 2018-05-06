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
    GIT_BRANCH = sh(returnStdout: true, script: 'git rev-parse --abbrev-ref HEAD').trim()
    LAST_COMMIT = sh(returnStdout: true, script: "git log --name-status HEAD^..HEAD --pretty=oneline | awk 'FNR==1 {print $1}'").trim()
    PREV_MERGE_COMMIT = sh(returnStdout: true, script: "git rev-list --min-parents=2 --max-count=2 HEAD | awk 'FNR==2'").trim()
    CHANGED_FOLDERS = sh(returnStdout: true, script: "git diff --name-only $PREV_MERGE_COMMIT ^LAST_COMMIT | grep / | awk 'BEGIN {FS="/"} {print $1}' | uniq").trim()

  }
  stages{
    stage('Build Dependencies'){
      steps{
        sh "touch ~/gcloud-service-key.json"
        sh "echo ${env.GCLOUD_SERVICE_KEY} | base64 --decode >> ~/gcloud-service-key.json"
        echo '---decoding service key--'
        sh "gcloud auth activate-service-account --key-file ~/gcloud-service-key.json"
        sh "gcloud config set project ${env.GCLOUD_PROJECT_NAME}"
        sh "gcloud --quiet config set container/cluster ${env.GCLOUD_CLUSTER_NAME}"
        sh "gcloud config set compute/zone ${env.CLOUDSDK_COMPUTE_ZONE}"
        sh "gcloud --quiet container clusters get-credentials ${env.GCLOUD_CLUSTER_NAME}"
        echo 'credentials set'
      }
    }
    stage('Build Changes'){
      when{
        expression{
          return GIT_BRANCH == "develop" || GIT_BRANCH =="master"
        }
      }
      steps{
        echo "Current branch is ${env.GIT_BRANCH}"
        echo "Checking out commit:${env.GIT_COMMIT}"
        checkout scm
        sh 'ls'
        echo "Changed folders: ${env.CHANGED_FOLDERS}"
      }
    }
    // stage('Deploy Changes'){
      // when{
      //   expression{
      //     return GIT_BRANCH == "develop" || GIT_BRANCH =="master"
      //   }
      // }
    // }
  }
}
