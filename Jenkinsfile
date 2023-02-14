pipeline {
    agent any
    stages {
        stage('Example Build') {
            steps {
               sh(returnStdout: true, script: "git tag --contains | head -1").trim()
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
