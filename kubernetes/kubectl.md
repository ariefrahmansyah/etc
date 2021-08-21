tags:
- #kubernetes

---

```
alias k="kubectl"
alias kcx="kubectx"
alias kns="kubens"
```

```
echo "[[ $commands[kubectl] ]] && source <(kubectl completion zsh)" >> ~/.zshrc # add autocomplete permanently to your zsh shell
```

```
# Delete jobs, one by one
for job in $(k get job -o name); do kubectl delete $job; done
```
