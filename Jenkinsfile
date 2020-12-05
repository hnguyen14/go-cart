pipeline {
	agent {
		docker {
			image '1.15.6-alpine3.12'
		}
	}
	stages {
		stage('Build') {
			steps {
				sh 'make docker.build'
			}
		}
	}
}
