# rename-projects
rename ugly all caps projects

Changes 
```
├── UI.COMMON
├── WEBUI.COMMON
├── WEBUI.APP
├── WEBUI.APP.FACADE
├── WEBUI.APP.MOBILE
├── WEBUI.APP.MOBILE.FACADE
```
into 
```
├── ui-common
├── webui-common
├── webui-app
├── webui-app-facade
├── webui-app-mobile
└── webui-app-mobile-facade
```

And searches the tree for `build.gradle` files and updates them to the new name.
