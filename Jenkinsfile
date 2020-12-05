pipeline {
	agent any
	stages {
		stage('SonarQube') {
			environment {
				SCANNER_HOME = tool "GocartScanner"
				PROJECT_NAME = "GocartSonar"
			}
			steps {
				withSonarQubeEnv('LocalSonarQube') {
					sh "$SCANNER_HOME/bin/sonar-scanner -Dsonar.projectKey=$PROJECT_NAME"
				}
			}
		}
	}
}
