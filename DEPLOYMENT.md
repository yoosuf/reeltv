# Deployment Guide

This guide covers deploying Reel TV to production environments, including both the backend (Go) and frontend (Next.js) applications.

## Prerequisites

- Docker and Docker Compose
- Kubernetes cluster (for K8s deployment)
- PostgreSQL 15+ database
- Redis 7+ cache
- S3-compatible object storage
- Domain name with SSL certificate
- Environment variables configured
- Node.js 18+ (for frontend build)

## Environment Variables

### Backend Environment Variables

Required environment variables (see `.env.example`):

```bash
# Application
APP_ENV=production
APP_PORT=8080

# Database
DB_HOST=your-db-host
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=reeltv
DB_SSLMODE=require

# Redis
REDIS_HOST=your-redis-host
REDIS_PORT=6379
REDIS_PASSWORD=your-redis-password
REDIS_DB=0

# JWT
JWT_SECRET=your-super-secret-jwt-key
JWT_ACCESS_EXPIRATION=15m
JWT_REFRESH_EXPIRATION=168h

# S3 (if using object storage)
S3_ENDPOINT=https://your-s3-endpoint
S3_ACCESS_KEY=your-access-key
S3_SECRET_KEY=your-secret-key
S3_BUCKET=reeltv-media
S3_REGION=us-east-1
S3_USE_SSL=true
```

### Frontend Environment Variables

Required environment variables for frontend (see `frontend/.env.example`):

```bash
NEXT_PUBLIC_API_URL=https://api.reeltv.com/api/v1
```

## Docker Deployment

### Backend Docker Deployment

#### Build the Application

```bash
# Build production image
docker build -t yoosuf/reeltv:latest -f deployments/Dockerfile .

# Or use the Makefile
make build
```

### Docker Compose Deployment

```bash
# Use the production compose file
docker-compose -f deployments/docker-compose.prod.yml up -d

# View logs
docker-compose -f deployments/docker-compose.prod.yml logs -f
```

### Docker Run

```bash
docker run -d \
  --name reeltv \
  -p 8080:8080 \
  --env-file .env \
  --network reeltv-network \
  yoosuf/reeltv:latest
```

### Frontend Docker Deployment

#### Build the Frontend

```bash
cd frontend

# Build production image
docker build -t yoosuf/reeltv-frontend:latest .

# Or use Docker Compose
docker build -t yoosuf/reeltv-frontend:latest -f Dockerfile .
```

#### Frontend Dockerfile

Create `frontend/Dockerfile`:

```dockerfile
# Stage 1: Build
FROM node:18-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build

# Stage 2: Runtime
FROM node:18-alpine

WORKDIR /app

COPY package*.json ./
RUN npm ci --production

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/next.config.ts ./

EXPOSE 3000

CMD ["npm", "start"]
```

#### Frontend Docker Compose

```bash
docker run -d \
  --name reeltv-frontend \
  -p 3000:3000 \
  --env-file .env.local \
  yoosuf/reeltv-frontend:latest
```

## Kubernetes Deployment

### Create ConfigMap

```bash
kubectl create configmap reeltv-config \
  --from-env-file=.env \
  --namespace=reeltv
```

### Create Secrets

```bash
kubectl create secret generic reeltv-secrets \
  --from-literal=db-password=your-db-password \
  --from-literal=jwt-secret=your-jwt-secret \
  --from-literal=redis-password=your-redis-password \
  --namespace=reeltv
```

### Deploy

```bash
# Apply Kubernetes manifests
kubectl apply -f deployments/k8s/
```

### Example Kubernetes Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: reeltv
  namespace: reeltv
spec:
  replicas: 3
  selector:
    matchLabels:
      app: reeltv
  template:
    metadata:
      labels:
        app: reeltv
    spec:
      containers:
      - name: reeltv
        image: yoosuf/reeltv:latest
        ports:
        - containerPort: 8080
        env:
        - name: DB_HOST
          valueFrom:
            configMapKeyRef:
              name: reeltv-config
              key: DB_HOST
        - name: DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: reeltv-secrets
              key: db-password
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
---
apiVersion: v1
kind: Service
metadata:
  name: reeltv
  namespace: reeltv
spec:
  selector:
    app: reeltv
  ports:
  - port: 80
    targetPort: 8080
  type: LoadBalancer
