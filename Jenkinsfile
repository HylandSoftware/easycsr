pipeline {
    agent {
        kubernetes {
            label 'easycsr-build'
            yamlFile './build-spec.yml'
        }
    }
    stages {
        stage ("Calculate Version") {
            environment {
                IGNORE_NORMALISATION_GIT_HEAD_MOVE = 1
            }
            steps {
                container('gitversion') {
                    script {
                        env.IMAGE_VERSION = sh(script: 'dotnet /app/GitVersion.dll /output json /showvariable NuGetVersionV2', returnStdout: true).trim()
                    }
                }
            }
        }
        stage ("Build") {
            steps {
                container('docker') {
                    sh 'docker build . -t hcr.io/do/easycsr:${IMAGE_VERSION}'
                }
            }
        }
        stage ("Push") {
            when {
                branch 'master'
            }
            steps {
                container('docker') {
                    withDockerRegistry([credentialsId: 'hcr-tfsbuild', url: 'https://hcr.io']) {
                        sh '''
                        docker tag hcr.io/do/easycsr:${IMAGE_VERSION} hcr.io/do/easycsr:latest
                        docker push hcr.io/do/easycsr:${IMAGE_VERSION}
                        docker push hcr.io/do/easycsr:latest
                        '''
                    }
                }
            }
        }
    }
}