tags:
- #cicd
- #gitlab
- #mac

---

Install Gitlab Runner on Mac:
```sh
brew install gitlab-runner
```

Run as service:
```sh
brew services start gitlab-runner
```

Run Gitlab Job locally:
```sh
gitlab-runner exec docker job-name
```