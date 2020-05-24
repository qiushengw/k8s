kubectl run curl --image=yauritux/busybox-curl -i --rm --restart=Never -- curl -v http://weather-service:8080


