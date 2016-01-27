zpipe is a small tool to play with git objects. You can look directly into git onbjects. Why?
For “fun”, like understanding git better. Using it in git for trainings for demonstrating the simplicity of git.

## Install

Installing the single binary is a one-liner:
```
curl -L https://github.com/lalyos/zpipe/releases/download/v1.0.0/zpipe_1.0.0_$(uname)_x86_64.tgz | tar -xz -C /usr/local/bin
```

## Usage

You can mimic the functio of `git cat-file -p` as:
```
cat .git/objects/b9/9a0c706a535e72efa2feda44551a7f33b44ddb | zpipe -d
```
## tldr

```
gadd() {
  msg="$@"
  len=$(echo -en 'lofasz' | wc -c)
  echo -en "blob $len\0$@" | shasum
}
```
