FROM golang:1.16

# google cloud SDK
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg  add - && apt-get update -y && apt-get install google-cloud-sdk google-cloud-sdk-app-engine-go -y
    
WORKDIR /app

COPY . ./

RUN gcloud auth activate-service-account --key-file key.json
RUN gcloud config set project experimentation-lab-60c88
RUN go mod download

RUN go build -o /pythia
CMD [ "/pythia" ]

