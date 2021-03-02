https://blog.logrocket.com/creating-a-web-server-with-golang/

========================	
	Local Commands		\
=========================================================
go build
	update the application

docker build -t go-kubernetes .
	update the images
	
docker tag go-kubernetes rturnbow/prj01-website:1.0.0
	Tag the image

------------------------
=> Update docker image
go build
docker build -t prj01-website .
docker tag prj01-website rturnbow/prj01-website:1.0.0

========================	
	Kubernetes			\
=========================================================
kubectl apply -f k8s-deployment.yml
	update the K8s deployment/service

------------------------
=> Set context / namespace
kubectl config use-context project01
kubectl config use-context docker-desktop
		
------------------------
=> Update the deployment	
kubectl delete service -n project01 prj01-website-service
kubectl delete deployment -n project01 prj01-website
kubectl apply -f k8s-deployment.yml