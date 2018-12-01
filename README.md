### Step by step guide to deploy Golang Application on Azure Web App

*Azure App Service Web Apps* (or just Web Apps) is a service for hosting web
applications and REST APIs. You can develop web app in your favorite language,
for example .NET, .NET Core, Golang .NET, .NET Core, Golang, Java, Ruby,
Node.js, PHP, or Python. You can run and scale apps with ease on Windows or
Linux VMs.

Azure web apps also provide DevOps capability by continuous deployment from
VSTS, GitHub, Docker Hub and other sources.

This article will give you a quick introduction on how to develop a Golang web
application on Azure and automatically deploy an updated web app.

We’ll cover:

* Prerequisites for this tutorial
* Simple Golang Web Application — Ping Request
* Build script to create compiled executables
* Publish image on Docker Hub
* Publish Docker image on Azure
* Azure web app configuration (Optional)
* Basic Troubleshooting
* Azure Auto Deploy Newly published docker image

*****

### Prerequisites for the Golang tutorial

1.  Basic knowledge of [Golang ](https://golang.org/)language
1.  IDE — [GoLand](https://www.jetbrains.com/go/) by
[Jetbrains](https://www.jetbrains.com/) or [Visual Studio
Code](https://code.visualstudio.com/) by Microsoft or [Atom](https://atom.io/)

### Simple Golang Web Application — Ping Request

It’s time to get our hands dirty. Open your favorite editor (GoLand, VS Code or
Atom). For this article, I will use GoLand editor.

1.  Create AzureGo folder inside GOROOT\src folder.
1.  Create a main.go file in the AzureGo folder, and paste in the following code
snippet (code is explained in comments).

3. Run the application using `go run main.go`

4. Open a browser and hit URL `http://localhost:3005/` and you will see `hello
world!` message.

5. The code snippet above exposes 2 APIs:

* `/` will return “h*ello world!” message*
* `/ping` will return *Pong message with the current *date and time.

### Compile packages and create executable

For compiling packages, we can use `go build` or `gox`. **Gox** is a simple,
no-frills tool for Go cross compilation that behaves a lot like a standard `go
build`. Gox can parallelize builds for multiple platforms.

To install Gox, use `go get`


To compile packages and create an executable for the Linux environment, use the
following command. It will create an *amd64* arch format *Linux* executable,
named app, in the build folder.

    -osarch="linux/amd64" --output="build/app"

### Docker Image of Golang Application

Please refer to the following *Dockerfile* code snippet :

I have used [Alpine Linux](http://alpinelinux.org/) as the base docker image. It
is a Linux distribution built around [musl libc](http://www.musl-libc.org/) and
[BusyBox](http://www.busybox.net/). The image is only 5 MB in size.

To create a Docker image, use the following command:

    docker build -t durgaprasad-budhwani/azurego .

It will create a Docker image named `durgaprasad-budhwani/azurego:latest`.

![](https://cdn-images-1.medium.com/max/1600/1*7TnA7Yr_LlLFRHOyQg6nLg.png)
<span class="figcaption_hack">Docker image build logs</span>

To see a list of all images, use the command `docker image ls`

![](https://cdn-images-1.medium.com/max/1600/1*6ZMqMqj8Tr-XP-SuxIx3IQ.png)
<span class="figcaption_hack">List of docker images</span>

To run a *Docker* image locally, use the following command:

    docker run -p 5005:80 --env PORT=80 durgaprasad-budhwani/azurego

You can launch a browser with [http://localhost:5005/](http://localhost:5005/).
It will show the “*hello, world!” *message.

### Publish the image on Docker Hub

[Docker Hub](https://hub.docker.com/) is a cloud-based registry service that
allows you to link to code repositories, build your images and test them, store
manually pushed images, and link to [Docker
Cloud](https://docs.docker.com/docker-cloud/) so you can deploy images to your
hosts. It provides a centralized resource for container image discovery,
distribution and change management, [user and team
collaboration](https://docs.docker.com/docker-hub/orgs/), and workflow
automation throughout the development pipeline.

To publish an image on Docker Hub, use the following command:

    docker login -u 'Docker Username' -p $'Docker Password'

    // push docker image
    docker push durgaprasad-budhwani/azurego

You can login to [Docker Hub](https://hub.docker.com) and refer to the pushed
image:
[https://hub.docker.com/r/durgaprasadbudhwani/azurego/](https://hub.docker.com/r/durgaprasadbudhwani/azurego/)

### Publish Docker Image to Azure Web App

To host a Docker based web app in Azure, we need to perform the following steps:

* Sign in to the [Azure portal](https://portal.azure.com/). (If you don’t have an
Azure account, you can use a 12 month trial, worth $200.)
* Click on the cloud shell button as shown below — it will open a cloud PowerShell
window.

![](https://cdn-images-1.medium.com/max/1600/1*PTdDz_qzDXgtzmYoZGVeOg.png)

* To deploy an application in Azure, we need to create a resource group. Consider
a resource group to be a container for all resources.<br> To create a resource
group, use the `az group create` command. The command uses the **name**
parameter to specify a name for the resource group and the **location**
parameter to specify its location. In this tutorial, we will create a resource
group with name **AzureGoRG** and location **South Central US**.

    az group create --name AzureGoRG --location "South Central US"

![](https://cdn-images-1.medium.com/max/1600/1*kTdUWZ1uxR6PZkCR9eMdJA.png)

* To create a Linux based app service plan which will use a Linux worker to host
the Docker app, use the az appservice plan create command. You can find more
information at
[https://docs.microsoft.com/en-us/cli/azure/appservice/plan?view=azure-cli-latest](https://docs.microsoft.com/en-us/cli/azure/appservice/plan?view=azure-cli-latest).
In this tutorial, we will create a service plan with name **AzureGoSP** and
resource group **AzureGoRG**.

    az appservice plan create --name AzureGoSP --resource-group AzureGoRG --sku S1 --is-linux

* To create a Docker based web app in Azure, use the command `az webapp create
--deployment-container-image-name` command. This link
[https://docs.microsoft.com/en-us/cli/azure/webapp?view=azure-cli-latest#az_webapp_create](https://docs.microsoft.com/en-us/cli/azure/webapp?view=azure-cli-latest#az_webapp_create)
provides more information on each command. Here, we will create **AzureGoApp**
from the Docker image that we had pushed on Docker Hub,
**durgaprasadbudhwani/azurego:latest** .

    az webapp create --resource-group AzureGoRG --plan AzureGoSP  --name AzureGoApp --deployment-container-image-name durgaprasadbudhwani/azurego:latest

![](https://cdn-images-1.medium.com/max/1600/1*3416IBUZ2iFAHOlcN43y7A.png)
<span class="figcaption_hack">Result of webapp create command</span>

If you did not run into any problems, your docker image is successfully hosted
in the cloud. You will see host information in the json results in the cloud
shell window with the tag **defaultHostName:azuregoapp.azurewebsites.net*** *,
where **azuregoapp** is the app name.**You can launch your application with this
URL,
[https://azuregoapp.azurewebsites.net/](https://azuregoapp.azurewebsites.net/),
and it will show you the “hello world!” message. Azure web app configuration
(Optional)

Sometimes, you may need to set an environment variable for your Docker image
during launch, e.g. to set the port number of the Docker image, or to pass a
connection string as an environment variable. You can do this by using the az
webapp config `az webapp config appsettings set` command.

    az webapp config appsettings set --resource-group AzureGoRG --name AzureGoApp --settings PORT=80 connection_string="connection_string"

### Basic Troubleshooting

For troubleshooting, you can use
[https://azuregoapp.scm.azurewebsites.net/](https://azuregoapp.scm.azurewebsites.net/)
(substitute your app name for azuregoapp). It will open
[Kudu](https://github.com/projectkudu/kudu) engine. You can check Docker logs,
and you will have a bash/shell window where you can do more troubleshooting.

### Azure Auto Deploy Newly Published Docker image

To auto-deploy a newly published Docker image to an Azure web app, follow these
steps:

1.  Open [Azure portal](https://portal.azure.com/)
1.  Navigate to your docker container -> App Services -> Your app name -> Docker
Container (as shown below).

![](https://cdn-images-1.medium.com/max/1600/1*AP8_od71l27F31DArQOvZw.png)

3. Enable Continuous Deployment and copy the webhook URL, as shown below:

![](https://cdn-images-1.medium.com/max/1600/1*8AOAUAHc8ywZpX2TRWrZ-A.png)

4. Login to [Docker Hub](https://hub.docker.com/) site

5. Go to docker image webhook page

6. Add a webhook

![](https://cdn-images-1.medium.com/max/1600/1*HOMj4f09z-6ihjNoRyFj6Q.png)

7. Publish a new image and it will auto-deploy the new Docker image to the Azure
web app.

*****

### Source Code

Please have a look at the entire source code at GitHub.
