pipeline {
  agent any
  stages {
    stage('Build') {
      when {
        branch 'master'
      }
      steps {
        sh 'echo "Building tag ${env.GIT_TAG}"'
      }
    }
  }
}