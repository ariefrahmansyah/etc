tags:
- #java
- #mac

---

## Java 11

```sh
brew install java
```

```sh
sudo ln -sfn /usr/local/opt/openjdk@11/libexec/openjdk.jdk /Library/Java/JavaVirtualMachines/openjdk-11.jdk
```

## Java 8

```sh
brew install openjdk@8
```

```sh
sudo ln -sfn /usr/local/opt/openjdk@8/libexec/openjdk.jdk /Library/Java/JavaVirtualMachines/openjdk-8.jdk
```

## Switch between JDK versions

In `~/.zshrc` or `~/.bashrc`:

```sh
export PATH="/usr/local/opt/openjdk@8/bin:$PATH"'
export JAVA_HOME=/Library/Java/JavaVirtualMachines/openjdk-8.jdk/Contents/Home

jdk() {
      version=$1
      unset JAVA_HOME;
      export JAVA_HOME=$(/usr/libexec/java_home -v"$version");
      java -version
}
```
