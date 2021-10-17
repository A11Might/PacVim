# PacVim in Go

> PacVim is a game that teaches you vim commands.
> 
> You must move pacman (the green cursor) to highlight each word on the gameboard while avoiding the ghosts (in red).

## 安装

- 新用户

`go get github.com/A11Might/PacVim`

- 老用户

`go get github.com/A11Might/PacVim@v0.0.4`

## 使用

键入 `pacvim` 来开始游戏

## 游玩方法

> The objective of PacVim is very similar to PacMan.
> 
> You must run over all the characters on the screen while avoiding the ghosts (red `G`).
> 
> PacVim has two special obstacles:
>
> 1. You cannot move into the walls (yellow color).  You must use vim motions to jump over them.
>
> 2. If you step on a tilde character (cyan `~`), you lose!

<h4>List of Implemented Commands</h4>

| key | what it does |
| --- | --- |
| h   | move left |
| j   | move down |
| k   | move up |
| l   | move right |
| w   | move forward to next word beginning |
| e   | move forward to next word ending |
| b   | move backward to next word beginning |
| $   | move to the end of the line |
| 0   | move to the beginning of the line |
| gg  | move to the beginning of the first line |
| G   | move to the beginning of the last line |

## 开发日志

- 20211015 地图上的玩家（P）可以动啦

- 20211016 
  - 地图上的幽灵（G）可以动啦
  - 新增 h,j,k,l,w,e,b,$,0,gg,G，注意 0,$ 可以穿墙
  - 走过所有字符可以赢得游戏啦

- 20211017 
  - 重构部分代码，看起来更面向对象
  - 今天一个人吃海底捞，难过:(

## 碎碎恋

参考项目 [PacVim](https://github.com/jmoon018/PacVim)，还没抄完，开发中...

发现代码有明显逻辑`bug`，但它就是能正常运行，我:)