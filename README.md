# go-eventgrid-webhook

This sample is an Event Grid Web Hook written in Go that supports receiving and processing a sample message type.

A complete walkthrough is available here: <https://github.com/peteroden/event-grid-go-quickstart/blob/master/event-grid-webhook-go-quickstart.md>

## Developer Prerequisites

* Git
* Go 1.11
* optional
    * Azure DevOps access to import CI/CD pipeline

## Repo Organization

* src
    * contains a sample Web Hook implementation

* azuredevops-import.json
    * Azure DevOps build pipeline import file
        * from Azure DevOps / Pipelines / Build, click "new" and choose import a pipeline then select azuredevops-import.json
            * Note that you will have to setup credentials for dockerhub and github

* sendmsg.sh
    * a simple shell script for sending a message to event grid
