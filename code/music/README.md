为了演示Go语言的面向对象编程特性，本章中设计并实现一个音乐播放器，命令行程序

关键流程
* 音乐库功能，使用者可以查看、添加和删除音乐曲目
* 播放音乐
* 支持MP3和WAV，但也能随时扩展以支持更多的音乐类型
* 退出程序

该程序在运行后进入一个循环，用于监听命令输入的状态。可以接受以下命令：
* 音乐库管理命令:lib，包括list/add/remove命令
* 播放管理：play命令，play后带歌曲名参数
* 退出程序：q命令

将程序路径添加到GOPATH
```bash
export GOPATH=XXX
```

```bash 
$ go run mplayer.go

		Enter following commands to control the player:
		lib list -- View the existing music lib
		lib add <name><artist><source><type> --Add a music to the music lib
		lib remove <id> --Remove the specified music from the lib
		play <name> -- Play the specified music
	
Enter command -> lib add gelka MJ gelka.mp3 MP3
Enter command -> play gelka
Playing mp3 music  gelka.mp3
..........
Finished playing gelka.mp3
Enter command -> lib list
1 : gelka gelka.mp3 MP3
Enter command -> lib remove 0
Enter command -> lib list
Enter command -> q
```

遗留问题：
* 多任务
音乐在播放时，不能进行其他操作。需要线程：用户界面、音乐播放和视频播放
可将play()作为一个独立的goroutine运行
* 控制播放
对上述goroutine进行控制，使用channel来实现。