```

## Database Setup

### Run Migrations

```bash
# Using Docker
docker run --rm \
  -e DB_HOST=your-db-host \
  -e DB_USER=your-db-user \
  -e DB_PASSWORD=your-db-password \
  -e DB_NAME=reeltv \
  yoosuf/reeltv:latest \
  /app migrate up

# Or using the application
make migrate-up
```

### Seed Reference Data

```bash
docker run --rm \
  -e DB_HOST=your-db-host \
  -e DB_USER=your-db-user \
  -e DB_PASSWORD=your-db-password \
  -e DB_NAME=reeltv \
  yoosuf/reeltv:latest \
  /app seed
```

## SSL/TLS Configuration

### Using Nginx Reverse Proxy

```nginx
server {
    listen 443 ssl http2;
    server_name api.reeltv.com;

    ssl_certificate /etc/nginx/ssl/fullchain.pem;
    ssl_certificate_key /etc/nginx/ssl/privkey.pem;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}

server {
    listen 80;
    server_name api.reeltv.com;
    return 301 https://$server_name$request_uri;
}
```

### Using Traefik

Add labels to your Docker container:
```yaml
labels:
  - "traefik.enable=true"
  - "traefik.http.routers.reeltv.rule=Host(`api.reeltv.com`)"
  - "traefik.http.routers.reeltv.tls=true"
  - "traefik.http.routers.reeltv.tls.certresolver=letsencrypt"
```

## Monitoring

### Health Check

```bash
curl https://api.reeltv.com/health
```

### Logs

```bash
# Docker logs
docker logs -f reeltv

# Kubernetes logs
kubectl logs -f deployment/reeltv -n reeltv
```

### Metrics

Consider integrating with:
- Prometheus for metrics collection
- Grafana for visualization
- ELK Stack for log aggregation

## Backup Strategy

### Database Backups

```bash
# Automated daily backup
0 2 * * * pg_dump -h $DB_HOST -U $DB_USER $DB_NAME | gzip > /backup/reeltv-$(date +\%Y\%m\%d).sql.gz
```

### Redis Backup

```bash
# Redis RDB backup
redis-cli --rdb /backup/dump.rdb
```

## Scaling

### Horizontal Scaling

- Increase replicas in Kubernetes
- Use Docker Swarm mode
- Deploy behind load balancer

### Vertical Scaling

- Increase CPU/memory limits
- Optimize database queries
- Add Redis caching

## Rollback Strategy

### Docker Rollback

```bash
# Stop current version
docker stop reeltv

# Start previous version
docker run -d --name reeltv \
  -p 8080:8080 \
  --env-file .env \
  yoosuf/reeltv:v1.0.0
```

### Kubernetes Rollback

```bash
kubectl rollout undo deployment/reeltv -n reeltv
```

### Database Rollback

```bash
# Rollback migration
docker run --rm \
  -e DB_HOST=your-db-host \
  -e DB_USER=your-db-user \
  -e DB_PASSWORD=your-db-password \
  -e DB_NAME=reeltv \
  yoosuf/reeltv:latest \
  /app migrate down
```

## Troubleshooting

### Container Won't Start

```bash
# Check logs
docker logs reeltv

# Check environment variables
docker inspect reeltv

# Test database connectivity
docker exec reeltv ping -c 1 $DB_HOST
```

### Database Connection Issues

```bash
# Test from container
docker exec reeltv psql -h $DB_HOST -U $DB_USER -d $DB_NAME

# Check network
docker network inspect reeltv-network
```

### High Memory Usage

```bash
# Check container stats
docker stats reeltv

# Reduce connections in connection pool
# Add more replicas instead of increasing container size
```

## Security Checklist

- [ ] Change default JWT secret
- [ ] Use strong database passwords
- [ ] Enable SSL/TLS
- [ ] Configure firewall rules
- [ ] Enable rate limiting
- [ ] Set up log aggregation
- [ ] Configure backup automation
- [ ] Enable audit logging
- [ ] Regular security updates
- [ ] Review access controls

## Post-Deployment

1. Verify health check endpoint
2. Test authentication flow
3. Test critical API endpoints
4. Monitor error rates
5. Check database connection pool
6. Verify Redis connectivity
7. Test file uploads (if applicable)
8. Load test the application
9. Set up monitoring alerts
10. Document any issues

## Support

For deployment issues, see [SUPPORT.md](SUPPORT.md).
