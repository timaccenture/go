## RPS with Docker and Kubernetes

- run <code> docker build -t rps:local .</code> to build a docker image of the game

- make sure to point your shell to minikube's docker deamon by running <code>minikube docker-env</code> and following the given instructions

- to deploy the image with minikube run <code>kubectl apply -f deployment.yaml</code>

- check if successful by running 
    `kubectl get pods`, `kubectl get deployments` and `kubectl get svc`

- to expose the deployment ports run `kubectl expose deployment rps-deployment --type=NodePort`

- to get minikube ip run `minikube ip` and the port by running `kubectl get svc`

- test with `curl {minikube ip}:Port`


useful cli command: 
`kubectl port-forward deployment.apps/rps-deployment 8080:8080`