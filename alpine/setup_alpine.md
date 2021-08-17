tags:
- #gcloud 
- #python

---

```
apk add bash

apk add curl

apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python

# Install gcloud
curl -sSL https://sdk.cloud.google.com | bash
```