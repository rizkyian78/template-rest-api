pipeline {
    agent any
    stages {
        stage('Example Build') {
            steps {
                echo git tag --sort=-creatordate | head -n 1
                echo "><<<<<"
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
