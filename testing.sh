
#!/bin/sh
curl -vvv -H "Content-Type: application/json" \
    -H "X-GitHub-Event: deployment_status" \
    -H "X-GitHub-Delivery: `uuid -v4`" \
    http://localhost:8080/ \
    -d "@payload.json"
