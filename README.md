$git commit -m 'add all'
On branch master
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)
Changes not staged for commit:
	modified:   json-tutorial (modified content)

方法一：
我是clone了一个文件夹下来，然后想把这个文件夹添加上去，使用了git add .，但是一直报Changes not staged for commit的错误，并且那个文件夹后面有(modified content)，解决如下：
因为是clone下来的文件，所以有.git文件，将.git文件删除就ok啦

方法二：
git commit -am "modified Cheatsheet.html"

-a 表示 add

此方法不行

$rm -rf json-tutorial/.gitignore
$rm -rf json-tutorial/.git
$git add *
$git commit -am 'add all'
问题还是未解决

mv json-tutorial/ tutorial
问题解决


https://github.com/OhBonsai/reading-notes/tree/master/writing-an-interpreter-in-Go/lang
https://www.jianshu.com/p/428d663cb2d8


