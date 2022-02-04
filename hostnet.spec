apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
kind: Deployment
metadata:
  name: toolbox-agent-hostnetwork
spec:
  selector:
    matchLabels:
      app: toolbox
  replicas: 1 # tells deployment to run 2 pods matching the template
  template:
    metadata:
      labels:
        app: toolbox
    spec:
      hostNetwork: true
      containers:
        - name: toolbox-agent-hostnetwork
          image: matmerr/toolbox:v10
          env:
            - name: HTTP_PORT
              value: "8083"
          securityContext:
            privileged: true
          volumeMounts:
            - name: cni-bin
              mountPath: /opt/cni/bin
            - name: xtables-lock
              mountPath: /run/xtables.lock
            - name: log
              mountPath: /var/log
            - name: varrun
              mountPath: /tmp/var/run/
            - name: protocols
              mountPath: /tmp/etc/protocols
            - name: cnietc
              mountPath: /tmp/etc/cni
      volumes:
        - name: cni-bin
          hostPath:
            path: /opt/cni/bin
        - name: log
          hostPath:
            path: /var/log
            type: Directory
        - name: xtables-lock
          hostPath:
            path: /run/xtables.lock
            type: File
        - name: varrun
          hostPath:
            path: /var/run
            type: Directory
        - name: protocols
          hostPath:
            path: /etc/protocols
            type: File
        - name: cnietc
          hostPath:
            path: /etc/cni
            type: Directory
