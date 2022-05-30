.DEFAULT_GOAL := help
.PHONY: build clean post help setup_new_site setup_dummy_dirs clean_site 

# setup_new_site: ## Creates a new hugo site titled "Awesome Inc." & a post using the ananke theme
# 	hugo new site --force ./
# 	hugo mod init github.com/ManojRamchandani/validation/module1_task1
# 	echo theme = [\"github.com/theNewDynamic/gohugo-theme-ananke\"] >> config.toml
# 	echo publishdir = \'./dist/\' >> config.toml
# 	#echo title = \"Awesome Inc.\" >> config.toml
# 	hugo new posts/welcome.md 
# 	echo title: \"Welcome to Awesome Inc.\" >> ./content/posts/welcome.md
# 	#echo draft: false >> ./content/posts/welcome.md

# setup_dummy_dirs: setup_new_site
# 	touch ./archetypes/foo.txt
# 	echo "Dummy text" >> ./archetypes/foo.txt

# 	touch ./content/foo.txt
# 	echo "Dummy text" >> ./content/foo.txt
	
# 	touch ./data/foo.txt
# 	echo "Dummy text" >> ./data/foo.txt
	
# 	touch ./layouts/foo.txt
# 	echo "Dummy text" >> ./layouts/foo.txt
	
# 	touch ./public/foo.txt
# 	echo "Dummy text" >> ./public/foo.txt
	
# 	touch ./static/foo.txt
# 	echo "Dummy text" >> ./static/foo.txt
	
# 	touch ./themes/foo.txt
# 	echo "Dummy text" >> ./themes/foo.txt

# 	touch ./resources/foo.txt
# 	echo "Dummy text" >> ./resources/foo.txt

build: ## Builds the site by invoking hugo
	echo 'Im building the website'
	hugo

clean: ## Cleaning the dist directory by removig all files in ./dist folder
	echo 'Im cleaning'
	rm -rf ./dist/*
			

# clean_site: ## Completly deletes the site
# 	echo 'Im cleaning the entire site'
# 	rm -rf ./dist
# 	rm -rf ./archetypes
# 	rm -rf ./content
# 	rm -rf ./data
# 	rm -rf ./layouts
# 	rm -rf ./public
# 	rm -rf ./static
# 	rm -rf ./themes
# 	rm -rf ./resources
# 	rm -f ./config.toml
# 	rm -f ./go.*
# 	rm -f ./.hugo_build_lock
	
post: ## Creates a new post using shell variable POST_NAME & titled POST_TITLE
	echo 'creating new post:  '$(POST_NAME).md
	hugo new posts/$(POST_NAME).md
	echo 'Adding title: '$(POST_TITLE) 'to '$(POST_NAME).md
	echo title: \"$(POST_TITLE)\" >> ./content/posts/$(POST_NAME).md
 
help: ## Prints the make targets available to invoke
	sed -n 's/##\([^\(]\)/ /p' ./Makefile 	\
	| column -t -s ':'
