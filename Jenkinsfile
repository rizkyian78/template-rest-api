pipeline {
    agent any
    environment {
        DOCKER_ACCOUNT="rizkyian78"
        DOCKER_PASSWORD="yukaritakeba12"
    }
    stages {
        stage('Building Docker') {

            steps {
                echo "${DOCKER_PASSWORD}"
                echo "${PATH}"
                sh "docker login --username ${DOCKER_ACCOUNT} --password-stdin ${DOCKER_PASSWORD}"
            }
        }
    }
}
