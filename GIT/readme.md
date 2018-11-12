## git学习指南 
[书籍链接](!https://item.jd.com/12023485.html)

### 提交究竟是什么？
- 三个概念
    - 版本库：通常留驻在.git目录，每一份clone都保留所有的版本信息。
    - 提交：每一次commit都称为一次提交，每次提交都包含时间，作者，修改信息等。这些关系会构成项目版本图，使用git log 可以显示出来。
    - 提交散列值：用于表示标识唯一一次提交，是根据修改信息和时间/作者等，可用来校验软件存储的完整性。通常一次提交长度为40个字符。

- git diff 提交之间的差异<br>
```
# 查看最近的三次提交记录
git log -n 3

# 比较两次记录 (只需要写部分commitID)
git diff 55e73 b51c9

# 与上一次提交比较
git diff 55e73^!

# 限制比较 某些文件或者目录差异
git diff 55e73^! -- leetcode/readme.MD

# 统计修改情况
git diff --stat 55e73^!
```
- git log 查看提交历史<br>
```
# 查看最近几次提交历史
git log -n 3

# 每次历史仅展示一行信息
git log --oneline

# 查看提交统计信息
# stat 显示被修改的文件
git log --stat
# dirstat 显示包含被修改文件的目录
git log --dirstat
# shortstat 显示项目中被修改 以及 新增删除 的文件
git log --shortstat

# 显示各个提交之间的关系
git log --graph --oneline
```

### 多次提交
- 三个概念
    - 工作区：当前版本的本地修改
    - 暂存区：本地修改后，git add 将本地修改由工作区添加至暂存区
    - 版本库：暂存区的修改，git commit 将暂存文件由暂存区提交到版本库
- git status<br> 查看当前工作区所发生的修改 
- git diff 
    - git diff --staged <br> 当前版本库与暂存区之间的区别
    - git diff <br> 暂存区与工作区区别
- git reset HEAD . <br> 重置某些目录的暂存区 
- git stash 储藏
    - git stash <br> 将工作区和暂存区中修改保存到 存储栈
    - git stash pop <br> 恢复栈顶的储藏修改
    - git stash lish <br> 查看堆栈中的修改内容
    - git stash pop stash@{1} <br> 恢复某一次修改
