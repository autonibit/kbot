pipeline {
    agent any

    parameters {
        choice(name: 'OS', choices: ['linux', 'macos', 'windows'], description: 'Pick OS')
        choice(name: 'ARCH', choices: ['amd64','arm', 'arm64'], description: 'Pick ARCH')
    }

    environment {
	GITHUB_TOKEN=credentials('github-token')
	REPO = 'https://github.com/autonibit/kbot'
	REGISTRY = 'ghcr.io/autonibit'
	BRANCH = 'develop'
	TARGETOS = "${params.OS}"
	TARGETARCH = "${params.ARCH}"
    }

    stages {

        stage('clone') {
            steps {
                echo 'Clone Repository'
                git branch: "${BRANCH}", url: "${REPO}"
            }
        }

        stage('test') {
            steps {
                echo 'Testing started'
                sh "make test"
            }
        }

        stage('build') {
            steps {
                echo "Building binary for platform ${params.OS} on ${params.ARCH} started"
                sh "make ${params.OS} ${params.ARCH}"
            }
        }

        stage('image') {
            steps {
                echo "Building image for platform ${params.OS} on ${params.ARCH} started"
                sh "make image-${params.OS} ${params.ARCH}"
            }
        } 

	stage('login to GHCR') {
            steps {
                sh "echo $GITHUB_TOKEN_PSW | docker login ghcr.io -u $GITHUB_TOKEN_USR --password-stdin"
            }
        }

	stage('push image') {
            steps {
                sh "make push"
            }
        }
    }
}
