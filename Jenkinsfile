pipeline {
  agent any
  stages {
    stage('Build') {
      when {
        expression { env.GIT_TAG != null }
      }
      steps {
        sh 'echo "Building tag ${env.GIT_TAG}"'
        // add build steps here
      }
    }
  }
}