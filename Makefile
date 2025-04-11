GITCMD = git

deploy-api:
	@echo "Deploying API to Heroku"
	@echo "Pushing to Heroku"
	${GITCMD} subtree push --prefix api heroku-api main

deploy-web:
	@echo "Deploying Web to Heroku"
	@echo "Pushing to Heroku"
	${GITCMD} subtree push --prefix web heroku-web main
