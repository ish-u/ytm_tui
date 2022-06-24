# ytm

A TUI to play music from YouTube made using [bubbletea](https://github.com/charmbracelet/bubbletea), [beep](https://github.com/faiface/beep) and [youtube-dl golang](https://github.com/kkdai/youtube)

### Run
- Install GoLang https://go.dev/doc/install

- Clone this repository
    ```bash
    git clone https://github.com/ish-u/ytm_tui.git
    ```

- Install Required Go Packages
    ```
    go mod tidy 
    ```

- Run
    ```
    go run main.go
    ```

### Demo

![](ytm.gif)

### Issues
- High Memory Usage - b/c YouTube doesn't provide streams in `mp3` format (which is a requirement of [bubbletea](https://github.com/charmbracelet/bubbletea), [beep](https://github.com/faiface/beep) we have to convert `m4a` stream from YouTube using [ffmpeg](https://github.com/modfy/go-fluent-ffmpeg). 
- `audio` folder holds the `mp3` files that are downloaded during streaming which are not deleted after the qutting the application.
- _Looks kinda ugly_

### Things I want to Improve/Add
- A Audio Visualizer - using the samples read during the streaming by [beep](https://github.com/faiface/beep).
- Ability to directly stream without the need to save the stream in a file and then playing it i.e. play the directly play the converted stream from [ffmpeg](https://github.com/modfy/go-fluent-ffmpeg).
- Improve the look of Queue using [lipgloss](https://github.com/charmbracelet/lipgloss).   


###### _have an awesome day stranger_.