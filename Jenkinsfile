pipeline {
	agent any
	stages {
		stage('SonarQube') {
			environment {
				scannerHome = tool 'GocartScanner'
			}
			steps {
				withSonarQubeEnv('sonarqube') {
					sh "${scannerHome}/bin/sonar_scanner"
				}
			}
		}
	}
}
