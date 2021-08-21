tags:
- #kubernetes

---

```sh
alias k="kubectl"
alias kcx="kubectx"
alias kns="kubens"
```

```sh
echo "[[ $commands[kubectl] ]] && source <(kubectl completion zsh)" >> ~/.zshrc # add autocomplete permanently to your zsh shell
```

```sh
# Delete jobs, one by one
for job in $(k get job -o name); do kubectl delete $job; done
```
