#java 
#mac

## Java 11

```
brew install java
```

```
sudo ln -sfn /usr/local/opt/openjdk@11/libexec/openjdk.jdk /Library/Java/JavaVirtualMachines/openjdk-11.jdk
```

## Java 8

```
brew install openjdk@8
```

```
sudo ln -sfn /usr/local/opt/openjdk@8/libexec/openjdk.jdk /Library/Java/JavaVirtualMachines/openjdk-8.jdk
```

## Switch between JDK versions

In `~/.zshrc` or `~/.bashrc`:

```
export PATH="/usr/local/opt/openjdk@8/bin:$PATH"'
export JAVA_HOME=/Library/Java/JavaVirtualMachines/openjdk-8.jdk/Contents/Home

jdk() {
      version=$1
      unset JAVA_HOME;
      export JAVA_HOME=$(/usr/libexec/java_home -v"$version");
      java -version
}
```