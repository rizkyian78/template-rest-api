pipeline {
    agent any
    stages {
        stage('Example Build') {
            steps {
                echo env.GIT_TAG
                echo "sampek sini"
            }
        }
        stage('Example Deploy') {
            when {
                branch 'production'
            }
            steps {
                echo 'Deploying'
            }
        }
    }
}
