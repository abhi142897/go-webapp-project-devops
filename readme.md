# DevOpsified Project

Welcome to the **DevOpsified Project**! This is a simple Go application that serves static content on port **8080**. We will completely devospify the application.

## Features

- Serves a welcome message at the endpoint `/home`.
- Listens on port **8080**.

## Getting Started

### Prerequisites

- Go **1.21** or later installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

### Installation
Make sure go is installed.

Clone the repository -> cd go-webapp-project -> run #go run main.go -> access the application using http://localhost:8080/home


Devopsified version of go-webapp-project: 

Steps to Devopsify a project -

1) Containerize the application using Dockerfile.
2) Build, test and push image to docker hub.
3) Write deployment and service manifests to deploy the application on k8s cluster.
4) Create helm chart with k8s manifests.
5) Create .github/workflows and write ci-cd.yaml for implmenting ci.
6) Install argocd and add a new application.
7) Enable auto trigger in argocd.