# FinBank: FinTech Microservices Project

FinBank is a realistic FinTech backend architecture implemented using Go. It demonstrates core industry concepts like microservices, Kafka messaging, PostgreSQL, Redis, Kubernetes, Prometheus/Grafana monitoring, and more.

---

## ✅ Use Case
- Handle **1 million API requests**
- Store **10 million records** efficiently
- Use **goroutines & channels** for concurrency
- Detect fraud using a **Kafka-based pipeline**
- Provide alerts and monitoring through **Prometheus + Grafana**

---

## 🧱 Project Structure
```
finbank/
├── transaction-service/         # Handles transaction creation
├── fraud-service/               # Detects fraud via Kafka
├── kafka/topics-config/         # Kafka topic definition
├── deployments/
│   ├── k8s/                     # Kubernetes YAML files
│   └── helm/                    # Helm charts
├── monitoring/                  # Prometheus & Grafana configs
└── README.md                    # Documentation
```

---

## 🔧 Prerequisites
- Go 1.21+
- Docker + Docker Compose
- Kubernetes (minikube, kind, or cloud)
- Helm

---

## 🚀 How to Run

### 1. **Clone the Repo**
```bash
git clone https://github.com/yourname/finbank.git
cd finbank
```

### 2. **Build Go Services**
```bash
cd transaction-service
go build -o transaction-service

cd ../fraud-service
go build -o fraud-service
```

### 3. **Run Locally (Optional)**
```bash
# Start DB & Kafka (use docker-compose or manual)
./transaction-service
./fraud-service
```

### 4. **Deploy to Kubernetes**
```bash
kubectl apply -f deployments/k8s/
```

### 5. **Install via Helm**
```bash
cd deployments/helm
helm install finbank .
```

---

## 📊 Monitoring

### Prometheus Target:
```
http://<prometheus-ip>:9090/targets
```

### Grafana Dashboard:
```
Import dashboard.json from monitoring/grafana/ into Grafana
```

---

## 🧪 Testing & Benchmarking
```bash
cd transaction-service/tests
go test -v

go test -bench=. -benchmem
```

```bash
cd fraud-service/tests
go test -v
```

---

## 🔐 Security & Scaling
- Load balancing via Kubernetes Service
- Autoscaling with HPA YAMLs
- Secrets and ConfigMaps can be added for DB/Kafka

---

## 📂 Kafka Topics
- `transactions`: New transaction events

---

## 📬 Contributing
Feel free to raise an issue or create a PR if you'd like to add modules like JWT auth, GraphQL APIs, or gRPC.

---

## 📃 License
MIT License. © 2025 FinBank Contributors
