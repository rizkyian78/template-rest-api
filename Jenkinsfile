pipeline {
   agent {
        docker { image 'node:16.13.1-alpine' }
    }
    environment {
        DOCKER_ACCOUNT="rizkyian78"
        DOCKER_PASSWORD="yukaritakeba12"
    }
    stages {
        stage('Building Docker') {
            steps {
                 sh 'node --version'
                echo "${DOCKER_PASSWORD}"
                sh "docker login --username ${DOCKER_ACCOUNT} --password-stdin ${DOCKER_PASSWORD}"
            }
        }
    }
}
