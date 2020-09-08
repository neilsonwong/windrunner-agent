# Windrunner Agent
A desktop program written in golang to open remotely stored (video) files on your desktop.

## Purpose
The program was written primarily to facilitate the opening on video files from a web page. My original intent was to utilize a custom protocol (linked from the [windrunner website](https://github.com/neilsonwong/windrunner-ui)) to facilitate this. However, the issue with custom protocols involves browser security and while (I believe) is the more elegant solution, would not work well at the time of creation (2018-2020). Thus the windrunner agent was born. This application allows the launching of files with a defined path exposed through a web api. The design currently expects to open files from a relative path on a mounted folder.

## Architecture
The agent primarily integrates with the windrunner ui through web calls (when videos are launched). It is cross platform (supporting windows, macOS, and linux that has xdg-open). Currently, the application will handle the mounting of the shared drive (if not already mounted).
The agent also functions as a local proxy (for the web ui) to proxy server calls to. Due to the agent will have network access to the [windrunner server](https://github.com/neilsonwong/windrunner), a function was added into the agent to proxy web calls to the server. This helps save on web traffic (since it becomes all local). Https is maintained as our address is bound to the host address (127.0.0.1). Thus on any funky host ip configurations where 127.0.0.1 is not resolvable, the proxy functions won't maintain https. 

### As a service
Originally, this made use of the [service package](https://github.com/kardianos/service) by [kardianos](https://github.com/kardianos), I was never able to get it to work consistently enough for each platform. In the end, in order to achieve the likeness of a true service, I have tweaked the implementation (via the installation method) for each platform.

On windows, the desired functionality can be achieved using a startup item (much simpler).
On macOS, launchctl is used along with a plist file.
On linux, a user service created and setup.

### Configuration
Configuration is done through a config.json file. The details on configuration are below.  
*Comments inserted for clarity*

```javascript
{
    // version number
    "version": "2.0.0",

    // SAMBA Configuration
    // name of the samba share
    "share_server": "MY_SHARE",
    // ip of the samba share
    "share_server_addr": "192.168.0.123",
    // folder being shared
    "share_folder": "VIDEOS",
    // osx mount folder
	"osx_mount": "/my/mount/point",

    // Proxy Server Configuration
    // address of the windrunner server
	"listing_server": "http://192.168.0.123:9876",
    // proxy all calls that match the following prefix
    "proxy_prefix": "/proxy",
    // agent api port
	"server_port": 8080
}
```

## API Offerings

- Play File: `POST /api/play `  
**Data**: file: path of file relative to share root  
*Plays the video file*

- Heartbeat: `GET /api/doki`  
*Responds with a 200* **ドキドキ**

- Get Agent Config: `GET /api/config`  
*Retrieves the config loaded by the agent*  

- Update Agent Config: `PUT /api/config`  
**Data**: updated config value  
*Update the config.json file and reload it, if the updated configuration is not valid, an error will be returned*  

- Proxy Calls: `GET|POST|PUT|DELETE /proxy/**`  
**Data**: any if required by the original call  
*Proxies the web request to the configured windrunner server*  

## Installation
Install scripts are bundled with each of the packages.  
For Windows, a `.bat` script is bundled. This extracts the files to the current directory and adds the extracted windrunner.exe to the startup.  
For macOS and linux, a `.sh` script is bundled. This should install to the `/opt/Windrunner` directory. A service will be created in `launchctl/systemctl`.  
After installing, the agent should startup right away.

## How to get the project up and running
Ensure that go is installed. Since this is setup as a go module, running `go build` or `go test` should automatically pull all the dependencies.

To start the project, you can then run `go run windrunner.go`.

*Ensure that this project is in your GOPATH, such that recursive references to other parts of the project will be automatically resolved!*

### How to build releases
Building releases can be done by using the `packageAll.sh` script. This will run the `buildAll.sh` to compile each of the different platform/architecture combinations then package the appropriate install scripts as well.

*the versioning of the output is controlled by the version variable within the `packageAll.sh` script*

## Future Enhancements
- Auto updates
- API Security