PROJECT_ID = go-cloudrun
SERVICE_NAME = go-cloudrun


.PHONY: build
## build: Container Build
build:
	gcloud builds submit --tag gcr.io/$(PROJECT_ID)/$(SERVICE_NAME)

.PHONY: deploy
## deploy: Deploy to Cloud Run
deploy:
	gcloud run deploy --image gcr.io/$(PROJECT_ID)/$(SERVICE_NAME) --platform managed
