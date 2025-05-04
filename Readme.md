# Squater Notes
Squatcher is a TUI for keeping and managing notes using neovim and Github.

The idea is that you keep a main directory of all your notes and folders in a github repo
and use Squatcher to edit and create those notes. 

My dream is to support a markdown reading mode so the md will render in the TUI
for more basic reading locally (if you're not into that sweet sweet raw markdown).

## Theoretical Commands
:Squatcher toggle - Squatch Toggle On/Off
:Squatcher on - Pull up TUI
:Squatcher off - Close TUI
:Squatcher sync - Sync Push/Pull notes with Github
:Squatcher delete - Delete Note or Directory
:Squatcher push - Push to github
:Squatcher pull - Pull from github
:Squatcher tmd - Toggle Markdown


### Nice to have features for my TODOs:
- Easy push/pull via github API when you save and quit or open squatcher (probably a preference in config)
- UI to update local address for notes repo
- Toggling between many repos for note bases
- Popup buffer for quick notes
- Sidebar for changing between folders/notes



⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣾⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣴⣿⣿⣿⣿⣿⣿⣆⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢀⣤⣶⣶⣶⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣟⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠋⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⣾⣿⣿⣿⣿⠛⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⢰⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀
⠀⣾⣿⣿⣿⡟⠁⠈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀
⢀⣿⣿⣿⡟⠀⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁⠙⠛⠿⣿⣿⣿⣷⡄⠀⠀⠀⠀
⢸⣿⣿⣿⣷⡄⠀⠹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⡀⠀⠀⠈⢻⣿⣿⣿⣦⠀⠀⠀
⠘⣿⣿⣿⠙⠇⠀⠀⠘⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣆⠀⠀⠀⢿⠉⢹⣿⡇⠀⠀
⠀⠀⠙⠛⠛⠂⠀⠀⠀⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡄⠀⠀⠀⠸⠛⠁⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⣿⣿⣿⣿⡿⢿⣿⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣠⣤⣶⣿⣿⣿⣿⣿⠃⠀⠈⢻⣿⣿⣿⣿⡄⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⡿⠋⠀⠀⠀⢸⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣠⣿⣿⣿⣿⣿⣿⡿⠋⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀
⠀⢀⣤⣾⣿⣿⣿⣿⡿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⢿⣿⣿⣿⣿⡆⠀⠀⠀⠀⠀
⣴⣿⣿⣿⣿⣿⠟⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢻⣿⣿⣿⣧⣤⣤⣀⣠⣤
⠻⣿⣿⣿⣿⣿⣿⣦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⠏
⠀⠀⠀⠀⠙⠛⠿⠿⠛⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠻⣿⣿⠿⠟⠛⠛⠉⠀
