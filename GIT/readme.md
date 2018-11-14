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


### 版本库
Git 主要由两个层面构成，其顶层结构就是我们所用的命令，例如log，reset commit等。这些命令使用起来很方便，也有很多可调用的选项，git的开发者们称之为**瓷质命令**(porcelain command)<br>
其底层结构，被称之为**管道**(plumbing)。他们是一组带有少量选项的简单命令，瓷质命令就是以此为基础被构建出来的。
- git 的核心是一个对象数据库<br>
    该对象数据库可用来存储文本或者二进制，存储后，会返回一个40个字符的代码，这个是被存储对象的键值。<br>
    git 对象数据库是一个非常高效的实现，对于一些大型项目(例如 linux 有两万次提交和两百万个对象)，访问操作可以几乎瞬间完成。**其性能瓶颈只有在总数据量非常巨大的时候才显现出来。**
    ```
    # git hash-object -w 表示将该文件作为一条记录插入对象服务器中
    git hash-object -w readme.md

    # 使用cat-file 来访问该对象 -p 表示打印
    git cat-file -p 124ad2
    ```
- 存储目录：Blob与Tree<br>
    在文件和目录存储上，git使用一种包含两种节点的简单树结构-blob对象与tree对象。blob即文件对象，tree对象包含了文件和子目录，包含访问权限，类型以及包含目录下文件的内容生成的散列值。
    ```
    # 查看
    git cat-file -p master^{tree}
    ```
- 相同的数据只存储一次<br>
    为了节省存储空间，git对于相同的数据将只存储一次。
    ```
    # 复制 两份一样的文件
    cp test1.txt test2.txt

    # 添加至 对象数据库
    git hash-object -w test1.txt
    git hash-object -w test2.txt

    # 发现返回的散列值是相同的，同时在git中，比较操作等都是比较相关的散列值，而不看其实际数据。
    ```
- git gc <br>
    git 对于相似内容还可以采取增量方法来存储。一般来说，使用gc命令，对多余的提交等进行删除和压缩。<br>
    一般来说,git 会定时运行 `auto gc`命令，定期清理压缩。
- git 散列冲突 <br>
    git 散列有2^160种，虽然理论上SHA1算法是有缺陷的，但是概率极低，可以认为其为安全的。
- git 提交对象 <br>
     `git cat-file -p hash` git的提交对象也存储在对象数据库中，提交对象包括作者，注释等信息，还有tree 对象表示提交内容，parent对象，指向其上一次提交。
- 提交对象重用 <br>
    即对于一次提交中，对于没有发生改变的修改, 则重用之前的对象
- 重命名 移动 复制 <br>
    在git中，重命名 移动 复制操作等操作会首先采用一个重命名检测算法，检测该文件与添加的文件是否有相同或者相似的地方，如果有的话，则重用该对象。
    ```
    # 获取一份摘要 -M激活检测算法
    git log --summary -M90%
    # 跟踪被移动文件历史
    git log --follow test.txt
    # 跟踪被复制的数据
    git log --summary -C90%
    # 打印源头信息
    git blame -M -C -C -C test.txt
    ```