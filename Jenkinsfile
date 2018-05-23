pipeline {
    agent { label 'linux && docker' }
    options {
        buildDiscarder(logRotator(daysToKeepStr: "30", artifactDaysToKeepStr: "30"))
    }
    stages {
        stage("Run Tests") {
            agent {
                docker {
                    image 'golang'
                    label 'linux && docker'
                    reuseNode true
                }
            }
            steps {
                sh '''
                mkdir -p $GOPATH/src/bitbucket.hylandqa.net/do
                ln -nsf $WORKSPACE $GOPATH/src/bitbucket.hylandqa.net/do/easycsr
                cd $GOPATH/src/bitbucket.hylandqa.net/do/easycsr
                make test
                '''
            }
        }
        stage ("Build") {
            agent {
                docker {
                    image 'golang'
                    label 'linux && docker'
                    reuseNode true
                }
            }
            steps {
                sh '''
                mkdir -p $GOPATH/src/bitbucket.hylandqa.net/do
                ln -nsf $WORKSPACE $GOPATH/src/bitbucket.hylandqa.net/do/easycsr
                cd $GOPATH/src/bitbucket.hylandqa.net/do/easycsr
                make build
                '''
            }
            post {
                success {
                    archiveArtifacts 'dist/*'
                }
            }
        }
    }
}