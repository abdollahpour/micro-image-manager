Kubernetes has all the things you want to run `micro-image-manager` at scale. You can create a file from the following template and name it `micro-image-manager.yml`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: micro-image-manager
  labels:
    app: micro-image-manager
spec:
  replicas: 3 # Choose the number you need
  selector:
    matchLabels:
      app: micro-image-manager
  template:
    metadata:
      labels:
        app: micro-image-manager
    spec:
      containers:
        - name: micro-image-manager
          image: abdollahpour/micro-image-manager:1.0
          ports:
            - name: http
              containerPort: 8080
          volumeMounts:
            - name: data
              mountPath: /mnt/images
      volumes:
        - name: data
          persistentVolumeClaim:
            claimName: micro-image-manager-data
---
apiVersion: v1
kind: Service
metadata:
  name: micro-image-manager
spec:
  selector:
    app: micro-image-manager
  ports:
    - protocol: TCP
      port: 80
      targetPort: http
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: micro-image-manager
  # For SSL you need cert-manager and letsencrypt-prod ClusterIssuer
  # annotations:
  #  cert-manager.io/cluster-issuer: letsencrypt-prod
spec:
  # tls:
  #  - hosts:
  #      - your-domain.com
  #    secretName: micro-pdf-generator-cert
  rules:
    - host: your-domain.com
      http:
        paths:
          # FOR TESTS ONLY
          # You should not expose API endpoint publicly
          # This API is for other microservices to write image
          # - path: /api/v1/images
          #  pathType: Exact
          #  backend:
          #    service:
          #      name: micro-image-manager
          #      port:
          #        number: 80
          - path: /images
            pathType: Prefix
            backend:
              service:
                name: micro-image-manager
                port:
                  number: 80
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: micro-image-manager-data
spec:
  storageClassName: "" # Default
  volumeName: your-persistant-volume
```

Please update:

- `your-domain.com` with your domain
- `your-persistant-volume` with persistant volume that you want to use (you can read more about persistant volume here)

Then run: `kubectl apply -f micro-image-manager.yml`
