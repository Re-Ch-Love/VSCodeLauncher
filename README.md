# VSCode Launcher

## 使用方式

1. 进入Visual Studio Code的安装目录，将`Code.exe`改成其它名字，如`_Code.exe`

2. 从release中下载或自行编译该软件，把可执行文件（`Code.exe`/`Code`）放在安装目录下。

3. 在家目录下新建.vscode_launcher.conf文件，第一行写改好名的Code.exe的路径，第二行写启动参数。
    如
    ```
    D:\Microsoft VS Code\_Code.exe
    --extensions-dir "E:\vscode_extensions"
    ```

## 背景

在使用vscode时，发现存放插件的文件太大了，而且又默认放在C盘，所以我想把插件文件夹迁移到其它盘。查阅资料后我发现，vscode不能直接设置，需要在启动时添加参数才行。要是直接从桌面上启动倒还好，可以给快捷方式加启动参数，但是如果是点击文件启动vscode这种就麻烦了，vscode会去默认文件夹加载插件，结果啥都加载不到。为了解决这个问题，便有了这个软件。

## Q&A

### 图标是如何添加的？

先写Code.exe.manifest文件，再放一个icon.ico文件，然后依次执行以下命令。

`go get github.com/akavel/rsrc`

`rsrc -manifest Code.exe.manifest -ico ico.ico -o Code.syso`

### 为什么不做成通用启动器？

因为这样，应用程序就没有图标了（大概吧，没有试过，如果可以记得踹我或者发PR）。

### 如何编译？

执行以下命令

`go build -ldflags "-w -s -H windowsgui" -o Code.exe`

### 如何排查错误

家目录下的`.vscode_launcher.log`文件