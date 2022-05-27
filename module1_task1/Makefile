.PHONY: Build Clean Post setup_new_site clean_site

setup_new_site: 
	hugo new site --force ./
	hugo mod init github.com/ManojRamchandani/validation/module1_task1
	echo theme = [\"github.com/theNewDynamic/gohugo-theme-ananke\"] >> config.toml
	echo publishdir = \'./dist/\' >> config.toml
	#echo title = \"Awesome Inc.\" >> config.toml
	hugo new posts/welcome.md 
	echo title: \"Welcome to Awesome Inc.\" >> ./content/posts/welcome.md
	#echo draft: false >> ./content/posts/welcome.md


Build:
	echo 'Im building the website'
	hugo

Clean: 
	echo 'Im cleaning'
	rm -rf ./dist/*
			

clean_site:
	echo 'Im cleaning the entire site'
	rm -rf ./dist
	rm -rf ./archetypes
	rm -rf ./content
	rm -rf ./data
	rm -rf ./layouts
	rm -rf ./public
	rm -rf ./static
	rm -rf ./themes
	rm -rf ./resources
	rm -f ./config.toml
	rm -f ./go.*
	rm -f ./.hugo_build_lock
	
Post: 
	echo 'creating new post:  '$(POST_NAME).md
	hugo new posts/$(POST_NAME).md
	echo 'Adding title: '$(POST_TITLE) 'to '$(POST_NAME).md
	echo title: \"$(POST_TITLE)\" >> ./content/posts/$(POST_NAME).md
 
