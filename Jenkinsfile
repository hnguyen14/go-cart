pipeline {
	agent any
	stages {
		stage('SonarQube') {
			environment {
				scannerHome = tool 'GocartScanner'
			}
			steps {
				withSonarQubeEnv('LocalSonarQube') {
					sh "${scannerHome}/bin/sonar_scanner"
				}
			}
		}
	}
}
