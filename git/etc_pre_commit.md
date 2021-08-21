tags:
- #git

---

Copy this snippet to .git/hooks/pre-commit:

```sh
#!/bin/sh

./.helper/readme-generator/readme-generator
git add README.md
```
