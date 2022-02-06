/* groovylint-disable CompileStatic */
CONTAINER_NAME = 'go-http-example'

pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh "docker build -t $CONTAINER_NAME -f Dockerfile.test ."
            }
        }

        stage('Test') {
            steps {
                sh """
                docker run --name $CONTAINER_NAME $CONTAINER_NAME sh -c 'go test -coverprofile=cover.out -v ./... | go2xunit; gocover-cobertura < cover.out > coverage.xml' > tests.xml
                docker cp $CONTAINER_NAME:/go/src/github.com/jamiefdhurst/go-http-example/coverage.xml coverage.xml
                """
                junit 'tests.xml'
                step([$class: 'CoberturaPublisher', coberturaReportFile: 'coverage.xml'])
            }
        }
    }

    post {
        always {
            sh """
            docker stop $CONTAINER_NAME
            docker rm $CONTAINER_NAME
            """
        }
    }
}
