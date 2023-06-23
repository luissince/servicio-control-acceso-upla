mkdir logs

docker stop servicio-control-acceso && docker rm servicio-control-acceso

docker image rm servicio-control-acceso

docker build -t servicio-control-acceso .

docker run -d \
--restart always \
--name servicio-control-acceso \
--net=upla \
-p 8893:80 \
-v $(pwd)/logs:/etc/logs \
servicio-app-movil