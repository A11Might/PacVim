# PacVim in Go

> PacVim is a game that teaches you vim commands.
> 
> You must move pacman (the green cursor) to highlight each word on the gameboard while avoiding the ghosts (in red).

## 安装

`go get github.com/A11Might/PacVim`

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
| e   | move forward to next word ending |
| $   | move to the end of the line |
| 0   | move to the beginning of the line |

目前就支持上下左右，:(

- 20211016 新增 h,j,k,l,e,$,0，注意 0,$ 可以穿墙

## 碎碎恋

参考项目 [PacVim](https://github.com/jmoon018/PacVim)，还没抄完，开发中...