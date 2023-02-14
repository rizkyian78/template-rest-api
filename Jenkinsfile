pipeline {
    agent any
    stages {
        stage('Example Build') {
            steps {
                echo $BRANCH_NAME
                echo $TAG_NAME
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
