# Service Environment Configuration

This service requires a few environment variables to be set before it can run.  
You should place them in a `.env` file inside the service directory.

---

## 📌 Required Environment Variables

### **1. PORT**
Defines the port the service will listen on.

```env
PORT=":4445"

```DB_URL
DB_URL="postgres://postgres:123456@host.docker.internal/t_hex?sslmode=disable"

```CACHE_TTL
CACHE_TTL="5"

```SELENIUM_HUB_URL
SELENIUM_HUB_URL="http://selenium-hub:4444"
